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

// BACnetConstructedDataUsesRemaining is the data-structure of this message
type BACnetConstructedDataUsesRemaining struct {
	*BACnetConstructedData
	UsesRemaining *BACnetApplicationTagSignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataUsesRemaining is the corresponding interface of BACnetConstructedDataUsesRemaining
type IBACnetConstructedDataUsesRemaining interface {
	IBACnetConstructedData
	// GetUsesRemaining returns UsesRemaining (property field)
	GetUsesRemaining() *BACnetApplicationTagSignedInteger
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

func (m *BACnetConstructedDataUsesRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataUsesRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USES_REMAINING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataUsesRemaining) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataUsesRemaining) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataUsesRemaining) GetUsesRemaining() *BACnetApplicationTagSignedInteger {
	return m.UsesRemaining
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUsesRemaining factory function for BACnetConstructedDataUsesRemaining
func NewBACnetConstructedDataUsesRemaining(usesRemaining *BACnetApplicationTagSignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataUsesRemaining {
	_result := &BACnetConstructedDataUsesRemaining{
		UsesRemaining:         usesRemaining,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataUsesRemaining(structType interface{}) *BACnetConstructedDataUsesRemaining {
	if casted, ok := structType.(BACnetConstructedDataUsesRemaining); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUsesRemaining); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataUsesRemaining(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataUsesRemaining(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataUsesRemaining) GetTypeName() string {
	return "BACnetConstructedDataUsesRemaining"
}

func (m *BACnetConstructedDataUsesRemaining) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataUsesRemaining) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (usesRemaining)
	lengthInBits += m.UsesRemaining.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataUsesRemaining) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUsesRemainingParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataUsesRemaining, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUsesRemaining"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (usesRemaining)
	if pullErr := readBuffer.PullContext("usesRemaining"); pullErr != nil {
		return nil, pullErr
	}
	_usesRemaining, _usesRemainingErr := BACnetApplicationTagParse(readBuffer)
	if _usesRemainingErr != nil {
		return nil, errors.Wrap(_usesRemainingErr, "Error parsing 'usesRemaining' field")
	}
	usesRemaining := CastBACnetApplicationTagSignedInteger(_usesRemaining)
	if closeErr := readBuffer.CloseContext("usesRemaining"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUsesRemaining"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataUsesRemaining{
		UsesRemaining:         CastBACnetApplicationTagSignedInteger(usesRemaining),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataUsesRemaining) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUsesRemaining"); pushErr != nil {
			return pushErr
		}

		// Simple Field (usesRemaining)
		if pushErr := writeBuffer.PushContext("usesRemaining"); pushErr != nil {
			return pushErr
		}
		_usesRemainingErr := m.UsesRemaining.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("usesRemaining"); popErr != nil {
			return popErr
		}
		if _usesRemainingErr != nil {
			return errors.Wrap(_usesRemainingErr, "Error serializing 'usesRemaining' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUsesRemaining"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataUsesRemaining) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
