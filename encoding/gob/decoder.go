// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/reflect"
	"github.com/shogo82148/std/sync"
)

// A Decoder manages the receipt of type and data information read from the
// remote side of a connection.  It is safe for concurrent use by multiple
// goroutines.
//
// The Decoder does only basic sanity checking on decoded input sizes,
// and its limits are not configurable. Take caution when decoding gob data
// from untrusted sources.
type Decoder struct {
	mutex        sync.Mutex
	r            io.Reader
	buf          decBuffer
	wireType     map[typeId]*wireType
	decoderCache map[reflect.Type]map[typeId]**decEngine
	ignorerCache map[typeId]**decEngine
	freeList     *decoderState
	countBuf     []byte
	err          error
	// ignoreDepth tracks the depth of recursively parsed ignored fields
	ignoreDepth int
}

// NewDecoder returns a new decoder that reads from the [io.Reader].
// If r does not also implement [io.ByteReader], it will be wrapped in a
// [bufio.Reader].
func NewDecoder(r io.Reader) *Decoder

// Decode reads the next value from the input stream and stores
// it in the data represented by the empty interface value.
// If e is nil, the value will be discarded. Otherwise,
// the value underlying e must be a pointer to the
// correct type for the next data item received.
// If the input is at EOF, Decode returns [io.EOF] and
// does not modify e.
func (dec *Decoder) Decode(e any) error

// DecodeValue reads the next value from the input stream.
// If v is the zero reflect.Value (v.Kind() == Invalid), DecodeValue discards the value.
// Otherwise, it stores the value into v. In that case, v must represent
// a non-nil pointer to data or be an assignable reflect.Value (v.CanSet())
// If the input is at EOF, DecodeValue returns [io.EOF] and
// does not modify v.
func (dec *Decoder) DecodeValue(v reflect.Value) error
