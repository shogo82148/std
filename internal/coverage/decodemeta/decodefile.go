// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package decodemeta

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/internal/coverage"
	"github.com/shogo82148/std/internal/coverage/stringtab"
	"github.com/shogo82148/std/os"
)

// CoverageMetaFileReader provides state and methods for reading
// a meta-data file from a code coverage run.
type CoverageMetaFileReader struct {
	f          *os.File
	hdr        coverage.MetaFileHeader
	tmp        []byte
	pkgOffsets []uint64
	pkgLengths []uint64
	strtab     *stringtab.Reader
	fileRdr    *bufio.Reader
	fileView   []byte
	debug      bool
}

// NewCoverageMetaFileReader returns a new helper object for reading
// the coverage meta-data output file 'f'. The param 'fileView' is a
// read-only slice containing the contents of 'f' obtained by mmap'ing
// the file read-only; 'fileView' may be nil, in which case the helper
// will read the contents of the file using regular file Read
// operations.
func NewCoverageMetaFileReader(f *os.File, fileView []byte) (*CoverageMetaFileReader, error)

// NumPackages returns the number of packages for which this file
// contains meta-data.
func (r *CoverageMetaFileReader) NumPackages() uint64

// CounterMode returns the counter mode (set, count, atomic) used
// when building for coverage for the program that produce this
// meta-data file.
func (r *CoverageMetaFileReader) CounterMode() coverage.CounterMode

// CounterGranularity returns the counter granularity (single counter per
// function, or counter per block) selected when building for coverage
// for the program that produce this meta-data file.
func (r *CoverageMetaFileReader) CounterGranularity() coverage.CounterGranularity

// FileHash returns the hash computed for all of the package meta-data
// blobs. Coverage counter data files refer to this hash, and the
// hash will be encoded into the meta-data file name.
func (r *CoverageMetaFileReader) FileHash() [16]byte

// GetPackageDecoder requests a decoder object for the package within
// the meta-data file whose index is 'pkIdx'. If the
// CoverageMetaFileReader was set up with a read-only file view, a
// pointer into that file view will be returned, otherwise the buffer
// 'payloadbuf' will be written to (or if it is not of sufficient
// size, a new buffer will be allocated). Return value is the decoder,
// a byte slice with the encoded meta-data, and an error.
func (r *CoverageMetaFileReader) GetPackageDecoder(pkIdx uint32, payloadbuf []byte) (*CoverageMetaDataDecoder, []byte, error)

// GetPackagePayload returns the raw (encoded) meta-data payload for the
// package with index 'pkIdx'. As with GetPackageDecoder, if the
// CoverageMetaFileReader was set up with a read-only file view, a
// pointer into that file view will be returned, otherwise the buffer
// 'payloadbuf' will be written to (or if it is not of sufficient
// size, a new buffer will be allocated). Return value is the decoder,
// a byte slice with the encoded meta-data, and an error.
func (r *CoverageMetaFileReader) GetPackagePayload(pkIdx uint32, payloadbuf []byte) ([]byte, error)
