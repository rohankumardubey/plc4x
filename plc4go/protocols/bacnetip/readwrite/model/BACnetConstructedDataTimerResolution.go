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

// BACnetConstructedDataTimerResolution is the data-structure of this message
type BACnetConstructedDataTimerResolution struct {
	*BACnetConstructedData
	Resolution *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataTimerResolution is the corresponding interface of BACnetConstructedDataTimerResolution
type IBACnetConstructedDataTimerResolution interface {
	IBACnetConstructedData
	// GetResolution returns Resolution (property field)
	GetResolution() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataTimerResolution) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMER
}

func (m *BACnetConstructedDataTimerResolution) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESOLUTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTimerResolution) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTimerResolution) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTimerResolution) GetResolution() *BACnetApplicationTagUnsignedInteger {
	return m.Resolution
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimerResolution factory function for BACnetConstructedDataTimerResolution
func NewBACnetConstructedDataTimerResolution(resolution *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataTimerResolution {
	_result := &BACnetConstructedDataTimerResolution{
		Resolution:            resolution,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTimerResolution(structType interface{}) *BACnetConstructedDataTimerResolution {
	if casted, ok := structType.(BACnetConstructedDataTimerResolution); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerResolution); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimerResolution(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimerResolution(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTimerResolution) GetTypeName() string {
	return "BACnetConstructedDataTimerResolution"
}

func (m *BACnetConstructedDataTimerResolution) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTimerResolution) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (resolution)
	lengthInBits += m.Resolution.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTimerResolution) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimerResolutionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataTimerResolution, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerResolution"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (resolution)
	if pullErr := readBuffer.PullContext("resolution"); pullErr != nil {
		return nil, pullErr
	}
	_resolution, _resolutionErr := BACnetApplicationTagParse(readBuffer)
	if _resolutionErr != nil {
		return nil, errors.Wrap(_resolutionErr, "Error parsing 'resolution' field")
	}
	resolution := CastBACnetApplicationTagUnsignedInteger(_resolution)
	if closeErr := readBuffer.CloseContext("resolution"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerResolution"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTimerResolution{
		Resolution:            CastBACnetApplicationTagUnsignedInteger(resolution),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTimerResolution) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerResolution"); pushErr != nil {
			return pushErr
		}

		// Simple Field (resolution)
		if pushErr := writeBuffer.PushContext("resolution"); pushErr != nil {
			return pushErr
		}
		_resolutionErr := m.Resolution.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("resolution"); popErr != nil {
			return popErr
		}
		if _resolutionErr != nil {
			return errors.Wrap(_resolutionErr, "Error serializing 'resolution' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerResolution"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTimerResolution) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
