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

public class NLMUpdateKeyDistributionKey extends NLM implements Message {

  // Accessors for discriminator values.
  public Short getMessageType() {
    return (short) 0x0F;
  }

  // Properties.
  protected final byte keyRevision;
  protected final NLMUpdateKeyUpdateKeyEntry key;

  // Arguments.
  protected final Integer apduLength;

  public NLMUpdateKeyDistributionKey(
      byte keyRevision, NLMUpdateKeyUpdateKeyEntry key, Integer apduLength) {
    super(apduLength);
    this.keyRevision = keyRevision;
    this.key = key;
    this.apduLength = apduLength;
  }

  public byte getKeyRevision() {
    return keyRevision;
  }

  public NLMUpdateKeyUpdateKeyEntry getKey() {
    return key;
  }

  @Override
  protected void serializeNLMChild(WriteBuffer writeBuffer) throws SerializationException {
    PositionAware positionAware = writeBuffer;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();
    int startPos = positionAware.getPos();
    writeBuffer.pushContext("NLMUpdateKeyDistributionKey");

    // Simple Field (keyRevision)
    writeSimpleField("keyRevision", keyRevision, writeByte(writeBuffer, 8));

    // Simple Field (key)
    writeSimpleField("key", key, new DataWriterComplexDefault<>(writeBuffer));

    writeBuffer.popContext("NLMUpdateKeyDistributionKey");
  }

  @Override
  public int getLengthInBytes() {
    return (int) Math.ceil((float) getLengthInBits() / 8.0);
  }

  @Override
  public int getLengthInBits() {
    int lengthInBits = super.getLengthInBits();
    NLMUpdateKeyDistributionKey _value = this;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    // Simple field (keyRevision)
    lengthInBits += 8;

    // Simple field (key)
    lengthInBits += key.getLengthInBits();

    return lengthInBits;
  }

  public static NLMBuilder staticParseNLMBuilder(ReadBuffer readBuffer, Integer apduLength)
      throws ParseException {
    readBuffer.pullContext("NLMUpdateKeyDistributionKey");
    PositionAware positionAware = readBuffer;
    int startPos = positionAware.getPos();
    int curPos;
    boolean _lastItem = ThreadLocalHelper.lastItemThreadLocal.get();

    byte keyRevision = readSimpleField("keyRevision", readByte(readBuffer, 8));

    NLMUpdateKeyUpdateKeyEntry key =
        readSimpleField(
            "key",
            new DataReaderComplexDefault<>(
                () -> NLMUpdateKeyUpdateKeyEntry.staticParse(readBuffer), readBuffer));

    readBuffer.closeContext("NLMUpdateKeyDistributionKey");
    // Create the instance
    return new NLMUpdateKeyDistributionKeyBuilderImpl(keyRevision, key, apduLength);
  }

  public static class NLMUpdateKeyDistributionKeyBuilderImpl implements NLM.NLMBuilder {
    private final byte keyRevision;
    private final NLMUpdateKeyUpdateKeyEntry key;
    private final Integer apduLength;

    public NLMUpdateKeyDistributionKeyBuilderImpl(
        byte keyRevision, NLMUpdateKeyUpdateKeyEntry key, Integer apduLength) {
      this.keyRevision = keyRevision;
      this.key = key;
      this.apduLength = apduLength;
    }

    public NLMUpdateKeyDistributionKey build(Integer apduLength) {

      NLMUpdateKeyDistributionKey nLMUpdateKeyDistributionKey =
          new NLMUpdateKeyDistributionKey(keyRevision, key, apduLength);
      return nLMUpdateKeyDistributionKey;
    }
  }

  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (!(o instanceof NLMUpdateKeyDistributionKey)) {
      return false;
    }
    NLMUpdateKeyDistributionKey that = (NLMUpdateKeyDistributionKey) o;
    return (getKeyRevision() == that.getKeyRevision())
        && (getKey() == that.getKey())
        && super.equals(that)
        && true;
  }

  @Override
  public int hashCode() {
    return Objects.hash(super.hashCode(), getKeyRevision(), getKey());
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
