// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package testing provides support for automated testing of Go packages.
// It is intended to be used in concert with the “go test” command, which automates
// execution of any function of the form
//
//	func TestXxx(*testing.T)
//
// where Xxx can be any alphanumeric string (but the first letter must not be in
// [a-z]) and serves to identify the test routine.
// These TestXxx routines should be declared within the package they are testing.
//
// Functions of the form
//
//	func BenchmarkXxx(*testing.B)
//
// are considered benchmarks, and are executed by the "go test" command when
// the -test.bench flag is provided.
//
// A sample benchmark function looks like this:
//
//	func BenchmarkHello(b *testing.B) {
//	    for i := 0; i < b.N; i++ {
//	        fmt.Sprintf("hello")
//	    }
//	}
//
// The benchmark package will vary b.N until the benchmark function lasts
// long enough to be timed reliably.  The output
//
//	testing.BenchmarkHello    10000000    282 ns/op
//
// means that the loop ran 10000000 times at a speed of 282 ns per loop.
//
// If a benchmark needs some expensive setup before running, the timer
// may be stopped:
//
//	func BenchmarkBigLen(b *testing.B) {
//	    b.StopTimer()
//	    big := NewBig()
//	    b.StartTimer()
//	    for i := 0; i < b.N; i++ {
//	        big.Len()
//	    }
//	}
//
// The package also runs and verifies example code. Example functions may
// include a concluding comment that begins with "Output:" and is compared with
// the standard output of the function when the tests are run, as in these
// examples of an example:
//
//	func ExampleHello() {
//	        fmt.Println("hello")
//	        // Output: hello
//	}
//
//	func ExampleSalutations() {
//	        fmt.Println("hello, and")
//	        fmt.Println("goodbye")
//	        // Output:
//	        // hello, and
//	        // goodbye
//	}
//
// Example functions without output comments are compiled but not executed.
//
// The naming convention to declare examples for a function F, a type T and
// method M on type T are:
//
//	func ExampleF() { ... }
//	func ExampleT() { ... }
//	func ExampleT_M() { ... }
//
// Multiple example functions for a type/function/method may be provided by
// appending a distinct suffix to the name. The suffix must start with a
// lower-case letter.
//
//	func ExampleF_suffix() { ... }
//	func ExampleT_suffix() { ... }
//	func ExampleT_M_suffix() { ... }
//
// The entire test file is presented as the example when it contains a single
// example function, at least one other function, type, variable, or constant
// declaration, and no test or benchmark functions.
package testing

// common holds the elements common between T and B and
// captures common methods such as Errorf.

// Short reports whether the -test.short flag is set.
func Short() bool

// T is a type passed to Test functions to manage test state and support formatted test logs.
// Logs are accumulated during execution and dumped to standard error when done.
type T struct {
	common
	name          string
	startParallel chan bool
}

// Parallel signals that this test is to be run in parallel with (and only with)
// other parallel tests in this CPU group.
func (t *T) Parallel()

// An internal type but exported because it is cross-package; part of the implementation
// of the "go test" command.
type InternalTest struct {
	Name string
	F    func(*T)
}

// An internal function but exported because it is cross-package; part of the implementation
// of the "go test" command.
func Main(matchString func(pat, str string) (bool, error), tests []InternalTest, benchmarks []InternalBenchmark, examples []InternalExample)

func RunTests(matchString func(pat, str string) (bool, error), tests []InternalTest) (ok bool)
