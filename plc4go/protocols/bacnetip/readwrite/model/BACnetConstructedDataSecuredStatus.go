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

// BACnetConstructedDataSecuredStatus is the data-structure of this message
type BACnetConstructedDataSecuredStatus struct {
	*BACnetConstructedData
	SecuredStatus *BACnetDoorSecuredStatusTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataSecuredStatus is the corresponding interface of BACnetConstructedDataSecuredStatus
type IBACnetConstructedDataSecuredStatus interface {
	IBACnetConstructedData
	// GetSecuredStatus returns SecuredStatus (property field)
	GetSecuredStatus() *BACnetDoorSecuredStatusTagged
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

func (m *BACnetConstructedDataSecuredStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSecuredStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SECURED_STATUS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSecuredStatus) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSecuredStatus) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSecuredStatus) GetSecuredStatus() *BACnetDoorSecuredStatusTagged {
	return m.SecuredStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSecuredStatus factory function for BACnetConstructedDataSecuredStatus
func NewBACnetConstructedDataSecuredStatus(securedStatus *BACnetDoorSecuredStatusTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataSecuredStatus {
	_result := &BACnetConstructedDataSecuredStatus{
		SecuredStatus:         securedStatus,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSecuredStatus(structType interface{}) *BACnetConstructedDataSecuredStatus {
	if casted, ok := structType.(BACnetConstructedDataSecuredStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSecuredStatus); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSecuredStatus(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSecuredStatus(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSecuredStatus) GetTypeName() string {
	return "BACnetConstructedDataSecuredStatus"
}

func (m *BACnetConstructedDataSecuredStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSecuredStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (securedStatus)
	lengthInBits += m.SecuredStatus.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataSecuredStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSecuredStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataSecuredStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSecuredStatus"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (securedStatus)
	if pullErr := readBuffer.PullContext("securedStatus"); pullErr != nil {
		return nil, pullErr
	}
	_securedStatus, _securedStatusErr := BACnetDoorSecuredStatusTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _securedStatusErr != nil {
		return nil, errors.Wrap(_securedStatusErr, "Error parsing 'securedStatus' field")
	}
	securedStatus := CastBACnetDoorSecuredStatusTagged(_securedStatus)
	if closeErr := readBuffer.CloseContext("securedStatus"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSecuredStatus"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSecuredStatus{
		SecuredStatus:         CastBACnetDoorSecuredStatusTagged(securedStatus),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSecuredStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSecuredStatus"); pushErr != nil {
			return pushErr
		}

		// Simple Field (securedStatus)
		if pushErr := writeBuffer.PushContext("securedStatus"); pushErr != nil {
			return pushErr
		}
		_securedStatusErr := m.SecuredStatus.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("securedStatus"); popErr != nil {
			return popErr
		}
		if _securedStatusErr != nil {
			return errors.Wrap(_securedStatusErr, "Error serializing 'securedStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSecuredStatus"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSecuredStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
