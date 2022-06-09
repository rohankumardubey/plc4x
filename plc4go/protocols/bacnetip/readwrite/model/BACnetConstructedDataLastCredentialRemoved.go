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

// BACnetConstructedDataLastCredentialRemoved is the data-structure of this message
type BACnetConstructedDataLastCredentialRemoved struct {
	*BACnetConstructedData
	LastCredentialRemoved *BACnetDeviceObjectReference

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataLastCredentialRemoved is the corresponding interface of BACnetConstructedDataLastCredentialRemoved
type IBACnetConstructedDataLastCredentialRemoved interface {
	IBACnetConstructedData
	// GetLastCredentialRemoved returns LastCredentialRemoved (property field)
	GetLastCredentialRemoved() *BACnetDeviceObjectReference
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

func (m *BACnetConstructedDataLastCredentialRemoved) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLastCredentialRemoved) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_CREDENTIAL_REMOVED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLastCredentialRemoved) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLastCredentialRemoved) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLastCredentialRemoved) GetLastCredentialRemoved() *BACnetDeviceObjectReference {
	return m.LastCredentialRemoved
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastCredentialRemoved factory function for BACnetConstructedDataLastCredentialRemoved
func NewBACnetConstructedDataLastCredentialRemoved(lastCredentialRemoved *BACnetDeviceObjectReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataLastCredentialRemoved {
	_result := &BACnetConstructedDataLastCredentialRemoved{
		LastCredentialRemoved: lastCredentialRemoved,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLastCredentialRemoved(structType interface{}) *BACnetConstructedDataLastCredentialRemoved {
	if casted, ok := structType.(BACnetConstructedDataLastCredentialRemoved); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCredentialRemoved); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastCredentialRemoved(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLastCredentialRemoved(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLastCredentialRemoved) GetTypeName() string {
	return "BACnetConstructedDataLastCredentialRemoved"
}

func (m *BACnetConstructedDataLastCredentialRemoved) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLastCredentialRemoved) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lastCredentialRemoved)
	lengthInBits += m.LastCredentialRemoved.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLastCredentialRemoved) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLastCredentialRemovedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataLastCredentialRemoved, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCredentialRemoved"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastCredentialRemoved)
	if pullErr := readBuffer.PullContext("lastCredentialRemoved"); pullErr != nil {
		return nil, pullErr
	}
	_lastCredentialRemoved, _lastCredentialRemovedErr := BACnetDeviceObjectReferenceParse(readBuffer)
	if _lastCredentialRemovedErr != nil {
		return nil, errors.Wrap(_lastCredentialRemovedErr, "Error parsing 'lastCredentialRemoved' field")
	}
	lastCredentialRemoved := CastBACnetDeviceObjectReference(_lastCredentialRemoved)
	if closeErr := readBuffer.CloseContext("lastCredentialRemoved"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCredentialRemoved"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLastCredentialRemoved{
		LastCredentialRemoved: CastBACnetDeviceObjectReference(lastCredentialRemoved),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLastCredentialRemoved) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCredentialRemoved"); pushErr != nil {
			return pushErr
		}

		// Simple Field (lastCredentialRemoved)
		if pushErr := writeBuffer.PushContext("lastCredentialRemoved"); pushErr != nil {
			return pushErr
		}
		_lastCredentialRemovedErr := m.LastCredentialRemoved.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lastCredentialRemoved"); popErr != nil {
			return popErr
		}
		if _lastCredentialRemovedErr != nil {
			return errors.Wrap(_lastCredentialRemovedErr, "Error serializing 'lastCredentialRemoved' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCredentialRemoved"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLastCredentialRemoved) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
