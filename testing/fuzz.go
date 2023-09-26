// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

// fuzzWorkerExitCode is used as an exit code by fuzz worker processes after an
// internal error. This distinguishes internal errors from uncontrolled panics
// and other failiures. Keep in sync with internal/fuzz.workerExitCode.

// InternalFuzzTarget is an internal type but exported because it is
// cross-package; it is part of the implementation of the "go test" command.
type InternalFuzzTarget struct {
	Name string
	Fn   func(f *F)
}

// F is a type passed to fuzz tests.
//
// Fuzz tests run generated inputs against a provided fuzz target, which can
// find and report potential bugs in the code being tested.
//
// A fuzz test runs the seed corpus by default, which includes entries provided
// by (*F).Add and entries in the testdata/fuzz/<FuzzTestName> directory. After
// any necessary setup and calls to (*F).Add, the fuzz test must then call
// (*F).Fuzz to provide the fuzz target. See the testing package documentation
// for an example, and see the F.Fuzz and F.Add method documentation for
// details.
//
// *F methods can only be called before (*F).Fuzz. Once the test is
// executing the fuzz target, only (*T) methods can be used. The only *F methods
// that are allowed in the (*F).Fuzz function are (*F).Failed and (*F).Name.
type F struct {
	common
	fuzzContext *fuzzContext
	testContext *testContext

	inFuzzFn bool

	corpus []corpusEntry

	result     fuzzResult
	fuzzCalled bool
}

var _ TB = (*F)(nil)

// corpusEntry is an alias to the same type as internal/fuzz.CorpusEntry.
// We use a type alias because we don't want to export this type, and we can't
// import internal/fuzz from testing.

// Helper marks the calling function as a test helper function.
// When printing file and line information, that function will be skipped.
// Helper may be called simultaneously from multiple goroutines.
func (f *F) Helper()

// Fail marks the function as having failed but continues execution.
func (f *F) Fail()

// Skipped reports whether the test was skipped.
func (f *F) Skipped() bool

// Add will add the arguments to the seed corpus for the fuzz test. This will be
// a no-op if called after or within the fuzz target, and args must match the
// arguments for the fuzz target.
func (f *F) Add(args ...any)

// supportedTypes represents all of the supported types which can be fuzzed.

// Fuzz runs the fuzz function, ff, for fuzz testing. If ff fails for a set of
// arguments, those arguments will be added to the seed corpus.
//
// ff must be a function with no return value whose first argument is *T and
// whose remaining arguments are the types to be fuzzed.
// For example:
//
//	f.Fuzz(func(t *testing.T, b []byte, i int) { ... })
//
// The following types are allowed: []byte, string, bool, byte, rune, float32,
// float64, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64.
// More types may be supported in the future.
//
// ff must not call any *F methods, e.g. (*F).Log, (*F).Error, (*F).Skip. Use
// the corresponding *T method instead. The only *F methods that are allowed in
// the (*F).Fuzz function are (*F).Failed and (*F).Name.
//
// This function should be fast and deterministic, and its behavior should not
// depend on shared state. No mutatable input arguments, or pointers to them,
// should be retained between executions of the fuzz function, as the memory
// backing them may be mutated during a subsequent invocation. ff must not
// modify the underlying data of the arguments provided by the fuzzing engine.
//
// When fuzzing, F.Fuzz does not return until a problem is found, time runs out
// (set with -fuzztime), or the test process is interrupted by a signal. F.Fuzz
// should be called exactly once, unless F.Skip or F.Fail is called beforehand.
func (f *F) Fuzz(ff any)

// fuzzResult contains the results of a fuzz run.

// fuzzCrashError is satisfied by a failing input detected while fuzzing.
// These errors are written to the seed corpus and can be re-run with 'go test'.
// Errors within the fuzzing framework (like I/O errors between coordinator
// and worker processes) don't satisfy this interface.

// fuzzContext holds fields common to all fuzz tests.
