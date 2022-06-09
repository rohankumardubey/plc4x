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

// BACnetConstructedDataAuthorizationMode is the data-structure of this message
type BACnetConstructedDataAuthorizationMode struct {
	*BACnetConstructedData
	AuthorizationMode *BACnetAuthorizationModeTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataAuthorizationMode is the corresponding interface of BACnetConstructedDataAuthorizationMode
type IBACnetConstructedDataAuthorizationMode interface {
	IBACnetConstructedData
	// GetAuthorizationMode returns AuthorizationMode (property field)
	GetAuthorizationMode() *BACnetAuthorizationModeTagged
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

func (m *BACnetConstructedDataAuthorizationMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAuthorizationMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHORIZATION_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAuthorizationMode) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAuthorizationMode) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAuthorizationMode) GetAuthorizationMode() *BACnetAuthorizationModeTagged {
	return m.AuthorizationMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAuthorizationMode factory function for BACnetConstructedDataAuthorizationMode
func NewBACnetConstructedDataAuthorizationMode(authorizationMode *BACnetAuthorizationModeTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataAuthorizationMode {
	_result := &BACnetConstructedDataAuthorizationMode{
		AuthorizationMode:     authorizationMode,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAuthorizationMode(structType interface{}) *BACnetConstructedDataAuthorizationMode {
	if casted, ok := structType.(BACnetConstructedDataAuthorizationMode); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthorizationMode); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAuthorizationMode(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAuthorizationMode(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAuthorizationMode) GetTypeName() string {
	return "BACnetConstructedDataAuthorizationMode"
}

func (m *BACnetConstructedDataAuthorizationMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAuthorizationMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (authorizationMode)
	lengthInBits += m.AuthorizationMode.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataAuthorizationMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAuthorizationModeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataAuthorizationMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthorizationMode"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (authorizationMode)
	if pullErr := readBuffer.PullContext("authorizationMode"); pullErr != nil {
		return nil, pullErr
	}
	_authorizationMode, _authorizationModeErr := BACnetAuthorizationModeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _authorizationModeErr != nil {
		return nil, errors.Wrap(_authorizationModeErr, "Error parsing 'authorizationMode' field")
	}
	authorizationMode := CastBACnetAuthorizationModeTagged(_authorizationMode)
	if closeErr := readBuffer.CloseContext("authorizationMode"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthorizationMode"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAuthorizationMode{
		AuthorizationMode:     CastBACnetAuthorizationModeTagged(authorizationMode),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAuthorizationMode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthorizationMode"); pushErr != nil {
			return pushErr
		}

		// Simple Field (authorizationMode)
		if pushErr := writeBuffer.PushContext("authorizationMode"); pushErr != nil {
			return pushErr
		}
		_authorizationModeErr := m.AuthorizationMode.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("authorizationMode"); popErr != nil {
			return popErr
		}
		if _authorizationModeErr != nil {
			return errors.Wrap(_authorizationModeErr, "Error serializing 'authorizationMode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthorizationMode"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAuthorizationMode) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
