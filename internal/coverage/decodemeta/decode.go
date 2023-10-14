// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decodemeta

import (
	"github.com/shogo82148/std/internal/coverage"
	"github.com/shogo82148/std/internal/coverage/slicereader"
	"github.com/shogo82148/std/internal/coverage/stringtab"
)

type CoverageMetaDataDecoder struct {
	r      *slicereader.Reader
	hdr    coverage.MetaSymbolHeader
	strtab *stringtab.Reader
	tmp    []byte
	debug  bool
}

func NewCoverageMetaDataDecoder(b []byte, readonly bool) (*CoverageMetaDataDecoder, error)

func (d *CoverageMetaDataDecoder) PackagePath() string

func (d *CoverageMetaDataDecoder) PackageName() string

func (d *CoverageMetaDataDecoder) ModulePath() string

func (d *CoverageMetaDataDecoder) NumFuncs() uint32

// ReadFunc reads the coverage meta-data for the function with index
// 'findex', filling it into the FuncDesc pointed to by 'f'.
func (d *CoverageMetaDataDecoder) ReadFunc(fidx uint32, f *coverage.FuncDesc) error
