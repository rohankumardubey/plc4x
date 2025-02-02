#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

from dataclasses import dataclass

from ctypes import c_bool
from ctypes import c_uint8
from plc4py.api.messages.PlcMessage import PlcMessage
from plc4py.protocols.modbus.readwrite.ModbusDeviceInformationConformityLevel import (
    ModbusDeviceInformationConformityLevel,
)
from plc4py.protocols.modbus.readwrite.ModbusDeviceInformationLevel import (
    ModbusDeviceInformationLevel,
)
from plc4py.protocols.modbus.readwrite.ModbusDeviceInformationMoreFollows import (
    ModbusDeviceInformationMoreFollows,
)
from plc4py.protocols.modbus.readwrite.ModbusDeviceInformationObject import (
    ModbusDeviceInformationObject,
)
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDU
from plc4py.protocols.modbus.readwrite.ModbusPDU import ModbusPDUBuilder
from typing import List
import math


@dataclass
class ModbusPDUReadDeviceIdentificationResponse(PlcMessage, ModbusPDU):
    level: ModbusDeviceInformationLevel
    individual_access: c_bool
    conformity_level: ModbusDeviceInformationConformityLevel
    more_follows: ModbusDeviceInformationMoreFollows
    next_object_id: c_uint8
    objects: List[ModbusDeviceInformationObject]
    MEITYPE: c_uint8 = 0x0E
    # Accessors for discriminator values.
    error_flag: c_bool = False
    function_flag: c_uint8 = 0x2B
    response: c_bool = True

    def __post_init__(self):
        super().__init__()

    def serialize_modbus_pdu_child(self, write_buffer: WriteBuffer):
        position_aware: PositionAware = write_buffer
        start_pos: int = position_aware.get_pos()
        write_buffer.push_context("ModbusPDUReadDeviceIdentificationResponse")

        # Const Field (meiType)
        write_const_field(
            "meiType", self.mei_type, write_unsigned_short(write_buffer, 8)
        )

        # Simple Field (level)
        write_simple_enum_field(
            "level",
            "ModbusDeviceInformationLevel",
            self.level,
            DataWriterEnumDefault(
                ModbusDeviceInformationLevel.value,
                ModbusDeviceInformationLevel.name,
                write_unsigned_short(write_buffer, 8),
            ),
        )

        # Simple Field (individualAccess)
        write_simple_field(
            "individualAccess", self.individual_access, write_boolean(write_buffer)
        )

        # Simple Field (conformityLevel)
        write_simple_enum_field(
            "conformityLevel",
            "ModbusDeviceInformationConformityLevel",
            self.conformity_level,
            DataWriterEnumDefault(
                ModbusDeviceInformationConformityLevel.value,
                ModbusDeviceInformationConformityLevel.name,
                write_unsigned_short(write_buffer, 7),
            ),
        )

        # Simple Field (moreFollows)
        write_simple_enum_field(
            "moreFollows",
            "ModbusDeviceInformationMoreFollows",
            self.more_follows,
            DataWriterEnumDefault(
                ModbusDeviceInformationMoreFollows.value,
                ModbusDeviceInformationMoreFollows.name,
                write_unsigned_short(write_buffer, 8),
            ),
        )

        # Simple Field (nextObjectId)
        write_simple_field(
            "nextObjectId", self.next_object_id, write_unsigned_short(write_buffer, 8)
        )

        # Implicit Field (number_of_objects) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
        number_of_objects: c_uint8 = c_uint8((COUNT(self.objects())))
        write_implicit_field(
            "numberOfObjects", number_of_objects, write_unsigned_short(write_buffer, 8)
        )

        # Array Field (objects)
        write_complex_type_array_field("objects", self.objects, write_buffer)

        write_buffer.pop_context("ModbusPDUReadDeviceIdentificationResponse")

    def length_in_bytes(self) -> int:
        return int(math.ceil(float(self.get_length_in_bits() / 8.0)))

    def get_length_in_bits(self) -> int:
        length_in_bits: int = super().get_length_in_bits()
        _value: ModbusPDUReadDeviceIdentificationResponse = self

        # Const Field (meiType)
        length_in_bits += 8

        # Simple field (level)
        length_in_bits += 8

        # Simple field (individualAccess)
        length_in_bits += 1

        # Simple field (conformityLevel)
        length_in_bits += 7

        # Simple field (moreFollows)
        length_in_bits += 8

        # Simple field (nextObjectId)
        length_in_bits += 8

        # Implicit Field (numberOfObjects)
        length_in_bits += 8

        # Array field
        if self.objects is not None:
            i: int = 0
            for element in self.objects:
                last: bool = ++i >= self.objects.size()
                length_in_bits += element.get_length_in_bits()

        return length_in_bits

    @staticmethod
    def static_parse_builder(read_buffer: ReadBuffer, response: c_bool):
        read_buffer.pull_context("ModbusPDUReadDeviceIdentificationResponse")
        position_aware: PositionAware = read_buffer
        start_pos: int = position_aware.get_pos()
        cur_pos: int = 0

        mei_type: c_uint8 = read_const_field(
            "meiType",
            read_unsigned_short(read_buffer, 8),
            ModbusPDUReadDeviceIdentificationResponse.MEITYPE,
        )

        level: ModbusDeviceInformationLevel = read_enum_field(
            "level",
            "ModbusDeviceInformationLevel",
            DataReaderEnumDefault(
                ModbusDeviceInformationLevel.enumForValue,
                read_unsigned_short(read_buffer, 8),
            ),
        )

        individual_access: c_bool = read_simple_field(
            "individualAccess", read_boolean(read_buffer)
        )

        conformity_level: ModbusDeviceInformationConformityLevel = read_enum_field(
            "conformityLevel",
            "ModbusDeviceInformationConformityLevel",
            DataReaderEnumDefault(
                ModbusDeviceInformationConformityLevel.enumForValue,
                read_unsigned_short(read_buffer, 7),
            ),
        )

        more_follows: ModbusDeviceInformationMoreFollows = read_enum_field(
            "moreFollows",
            "ModbusDeviceInformationMoreFollows",
            DataReaderEnumDefault(
                ModbusDeviceInformationMoreFollows.enumForValue,
                read_unsigned_short(read_buffer, 8),
            ),
        )

        next_object_id: c_uint8 = read_simple_field(
            "nextObjectId", read_unsigned_short(read_buffer, 8)
        )

        number_of_objects: c_uint8 = read_implicit_field(
            "numberOfObjects", read_unsigned_short(read_buffer, 8)
        )

        objects: List[ModbusDeviceInformationObject] = read_count_array_field(
            "objects",
            DataReaderComplexDefault(
                ModbusDeviceInformationObject.static_parse(read_buffer), read_buffer
            ),
            number_of_objects,
        )

        read_buffer.close_context("ModbusPDUReadDeviceIdentificationResponse")
        # Create the instance
        return ModbusPDUReadDeviceIdentificationResponseBuilder(
            level,
            individual_access,
            conformity_level,
            more_follows,
            next_object_id,
            objects,
        )

    def equals(self, o: object) -> bool:
        if self == o:
            return True

        if not isinstance(o, ModbusPDUReadDeviceIdentificationResponse):
            return False

        that: ModbusPDUReadDeviceIdentificationResponse = (
            ModbusPDUReadDeviceIdentificationResponse(o)
        )
        return (
            (self.level == that.level)
            and (self.individual_access == that.individual_access)
            and (self.conformity_level == that.conformity_level)
            and (self.more_follows == that.more_follows)
            and (self.next_object_id == that.next_object_id)
            and (self.objects == that.objects)
            and super().equals(that)
            and True
        )

    def hash_code(self) -> int:
        return hash(self)

    def __str__(self) -> str:
        write_buffer_box_based: WriteBufferBoxBased = WriteBufferBoxBased(True, True)
        try:
            write_buffer_box_based.writeSerializable(self)
        except SerializationException as e:
            raise RuntimeException(e)

        return "\n" + str(write_buffer_box_based.get_box()) + "\n"


@dataclass
class ModbusPDUReadDeviceIdentificationResponseBuilder(ModbusPDUBuilder):
    level: ModbusDeviceInformationLevel
    individualAccess: c_bool
    conformityLevel: ModbusDeviceInformationConformityLevel
    moreFollows: ModbusDeviceInformationMoreFollows
    nextObjectId: c_uint8
    objects: List[ModbusDeviceInformationObject]

    def __post_init__(self):
        pass

    def build(
        self,
    ) -> ModbusPDUReadDeviceIdentificationResponse:
        modbus_pdu_read_device_identification_response: ModbusPDUReadDeviceIdentificationResponse = ModbusPDUReadDeviceIdentificationResponse(
            self.level,
            self.individual_access,
            self.conformity_level,
            self.more_follows,
            self.next_object_id,
            self.objects,
        )
        return modbus_pdu_read_device_identification_response
