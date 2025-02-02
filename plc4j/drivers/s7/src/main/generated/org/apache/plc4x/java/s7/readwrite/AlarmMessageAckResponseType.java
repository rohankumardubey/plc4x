/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.s7.readwrite;

import static org.apache.plc4x.java.spi.codegen.fields.FieldReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.fields.FieldWriterFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataReaderFactory.*;
import static org.apache.plc4x.java.spi.codegen.io.DataWriterFactory.*;
import static org.apache.plc4x.java.spi.generation.StaticHelper.*;

import java.time.*;
import java.util.*;
import org.apache.plc4x.java.api.exceptions.*;
import org.apache.plc4x.java.api.value.*;
import org.apache.plc4x.java.spi.codegen.*;
import org.apache.plc4x.java.spi.codegen.fields.*;
import org.apache.plc4x.java.spi.codegen.io.*;
import org.apache.plc4x.java.spi.generation.*;

// Code generated by code-generation. DO NOT EDIT.

public class AlarmMessageAckResponseType implements Message {

  // Properties.
  protected final short functionId;
  protected final short numberOfObjects;
  protected final List<Short> messageObjects;

  public AlarmMessageAckResponseType(
      short functionId, short numberOfObjects, List<Short> messageObjects) {
    super();
    this.functionId = functionId;
    this.numberOfObjects = numberOfObjects;
    this.messageObjects = messageObjects;
  }

  public short getFunctionId() {
    return functionId;
  }

  public short getNumberOfObjects() {
    return numberOfObjects;
  }

  public List<Short> getMessageObjects() {
    return messageObjects;
  }

  public void serialize(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AlarmMessageAckResponseType");

    // Simple Field (functionId)
    writeSimpleField("functionId", functionId, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (numberOfObjects)
    writeSimpleField("numberOfObjects", numberOfObjects, writeUnsignedShort(writeBuffer, 8));

    // Array Field (messageObjects)
    writeSimpleTypeArrayField("messageObjects", messageObjects, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("AlarmMessageAckResponseType");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = 0;
    AlarmMessageAckResponseType _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (functionId)
    lengthInBits += 8;

    // Simple field (numberOfObjects)
    lengthInBits += 8;

    // Array field
    if (messageObjects != null) {
      lengthInBits += 8 * messageObjects.size();
    }

    return lengthInBits;
  }

  public static AlarmMessageAckResponseType staticParse(ReadBuffer readBuffer, Object... args)
      throws ParseException {
    PositionAware positionAware = readBuffer;
    return staticParse(readBuffer);
  }

  public static AlarmMessageAckResponseType staticParse(ReadBuffer readBuffer)
      throws ParseException {
    readBuffer.pullContext("AlarmMessageAckResponseType");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short functionId = readSimpleField("functionId", readUnsignedShort(readBuffer, 8));

    short numberOfObjects = readSimpleField("numberOfObjects", readUnsignedShort(readBuffer, 8));

    List<Short> messageObjects =
        readCountArrayField("messageObjects", readUnsignedShort(readBuffer, 8), numberOfObjects);

    readBuffer.closeContext("AlarmMessageAckResponseType");
    // Create the instance
    AlarmMessageAckResponseType _alarmMessageAckResponseType;
    _alarmMessageAckResponseType =
        new AlarmMessageAckResponseType(functionId, numberOfObjects, messageObjects);
    return _alarmMessageAckResponseType;
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AlarmMessageAckResponseType)) {
      return false;
    }
    AlarmMessageAckResponseType that = (AlarmMessageAckResponseType) o;
    return (getFunctionId() == that.getFunctionId())
        && (getNumberOfObjects() == that.getNumberOfObjects())
        && (getMessageObjects() == that.getMessageObjects())
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(getFunctionId(), getNumberOfObjects(), getMessageObjects());
  }

  @Override
  public String toString() {
    WriteBufferBoxBased writeBufferBoxBased = new WriteBufferBoxBased(true, true);
    try {
      writeBufferBoxBased.writeSerializable(this);
    } catch (SerializationException e) {
      throw new RuntimeException(e);
    }
    return "\n" + writeBufferBoxBased.getBox().toString() + "\n";
  }
}
