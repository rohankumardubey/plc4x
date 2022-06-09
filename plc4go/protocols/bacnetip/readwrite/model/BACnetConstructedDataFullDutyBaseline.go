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

// BACnetConstructedDataFullDutyBaseline is the data-structure of this message
type BACnetConstructedDataFullDutyBaseline struct {
	*BACnetConstructedData
	FullDutyBaseLine *BACnetApplicationTagReal

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataFullDutyBaseline is the corresponding interface of BACnetConstructedDataFullDutyBaseline
type IBACnetConstructedDataFullDutyBaseline interface {
	IBACnetConstructedData
	// GetFullDutyBaseLine returns FullDutyBaseLine (property field)
	GetFullDutyBaseLine() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataFullDutyBaseline) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataFullDutyBaseline) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FULL_DUTY_BASELINE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataFullDutyBaseline) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataFullDutyBaseline) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataFullDutyBaseline) GetFullDutyBaseLine() *BACnetApplicationTagReal {
	return m.FullDutyBaseLine
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFullDutyBaseline factory function for BACnetConstructedDataFullDutyBaseline
func NewBACnetConstructedDataFullDutyBaseline(fullDutyBaseLine *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataFullDutyBaseline {
	_result := &BACnetConstructedDataFullDutyBaseline{
		FullDutyBaseLine:      fullDutyBaseLine,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataFullDutyBaseline(structType interface{}) *BACnetConstructedDataFullDutyBaseline {
	if casted, ok := structType.(BACnetConstructedDataFullDutyBaseline); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFullDutyBaseline); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataFullDutyBaseline(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataFullDutyBaseline(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataFullDutyBaseline) GetTypeName() string {
	return "BACnetConstructedDataFullDutyBaseline"
}

func (m *BACnetConstructedDataFullDutyBaseline) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataFullDutyBaseline) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fullDutyBaseLine)
	lengthInBits += m.FullDutyBaseLine.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataFullDutyBaseline) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataFullDutyBaselineParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataFullDutyBaseline, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFullDutyBaseline"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fullDutyBaseLine)
	if pullErr := readBuffer.PullContext("fullDutyBaseLine"); pullErr != nil {
		return nil, pullErr
	}
	_fullDutyBaseLine, _fullDutyBaseLineErr := BACnetApplicationTagParse(readBuffer)
	if _fullDutyBaseLineErr != nil {
		return nil, errors.Wrap(_fullDutyBaseLineErr, "Error parsing 'fullDutyBaseLine' field")
	}
	fullDutyBaseLine := CastBACnetApplicationTagReal(_fullDutyBaseLine)
	if closeErr := readBuffer.CloseContext("fullDutyBaseLine"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFullDutyBaseline"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataFullDutyBaseline{
		FullDutyBaseLine:      CastBACnetApplicationTagReal(fullDutyBaseLine),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataFullDutyBaseline) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFullDutyBaseline"); pushErr != nil {
			return pushErr
		}

		// Simple Field (fullDutyBaseLine)
		if pushErr := writeBuffer.PushContext("fullDutyBaseLine"); pushErr != nil {
			return pushErr
		}
		_fullDutyBaseLineErr := m.FullDutyBaseLine.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("fullDutyBaseLine"); popErr != nil {
			return popErr
		}
		if _fullDutyBaseLineErr != nil {
			return errors.Wrap(_fullDutyBaseLineErr, "Error serializing 'fullDutyBaseLine' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFullDutyBaseline"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataFullDutyBaseline) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
