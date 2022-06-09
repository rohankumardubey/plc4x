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

// BACnetConstructedDataAverageValue is the data-structure of this message
type BACnetConstructedDataAverageValue struct {
	*BACnetConstructedData
	AverageValue *BACnetApplicationTagReal

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataAverageValue is the corresponding interface of BACnetConstructedDataAverageValue
type IBACnetConstructedDataAverageValue interface {
	IBACnetConstructedData
	// GetAverageValue returns AverageValue (property field)
	GetAverageValue() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataAverageValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAverageValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AVERAGE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAverageValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAverageValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAverageValue) GetAverageValue() *BACnetApplicationTagReal {
	return m.AverageValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAverageValue factory function for BACnetConstructedDataAverageValue
func NewBACnetConstructedDataAverageValue(averageValue *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataAverageValue {
	_result := &BACnetConstructedDataAverageValue{
		AverageValue:          averageValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAverageValue(structType interface{}) *BACnetConstructedDataAverageValue {
	if casted, ok := structType.(BACnetConstructedDataAverageValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAverageValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAverageValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAverageValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAverageValue) GetTypeName() string {
	return "BACnetConstructedDataAverageValue"
}

func (m *BACnetConstructedDataAverageValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAverageValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (averageValue)
	lengthInBits += m.AverageValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAverageValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAverageValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataAverageValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAverageValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (averageValue)
	if pullErr := readBuffer.PullContext("averageValue"); pullErr != nil {
		return nil, pullErr
	}
	_averageValue, _averageValueErr := BACnetApplicationTagParse(readBuffer)
	if _averageValueErr != nil {
		return nil, errors.Wrap(_averageValueErr, "Error parsing 'averageValue' field")
	}
	averageValue := CastBACnetApplicationTagReal(_averageValue)
	if closeErr := readBuffer.CloseContext("averageValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAverageValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAverageValue{
		AverageValue:          CastBACnetApplicationTagReal(averageValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAverageValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAverageValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (averageValue)
		if pushErr := writeBuffer.PushContext("averageValue"); pushErr != nil {
			return pushErr
		}
		_averageValueErr := m.AverageValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("averageValue"); popErr != nil {
			return popErr
		}
		if _averageValueErr != nil {
			return errors.Wrap(_averageValueErr, "Error serializing 'averageValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAverageValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAverageValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
