// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encodemeta

import (
	"github.com/shogo82148/std/hash"
	"github.com/shogo82148/std/internal/coverage"
	"github.com/shogo82148/std/internal/coverage/stringtab"
	"github.com/shogo82148/std/io"
)

type CoverageMetaDataBuilder struct {
	stab    stringtab.Writer
	funcs   []funcDesc
	tmp     []byte
	h       hash.Hash
	pkgpath uint32
	pkgname uint32
	modpath uint32
	debug   bool
	werr    error
}

func NewCoverageMetaDataBuilder(pkgpath string, pkgname string, modulepath string) (*CoverageMetaDataBuilder, error)

// AddFunc registers a new function with the meta data builder.
func (b *CoverageMetaDataBuilder) AddFunc(f coverage.FuncDesc) uint

// Emit writes the meta-data accumulated so far in this builder to 'w'.
// Returns a hash of the meta-data payload and an error.
func (b *CoverageMetaDataBuilder) Emit(w io.WriteSeeker) ([16]byte, error)

// HashFuncDesc computes an md5 sum of a coverage.FuncDesc and returns
// a digest for it.
func HashFuncDesc(f *coverage.FuncDesc) [16]byte
