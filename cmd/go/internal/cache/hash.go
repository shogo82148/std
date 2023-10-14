// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/hash"
)

// HashSize is the number of bytes in a hash.
const HashSize = 32

// A Hash provides access to the canonical hash function used to index the cache.
// The current implementation uses salted SHA256, but clients must not assume this.
type Hash struct {
	h    hash.Hash
	name string
	buf  *bytes.Buffer
}

// Subkey returns an action ID corresponding to mixing a parent
// action ID with a string description of the subkey.
func Subkey(parent ActionID, desc string) ActionID

// NewHash returns a new Hash.
// The caller is expected to Write data to it and then call Sum.
func NewHash(name string) *Hash

// Write writes data to the running hash.
func (h *Hash) Write(b []byte) (int, error)

// Sum returns the hash of the data written previously.
func (h *Hash) Sum() [HashSize]byte

// FileHash returns the hash of the named file.
// It caches repeated lookups for a given file,
// and the cache entry for a file can be initialized
// using SetFileHash.
// The hash used by FileHash is not the same as
// the hash used by NewHash.
func FileHash(file string) ([HashSize]byte, error)

// SetFileHash sets the hash returned by FileHash for file.
func SetFileHash(file string, sum [HashSize]byte)
