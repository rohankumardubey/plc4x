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

// BACnetConstructedDataLowerDeck is the data-structure of this message
type BACnetConstructedDataLowerDeck struct {
	*BACnetConstructedData
	LowerDeck *BACnetApplicationTagObjectIdentifier

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataLowerDeck is the corresponding interface of BACnetConstructedDataLowerDeck
type IBACnetConstructedDataLowerDeck interface {
	IBACnetConstructedData
	// GetLowerDeck returns LowerDeck (property field)
	GetLowerDeck() *BACnetApplicationTagObjectIdentifier
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

func (m *BACnetConstructedDataLowerDeck) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLowerDeck) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOWER_DECK
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLowerDeck) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLowerDeck) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLowerDeck) GetLowerDeck() *BACnetApplicationTagObjectIdentifier {
	return m.LowerDeck
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLowerDeck factory function for BACnetConstructedDataLowerDeck
func NewBACnetConstructedDataLowerDeck(lowerDeck *BACnetApplicationTagObjectIdentifier, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataLowerDeck {
	_result := &BACnetConstructedDataLowerDeck{
		LowerDeck:             lowerDeck,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLowerDeck(structType interface{}) *BACnetConstructedDataLowerDeck {
	if casted, ok := structType.(BACnetConstructedDataLowerDeck); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLowerDeck); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLowerDeck(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLowerDeck(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLowerDeck) GetTypeName() string {
	return "BACnetConstructedDataLowerDeck"
}

func (m *BACnetConstructedDataLowerDeck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLowerDeck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lowerDeck)
	lengthInBits += m.LowerDeck.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLowerDeck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLowerDeckParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataLowerDeck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLowerDeck"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lowerDeck)
	if pullErr := readBuffer.PullContext("lowerDeck"); pullErr != nil {
		return nil, pullErr
	}
	_lowerDeck, _lowerDeckErr := BACnetApplicationTagParse(readBuffer)
	if _lowerDeckErr != nil {
		return nil, errors.Wrap(_lowerDeckErr, "Error parsing 'lowerDeck' field")
	}
	lowerDeck := CastBACnetApplicationTagObjectIdentifier(_lowerDeck)
	if closeErr := readBuffer.CloseContext("lowerDeck"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLowerDeck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLowerDeck{
		LowerDeck:             CastBACnetApplicationTagObjectIdentifier(lowerDeck),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLowerDeck) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLowerDeck"); pushErr != nil {
			return pushErr
		}

		// Simple Field (lowerDeck)
		if pushErr := writeBuffer.PushContext("lowerDeck"); pushErr != nil {
			return pushErr
		}
		_lowerDeckErr := m.LowerDeck.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lowerDeck"); popErr != nil {
			return popErr
		}
		if _lowerDeckErr != nil {
			return errors.Wrap(_lowerDeckErr, "Error serializing 'lowerDeck' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLowerDeck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLowerDeck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
