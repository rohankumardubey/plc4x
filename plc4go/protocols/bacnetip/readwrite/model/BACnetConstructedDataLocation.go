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

// BACnetConstructedDataLocation is the data-structure of this message
type BACnetConstructedDataLocation struct {
	*BACnetConstructedData
	Location *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataLocation is the corresponding interface of BACnetConstructedDataLocation
type IBACnetConstructedDataLocation interface {
	IBACnetConstructedData
	// GetLocation returns Location (property field)
	GetLocation() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataLocation) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataLocation) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataLocation) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataLocation) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataLocation) GetLocation() *BACnetApplicationTagCharacterString {
	return m.Location
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLocation factory function for BACnetConstructedDataLocation
func NewBACnetConstructedDataLocation(location *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataLocation {
	_result := &BACnetConstructedDataLocation{
		Location:              location,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataLocation(structType interface{}) *BACnetConstructedDataLocation {
	if casted, ok := structType.(BACnetConstructedDataLocation); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLocation); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataLocation(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataLocation(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataLocation) GetTypeName() string {
	return "BACnetConstructedDataLocation"
}

func (m *BACnetConstructedDataLocation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataLocation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (location)
	lengthInBits += m.Location.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataLocation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLocationParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataLocation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLocation"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (location)
	if pullErr := readBuffer.PullContext("location"); pullErr != nil {
		return nil, pullErr
	}
	_location, _locationErr := BACnetApplicationTagParse(readBuffer)
	if _locationErr != nil {
		return nil, errors.Wrap(_locationErr, "Error parsing 'location' field")
	}
	location := CastBACnetApplicationTagCharacterString(_location)
	if closeErr := readBuffer.CloseContext("location"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLocation"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataLocation{
		Location:              CastBACnetApplicationTagCharacterString(location),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataLocation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLocation"); pushErr != nil {
			return pushErr
		}

		// Simple Field (location)
		if pushErr := writeBuffer.PushContext("location"); pushErr != nil {
			return pushErr
		}
		_locationErr := m.Location.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("location"); popErr != nil {
			return popErr
		}
		if _locationErr != nil {
			return errors.Wrap(_locationErr, "Error serializing 'location' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLocation"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataLocation) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
