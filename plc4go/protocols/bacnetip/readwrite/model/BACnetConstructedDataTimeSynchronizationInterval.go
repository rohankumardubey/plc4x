/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package model

import (
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataTimeSynchronizationInterval is the data-structure of this message
type BACnetConstructedDataTimeSynchronizationInterval struct {
	*BACnetConstructedData
	TimeSynchronization *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataTimeSynchronizationInterval is the corresponding interface of BACnetConstructedDataTimeSynchronizationInterval
type IBACnetConstructedDataTimeSynchronizationInterval interface {
	IBACnetConstructedData
	// GetTimeSynchronization returns TimeSynchronization (property field)
	GetTimeSynchronization() *BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataTimeSynchronizationInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataTimeSynchronizationInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_SYNCHRONIZATION_INTERVAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTimeSynchronizationInterval) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTimeSynchronizationInterval) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTimeSynchronizationInterval) GetTimeSynchronization() *BACnetApplicationTagUnsignedInteger {
	return m.TimeSynchronization
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeSynchronizationInterval factory function for BACnetConstructedDataTimeSynchronizationInterval
func NewBACnetConstructedDataTimeSynchronizationInterval(timeSynchronization *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataTimeSynchronizationInterval {
	_result := &BACnetConstructedDataTimeSynchronizationInterval{
		TimeSynchronization:   timeSynchronization,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTimeSynchronizationInterval(structType interface{}) *BACnetConstructedDataTimeSynchronizationInterval {
	if casted, ok := structType.(BACnetConstructedDataTimeSynchronizationInterval); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeSynchronizationInterval); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeSynchronizationInterval(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeSynchronizationInterval(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTimeSynchronizationInterval) GetTypeName() string {
	return "BACnetConstructedDataTimeSynchronizationInterval"
}

func (m *BACnetConstructedDataTimeSynchronizationInterval) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTimeSynchronizationInterval) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeSynchronization)
	lengthInBits += m.TimeSynchronization.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTimeSynchronizationInterval) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimeSynchronizationIntervalParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataTimeSynchronizationInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeSynchronizationInterval"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeSynchronization)
	if pullErr := readBuffer.PullContext("timeSynchronization"); pullErr != nil {
		return nil, pullErr
	}
	_timeSynchronization, _timeSynchronizationErr := BACnetApplicationTagParse(readBuffer)
	if _timeSynchronizationErr != nil {
		return nil, errors.Wrap(_timeSynchronizationErr, "Error parsing 'timeSynchronization' field")
	}
	timeSynchronization := CastBACnetApplicationTagUnsignedInteger(_timeSynchronization)
	if closeErr := readBuffer.CloseContext("timeSynchronization"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeSynchronizationInterval"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTimeSynchronizationInterval{
		TimeSynchronization:   CastBACnetApplicationTagUnsignedInteger(timeSynchronization),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTimeSynchronizationInterval) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeSynchronizationInterval"); pushErr != nil {
			return pushErr
		}

		// Simple Field (timeSynchronization)
		if pushErr := writeBuffer.PushContext("timeSynchronization"); pushErr != nil {
			return pushErr
		}
		_timeSynchronizationErr := m.TimeSynchronization.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeSynchronization"); popErr != nil {
			return popErr
		}
		if _timeSynchronizationErr != nil {
			return errors.Wrap(_timeSynchronizationErr, "Error serializing 'timeSynchronization' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeSynchronizationInterval"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTimeSynchronizationInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
