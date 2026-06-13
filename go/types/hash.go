// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/hash/maphash"
)

type (
	// Hasher defines a hash function and equivalence relation
	// for [Types] that is consistent with [Identical].
	// Hashers are stateless.
	Hasher struct{}

	// HasherIgnoreTags is a variant of [Hasher] that is
	// consistent with [IdenticalIgnoreTags].
	HasherIgnoreTags struct{}
)

var (
	_ maphash.Hasher[Type] = Hasher{}
	_ maphash.Hasher[Type] = HasherIgnoreTags{}
)

func (Hasher) Hash(h *maphash.Hash, t Type)

func (HasherIgnoreTags) Hash(h *maphash.Hash, t Type)

func (Hasher) Equal(x, y Type) bool
func (HasherIgnoreTags) Equal(x, y Type) bool
