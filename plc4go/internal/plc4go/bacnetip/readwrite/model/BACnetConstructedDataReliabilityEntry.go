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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataReliabilityEntry is the data-structure of this message
type BACnetConstructedDataReliabilityEntry struct {
	RawData *BACnetApplicationTagEnumerated
}

// IBACnetConstructedDataReliabilityEntry is the corresponding interface of BACnetConstructedDataReliabilityEntry
type IBACnetConstructedDataReliabilityEntry interface {
	// GetRawData returns RawData (property field)
	GetRawData() *BACnetApplicationTagEnumerated
	// GetIsBACnetReliabilityProprietary returns IsBACnetReliabilityProprietary (virtual field)
	GetIsBACnetReliabilityProprietary() bool
	// GetReliability returns Reliability (virtual field)
	GetReliability() BACnetReliability
	// GetReliabilityProprietary returns ReliabilityProprietary (virtual field)
	GetReliabilityProprietary() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataReliabilityEntry) GetRawData() *BACnetApplicationTagEnumerated {
	return m.RawData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataReliabilityEntry) GetIsBACnetReliabilityProprietary() bool {
	return bool(bool((m.GetRawData().GetActualValue()) > (255)))
}

func (m *BACnetConstructedDataReliabilityEntry) GetReliability() BACnetReliability {
	return BACnetReliability(MapBACnetReliability(m.GetRawData(), m.GetIsBACnetReliabilityProprietary()))
}

func (m *BACnetConstructedDataReliabilityEntry) GetReliabilityProprietary() uint16 {
	return uint16(utils.InlineIf(m.GetIsBACnetReliabilityProprietary(), func() interface{} { return uint16(m.GetRawData().GetActualValue()) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataReliabilityEntry factory function for BACnetConstructedDataReliabilityEntry
func NewBACnetConstructedDataReliabilityEntry(rawData *BACnetApplicationTagEnumerated) *BACnetConstructedDataReliabilityEntry {
	return &BACnetConstructedDataReliabilityEntry{RawData: rawData}
}

func CastBACnetConstructedDataReliabilityEntry(structType interface{}) *BACnetConstructedDataReliabilityEntry {
	if casted, ok := structType.(BACnetConstructedDataReliabilityEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReliabilityEntry); ok {
		return casted
	}
	return nil
}

func (m *BACnetConstructedDataReliabilityEntry) GetTypeName() string {
	return "BACnetConstructedDataReliabilityEntry"
}

func (m *BACnetConstructedDataReliabilityEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataReliabilityEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (rawData)
	lengthInBits += m.RawData.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataReliabilityEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataReliabilityEntryParse(readBuffer utils.ReadBuffer) (*BACnetConstructedDataReliabilityEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReliabilityEntry"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (rawData)
	if pullErr := readBuffer.PullContext("rawData"); pullErr != nil {
		return nil, pullErr
	}
	_rawData, _rawDataErr := BACnetApplicationTagParse(readBuffer)
	if _rawDataErr != nil {
		return nil, errors.Wrap(_rawDataErr, "Error parsing 'rawData' field")
	}
	rawData := CastBACnetApplicationTagEnumerated(_rawData)
	if closeErr := readBuffer.CloseContext("rawData"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_isBACnetReliabilityProprietary := bool((rawData.GetActualValue()) > (255))
	isBACnetReliabilityProprietary := bool(_isBACnetReliabilityProprietary)
	_ = isBACnetReliabilityProprietary

	// Virtual field
	_reliability := MapBACnetReliability(rawData, isBACnetReliabilityProprietary)
	reliability := BACnetReliability(_reliability)
	_ = reliability

	// Virtual field
	_reliabilityProprietary := utils.InlineIf(isBACnetReliabilityProprietary, func() interface{} { return uint16(rawData.GetActualValue()) }, func() interface{} { return uint16(uint16(0)) }).(uint16)
	reliabilityProprietary := uint16(_reliabilityProprietary)
	_ = reliabilityProprietary

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReliabilityEntry"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetConstructedDataReliabilityEntry(rawData), nil
}

func (m *BACnetConstructedDataReliabilityEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetConstructedDataReliabilityEntry"); pushErr != nil {
		return pushErr
	}

	// Simple Field (rawData)
	if pushErr := writeBuffer.PushContext("rawData"); pushErr != nil {
		return pushErr
	}
	_rawDataErr := m.RawData.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("rawData"); popErr != nil {
		return popErr
	}
	if _rawDataErr != nil {
		return errors.Wrap(_rawDataErr, "Error serializing 'rawData' field")
	}
	// Virtual field
	if _isBACnetReliabilityProprietaryErr := writeBuffer.WriteVirtual("isBACnetReliabilityProprietary", m.GetIsBACnetReliabilityProprietary()); _isBACnetReliabilityProprietaryErr != nil {
		return errors.Wrap(_isBACnetReliabilityProprietaryErr, "Error serializing 'isBACnetReliabilityProprietary' field")
	}
	// Virtual field
	if _reliabilityErr := writeBuffer.WriteVirtual("reliability", m.GetReliability()); _reliabilityErr != nil {
		return errors.Wrap(_reliabilityErr, "Error serializing 'reliability' field")
	}
	// Virtual field
	if _reliabilityProprietaryErr := writeBuffer.WriteVirtual("reliabilityProprietary", m.GetReliabilityProprietary()); _reliabilityProprietaryErr != nil {
		return errors.Wrap(_reliabilityProprietaryErr, "Error serializing 'reliabilityProprietary' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConstructedDataReliabilityEntry"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConstructedDataReliabilityEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
