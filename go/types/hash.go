// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/hash/maphash"
)

type (
	// Hasher and HasherIgnoreTags define hash functions and
	// equivalence relations for [Types] that are consistent with
	// [Identical] and [IdenticalIgnoreTags], respectively.
	// They use the same hash function, which ignores tags;
	// only the Equal methods vary.
	//
	// Hashers are stateless.
	Hasher           struct{}
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
