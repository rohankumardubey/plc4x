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

public class S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse
    extends S7PayloadUserDataItem implements Message {

  // Accessors for discriminator values.
  public Byte getCpuFunctionType() {
    return (byte) 0x08;
  }

  public Short getCpuSubfunction() {
    return (short) 0x02;
  }

  public Integer getDataLength() {
    return (int) 0x02;
  }

  // Properties.
  protected final short result;
  protected final short reserved01;

  public S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(
      DataTransportErrorCode returnCode,
      DataTransportSize transportSize,
      short result,
      short reserved01) {
    super(returnCode, transportSize);
    this.result = result;
    this.reserved01 = reserved01;
  }

  public short getResult() {
    return result;
  }

  public short getReserved01() {
    return reserved01;
  }

  @Override
  protected void serializeS7PayloadUserDataItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse");

    // Simple Field (result)
    writeSimpleField("result", result, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (reserved01)
    writeSimpleField("reserved01", reserved01, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (result)
    lengthInBits += 8;

    // Simple field (reserved01)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static S7PayloadUserDataItemBuilder staticParseS7PayloadUserDataItemBuilder(
      ReadBuffer readBuffer, Byte cpuFunctionType, Short cpuSubfunction) throws ParseException {
    readBuffer.pullContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short result = readSimpleField("result", readUnsignedShort(readBuffer, 8));

    short reserved01 = readSimpleField("reserved01", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse");
    // Create the instance
    return new S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseBuilderImpl(
        result, reserved01);
  }

  public static class S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseBuilderImpl
      implements S7PayloadUserDataItem.S7PayloadUserDataItemBuilder {
    private final short result;
    private final short reserved01;

    public S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponseBuilderImpl(
        short result, short reserved01) {
      this.result = result;
      this.reserved01 = reserved01;
    }

    public S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse build(
        DataTransportErrorCode returnCode, DataTransportSize transportSize) {
      S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse
          s7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse =
              new S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse(
                  returnCode, transportSize, result, reserved01);
      return s7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse)) {
      return false;
    }
    S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse that =
        (S7PayloadUserDataItemCpuFunctionMsgSubscriptionSysResponse) o;
    return (getResult() == that.getResult())
        && (getReserved01() == that.getReserved01())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getResult(), getReserved01());
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
