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

// The data-structure of this message
type IdentifyReplyCommandManufacturer struct {
	*IdentifyReplyCommand
	ManufacturerName string
}

// The corresponding interface
type IIdentifyReplyCommandManufacturer interface {
	// GetManufacturerName returns ManufacturerName
	GetManufacturerName() string
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *IdentifyReplyCommandManufacturer) Attribute() Attribute {
	return Attribute_Manufacturer
}

func (m *IdentifyReplyCommandManufacturer) GetAttribute() Attribute {
	return Attribute_Manufacturer
}

func (m *IdentifyReplyCommandManufacturer) InitializeParent(parent *IdentifyReplyCommand) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *IdentifyReplyCommandManufacturer) GetManufacturerName() string {
	return m.ManufacturerName
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewIdentifyReplyCommandManufacturer factory function for IdentifyReplyCommandManufacturer
func NewIdentifyReplyCommandManufacturer(manufacturerName string) *IdentifyReplyCommand {
	child := &IdentifyReplyCommandManufacturer{
		ManufacturerName:     manufacturerName,
		IdentifyReplyCommand: NewIdentifyReplyCommand(),
	}
	child.Child = child
	return child.IdentifyReplyCommand
}

func CastIdentifyReplyCommandManufacturer(structType interface{}) *IdentifyReplyCommandManufacturer {
	castFunc := func(typ interface{}) *IdentifyReplyCommandManufacturer {
		if casted, ok := typ.(IdentifyReplyCommandManufacturer); ok {
			return &casted
		}
		if casted, ok := typ.(*IdentifyReplyCommandManufacturer); ok {
			return casted
		}
		if casted, ok := typ.(IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandManufacturer(casted.Child)
		}
		if casted, ok := typ.(*IdentifyReplyCommand); ok {
			return CastIdentifyReplyCommandManufacturer(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *IdentifyReplyCommandManufacturer) GetTypeName() string {
	return "IdentifyReplyCommandManufacturer"
}

func (m *IdentifyReplyCommandManufacturer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *IdentifyReplyCommandManufacturer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (manufacturerName)
	lengthInBits += 64

	return lengthInBits
}

func (m *IdentifyReplyCommandManufacturer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandManufacturerParse(readBuffer utils.ReadBuffer, attribute Attribute) (*IdentifyReplyCommand, error) {
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandManufacturer"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (manufacturerName)
	_manufacturerName, _manufacturerNameErr := readBuffer.ReadString("manufacturerName", uint32(64))
	if _manufacturerNameErr != nil {
		return nil, errors.Wrap(_manufacturerNameErr, "Error parsing 'manufacturerName' field")
	}
	manufacturerName := _manufacturerName

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandManufacturer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &IdentifyReplyCommandManufacturer{
		ManufacturerName:     manufacturerName,
		IdentifyReplyCommand: &IdentifyReplyCommand{},
	}
	_child.IdentifyReplyCommand.Child = _child
	return _child.IdentifyReplyCommand, nil
}

func (m *IdentifyReplyCommandManufacturer) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandManufacturer"); pushErr != nil {
			return pushErr
		}

		// Simple Field (manufacturerName)
		manufacturerName := string(m.ManufacturerName)
		_manufacturerNameErr := writeBuffer.WriteString("manufacturerName", uint32(64), "UTF-8", (manufacturerName))
		if _manufacturerNameErr != nil {
			return errors.Wrap(_manufacturerNameErr, "Error serializing 'manufacturerName' field")
		}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandManufacturer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *IdentifyReplyCommandManufacturer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
