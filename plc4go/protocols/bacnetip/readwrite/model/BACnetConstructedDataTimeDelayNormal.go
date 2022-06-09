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

// BACnetConstructedDataTimeDelayNormal is the data-structure of this message
type BACnetConstructedDataTimeDelayNormal struct {
	*BACnetConstructedData
	TimeDelayNormal *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataTimeDelayNormal is the corresponding interface of BACnetConstructedDataTimeDelayNormal
type IBACnetConstructedDataTimeDelayNormal interface {
	IBACnetConstructedData
	// GetTimeDelayNormal returns TimeDelayNormal (property field)
	GetTimeDelayNormal() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataTimeDelayNormal) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataTimeDelayNormal) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_DELAY_NORMAL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTimeDelayNormal) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTimeDelayNormal) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTimeDelayNormal) GetTimeDelayNormal() *BACnetApplicationTagUnsignedInteger {
	return m.TimeDelayNormal
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeDelayNormal factory function for BACnetConstructedDataTimeDelayNormal
func NewBACnetConstructedDataTimeDelayNormal(timeDelayNormal *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataTimeDelayNormal {
	_result := &BACnetConstructedDataTimeDelayNormal{
		TimeDelayNormal:       timeDelayNormal,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTimeDelayNormal(structType interface{}) *BACnetConstructedDataTimeDelayNormal {
	if casted, ok := structType.(BACnetConstructedDataTimeDelayNormal); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeDelayNormal); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeDelayNormal(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimeDelayNormal(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTimeDelayNormal) GetTypeName() string {
	return "BACnetConstructedDataTimeDelayNormal"
}

func (m *BACnetConstructedDataTimeDelayNormal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTimeDelayNormal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeDelayNormal)
	lengthInBits += m.TimeDelayNormal.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTimeDelayNormal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimeDelayNormalParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataTimeDelayNormal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeDelayNormal"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeDelayNormal)
	if pullErr := readBuffer.PullContext("timeDelayNormal"); pullErr != nil {
		return nil, pullErr
	}
	_timeDelayNormal, _timeDelayNormalErr := BACnetApplicationTagParse(readBuffer)
	if _timeDelayNormalErr != nil {
		return nil, errors.Wrap(_timeDelayNormalErr, "Error parsing 'timeDelayNormal' field")
	}
	timeDelayNormal := CastBACnetApplicationTagUnsignedInteger(_timeDelayNormal)
	if closeErr := readBuffer.CloseContext("timeDelayNormal"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeDelayNormal"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTimeDelayNormal{
		TimeDelayNormal:       CastBACnetApplicationTagUnsignedInteger(timeDelayNormal),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTimeDelayNormal) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeDelayNormal"); pushErr != nil {
			return pushErr
		}

		// Simple Field (timeDelayNormal)
		if pushErr := writeBuffer.PushContext("timeDelayNormal"); pushErr != nil {
			return pushErr
		}
		_timeDelayNormalErr := m.TimeDelayNormal.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeDelayNormal"); popErr != nil {
			return popErr
		}
		if _timeDelayNormalErr != nil {
			return errors.Wrap(_timeDelayNormalErr, "Error serializing 'timeDelayNormal' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeDelayNormal"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTimeDelayNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
