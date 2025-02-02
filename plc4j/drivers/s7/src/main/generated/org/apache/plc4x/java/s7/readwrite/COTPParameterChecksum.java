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

public class COTPParameterChecksum extends COTPParameter implements Message {

  // Accessors for discriminator values.
  public Short getParameterType() {
    return (short) 0xC3;
  }

  // Properties.
  protected final short crc;

  public COTPParameterChecksum(short crc) {
    super();
    this.crc = crc;
  }

  public short getCrc() {
    return crc;
  }

  @Override
  protected void serializeCOTPParameterChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("COTPParameterChecksum");

    // Simple Field (crc)
    writeSimpleField("crc", crc, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("COTPParameterChecksum");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    COTPParameterChecksum _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (crc)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static COTPParameterBuilder staticParseCOTPParameterBuilder(
      ReadBuffer readBuffer, Short rest) throws ParseException {
    readBuffer.pullContext("COTPParameterChecksum");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short crc = readSimpleField("crc", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("COTPParameterChecksum");
    // Create the instance
    return new COTPParameterChecksumBuilderImpl(crc);
  }

  public static class COTPParameterChecksumBuilderImpl
      implements COTPParameter.COTPParameterBuilder {
    private final short crc;

    public COTPParameterChecksumBuilderImpl(short crc) {
      this.crc = crc;
    }

    public COTPParameterChecksum build() {
      COTPParameterChecksum cOTPParameterChecksum = new COTPParameterChecksum(crc);
      return cOTPParameterChecksum;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof COTPParameterChecksum)) {
      return false;
    }
    COTPParameterChecksum that = (COTPParameterChecksum) o;
    return (getCrc() == that.getCrc()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCrc());
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
