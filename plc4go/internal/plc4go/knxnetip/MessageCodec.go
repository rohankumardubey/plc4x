//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.
//
package knxnetip

import (
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/knxnetip/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/plcerrors"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/transports"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"time"
)

type Expectation struct {
	expiration     time.Time
	acceptsMessage spi.AcceptsMessage
	handleMessage  spi.HandleMessage
	handleError    spi.HandleError
}

func (m Expectation) String() string {
	return fmt.Sprintf("Expectation(expires at %v)", m.expiration)
}

type MessageCodec struct {
	sequenceCounter               int32
	transportInstance             transports.TransportInstance
	messageInterceptor            func(message interface{})
	defaultIncomingMessageChannel chan interface{}
	expectations                  []Expectation
	running                       bool
}

func NewMessageCodec(transportInstance transports.TransportInstance, messageInterceptor func(message interface{})) *MessageCodec {
	codec := &MessageCodec{
		sequenceCounter:               0,
		transportInstance:             transportInstance,
		messageInterceptor:            messageInterceptor,
		defaultIncomingMessageChannel: make(chan interface{}),
		expectations:                  []Expectation{},
		running:                       true,
	}
	// TODO: should we better move this go func into Connect(). If not a better explanation why we start the worker so early
	// Start a worker that handles processing of responses
	go func() {
		for {
			work(codec)
			// TODO: smells like antipattern: why would the loop exit randomly? because of panic? better catch them in this case
			log.Warn().Msg("Restarting Work loop")
		}
	}()
	return codec
}

func (m *MessageCodec) Connect() error {
	log.Info().Msg("Connecting")
	// "connect" to the remote UDP server
	err := m.transportInstance.Connect()
	if err == nil {
		// TODO: smells like antipattern: why would the loop exit randomly? because of panic? better catch them in this case
		/*if !m.running {
			go work(m)
		}*/
		m.running = true
	}
	return err
}

func (m *MessageCodec) Disconnect() error {
	log.Info().Msg("Disconnecting")
	m.running = false
	return m.transportInstance.Close()
}

func (m *MessageCodec) Send(message interface{}) error {
	log.Trace().Msg("Sending message")
	// Cast the message to the correct type of struct
	knxMessage := model.CastKnxNetIpMessage(message)
	// Serialize the request
	wb := utils.NewWriteBuffer()
	err := knxMessage.Serialize(*wb)
	if err != nil {
		return errors.Wrap(err, "error serializing request")
	}

	// Send it to the PLC
	err = m.transportInstance.Write(wb.GetBytes())
	if err != nil {
		return errors.Wrap(err, "error sending request ")
	}
	return nil
}

func (m *MessageCodec) Expect(acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration) error {
	expectation := Expectation{
		expiration:     time.Now().Add(ttl),
		acceptsMessage: acceptsMessage,
		handleMessage:  handleMessage,
		handleError:    handleError,
	}
	m.expectations = append(m.expectations, expectation)
	return nil
}

func (m *MessageCodec) SendRequest(message interface{}, acceptsMessage spi.AcceptsMessage, handleMessage spi.HandleMessage, handleError spi.HandleError, ttl time.Duration) error {
	log.Trace().Msg("Sending request")
	// Send the actual message
	err := m.Send(message)
	if err != nil {
		return errors.Wrap(err, "Error sending the request")
	}
	return m.Expect(acceptsMessage, handleMessage, handleError, ttl)
}

func (m *MessageCodec) GetDefaultIncomingMessageChannel() chan interface{} {
	return m.defaultIncomingMessageChannel
}

func (m *MessageCodec) receive() (interface{}, error) {
	log.Trace().Msg("receiving")
	// We need at least 6 bytes in order to know how big the packet is in total
	if num, err := m.transportInstance.GetNumReadableBytes(); (err == nil) && (num >= 6) {
		log.Debug().Msgf("we got %d readable bytes", num)
		data, err := m.transportInstance.PeekReadableBytes(6)
		if err != nil {
			log.Warn().Err(err).Msg("error peeking")
			// TODO: Possibly clean up ...
			return nil, nil
		}
		// Get the size of the entire packet
		packetSize := (uint32(data[4]) << 8) + uint32(data[5])
		if num < packetSize {
			log.Debug().Msgf("Not enough bytes. Got: %d Need: %d\n", num, packetSize)
			return nil, nil
		}
		data, err = m.transportInstance.Read(packetSize)
		if err != nil {
			log.Warn().Err(err).Msg("error reading")
			// TODO: Possibly clean up ...
			return nil, nil
		}
		rb := utils.NewReadBuffer(data)
		knxMessage, err := model.KnxNetIpMessageParse(rb)
		if err != nil {
			log.Warn().Err(err).Msg("error parsing message")
			// TODO: Possibly clean up ...
			return nil, nil
		}
		return knxMessage, nil
	} else if err != nil {
		log.Warn().Err(err).Msg("Got error reading")
		return nil, nil
	}
	return nil, nil
}

