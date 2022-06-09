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

// BACnetConstructedDataOutOfService is the data-structure of this message
type BACnetConstructedDataOutOfService struct {
	*BACnetConstructedData
	OutOfService *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataOutOfService is the corresponding interface of BACnetConstructedDataOutOfService
type IBACnetConstructedDataOutOfService interface {
	IBACnetConstructedData
	// GetOutOfService returns OutOfService (property field)
	GetOutOfService() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataOutOfService) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataOutOfService) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OUT_OF_SERVICE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataOutOfService) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataOutOfService) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataOutOfService) GetOutOfService() *BACnetApplicationTagBoolean {
	return m.OutOfService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOutOfService factory function for BACnetConstructedDataOutOfService
func NewBACnetConstructedDataOutOfService(outOfService *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataOutOfService {
	_result := &BACnetConstructedDataOutOfService{
		OutOfService:          outOfService,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataOutOfService(structType interface{}) *BACnetConstructedDataOutOfService {
	if casted, ok := structType.(BACnetConstructedDataOutOfService); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOutOfService); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataOutOfService(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataOutOfService(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataOutOfService) GetTypeName() string {
	return "BACnetConstructedDataOutOfService"
}

func (m *BACnetConstructedDataOutOfService) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataOutOfService) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (outOfService)
	lengthInBits += m.OutOfService.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataOutOfService) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOutOfServiceParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataOutOfService, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOutOfService"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (outOfService)
	if pullErr := readBuffer.PullContext("outOfService"); pullErr != nil {
		return nil, pullErr
	}
	_outOfService, _outOfServiceErr := BACnetApplicationTagParse(readBuffer)
	if _outOfServiceErr != nil {
		return nil, errors.Wrap(_outOfServiceErr, "Error parsing 'outOfService' field")
	}
	outOfService := CastBACnetApplicationTagBoolean(_outOfService)
	if closeErr := readBuffer.CloseContext("outOfService"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOutOfService"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataOutOfService{
		OutOfService:          CastBACnetApplicationTagBoolean(outOfService),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataOutOfService) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOutOfService"); pushErr != nil {
			return pushErr
		}

		// Simple Field (outOfService)
		if pushErr := writeBuffer.PushContext("outOfService"); pushErr != nil {
			return pushErr
		}
		_outOfServiceErr := m.OutOfService.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("outOfService"); popErr != nil {
			return popErr
		}
		if _outOfServiceErr != nil {
			return errors.Wrap(_outOfServiceErr, "Error serializing 'outOfService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOutOfService"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataOutOfService) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
