// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/reflect"
	"github.com/shogo82148/std/sync"
)

// An Encoder manages the transmission of type and data information to the
// other side of a connection.
type Encoder struct {
	mutex      sync.Mutex
	w          []io.Writer
	sent       map[reflect.Type]typeId
	countState *encoderState
	freeList   *encoderState
	byteBuf    bytes.Buffer
	err        error
}

// Before we encode a message, we reserve space at the head of the
// buffer in which to encode its length. This means we can use the
// buffer to assemble the message without another allocation.

// NewEncoder returns a new encoder that will transmit on the io.Writer.
func NewEncoder(w io.Writer) *Encoder

// Encode transmits the data item represented by the empty interface value,
// guaranteeing that all necessary type information has been transmitted first.
func (enc *Encoder) Encode(e interface{}) error

// EncodeValue transmits the data item represented by the reflection value,
// guaranteeing that all necessary type information has been transmitted first.
func (enc *Encoder) EncodeValue(value reflect.Value) error
