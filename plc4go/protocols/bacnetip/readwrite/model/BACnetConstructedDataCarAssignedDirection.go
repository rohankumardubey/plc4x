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

// BACnetConstructedDataCarAssignedDirection is the data-structure of this message
type BACnetConstructedDataCarAssignedDirection struct {
	*BACnetConstructedData
	AssignedDirection *BACnetLiftCarDirectionTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataCarAssignedDirection is the corresponding interface of BACnetConstructedDataCarAssignedDirection
type IBACnetConstructedDataCarAssignedDirection interface {
	IBACnetConstructedData
	// GetAssignedDirection returns AssignedDirection (property field)
	GetAssignedDirection() *BACnetLiftCarDirectionTagged
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

func (m *BACnetConstructedDataCarAssignedDirection) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataCarAssignedDirection) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_ASSIGNED_DIRECTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataCarAssignedDirection) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataCarAssignedDirection) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataCarAssignedDirection) GetAssignedDirection() *BACnetLiftCarDirectionTagged {
	return m.AssignedDirection
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarAssignedDirection factory function for BACnetConstructedDataCarAssignedDirection
func NewBACnetConstructedDataCarAssignedDirection(assignedDirection *BACnetLiftCarDirectionTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataCarAssignedDirection {
	_result := &BACnetConstructedDataCarAssignedDirection{
		AssignedDirection:     assignedDirection,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataCarAssignedDirection(structType interface{}) *BACnetConstructedDataCarAssignedDirection {
	if casted, ok := structType.(BACnetConstructedDataCarAssignedDirection); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarAssignedDirection); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataCarAssignedDirection(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataCarAssignedDirection(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataCarAssignedDirection) GetTypeName() string {
	return "BACnetConstructedDataCarAssignedDirection"
}

func (m *BACnetConstructedDataCarAssignedDirection) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataCarAssignedDirection) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (assignedDirection)
	lengthInBits += m.AssignedDirection.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataCarAssignedDirection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCarAssignedDirectionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataCarAssignedDirection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarAssignedDirection"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (assignedDirection)
	if pullErr := readBuffer.PullContext("assignedDirection"); pullErr != nil {
		return nil, pullErr
	}
	_assignedDirection, _assignedDirectionErr := BACnetLiftCarDirectionTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _assignedDirectionErr != nil {
		return nil, errors.Wrap(_assignedDirectionErr, "Error parsing 'assignedDirection' field")
	}
	assignedDirection := CastBACnetLiftCarDirectionTagged(_assignedDirection)
	if closeErr := readBuffer.CloseContext("assignedDirection"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarAssignedDirection"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataCarAssignedDirection{
		AssignedDirection:     CastBACnetLiftCarDirectionTagged(assignedDirection),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataCarAssignedDirection) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarAssignedDirection"); pushErr != nil {
			return pushErr
		}

		// Simple Field (assignedDirection)
		if pushErr := writeBuffer.PushContext("assignedDirection"); pushErr != nil {
			return pushErr
		}
		_assignedDirectionErr := m.AssignedDirection.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("assignedDirection"); popErr != nil {
			return popErr
		}
		if _assignedDirectionErr != nil {
			return errors.Wrap(_assignedDirectionErr, "Error serializing 'assignedDirection' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarAssignedDirection"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataCarAssignedDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
