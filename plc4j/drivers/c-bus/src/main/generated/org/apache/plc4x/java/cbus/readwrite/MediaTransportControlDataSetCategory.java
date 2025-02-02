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
package org.apache.plc4x.java.cbus.readwrite;

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

public class MediaTransportControlDataSetCategory extends MediaTransportControlData
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final short categoryNumber;

  public MediaTransportControlDataSetCategory(
      MediaTransportControlCommandTypeContainer commandTypeContainer,
      byte mediaLinkGroup,
      short categoryNumber) {
    super(commandTypeContainer, mediaLinkGroup);
    this.categoryNumber = categoryNumber;
  }

  public short getCategoryNumber() {
    return categoryNumber;
  }

  @Override
  protected void serializeMediaTransportControlDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("MediaTransportControlDataSetCategory");

    // Simple Field (categoryNumber)
    writeSimpleField("categoryNumber", categoryNumber, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("MediaTransportControlDataSetCategory");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    MediaTransportControlDataSetCategory _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (categoryNumber)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static MediaTransportControlDataBuilder staticParseMediaTransportControlDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("MediaTransportControlDataSetCategory");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    short categoryNumber = readSimpleField("categoryNumber", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("MediaTransportControlDataSetCategory");
    // Create the instance
    return new MediaTransportControlDataSetCategoryBuilderImpl(categoryNumber);
  }

  public static class MediaTransportControlDataSetCategoryBuilderImpl
      implements MediaTransportControlData.MediaTransportControlDataBuilder {
    private final short categoryNumber;

    public MediaTransportControlDataSetCategoryBuilderImpl(short categoryNumber) {
      this.categoryNumber = categoryNumber;
    }

    public MediaTransportControlDataSetCategory build(
        MediaTransportControlCommandTypeContainer commandTypeContainer, byte mediaLinkGroup) {
      MediaTransportControlDataSetCategory mediaTransportControlDataSetCategory =
          new MediaTransportControlDataSetCategory(
              commandTypeContainer, mediaLinkGroup, categoryNumber);
      return mediaTransportControlDataSetCategory;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof MediaTransportControlDataSetCategory)) {
      return false;
    }
    MediaTransportControlDataSetCategory that = (MediaTransportControlDataSetCategory) o;
    return (getCategoryNumber() == that.getCategoryNumber()) && super.equals(that) && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getCategoryNumber());
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
