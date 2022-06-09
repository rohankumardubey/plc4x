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

// BACnetConstructedDataValueChangeTime is the data-structure of this message
type BACnetConstructedDataValueChangeTime struct {
	*BACnetConstructedData
	ValueChangeTime *BACnetDateTime

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataValueChangeTime is the corresponding interface of BACnetConstructedDataValueChangeTime
type IBACnetConstructedDataValueChangeTime interface {
	IBACnetConstructedData
	// GetValueChangeTime returns ValueChangeTime (property field)
	GetValueChangeTime() *BACnetDateTime
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

func (m *BACnetConstructedDataValueChangeTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataValueChangeTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VALUE_CHANGE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataValueChangeTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataValueChangeTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataValueChangeTime) GetValueChangeTime() *BACnetDateTime {
	return m.ValueChangeTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataValueChangeTime factory function for BACnetConstructedDataValueChangeTime
func NewBACnetConstructedDataValueChangeTime(valueChangeTime *BACnetDateTime, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataValueChangeTime {
	_result := &BACnetConstructedDataValueChangeTime{
		ValueChangeTime:       valueChangeTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataValueChangeTime(structType interface{}) *BACnetConstructedDataValueChangeTime {
	if casted, ok := structType.(BACnetConstructedDataValueChangeTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValueChangeTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataValueChangeTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataValueChangeTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataValueChangeTime) GetTypeName() string {
	return "BACnetConstructedDataValueChangeTime"
}

func (m *BACnetConstructedDataValueChangeTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataValueChangeTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (valueChangeTime)
	lengthInBits += m.ValueChangeTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataValueChangeTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataValueChangeTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataValueChangeTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValueChangeTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (valueChangeTime)
	if pullErr := readBuffer.PullContext("valueChangeTime"); pullErr != nil {
		return nil, pullErr
	}
	_valueChangeTime, _valueChangeTimeErr := BACnetDateTimeParse(readBuffer)
	if _valueChangeTimeErr != nil {
		return nil, errors.Wrap(_valueChangeTimeErr, "Error parsing 'valueChangeTime' field")
	}
	valueChangeTime := CastBACnetDateTime(_valueChangeTime)
	if closeErr := readBuffer.CloseContext("valueChangeTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValueChangeTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataValueChangeTime{
		ValueChangeTime:       CastBACnetDateTime(valueChangeTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataValueChangeTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValueChangeTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (valueChangeTime)
		if pushErr := writeBuffer.PushContext("valueChangeTime"); pushErr != nil {
			return pushErr
		}
		_valueChangeTimeErr := m.ValueChangeTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("valueChangeTime"); popErr != nil {
			return popErr
		}
		if _valueChangeTimeErr != nil {
			return errors.Wrap(_valueChangeTimeErr, "Error serializing 'valueChangeTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValueChangeTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataValueChangeTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
