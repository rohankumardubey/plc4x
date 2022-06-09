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

// BACnetConstructedDataNumberOfAuthenticationPolicies is the data-structure of this message
type BACnetConstructedDataNumberOfAuthenticationPolicies struct {
	*BACnetConstructedData
	NumberOfAuthenticationPolicies *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataNumberOfAuthenticationPolicies is the corresponding interface of BACnetConstructedDataNumberOfAuthenticationPolicies
type IBACnetConstructedDataNumberOfAuthenticationPolicies interface {
	IBACnetConstructedData
	// GetNumberOfAuthenticationPolicies returns NumberOfAuthenticationPolicies (property field)
	GetNumberOfAuthenticationPolicies() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NUMBER_OF_AUTHENTICATION_POLICIES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) GetNumberOfAuthenticationPolicies() *BACnetApplicationTagUnsignedInteger {
	return m.NumberOfAuthenticationPolicies
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNumberOfAuthenticationPolicies factory function for BACnetConstructedDataNumberOfAuthenticationPolicies
func NewBACnetConstructedDataNumberOfAuthenticationPolicies(numberOfAuthenticationPolicies *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataNumberOfAuthenticationPolicies {
	_result := &BACnetConstructedDataNumberOfAuthenticationPolicies{
		NumberOfAuthenticationPolicies: numberOfAuthenticationPolicies,
		BACnetConstructedData:          NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataNumberOfAuthenticationPolicies(structType interface{}) *BACnetConstructedDataNumberOfAuthenticationPolicies {
	if casted, ok := structType.(BACnetConstructedDataNumberOfAuthenticationPolicies); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNumberOfAuthenticationPolicies); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataNumberOfAuthenticationPolicies(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataNumberOfAuthenticationPolicies(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) GetTypeName() string {
	return "BACnetConstructedDataNumberOfAuthenticationPolicies"
}

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numberOfAuthenticationPolicies)
	lengthInBits += m.NumberOfAuthenticationPolicies.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNumberOfAuthenticationPoliciesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataNumberOfAuthenticationPolicies, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNumberOfAuthenticationPolicies"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numberOfAuthenticationPolicies)
	if pullErr := readBuffer.PullContext("numberOfAuthenticationPolicies"); pullErr != nil {
		return nil, pullErr
	}
	_numberOfAuthenticationPolicies, _numberOfAuthenticationPoliciesErr := BACnetApplicationTagParse(readBuffer)
	if _numberOfAuthenticationPoliciesErr != nil {
		return nil, errors.Wrap(_numberOfAuthenticationPoliciesErr, "Error parsing 'numberOfAuthenticationPolicies' field")
	}
	numberOfAuthenticationPolicies := CastBACnetApplicationTagUnsignedInteger(_numberOfAuthenticationPolicies)
	if closeErr := readBuffer.CloseContext("numberOfAuthenticationPolicies"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNumberOfAuthenticationPolicies"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataNumberOfAuthenticationPolicies{
		NumberOfAuthenticationPolicies: CastBACnetApplicationTagUnsignedInteger(numberOfAuthenticationPolicies),
		BACnetConstructedData:          &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNumberOfAuthenticationPolicies"); pushErr != nil {
			return pushErr
		}

		// Simple Field (numberOfAuthenticationPolicies)
		if pushErr := writeBuffer.PushContext("numberOfAuthenticationPolicies"); pushErr != nil {
			return pushErr
		}
		_numberOfAuthenticationPoliciesErr := m.NumberOfAuthenticationPolicies.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("numberOfAuthenticationPolicies"); popErr != nil {
			return popErr
		}
		if _numberOfAuthenticationPoliciesErr != nil {
			return errors.Wrap(_numberOfAuthenticationPoliciesErr, "Error serializing 'numberOfAuthenticationPolicies' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNumberOfAuthenticationPolicies"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataNumberOfAuthenticationPolicies) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
