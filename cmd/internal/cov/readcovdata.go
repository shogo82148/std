// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cov

import (
	"github.com/shogo82148/std/internal/coverage"
	"github.com/shogo82148/std/internal/coverage/decodecounter"
	"github.com/shogo82148/std/internal/coverage/decodemeta"
	"github.com/shogo82148/std/internal/coverage/pods"
)

// CovDataReader is a general-purpose helper/visitor object for
// reading coverage data files in a structured way. Clients create a
// CovDataReader to process a given collection of coverage data file
// directories, then pass in a visitor object with methods that get
// invoked at various important points. CovDataReader is intended
// to facilitate common coverage data file operations such as
// merging or intersecting data files, analyzing data files, or
// dumping data files.
type CovDataReader struct {
	vis            CovDataVisitor
	indirs         []string
	matchpkg       func(name string) bool
	flags          CovDataReaderFlags
	err            error
	verbosityLevel int
}

// MakeCovDataReader creates a CovDataReader object to process the
// given set of input directories. Here 'vis' is a visitor object
// providing methods to be invoked as we walk through the data,
// 'indirs' is the set of coverage data directories to examine,
// 'verbosityLevel' controls the level of debugging trace messages
// (zero for off, higher for more output), 'flags' stores flags that
// indicate what to do if errors are detected, and 'matchpkg' is a
// caller-provided function that can be used to select specific
// packages by name (if nil, then all packages are included).
func MakeCovDataReader(vis CovDataVisitor, indirs []string, verbosityLevel int, flags CovDataReaderFlags, matchpkg func(name string) bool) *CovDataReader

type CovDataVisitor interface {
	// Invoked at the start and end of a given pod (a pod here is a
	// specific coverage meta-data files with the counter data files
	// that correspond to it).
	BeginPod(p pods.Pod)
	EndPod(p pods.Pod)

	// Invoked when the reader is starting to examine the meta-data
	// file for a pod. Here 'mdf' is the path of the file, and 'mfr'
	// is an open meta-data reader.
	VisitMetaDataFile(mdf string, mfr *decodemeta.CoverageMetaFileReader)

	// Invoked when the reader processes a counter data file, first
	// the 'begin' method at the start, then the 'end' method when
	// we're done with the file.
	BeginCounterDataFile(cdf string, cdr *decodecounter.CounterDataReader, dirIdx int)
	EndCounterDataFile(cdf string, cdr *decodecounter.CounterDataReader, dirIdx int)

	// Invoked once for each live function in the counter data file.
	VisitFuncCounterData(payload decodecounter.FuncPayload)

	// Invoked when we've finished processing the counter files in a
	// POD (e.g. no more calls to VisitFuncCounterData).
	EndCounters()

	// Invoked for each package in the meta-data file for the pod,
	// first the 'begin' method when processing of the package starts,
	// then the 'end' method when we're done
	BeginPackage(pd *decodemeta.CoverageMetaDataDecoder, pkgIdx uint32)
	EndPackage(pd *decodemeta.CoverageMetaDataDecoder, pkgIdx uint32)

	// Invoked for each function  the package being visited.
	VisitFunc(pkgIdx uint32, fnIdx uint32, fd *coverage.FuncDesc)

	// Invoked when all counter + meta-data file processing is complete.
	Finish()
}

type CovDataReaderFlags uint32

const (
	CovDataReaderNoFlags CovDataReaderFlags = 0
	PanicOnError                            = 1 << iota
	PanicOnWarning
)

func (r *CovDataReader) Visit() error
