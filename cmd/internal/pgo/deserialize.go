// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pgo

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/io"
)

// IsSerialized returns true if r is a serialized Profile.
//
// IsSerialized only peeks at r, so seeking back after calling is not
// necessary.
func IsSerialized(r *bufio.Reader) (bool, error)

// FromSerialized parses a profile from serialization output of Profile.WriteTo.
func FromSerialized(r io.Reader) (*Profile, error)
