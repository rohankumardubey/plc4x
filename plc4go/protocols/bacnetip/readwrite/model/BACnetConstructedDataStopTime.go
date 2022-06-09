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

// BACnetConstructedDataStopTime is the data-structure of this message
type BACnetConstructedDataStopTime struct {
	*BACnetConstructedData
	StopTime *BACnetDateTime

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataStopTime is the corresponding interface of BACnetConstructedDataStopTime
type IBACnetConstructedDataStopTime interface {
	IBACnetConstructedData
	// GetStopTime returns StopTime (property field)
	GetStopTime() *BACnetDateTime
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

func (m *BACnetConstructedDataStopTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataStopTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_STOP_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataStopTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataStopTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataStopTime) GetStopTime() *BACnetDateTime {
	return m.StopTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataStopTime factory function for BACnetConstructedDataStopTime
func NewBACnetConstructedDataStopTime(stopTime *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataStopTime {
	_result := &BACnetConstructedDataStopTime{
		StopTime:              stopTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataStopTime(structType interface{}) *BACnetConstructedDataStopTime {
	if casted, ok := structType.(BACnetConstructedDataStopTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStopTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataStopTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataStopTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataStopTime) GetTypeName() string {
	return "BACnetConstructedDataStopTime"
}

func (m *BACnetConstructedDataStopTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataStopTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (stopTime)
	lengthInBits += m.StopTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataStopTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataStopTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataStopTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStopTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (stopTime)
	if pullErr := readBuffer.PullContext("stopTime"); pullErr != nil {
		return nil, pullErr
	}
	_stopTime, _stopTimeErr := BACnetDateTimeParse(readBuffer)
	if _stopTimeErr != nil {
		return nil, errors.Wrap(_stopTimeErr, "Error parsing 'stopTime' field")
	}
	stopTime := CastBACnetDateTime(_stopTime)
	if closeErr := readBuffer.CloseContext("stopTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStopTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataStopTime{
		StopTime:              CastBACnetDateTime(stopTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataStopTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStopTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (stopTime)
		if pushErr := writeBuffer.PushContext("stopTime"); pushErr != nil {
			return pushErr
		}
		_stopTimeErr := m.StopTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("stopTime"); popErr != nil {
			return popErr
		}
		if _stopTimeErr != nil {
			return errors.Wrap(_stopTimeErr, "Error serializing 'stopTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStopTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataStopTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
