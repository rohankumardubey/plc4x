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

// BACnetConstructedDataOccupancyCountAdjust is the data-structure of this message
type BACnetConstructedDataOccupancyCountAdjust struct {
	*BACnetConstructedData
	OccupancyCountAdjust *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataOccupancyCountAdjust is the corresponding interface of BACnetConstructedDataOccupancyCountAdjust
type IBACnetConstructedDataOccupancyCountAdjust interface {
	IBACnetConstructedData
	// GetOccupancyCountAdjust returns OccupancyCountAdjust (property field)
	GetOccupancyCountAdjust() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataOccupancyCountAdjust) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataOccupancyCountAdjust) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_COUNT_ADJUST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataOccupancyCountAdjust) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataOccupancyCountAdjust) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataOccupancyCountAdjust) GetOccupancyCountAdjust() *BACnetApplicationTagBoolean {
	return m.OccupancyCountAdjust
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOccupancyCountAdjust factory function for BACnetConstructedDataOccupancyCountAdjust
func NewBACnetConstructedDataOccupancyCountAdjust(occupancyCountAdjust *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataOccupancyCountAdjust {
	_result := &BACnetConstructedDataOccupancyCountAdjust{
		OccupancyCountAdjust:  occupancyCountAdjust,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataOccupancyCountAdjust(structType interface{}) *BACnetConstructedDataOccupancyCountAdjust {
	if casted, ok := structType.(BACnetConstructedDataOccupancyCountAdjust); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyCountAdjust); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataOccupancyCountAdjust(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataOccupancyCountAdjust(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataOccupancyCountAdjust) GetTypeName() string {
	return "BACnetConstructedDataOccupancyCountAdjust"
}

func (m *BACnetConstructedDataOccupancyCountAdjust) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataOccupancyCountAdjust) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (occupancyCountAdjust)
	lengthInBits += m.OccupancyCountAdjust.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataOccupancyCountAdjust) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOccupancyCountAdjustParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataOccupancyCountAdjust, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyCountAdjust"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (occupancyCountAdjust)
	if pullErr := readBuffer.PullContext("occupancyCountAdjust"); pullErr != nil {
		return nil, pullErr
	}
	_occupancyCountAdjust, _occupancyCountAdjustErr := BACnetApplicationTagParse(readBuffer)
	if _occupancyCountAdjustErr != nil {
		return nil, errors.Wrap(_occupancyCountAdjustErr, "Error parsing 'occupancyCountAdjust' field")
	}
	occupancyCountAdjust := CastBACnetApplicationTagBoolean(_occupancyCountAdjust)
	if closeErr := readBuffer.CloseContext("occupancyCountAdjust"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyCountAdjust"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataOccupancyCountAdjust{
		OccupancyCountAdjust:  CastBACnetApplicationTagBoolean(occupancyCountAdjust),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataOccupancyCountAdjust) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyCountAdjust"); pushErr != nil {
			return pushErr
		}

		// Simple Field (occupancyCountAdjust)
		if pushErr := writeBuffer.PushContext("occupancyCountAdjust"); pushErr != nil {
			return pushErr
		}
		_occupancyCountAdjustErr := m.OccupancyCountAdjust.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("occupancyCountAdjust"); popErr != nil {
			return popErr
		}
		if _occupancyCountAdjustErr != nil {
			return errors.Wrap(_occupancyCountAdjustErr, "Error serializing 'occupancyCountAdjust' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyCountAdjust"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataOccupancyCountAdjust) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
