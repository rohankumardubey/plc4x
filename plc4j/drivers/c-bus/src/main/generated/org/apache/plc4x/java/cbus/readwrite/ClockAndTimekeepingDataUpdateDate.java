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

public class ClockAndTimekeepingDataUpdateDate extends ClockAndTimekeepingData implements Message {

  // Accessors for discriminator values.

  // Properties.
  protected final byte yearHigh;
  protected final byte yearLow;
  protected final short month;
  protected final short day;
  protected final short dayOfWeek;

  public ClockAndTimekeepingDataUpdateDate(
      ClockAndTimekeepingCommandTypeContainer commandTypeContainer,
      byte argument,
      byte yearHigh,
      byte yearLow,
      short month,
      short day,
      short dayOfWeek) {
    super(commandTypeContainer, argument);
    this.yearHigh = yearHigh;
    this.yearLow = yearLow;
    this.month = month;
    this.day = day;
    this.dayOfWeek = dayOfWeek;
  }

  public byte getYearHigh() {
    return yearHigh;
  }

  public byte getYearLow() {
    return yearLow;
  }

  public short getMonth() {
    return month;
  }

  public short getDay() {
    return day;
  }

  public short getDayOfWeek() {
    return dayOfWeek;
  }

  @Override
  protected void serializeClockAndTimekeepingDataChild(WriteBuffer writeBuffer)
      throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("ClockAndTimekeepingDataUpdateDate");

    // Simple Field (yearHigh)
    writeSimpleField("yearHigh", yearHigh, writeByte(writeBuffer, 8));

    // Simple Field (yearLow)
    writeSimpleField("yearLow", yearLow, writeByte(writeBuffer, 8));

    // Simple Field (month)
    writeSimpleField("month", month, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (day)
    writeSimpleField("day", day, writeUnsignedShort(writeBuffer, 8));

    // Simple Field (dayOfWeek)
    writeSimpleField("dayOfWeek", dayOfWeek, writeUnsignedShort(writeBuffer, 8));

    writeBuffer.popContext("ClockAndTimekeepingDataUpdateDate");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    ClockAndTimekeepingDataUpdateDate _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (yearHigh)
    lengthInBits += 8;

    // Simple field (yearLow)
    lengthInBits += 8;

    // Simple field (month)
    lengthInBits += 8;

    // Simple field (day)
    lengthInBits += 8;

    // Simple field (dayOfWeek)
    lengthInBits += 8;

    return lengthInBits;
  }

  public static ClockAndTimekeepingDataBuilder staticParseClockAndTimekeepingDataBuilder(
      ReadBuffer readBuffer) throws ParseException {
    readBuffer.pullContext("ClockAndTimekeepingDataUpdateDate");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte yearHigh = readSimpleField("yearHigh", readByte(readBuffer, 8));

    byte yearLow = readSimpleField("yearLow", readByte(readBuffer, 8));

    short month = readSimpleField("month", readUnsignedShort(readBuffer, 8));

    short day = readSimpleField("day", readUnsignedShort(readBuffer, 8));

    short dayOfWeek = readSimpleField("dayOfWeek", readUnsignedShort(readBuffer, 8));

    readBuffer.closeContext("ClockAndTimekeepingDataUpdateDate");
    // Create the instance
    return new ClockAndTimekeepingDataUpdateDateBuilderImpl(
        yearHigh, yearLow, month, day, dayOfWeek);
  }

  public static class ClockAndTimekeepingDataUpdateDateBuilderImpl
      implements ClockAndTimekeepingData.ClockAndTimekeepingDataBuilder {
    private final byte yearHigh;
    private final byte yearLow;
    private final short month;
    private final short day;
    private final short dayOfWeek;

    public ClockAndTimekeepingDataUpdateDateBuilderImpl(
        byte yearHigh, byte yearLow, short month, short day, short dayOfWeek) {
      this.yearHigh = yearHigh;
      this.yearLow = yearLow;
      this.month = month;
      this.day = day;
      this.dayOfWeek = dayOfWeek;
    }

    public ClockAndTimekeepingDataUpdateDate build(
        ClockAndTimekeepingCommandTypeContainer commandTypeContainer, byte argument) {
      ClockAndTimekeepingDataUpdateDate clockAndTimekeepingDataUpdateDate =
          new ClockAndTimekeepingDataUpdateDate(
              commandTypeContainer, argument, yearHigh, yearLow, month, day, dayOfWeek);
      return clockAndTimekeepingDataUpdateDate;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof ClockAndTimekeepingDataUpdateDate)) {
      return false;
    }
    ClockAndTimekeepingDataUpdateDate that = (ClockAndTimekeepingDataUpdateDate) o;
    return (getYearHigh() == that.getYearHigh())
        && (getYearLow() == that.getYearLow())
        && (getMonth() == that.getMonth())
        && (getDay() == that.getDay())
        && (getDayOfWeek() == that.getDayOfWeek())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(
        super.hashCode(), getYearHigh(), getYearLow(), getMonth(), getDay(), getDayOfWeek());
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
