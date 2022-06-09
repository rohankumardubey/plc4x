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

// BACnetConstructedDataBACnetIPv6Mode is the data-structure of this message
type BACnetConstructedDataBACnetIPv6Mode struct {
	*BACnetConstructedData
	BacnetIpv6Mode *BACnetIPModeTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataBACnetIPv6Mode is the corresponding interface of BACnetConstructedDataBACnetIPv6Mode
type IBACnetConstructedDataBACnetIPv6Mode interface {
	IBACnetConstructedData
	// GetBacnetIpv6Mode returns BacnetIpv6Mode (property field)
	GetBacnetIpv6Mode() *BACnetIPModeTagged
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

func (m *BACnetConstructedDataBACnetIPv6Mode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBACnetIPv6Mode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IPV6_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBACnetIPv6Mode) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBACnetIPv6Mode) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBACnetIPv6Mode) GetBacnetIpv6Mode() *BACnetIPModeTagged {
	return m.BacnetIpv6Mode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPv6Mode factory function for BACnetConstructedDataBACnetIPv6Mode
func NewBACnetConstructedDataBACnetIPv6Mode(bacnetIpv6Mode *BACnetIPModeTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataBACnetIPv6Mode {
	_result := &BACnetConstructedDataBACnetIPv6Mode{
		BacnetIpv6Mode:        bacnetIpv6Mode,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBACnetIPv6Mode(structType interface{}) *BACnetConstructedDataBACnetIPv6Mode {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPv6Mode); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPv6Mode); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBACnetIPv6Mode(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBACnetIPv6Mode(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBACnetIPv6Mode) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPv6Mode"
}

func (m *BACnetConstructedDataBACnetIPv6Mode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBACnetIPv6Mode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bacnetIpv6Mode)
	lengthInBits += m.BacnetIpv6Mode.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBACnetIPv6Mode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBACnetIPv6ModeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataBACnetIPv6Mode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPv6Mode"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bacnetIpv6Mode)
	if pullErr := readBuffer.PullContext("bacnetIpv6Mode"); pullErr != nil {
		return nil, pullErr
	}
	_bacnetIpv6Mode, _bacnetIpv6ModeErr := BACnetIPModeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _bacnetIpv6ModeErr != nil {
		return nil, errors.Wrap(_bacnetIpv6ModeErr, "Error parsing 'bacnetIpv6Mode' field")
	}
	bacnetIpv6Mode := CastBACnetIPModeTagged(_bacnetIpv6Mode)
	if closeErr := readBuffer.CloseContext("bacnetIpv6Mode"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPv6Mode"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBACnetIPv6Mode{
		BacnetIpv6Mode:        CastBACnetIPModeTagged(bacnetIpv6Mode),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBACnetIPv6Mode) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPv6Mode"); pushErr != nil {
			return pushErr
		}

		// Simple Field (bacnetIpv6Mode)
		if pushErr := writeBuffer.PushContext("bacnetIpv6Mode"); pushErr != nil {
			return pushErr
		}
		_bacnetIpv6ModeErr := m.BacnetIpv6Mode.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("bacnetIpv6Mode"); popErr != nil {
			return popErr
		}
		if _bacnetIpv6ModeErr != nil {
			return errors.Wrap(_bacnetIpv6ModeErr, "Error serializing 'bacnetIpv6Mode' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPv6Mode"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBACnetIPv6Mode) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
