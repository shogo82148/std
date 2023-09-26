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

// A Decoder manages the receipt of type and data information read from the
// remote side of a connection.
type Decoder struct {
	mutex        sync.Mutex
	r            io.Reader
	buf          bytes.Buffer
	wireType     map[typeId]*wireType
	decoderCache map[reflect.Type]map[typeId]**decEngine
	ignorerCache map[typeId]**decEngine
	freeList     *decoderState
	countBuf     []byte
	tmp          []byte
	err          error
}

// NewDecoder returns a new decoder that reads from the io.Reader.
// If r does not also implement io.ByteReader, it will be wrapped in a
// bufio.Reader.
func NewDecoder(r io.Reader) *Decoder

// Decode reads the next value from the connection and stores
// it in the data represented by the empty interface value.
// If e is nil, the value will be discarded. Otherwise,
// the value underlying e must be a pointer to the
// correct type for the next data item received.
func (dec *Decoder) Decode(e interface{}) error

// DecodeValue reads the next value from the connection.
// If v is the zero reflect.Value (v.Kind() == Invalid), DecodeValue discards the value.
// Otherwise, it stores the value into v.  In that case, v must represent
// a non-nil pointer to data or be an assignable reflect.Value (v.CanSet())
func (dec *Decoder) DecodeValue(v reflect.Value) error

// If debug.go is compiled into the program , debugFunc prints a human-readable
// representation of the gob data read from r by calling that file's Debug function.
// Otherwise it is nil.
