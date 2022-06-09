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

// BACnetConstructedDataUserType is the data-structure of this message
type BACnetConstructedDataUserType struct {
	*BACnetConstructedData
	UserType *BACnetAccessUserTypeTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataUserType is the corresponding interface of BACnetConstructedDataUserType
type IBACnetConstructedDataUserType interface {
	IBACnetConstructedData
	// GetUserType returns UserType (property field)
	GetUserType() *BACnetAccessUserTypeTagged
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

func (m *BACnetConstructedDataUserType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataUserType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USER_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataUserType) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataUserType) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataUserType) GetUserType() *BACnetAccessUserTypeTagged {
	return m.UserType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUserType factory function for BACnetConstructedDataUserType
func NewBACnetConstructedDataUserType(userType *BACnetAccessUserTypeTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataUserType {
	_result := &BACnetConstructedDataUserType{
		UserType:              userType,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataUserType(structType interface{}) *BACnetConstructedDataUserType {
	if casted, ok := structType.(BACnetConstructedDataUserType); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUserType); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataUserType(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataUserType(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataUserType) GetTypeName() string {
	return "BACnetConstructedDataUserType"
}

func (m *BACnetConstructedDataUserType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataUserType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (userType)
	lengthInBits += m.UserType.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataUserType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUserTypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataUserType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUserType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (userType)
	if pullErr := readBuffer.PullContext("userType"); pullErr != nil {
		return nil, pullErr
	}
	_userType, _userTypeErr := BACnetAccessUserTypeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _userTypeErr != nil {
		return nil, errors.Wrap(_userTypeErr, "Error parsing 'userType' field")
	}
	userType := CastBACnetAccessUserTypeTagged(_userType)
	if closeErr := readBuffer.CloseContext("userType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUserType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataUserType{
		UserType:              CastBACnetAccessUserTypeTagged(userType),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataUserType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUserType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (userType)
		if pushErr := writeBuffer.PushContext("userType"); pushErr != nil {
			return pushErr
		}
		_userTypeErr := m.UserType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("userType"); popErr != nil {
			return popErr
		}
		if _userTypeErr != nil {
			return errors.Wrap(_userTypeErr, "Error serializing 'userType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUserType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataUserType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