func work(m *MessageCodec) {
	// If this method ever exits, start it again.
	// TODO: smells like antipattern: why would the loop exit randomly? because of panic? better catch them in this case
	defer work(m)
	// Start an endless loop
mainLoop:
	for m.running {
		// Check for any expired expectations.
		// (Doing this outside the loop lets us expire expectations even if no input is coming in)
		now := time.Now()

		// Guard against empty expectations
		if len(m.expectations) <= 0 {
			log.Trace().Msg("we got no expectation")
			// Sleep for 10ms
			time.Sleep(time.Millisecond * 10)
			continue mainLoop
		}
	expectationLoop:
		for index, expectation := range m.expectations {
			// Check if this expectation has expired.
			if now.After(expectation.expiration) {
				// Remove this expectation from the list.
				m.expectations = append(m.expectations[:index], m.expectations[index+1:]...)
				// Call the error handler.
				err := expectation.handleError(plcerrors.NewTimeoutError(now.Sub(expectation.expiration)))
				if err != nil {
					log.Error().Err(err).Msg("Got an error handling error on expectation")
				}
				continue expectationLoop
			}
		}

		// Check for incoming messages.
		message, err := m.receive()
		if err != nil {
			log.Error().Err(err).Msg("got an error reading from transport")
			time.Sleep(time.Millisecond * 10)
			continue mainLoop
		}
		if message == nil {
			// Sleep for 10ms before checking again, in order to not
			// consume 100% CPU Power.
			time.Sleep(time.Millisecond * 10)
			continue mainLoop
		}

		// If this message is a simple KNXNet/IP UDP Ack, ignore it for now
		tunnelingResponse := model.CastTunnelingResponse(message)
		if tunnelingResponse != nil {
			continue mainLoop
		}

		// If this is an incoming tunneling request, automatically send a tunneling ACK back to the gateway
		tunnelingRequest := model.CastTunnelingRequest(message)
		if tunnelingRequest != nil {
			response := model.NewTunnelingResponse(
				model.NewTunnelingResponseDataBlock(
					tunnelingRequest.TunnelingRequestDataBlock.CommunicationChannelId,
					tunnelingRequest.TunnelingRequestDataBlock.SequenceCounter,
					model.Status_NO_ERROR),
			)
			err = m.Send(response)
			if err != nil {
				log.Warn().Err(err).Msg("got an error sending ACK from transport")
			}
		}

		// Handle the packet itself
		// Give a message interceptor a chance to intercept
		if m.messageInterceptor != nil {
			m.messageInterceptor(message)
		}

		// Go through all expectations
		messageHandled := false
		for index, expectation := range m.expectations {
			// Check if the current message matches the expectations
			// If it does, let it handle the message.
			if accepts := expectation.acceptsMessage(message); accepts {
				log.Debug().Stringer("expectation", expectation).Msg("accepts message")
				err = expectation.handleMessage(message)
				if err == nil {
					messageHandled = true
					// If this is the last element of the list remove it differently than if it's before that
					if (index + 1) == len(m.expectations) {
						m.expectations = m.expectations[:index]
					} else if (index + 1) < len(m.expectations) {
						m.expectations = append(m.expectations[:index], m.expectations[index+1:]...)
					}
				} else {
					// Pass the error to the error handler.
					err := expectation.handleError(err)
					if err != nil {
						log.Error().Err(err).Msg("Got an error handling error on expectation")
					}
				}
			}
		}

		// If the message has not been handled and a default handler is provided, call this ...
		if !messageHandled {
			// TODO: how do we prevent endless blocking if there is no reader on this channel?
			m.defaultIncomingMessageChannel <- message
		}
	}
}

func (m MessageCodec) GetTransportInstance() transports.TransportInstance {
	return m.transportInstance
}
