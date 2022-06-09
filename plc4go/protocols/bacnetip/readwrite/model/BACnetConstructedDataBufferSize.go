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

// BACnetConstructedDataBufferSize is the data-structure of this message
type BACnetConstructedDataBufferSize struct {
	*BACnetConstructedData
	BufferSize *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataBufferSize is the corresponding interface of BACnetConstructedDataBufferSize
type IBACnetConstructedDataBufferSize interface {
	IBACnetConstructedData
	// GetBufferSize returns BufferSize (property field)
	GetBufferSize() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataBufferSize) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBufferSize) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BUFFER_SIZE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBufferSize) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBufferSize) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBufferSize) GetBufferSize() *BACnetApplicationTagUnsignedInteger {
	return m.BufferSize
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBufferSize factory function for BACnetConstructedDataBufferSize
func NewBACnetConstructedDataBufferSize(bufferSize *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataBufferSize {
	_result := &BACnetConstructedDataBufferSize{
		BufferSize:            bufferSize,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBufferSize(structType interface{}) *BACnetConstructedDataBufferSize {
	if casted, ok := structType.(BACnetConstructedDataBufferSize); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBufferSize); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBufferSize(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBufferSize(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBufferSize) GetTypeName() string {
	return "BACnetConstructedDataBufferSize"
}

func (m *BACnetConstructedDataBufferSize) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBufferSize) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (bufferSize)
	lengthInBits += m.BufferSize.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBufferSize) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBufferSizeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataBufferSize, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBufferSize"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bufferSize)
	if pullErr := readBuffer.PullContext("bufferSize"); pullErr != nil {
		return nil, pullErr
	}
	_bufferSize, _bufferSizeErr := BACnetApplicationTagParse(readBuffer)
	if _bufferSizeErr != nil {
		return nil, errors.Wrap(_bufferSizeErr, "Error parsing 'bufferSize' field")
	}
	bufferSize := CastBACnetApplicationTagUnsignedInteger(_bufferSize)
	if closeErr := readBuffer.CloseContext("bufferSize"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBufferSize"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBufferSize{
		BufferSize:            CastBACnetApplicationTagUnsignedInteger(bufferSize),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBufferSize) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBufferSize"); pushErr != nil {
			return pushErr
		}

		// Simple Field (bufferSize)
		if pushErr := writeBuffer.PushContext("bufferSize"); pushErr != nil {
			return pushErr
		}
		_bufferSizeErr := m.BufferSize.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("bufferSize"); popErr != nil {
			return popErr
		}
		if _bufferSizeErr != nil {
			return errors.Wrap(_bufferSizeErr, "Error serializing 'bufferSize' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBufferSize"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBufferSize) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
