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

// BACnetConstructedDataTimeOfStateCountReset is the data-structure of this message
type BACnetConstructedDataTimeOfStateCountReset struct {
	*BACnetConstructedData
	TimeOfStateCountReset *BACnetDateTime

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataTimeOfStateCountReset is the corresponding interface of BACnetConstructedDataTimeOfStateCountReset
type IBACnetConstructedDataTimeOfStateCountReset interface {
	IBACnetConstructedData
	// GetTimeOfStateCountReset returns TimeOfStateCountReset (property field)
	GetTimeOfStateCountReset() *BACnetDateTime
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

func (m *BACnetConstructedDataTimeOfStateCountReset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataTimeOfStateCountReset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_OF_STATE_COUNT_RESET
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTimeOfStateCountReset) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTimeOfStateCountReset) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTimeOfStateCountReset) GetTimeOfStateCountReset() *BACnetDateTime {
	return m.TimeOfStateCountReset
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeOfStateCountReset factory function for BACnetConstructedDataTimeOfStateCountReset
func NewBACnetConstructedDataTimeOfStateCountReset(timeOfStateCountReset *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataTimeOfStateCountReset {
	_result := &BACnetConstructedDataTimeOfStateCountReset{
		TimeOfStateCountReset: timeOfStateCountReset,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTimeOfStateCountReset(structType interface{}) *BACnetConstructedDataTimeOfStateCountReset {
	if casted, ok := structType.(BACnetConstructedDataTimeOfStateCountReset); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeOfStateCountReset); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeOfStateCountReset(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeOfStateCountReset(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTimeOfStateCountReset) GetTypeName() string {
	return "BACnetConstructedDataTimeOfStateCountReset"
}

func (m *BACnetConstructedDataTimeOfStateCountReset) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTimeOfStateCountReset) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeOfStateCountReset)
	lengthInBits += m.TimeOfStateCountReset.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTimeOfStateCountReset) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimeOfStateCountResetParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataTimeOfStateCountReset, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeOfStateCountReset"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeOfStateCountReset)
	if pullErr := readBuffer.PullContext("timeOfStateCountReset"); pullErr != nil {
		return nil, pullErr
	}
	_timeOfStateCountReset, _timeOfStateCountResetErr := BACnetDateTimeParse(readBuffer)
	if _timeOfStateCountResetErr != nil {
		return nil, errors.Wrap(_timeOfStateCountResetErr, "Error parsing 'timeOfStateCountReset' field")
	}
	timeOfStateCountReset := CastBACnetDateTime(_timeOfStateCountReset)
	if closeErr := readBuffer.CloseContext("timeOfStateCountReset"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeOfStateCountReset"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTimeOfStateCountReset{
		TimeOfStateCountReset: CastBACnetDateTime(timeOfStateCountReset),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTimeOfStateCountReset) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeOfStateCountReset"); pushErr != nil {
			return pushErr
		}

		// Simple Field (timeOfStateCountReset)
		if pushErr := writeBuffer.PushContext("timeOfStateCountReset"); pushErr != nil {
			return pushErr
		}
		_timeOfStateCountResetErr := m.TimeOfStateCountReset.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeOfStateCountReset"); popErr != nil {
			return popErr
		}
		if _timeOfStateCountResetErr != nil {
			return errors.Wrap(_timeOfStateCountResetErr, "Error serializing 'timeOfStateCountReset' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeOfStateCountReset"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTimeOfStateCountReset) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
