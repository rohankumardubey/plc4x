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

// BACnetConstructedDataSubordinateNodeTypes is the data-structure of this message
type BACnetConstructedDataSubordinateNodeTypes struct {
	*BACnetConstructedData
	SubordinateNodeTypes []*BACnetNodeTypeTagged

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataSubordinateNodeTypes is the corresponding interface of BACnetConstructedDataSubordinateNodeTypes
type IBACnetConstructedDataSubordinateNodeTypes interface {
	IBACnetConstructedData
	// GetSubordinateNodeTypes returns SubordinateNodeTypes (property field)
	GetSubordinateNodeTypes() []*BACnetNodeTypeTagged
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

func (m *BACnetConstructedDataSubordinateNodeTypes) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBORDINATE_NODE_TYPES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSubordinateNodeTypes) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSubordinateNodeTypes) GetSubordinateNodeTypes() []*BACnetNodeTypeTagged {
	return m.SubordinateNodeTypes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSubordinateNodeTypes factory function for BACnetConstructedDataSubordinateNodeTypes
func NewBACnetConstructedDataSubordinateNodeTypes(subordinateNodeTypes []*BACnetNodeTypeTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataSubordinateNodeTypes {
	_result := &BACnetConstructedDataSubordinateNodeTypes{
		SubordinateNodeTypes:  subordinateNodeTypes,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSubordinateNodeTypes(structType interface{}) *BACnetConstructedDataSubordinateNodeTypes {
	if casted, ok := structType.(BACnetConstructedDataSubordinateNodeTypes); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubordinateNodeTypes); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSubordinateNodeTypes(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSubordinateNodeTypes(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetTypeName() string {
	return "BACnetConstructedDataSubordinateNodeTypes"
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.SubordinateNodeTypes) > 0 {
		for _, element := range m.SubordinateNodeTypes {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataSubordinateNodeTypes) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSubordinateNodeTypesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataSubordinateNodeTypes, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubordinateNodeTypes"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (subordinateNodeTypes)
	if pullErr := readBuffer.PullContext("subordinateNodeTypes", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	subordinateNodeTypes := make([]*BACnetNodeTypeTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetNodeTypeTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'subordinateNodeTypes' field")
			}
			subordinateNodeTypes = append(subordinateNodeTypes, CastBACnetNodeTypeTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("subordinateNodeTypes", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubordinateNodeTypes"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSubordinateNodeTypes{
		SubordinateNodeTypes:  subordinateNodeTypes,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSubordinateNodeTypes) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubordinateNodeTypes"); pushErr != nil {
			return pushErr
		}

		// Array Field (subordinateNodeTypes)
		if m.SubordinateNodeTypes != nil {
			if pushErr := writeBuffer.PushContext("subordinateNodeTypes", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.SubordinateNodeTypes {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'subordinateNodeTypes' field")
				}
			}
			if popErr := writeBuffer.PopContext("subordinateNodeTypes", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubordinateNodeTypes"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSubordinateNodeTypes) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
