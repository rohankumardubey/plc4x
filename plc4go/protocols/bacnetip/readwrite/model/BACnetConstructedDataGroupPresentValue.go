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

// BACnetConstructedDataGroupPresentValue is the data-structure of this message
type BACnetConstructedDataGroupPresentValue struct {
	*BACnetConstructedData
	PresentValue []*BACnetReadAccessResult

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataGroupPresentValue is the corresponding interface of BACnetConstructedDataGroupPresentValue
type IBACnetConstructedDataGroupPresentValue interface {
	IBACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() []*BACnetReadAccessResult
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

func (m *BACnetConstructedDataGroupPresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_GROUP
}

func (m *BACnetConstructedDataGroupPresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataGroupPresentValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataGroupPresentValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataGroupPresentValue) GetPresentValue() []*BACnetReadAccessResult {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataGroupPresentValue factory function for BACnetConstructedDataGroupPresentValue
func NewBACnetConstructedDataGroupPresentValue(presentValue []*BACnetReadAccessResult, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataGroupPresentValue {
	_result := &BACnetConstructedDataGroupPresentValue{
		PresentValue:          presentValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataGroupPresentValue(structType interface{}) *BACnetConstructedDataGroupPresentValue {
	if casted, ok := structType.(BACnetConstructedDataGroupPresentValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupPresentValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataGroupPresentValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataGroupPresentValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataGroupPresentValue) GetTypeName() string {
	return "BACnetConstructedDataGroupPresentValue"
}

func (m *BACnetConstructedDataGroupPresentValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataGroupPresentValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.PresentValue) > 0 {
		for _, element := range m.PresentValue {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataGroupPresentValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataGroupPresentValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataGroupPresentValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupPresentValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (presentValue)
	if pullErr := readBuffer.PullContext("presentValue", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	presentValue := make([]*BACnetReadAccessResult, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetReadAccessResultParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'presentValue' field")
			}
			presentValue = append(presentValue, CastBACnetReadAccessResult(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("presentValue", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupPresentValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataGroupPresentValue{
		PresentValue:          presentValue,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataGroupPresentValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupPresentValue"); pushErr != nil {
			return pushErr
		}

		// Array Field (presentValue)
		if m.PresentValue != nil {
			if pushErr := writeBuffer.PushContext("presentValue", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.PresentValue {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'presentValue' field")
				}
			}
			if popErr := writeBuffer.PopContext("presentValue", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupPresentValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataGroupPresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
