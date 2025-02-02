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
package org.apache.plc4x.java.ads.readwrite;

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

public class AdsMultiRequestItemWrite extends AdsMultiRequestItem implements Message {

  // Accessors for discriminator values.
  public Long getIndexGroup() {
    return (long) 61569L;
  }

  // Properties.
  protected final long itemIndexGroup;
  protected final long itemIndexOffset;
  protected final long itemWriteLength;

  public AdsMultiRequestItemWrite(long itemIndexGroup, long itemIndexOffset, long itemWriteLength) {
    super();
    this.itemIndexGroup = itemIndexGroup;
    this.itemIndexOffset = itemIndexOffset;
    this.itemWriteLength = itemWriteLength;
  }

  public long getItemIndexGroup() {
    return itemIndexGroup;
  }

  public long getItemIndexOffset() {
    return itemIndexOffset;
  }

  public long getItemWriteLength() {
    return itemWriteLength;
  }

  @Override
  protected void serializeAdsMultiRequestItemChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("AdsMultiRequestItemWrite");

    // Simple Field (itemIndexGroup)
    writeSimpleField("itemIndexGroup", itemIndexGroup, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (itemIndexOffset)
    writeSimpleField("itemIndexOffset", itemIndexOffset, writeUnsignedLong(writeBuffer, 32));

    // Simple Field (itemWriteLength)
    writeSimpleField("itemWriteLength", itemWriteLength, writeUnsignedLong(writeBuffer, 32));

    writeBuffer.popContext("AdsMultiRequestItemWrite");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    AdsMultiRequestItemWrite _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (itemIndexGroup)
    lengthInBits += 32;

    // Simple field (itemIndexOffset)
    lengthInBits += 32;

    // Simple field (itemWriteLength)
    lengthInBits += 32;

    return lengthInBits;
  }

  public static AdsMultiRequestItemBuilder staticParseAdsMultiRequestItemBuilder(
      ReadBuffer readBuffer, Long indexGroup) throws ParseException {
    readBuffer.pullContext("AdsMultiRequestItemWrite");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    long itemIndexGroup = readSimpleField("itemIndexGroup", readUnsignedLong(readBuffer, 32));

    long itemIndexOffset = readSimpleField("itemIndexOffset", readUnsignedLong(readBuffer, 32));

    long itemWriteLength = readSimpleField("itemWriteLength", readUnsignedLong(readBuffer, 32));

    readBuffer.closeContext("AdsMultiRequestItemWrite");
    // Create the instance
    return new AdsMultiRequestItemWriteBuilderImpl(
        itemIndexGroup, itemIndexOffset, itemWriteLength);
  }

  public static class AdsMultiRequestItemWriteBuilderImpl
      implements AdsMultiRequestItem.AdsMultiRequestItemBuilder {
    private final long itemIndexGroup;
    private final long itemIndexOffset;
    private final long itemWriteLength;

    public AdsMultiRequestItemWriteBuilderImpl(
        long itemIndexGroup, long itemIndexOffset, long itemWriteLength) {
      this.itemIndexGroup = itemIndexGroup;
      this.itemIndexOffset = itemIndexOffset;
      this.itemWriteLength = itemWriteLength;
    }

    public AdsMultiRequestItemWrite build() {
      AdsMultiRequestItemWrite adsMultiRequestItemWrite =
          new AdsMultiRequestItemWrite(itemIndexGroup, itemIndexOffset, itemWriteLength);
      return adsMultiRequestItemWrite;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof AdsMultiRequestItemWrite)) {
      return false;
    }
    AdsMultiRequestItemWrite that = (AdsMultiRequestItemWrite) o;
    return (getItemIndexGroup() == that.getItemIndexGroup())
        && (getItemIndexOffset() == that.getItemIndexOffset())
        && (getItemWriteLength() == that.getItemWriteLength())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getItemIndexGroup(), getItemIndexOffset(), getItemWriteLength());
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
