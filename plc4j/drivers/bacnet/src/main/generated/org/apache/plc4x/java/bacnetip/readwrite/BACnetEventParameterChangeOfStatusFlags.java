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
package org.apache.plc4x.java.bacnetip.readwrite;

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

public class BACnetEventParameterChangeOfStatusFlags extends BACnetEventParameter
    implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final BACnetOpeningTag openingTag;
  protected final BACnetContextTagUnsignedInteger timeDelay;
  protected final BACnetStatusFlagsTagged selectedFlags;
  protected final BACnetClosingTag closingTag;

  public BACnetEventParameterChangeOfStatusFlags(
      BACnetTagHeader peekedTagHeader,
      BACnetOpeningTag openingTag,
      BACnetContextTagUnsignedInteger timeDelay,
      BACnetStatusFlagsTagged selectedFlags,
      BACnetClosingTag closingTag) {
    super(peekedTagHeader);
    this.openingTag = openingTag;
    this.timeDelay = timeDelay;
    this.selectedFlags = selectedFlags;
    this.closingTag = closingTag;
  }

  public BACnetOpeningTag getOpeningTag() {
    return openingTag;
  }

  public BACnetContextTagUnsignedInteger getTimeDelay() {
    return timeDelay;
  }

  public BACnetStatusFlagsTagged getSelectedFlags() {
    return selectedFlags;
  }

  public BACnetClosingTag getClosingTag() {
    return closingTag;
  }

  @Override
  protected void serializeBACnetEventParameterChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("BACnetEventParameterChangeOfStatusFlags");

    // Simple Field (openingTag)
    writeSimpleField("openingTag", openingTag, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (timeDelay)
    writeSimpleField("timeDelay", timeDelay, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (selectedFlags)
    writeSimpleField("selectedFlags", selectedFlags, new DataWriterComplexDefault<>(writeBuffer));

    // Simple Field (closingTag)
    writeSimpleField("closingTag", closingTag, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("BACnetEventParameterChangeOfStatusFlags");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    BACnetEventParameterChangeOfStatusFlags _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (openingTag)
    lengthInBits += openingTag.getLengthInBits();

    // Simple field (timeDelay)
    lengthInBits += timeDelay.getLengthInBits();

    // Simple field (selectedFlags)
    lengthInBits += selectedFlags.getLengthInBits();

    // Simple field (closingTag)
    lengthInBits += closingTag.getLengthInBits();

    return lengthInBits;
  }

  public static BACnetEventParameterBuilder staticParseBACnetEventParameterBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("BACnetEventParameterChangeOfStatusFlags");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    BACnetOpeningTag openingTag =
        readSimpleField(
            "openingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetOpeningTag.staticParse(readBuffer, (short) (18)), readBuffer));

    BACnetContextTagUnsignedInteger timeDelay =
        readSimpleField(
            "timeDelay",
            new DataReaderComplexDefault<>(
                () ->
                    (BACnetContextTagUnsignedInteger)
                        BACnetContextTag.staticParse(
                            readBuffer,
                            (short) (0),
                            (BACnetDataType) (BACnetDataType.UNSIGNED_INTEGER)),
                readBuffer));

    BACnetStatusFlagsTagged selectedFlags =
        readSimpleField(
            "selectedFlags",
            new DataReaderComplexDefault<>(
                () ->
                    BACnetStatusFlagsTagged.staticParse(
                        readBuffer, (short) (1), (TagClass) (TagClass.CONTEXT_SPECIFIC_TAGS)),
                readBuffer));

    BACnetClosingTag closingTag =
        readSimpleField(
            "closingTag",
            new DataReaderComplexDefault<>(
                () -> BACnetClosingTag.staticParse(readBuffer, (short) (18)), readBuffer));

    readBuffer.closeContext("BACnetEventParameterChangeOfStatusFlags");
    // Create the instance
    return new BACnetEventParameterChangeOfStatusFlagsBuilderImpl(
        openingTag, timeDelay, selectedFlags, closingTag);
  }

  public static class BACnetEventParameterChangeOfStatusFlagsBuilderImpl
      implements BACnetEventParameter.BACnetEventParameterBuilder {
    private final BACnetOpeningTag openingTag;
    private final BACnetContextTagUnsignedInteger timeDelay;
    private final BACnetStatusFlagsTagged selectedFlags;
    private final BACnetClosingTag closingTag;

    public BACnetEventParameterChangeOfStatusFlagsBuilderImpl(
        BACnetOpeningTag openingTag,
        BACnetContextTagUnsignedInteger timeDelay,
        BACnetStatusFlagsTagged selectedFlags,
        BACnetClosingTag closingTag) {
      this.openingTag = openingTag;
      this.timeDelay = timeDelay;
      this.selectedFlags = selectedFlags;
      this.closingTag = closingTag;
    }

    public BACnetEventParameterChangeOfStatusFlags build(BACnetTagHeader peekedTagHeader) {
      BACnetEventParameterChangeOfStatusFlags bACnetEventParameterChangeOfStatusFlags =
          new BACnetEventParameterChangeOfStatusFlags(
              peekedTagHeader, openingTag, timeDelay, selectedFlags, closingTag);
      return bACnetEventParameterChangeOfStatusFlags;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof BACnetEventParameterChangeOfStatusFlags)) {
      return false;
    }
    BACnetEventParameterChangeOfStatusFlags that = (BACnetEventParameterChangeOfStatusFlags) o;
    return (getOpeningTag() == that.getOpeningTag())
        && (getTimeDelay() == that.getTimeDelay())
        && (getSelectedFlags() == that.getSelectedFlags())
        && (getClosingTag() == that.getClosingTag())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getOpeningTag(), getTimeDelay(), getSelectedFlags(), getClosingTag());
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
