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
package org.apache.plc4x.java.knxnetip.readwrite;

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

public class ComObjectTableRealisationType2 extends ComObjectTable implements Message {

  // Accessors for discriminator values.
  public FirmwareType getFirmwareType() {
    return FirmwareType.SYSTEM_2;
  }

  // Properties.
  protected final short numEntries;
  protected final short ramFlagsTablePointer;
  protected final List<GroupObjectDescriptorRealisationType2> comObjectDescriptors;

  public ComObjectTableRealisationType2(
      short numEntries,
      short ramFlagsTablePointer,
      List<GroupObjectDescriptorRealisationType2> comObjectDescriptors) {
    super();
    this.numEntries = numEntries;
    this.ramFlagsTablePointer = ramFlagsTablePointer;
    this.comObjectDescriptors = comObjectDescriptors;
  }

  public short getNumEntries() {
    return numEntries;
  }

  public short getRamFlagsTablePointer() {
    return ramFlagsTablePointer;
  }

  public List<GroupObjectDescriptorRealisationType2> getComObjectDescriptors() {
    return comObjectDescriptors;
  }

  @Override
  protected void serializeComObjectTableChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ComObjectTableRealisationType2");

    // Simple Field (numEntries)
    writeSimpleField("numEntries", numEntries, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (ramFlagsTablePointer)
    writeSimpleField(
        "ramFlagsTablePointer", ramFlagsTablePointer, writeUnsignedShort(writeBuffer, 8));

    // Array Field (comObjectDescriptors)
    writeComplexTypeArrayField("comObjectDescriptors", comObjectDescriptors, writeBuffer);

    writeBuffer.popContext("ComObjectTableRealisationType2");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ComObjectTableRealisationType2 _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (numEntries)
    lengthInBits += 8;

    // Simple field (ramFlagsTablePointer)
    lengthInBits += 8;

    // Array field
    if (comObjectDescriptors != null) {
      int i = 0;
      for (GroupObjectDescriptorRealisationType2 element : comObjectDescriptors) {
        ThreadLocalHelper.lastItemThreadLocal.set(++i >= comObjectDescriptors.size());
        lengthInBits += element.getLengthInBits();
      }
    }

    return lengthInBits;
  }

  public static ComObjectTableBuilder staticParseComObjectTableBuilder(
      ReadBuffer readBuffer, FirmwareType firmwareType) throws ParseException {
    readBuffer.pullContext("ComObjectTableRealisationType2");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short numEntries = readSimpleField("numEntries", readUnsignedShort(readBuffer, 8));

    short ramFlagsTablePointer =
        readSimpleField("ramFlagsTablePointer", readUnsignedShort(readBuffer, 8));

    List<GroupObjectDescriptorRealisationType2> comObjectDescriptors =
        readCountArrayField(
            "comObjectDescriptors",
            new DataReaderComplexDefault<>(
                () -> GroupObjectDescriptorRealisationType2.staticParse(readBuffer), readBuffer),
            numEntries);

    readBuffer.closeContext("ComObjectTableRealisationType2");
    // Create the instance
    return new ComObjectTableRealisationType2BuilderImpl(
        numEntries, ramFlagsTablePointer, comObjectDescriptors);
  }

  public static class ComObjectTableRealisationType2BuilderImpl
      implements ComObjectTable.ComObjectTableBuilder {
    private final short numEntries;
    private final short ramFlagsTablePointer;
    private final List<GroupObjectDescriptorRealisationType2> comObjectDescriptors;

    public ComObjectTableRealisationType2BuilderImpl(
        short numEntries,
        short ramFlagsTablePointer,
        List<GroupObjectDescriptorRealisationType2> comObjectDescriptors) {
      this.numEntries = numEntries;
      this.ramFlagsTablePointer = ramFlagsTablePointer;
      this.comObjectDescriptors = comObjectDescriptors;
    }

    public ComObjectTableRealisationType2 build() {
      ComObjectTableRealisationType2 comObjectTableRealisationType2 =
          new ComObjectTableRealisationType2(
              numEntries, ramFlagsTablePointer, comObjectDescriptors);
      return comObjectTableRealisationType2;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ComObjectTableRealisationType2)) {
      return false;
    }
    ComObjectTableRealisationType2 that = (ComObjectTableRealisationType2) o;
    return (getNumEntries() == that.getNumEntries())
        && (getRamFlagsTablePointer() == that.getRamFlagsTablePointer())
        && (getComObjectDescriptors() == that.getComObjectDescriptors())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getNumEntries(), getRamFlagsTablePointer(), getComObjectDescriptors());
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
