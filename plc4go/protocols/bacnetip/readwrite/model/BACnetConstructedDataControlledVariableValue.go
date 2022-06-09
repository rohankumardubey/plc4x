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

// BACnetConstructedDataControlledVariableValue is the data-structure of this message
type BACnetConstructedDataControlledVariableValue struct {
	*BACnetConstructedData
	ControlledVariableValue *BACnetApplicationTagReal

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataControlledVariableValue is the corresponding interface of BACnetConstructedDataControlledVariableValue
type IBACnetConstructedDataControlledVariableValue interface {
	IBACnetConstructedData
	// GetControlledVariableValue returns ControlledVariableValue (property field)
	GetControlledVariableValue() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataControlledVariableValue) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataControlledVariableValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROLLED_VARIABLE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataControlledVariableValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataControlledVariableValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataControlledVariableValue) GetControlledVariableValue() *BACnetApplicationTagReal {
	return m.ControlledVariableValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataControlledVariableValue factory function for BACnetConstructedDataControlledVariableValue
func NewBACnetConstructedDataControlledVariableValue(controlledVariableValue *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataControlledVariableValue {
	_result := &BACnetConstructedDataControlledVariableValue{
		ControlledVariableValue: controlledVariableValue,
		BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataControlledVariableValue(structType interface{}) *BACnetConstructedDataControlledVariableValue {
	if casted, ok := structType.(BACnetConstructedDataControlledVariableValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlledVariableValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataControlledVariableValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataControlledVariableValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataControlledVariableValue) GetTypeName() string {
	return "BACnetConstructedDataControlledVariableValue"
}

func (m *BACnetConstructedDataControlledVariableValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataControlledVariableValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (controlledVariableValue)
	lengthInBits += m.ControlledVariableValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataControlledVariableValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataControlledVariableValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataControlledVariableValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlledVariableValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (controlledVariableValue)
	if pullErr := readBuffer.PullContext("controlledVariableValue"); pullErr != nil {
		return nil, pullErr
	}
	_controlledVariableValue, _controlledVariableValueErr := BACnetApplicationTagParse(readBuffer)
	if _controlledVariableValueErr != nil {
		return nil, errors.Wrap(_controlledVariableValueErr, "Error parsing 'controlledVariableValue' field")
	}
	controlledVariableValue := CastBACnetApplicationTagReal(_controlledVariableValue)
	if closeErr := readBuffer.CloseContext("controlledVariableValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlledVariableValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataControlledVariableValue{
		ControlledVariableValue: CastBACnetApplicationTagReal(controlledVariableValue),
		BACnetConstructedData:   &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataControlledVariableValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlledVariableValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (controlledVariableValue)
		if pushErr := writeBuffer.PushContext("controlledVariableValue"); pushErr != nil {
			return pushErr
		}
		_controlledVariableValueErr := m.ControlledVariableValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("controlledVariableValue"); popErr != nil {
			return popErr
		}
		if _controlledVariableValueErr != nil {
			return errors.Wrap(_controlledVariableValueErr, "Error serializing 'controlledVariableValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlledVariableValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataControlledVariableValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
