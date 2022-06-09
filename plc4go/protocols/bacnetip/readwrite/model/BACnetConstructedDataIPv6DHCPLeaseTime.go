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

// BACnetConstructedDataIPv6DHCPLeaseTime is the data-structure of this message
type BACnetConstructedDataIPv6DHCPLeaseTime struct {
	*BACnetConstructedData
	Ipv6DhcpLeaseTime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataIPv6DHCPLeaseTime is the corresponding interface of BACnetConstructedDataIPv6DHCPLeaseTime
type IBACnetConstructedDataIPv6DHCPLeaseTime interface {
	IBACnetConstructedData
	// GetIpv6DhcpLeaseTime returns Ipv6DhcpLeaseTime (property field)
	GetIpv6DhcpLeaseTime() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DHCP_LEASE_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) GetIpv6DhcpLeaseTime() *BACnetApplicationTagUnsignedInteger {
	return m.Ipv6DhcpLeaseTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6DHCPLeaseTime factory function for BACnetConstructedDataIPv6DHCPLeaseTime
func NewBACnetConstructedDataIPv6DHCPLeaseTime(ipv6DhcpLeaseTime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataIPv6DHCPLeaseTime {
	_result := &BACnetConstructedDataIPv6DHCPLeaseTime{
		Ipv6DhcpLeaseTime:     ipv6DhcpLeaseTime,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIPv6DHCPLeaseTime(structType interface{}) *BACnetConstructedDataIPv6DHCPLeaseTime {
	if casted, ok := structType.(BACnetConstructedDataIPv6DHCPLeaseTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DHCPLeaseTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPv6DHCPLeaseTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPv6DHCPLeaseTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) GetTypeName() string {
	return "BACnetConstructedDataIPv6DHCPLeaseTime"
}

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipv6DhcpLeaseTime)
	lengthInBits += m.Ipv6DhcpLeaseTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPv6DHCPLeaseTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataIPv6DHCPLeaseTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DHCPLeaseTime"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6DhcpLeaseTime)
	if pullErr := readBuffer.PullContext("ipv6DhcpLeaseTime"); pullErr != nil {
		return nil, pullErr
	}
	_ipv6DhcpLeaseTime, _ipv6DhcpLeaseTimeErr := BACnetApplicationTagParse(readBuffer)
	if _ipv6DhcpLeaseTimeErr != nil {
		return nil, errors.Wrap(_ipv6DhcpLeaseTimeErr, "Error parsing 'ipv6DhcpLeaseTime' field")
	}
	ipv6DhcpLeaseTime := CastBACnetApplicationTagUnsignedInteger(_ipv6DhcpLeaseTime)
	if closeErr := readBuffer.CloseContext("ipv6DhcpLeaseTime"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DHCPLeaseTime"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIPv6DHCPLeaseTime{
		Ipv6DhcpLeaseTime:     CastBACnetApplicationTagUnsignedInteger(ipv6DhcpLeaseTime),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DHCPLeaseTime"); pushErr != nil {
			return pushErr
		}

		// Simple Field (ipv6DhcpLeaseTime)
		if pushErr := writeBuffer.PushContext("ipv6DhcpLeaseTime"); pushErr != nil {
			return pushErr
		}
		_ipv6DhcpLeaseTimeErr := m.Ipv6DhcpLeaseTime.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("ipv6DhcpLeaseTime"); popErr != nil {
			return popErr
		}
		if _ipv6DhcpLeaseTimeErr != nil {
			return errors.Wrap(_ipv6DhcpLeaseTimeErr, "Error serializing 'ipv6DhcpLeaseTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DHCPLeaseTime"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIPv6DHCPLeaseTime) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
