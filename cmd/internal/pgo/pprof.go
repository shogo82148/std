// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package pgo contains the compiler-agnostic portions of PGO profile handling.
// Notably, parsing pprof profiles and serializing/deserializing from a custom
// intermediate representation.
package pgo

import (
	"github.com/shogo82148/std/io"
)

// FromPProf parses Profile from a pprof profile.
func FromPProf(r io.Reader) (*Profile, error)
