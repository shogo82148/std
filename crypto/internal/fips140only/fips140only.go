// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fips140only

import (
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/internal/godebug"
	"github.com/shogo82148/std/io"
)

// Enabled reports whether FIPS 140-only mode is enabled, in which non-approved
// cryptography returns an error or panics.
var Enabled = godebug.New("fips140").Value() == "only"

func ApprovedHash(h hash.Hash) bool

func ApprovedRandomReader(r io.Reader) bool
