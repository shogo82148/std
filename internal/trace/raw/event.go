// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package raw

import (
	"github.com/shogo82148/std/internal/trace/tracev2"
	"github.com/shogo82148/std/internal/trace/version"
)

// Event is a simple representation of a trace event.
//
// Note that this typically includes much more than just
// timestamped events, and it also represents parts of the
// trace format's framing. (But not interpreted.)
type Event struct {
	Version version.Version
	Ev      tracev2.EventType
	Args    []uint64
	Data    []byte
}

// String returns the canonical string representation of the event.
//
// This format is the same format that is parsed by the TextReader
// and emitted by the TextWriter.
func (e *Event) String() string

// EncodedSize returns the canonical encoded size of an event.
func (e *Event) EncodedSize() int
