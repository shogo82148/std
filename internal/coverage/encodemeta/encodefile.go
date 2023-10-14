// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package encodemeta

import (
	"github.com/shogo82148/std/bufio"
	"github.com/shogo82148/std/internal/coverage"
	"github.com/shogo82148/std/internal/coverage/stringtab"
	"github.com/shogo82148/std/io"
)

type CoverageMetaFileWriter struct {
	stab   stringtab.Writer
	mfname string
	w      *bufio.Writer
	tmp    []byte
	debug  bool
}

func NewCoverageMetaFileWriter(mfname string, w io.Writer) *CoverageMetaFileWriter

func (m *CoverageMetaFileWriter) Write(finalHash [16]byte, blobs [][]byte, mode coverage.CounterMode, granularity coverage.CounterGranularity) error
