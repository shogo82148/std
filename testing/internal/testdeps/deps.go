// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package testdeps provides access to dependencies needed by test execution.
//
// This package is imported by the generated main package, which passes
// TestDeps into testing.Main. This allows tests to use packages at run time
// without making those packages direct dependencies of package testing.
// Direct dependencies of package testing are harder to write tests for.
package testdeps

import (
	"github.com/shogo82148/std/internal/fuzz"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/reflect"
	"github.com/shogo82148/std/time"
)

// Cover indicates whether coverage is enabled.
var Cover bool

// TestDeps is an implementation of the testing.testDeps interface,
// suitable for passing to [testing.MainStart].
type TestDeps struct{}

func (TestDeps) MatchString(pat, str string) (result bool, err error)

func (TestDeps) StartCPUProfile(w io.Writer) error

func (TestDeps) StopCPUProfile()

func (TestDeps) WriteProfileTo(name string, w io.Writer, debug int) error

// ImportPath is the import path of the testing binary, set by the generated main function.
var ImportPath string

func (TestDeps) ImportPath() string

func (TestDeps) StartTestLog(w io.Writer)

func (TestDeps) StopTestLog() error

// SetPanicOnExit0 tells the os package whether to panic on os.Exit(0).
func (TestDeps) SetPanicOnExit0(v bool)

func (TestDeps) CoordinateFuzzing(
	timeout time.Duration,
	limit int64,
	minimizeTimeout time.Duration,
	minimizeLimit int64,
	parallel int,
	seed []fuzz.CorpusEntry,
	types []reflect.Type,
	corpusDir,
	cacheDir string) (err error)

func (TestDeps) RunFuzzWorker(fn func(fuzz.CorpusEntry) error) error

func (TestDeps) ReadCorpus(dir string, types []reflect.Type) ([]fuzz.CorpusEntry, error)

func (TestDeps) CheckCorpus(vals []any, types []reflect.Type) error

func (TestDeps) ResetCoverage()

func (TestDeps) SnapshotCoverage()

var CoverMode string
var Covered string

func (TestDeps) InitRuntimeCoverage() (mode string, tearDown func(string, string) (string, error), snapcov func() float64)
