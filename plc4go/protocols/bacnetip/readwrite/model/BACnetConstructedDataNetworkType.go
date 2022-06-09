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

// BACnetConstructedDataNetworkType is the data-structure of this message
type BACnetConstructedDataNetworkType struct {
	*BACnetConstructedData
	NetworkType *BACnetNetworkTypeTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataNetworkType is the corresponding interface of BACnetConstructedDataNetworkType
type IBACnetConstructedDataNetworkType interface {
	IBACnetConstructedData
	// GetNetworkType returns NetworkType (property field)
	GetNetworkType() *BACnetNetworkTypeTagged
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

func (m *BACnetConstructedDataNetworkType) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataNetworkType) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NETWORK_TYPE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataNetworkType) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataNetworkType) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataNetworkType) GetNetworkType() *BACnetNetworkTypeTagged {
	return m.NetworkType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNetworkType factory function for BACnetConstructedDataNetworkType
func NewBACnetConstructedDataNetworkType(networkType *BACnetNetworkTypeTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataNetworkType {
	_result := &BACnetConstructedDataNetworkType{
		NetworkType:           networkType,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataNetworkType(structType interface{}) *BACnetConstructedDataNetworkType {
	if casted, ok := structType.(BACnetConstructedDataNetworkType); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkType); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkType(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataNetworkType(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataNetworkType) GetTypeName() string {
	return "BACnetConstructedDataNetworkType"
}

func (m *BACnetConstructedDataNetworkType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataNetworkType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (networkType)
	lengthInBits += m.NetworkType.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataNetworkType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNetworkTypeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataNetworkType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkType"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkType)
	if pullErr := readBuffer.PullContext("networkType"); pullErr != nil {
		return nil, pullErr
	}
	_networkType, _networkTypeErr := BACnetNetworkTypeTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _networkTypeErr != nil {
		return nil, errors.Wrap(_networkTypeErr, "Error parsing 'networkType' field")
	}
	networkType := CastBACnetNetworkTypeTagged(_networkType)
	if closeErr := readBuffer.CloseContext("networkType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkType"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataNetworkType{
		NetworkType:           CastBACnetNetworkTypeTagged(networkType),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataNetworkType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkType"); pushErr != nil {
			return pushErr
		}

		// Simple Field (networkType)
		if pushErr := writeBuffer.PushContext("networkType"); pushErr != nil {
			return pushErr
		}
		_networkTypeErr := m.NetworkType.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("networkType"); popErr != nil {
			return popErr
		}
		if _networkTypeErr != nil {
			return errors.Wrap(_networkTypeErr, "Error serializing 'networkType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkType"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataNetworkType) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
