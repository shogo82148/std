// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package test

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdTest = &base.Command{
	CustomFlags: true,
	UsageLine:   testUsage,
	Short:       "test packages",
	Long: `
'Go test' automates testing the packages named by the import paths.
It prints a summary of the test results in the format:

	ok   archive/tar   0.011s
	FAIL archive/zip   0.022s
	ok   compress/gzip 0.033s
	...

followed by detailed output for each failed package.

'Go test' recompiles each package along with any files with names matching
the file pattern "*_test.go".
These additional files can contain test functions, benchmark functions, fuzz
tests and example functions. See 'go help testfunc' for more.
Each listed package causes the execution of a separate test binary.
Files whose names begin with "_" (including "_test.go") or "." are ignored.

Test files that declare a package with the suffix "_test" will be compiled as a
separate package, and then linked and run with the main test binary.

The go tool will ignore a directory named "testdata", making it available
to hold ancillary data needed by the tests.

As part of building a test binary, go test runs go vet on the package
and its test source files to identify significant problems. If go vet
finds any problems, go test reports those and does not run the test
binary. Only a high-confidence subset of the default go vet checks are
used. That subset is: atomic, bool, buildtags, directive, errorsas,
ifaceassert, nilfunc, printf, stringintconv, and tests. You can see
the documentation for these and other vet tests via "go doc cmd/vet".
To disable the running of go vet, use the -vet=off flag. To run all
checks, use the -vet=all flag.

All test output and summary lines are printed to the go command's
standard output, even if the test printed them to its own standard
error. (The go command's standard error is reserved for printing
errors building the tests.)

The go command places $GOROOT/bin at the beginning of $PATH
in the test's environment, so that tests that execute
'go' commands use the same 'go' as the parent 'go test' command.

Go test runs in two different modes:

The first, called local directory mode, occurs when go test is
invoked with no package arguments (for example, 'go test' or 'go
test -v'). In this mode, go test compiles the package sources and
tests found in the current directory and then runs the resulting
test binary. In this mode, caching (discussed below) is disabled.
After the package test finishes, go test prints a summary line
showing the test status ('ok' or 'FAIL'), package name, and elapsed
time.

The second, called package list mode, occurs when go test is invoked
with explicit package arguments (for example 'go test math', 'go
test ./...', and even 'go test .'). In this mode, go test compiles
and tests each of the packages listed on the command line. If a
package test passes, go test prints only the final 'ok' summary
line. If a package test fails, go test prints the full test output.
If invoked with the -bench or -v flag, go test prints the full
output even for passing package tests, in order to display the
requested benchmark results or verbose logging. After the package
tests for all of the listed packages finish, and their output is
printed, go test prints a final 'FAIL' status if any package test
has failed.

In package list mode only, go test caches successful package test
results to avoid unnecessary repeated running of tests. When the
result of a test can be recovered from the cache, go test will
redisplay the previous output instead of running the test binary
again. When this happens, go test prints '(cached)' in place of the
elapsed time in the summary line.

The rule for a match in the cache is that the run involves the same
test binary and the flags on the command line come entirely from a
restricted set of 'cacheable' test flags, defined as -benchtime, -cpu,
-list, -parallel, -run, -short, -timeout, -failfast, -fullpath and -v.
If a run of go test has any test or non-test flags outside this set,
the result is not cached. To disable test caching, use any test flag
or argument other than the cacheable flags. The idiomatic way to disable
test caching explicitly is to use -count=1. Tests that open files within
the package's module or that consult environment variables only
match future runs in which the files and environment variables are
unchanged. A cached test result is treated as executing in no time
at all, so a successful package test result will be cached and
reused regardless of -timeout setting.

In addition to the build flags, the flags handled by 'go test' itself are:

	-args
	    Pass the remainder of the command line (everything after -args)
	    to the test binary, uninterpreted and unchanged.
	    Because this flag consumes the remainder of the command line,
	    the package list (if present) must appear before this flag.

	-c
	    Compile the test binary to pkg.test in the current directory but do not run it
	    (where pkg is the last element of the package's import path).
	    The file name or target directory can be changed with the -o flag.

	-exec xprog
	    Run the test binary using xprog. The behavior is the same as
	    in 'go run'. See 'go help run' for details.

	-json
	    Convert test output to JSON suitable for automated processing.
	    See 'go doc test2json' for the encoding details.
	    Also emits build output in JSON. See 'go help buildjson'.

	-o file
	    Compile the test binary to the named file.
	    The test still runs (unless -c or -i is specified).
	    If file ends in a slash or names an existing directory,
	    the test is written to pkg.test in that directory.

The test binary also accepts flags that control execution of the test; these
flags are also accessible by 'go test'. See 'go help testflag' for details.

For more about build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build, go vet.
`,
}

var HelpTestflag = &base.Command{
	UsageLine: "testflag",
	Short:     "testing flags",
	Long: `
The 'go test' command takes both flags that apply to 'go test' itself
and flags that apply to the resulting test binary.

Several of the flags control profiling and write an execution profile
suitable for "go tool pprof"; run "go tool pprof -h" for more
information. The --alloc_space, --alloc_objects, and --show_bytes
options of pprof control how the information is presented.

The following flags are recognized by the 'go test' command and
control the execution of any test:

	-bench regexp
	    Run only those benchmarks matching a regular expression.
	    By default, no benchmarks are run.
	    To run all benchmarks, use '-bench .' or '-bench=.'.
	    The regular expression is split by unbracketed slash (/)
	    characters into a sequence of regular expressions, and each
	    part of a benchmark's identifier must match the corresponding
	    element in the sequence, if any. Possible parents of matches
	    are run with b.N=1 to identify sub-benchmarks. For example,
	    given -bench=X/Y, top-level benchmarks matching X are run
	    with b.N=1 to find any sub-benchmarks matching Y, which are
	    then run in full.

	-benchtime t
	    Run enough iterations of each benchmark to take t, specified
	    as a time.Duration (for example, -benchtime 1h30s).
	    The default is 1 second (1s).
	    The special syntax Nx means to run the benchmark N times
	    (for example, -benchtime 100x).

	-count n
	    Run each test, benchmark, and fuzz seed n times (default 1).
	    If -cpu is set, run n times for each GOMAXPROCS value.
	    Examples are always run once. -count does not apply to
	    fuzz tests matched by -fuzz.

	-cover
	    Enable coverage analysis.
	    Note that because coverage works by annotating the source
	    code before compilation, compilation and test failures with
	    coverage enabled may report line numbers that don't correspond
	    to the original sources.

	-covermode set,count,atomic
	    Set the mode for coverage analysis for the package[s]
	    being tested. The default is "set" unless -race is enabled,
	    in which case it is "atomic".
	    The values:
		set: bool: does this statement run?
		count: int: how many times does this statement run?
		atomic: int: count, but correct in multithreaded tests;
			significantly more expensive.
	    Sets -cover.

	-coverpkg pattern1,pattern2,pattern3
	    Apply coverage analysis in each test to packages whose import paths
	    match the patterns. The default is for each test to analyze only
	    the package being tested. See 'go help packages' for a description
	    of package patterns. Sets -cover.

	-cpu 1,2,4
	    Specify a list of GOMAXPROCS values for which the tests, benchmarks or
	    fuzz tests should be executed. The default is the current value
	    of GOMAXPROCS. -cpu does not apply to fuzz tests matched by -fuzz.

	-failfast
	    Do not start new tests after the first test failure.

	-fullpath
	    Show full file names in the error messages.

	-fuzz regexp
	    Run the fuzz test matching the regular expression. When specified,
	    the command line argument must match exactly one package within the
	    main module, and regexp must match exactly one fuzz test within
	    that package. Fuzzing will occur after tests, benchmarks, seed corpora
	    of other fuzz tests, and examples have completed. See the Fuzzing
	    section of the testing package documentation for details.

	-fuzztime t
	    Run enough iterations of the fuzz target during fuzzing to take t,
	    specified as a time.Duration (for example, -fuzztime 1h30s).
		The default is to run forever.
	    The special syntax Nx means to run the fuzz target N times
	    (for example, -fuzztime 1000x).

	-fuzzminimizetime t
	    Run enough iterations of the fuzz target during each minimization
	    attempt to take t, as specified as a time.Duration (for example,
	    -fuzzminimizetime 30s).
		The default is 60s.
	    The special syntax Nx means to run the fuzz target N times
	    (for example, -fuzzminimizetime 100x).

	-json
	    Log verbose output and test results in JSON. This presents the
	    same information as the -v flag in a machine-readable format.

	-list regexp
	    List tests, benchmarks, fuzz tests, or examples matching the regular
	    expression. No tests, benchmarks, fuzz tests, or examples will be run.
	    This will only list top-level tests. No subtest or subbenchmarks will be
	    shown.

	-parallel n
	    Allow parallel execution of test functions that call t.Parallel, and
	    fuzz targets that call t.Parallel when running the seed corpus.
	    The value of this flag is the maximum number of tests to run
	    simultaneously.
	    While fuzzing, the value of this flag is the maximum number of
	    subprocesses that may call the fuzz function simultaneously, regardless of
	    whether T.Parallel is called.
	    By default, -parallel is set to the value of GOMAXPROCS.
	    Setting -parallel to values higher than GOMAXPROCS may cause degraded
	    performance due to CPU contention, especially when fuzzing.
	    Note that -parallel only applies within a single test binary.
	    The 'go test' command may run tests for different packages
	    in parallel as well, according to the setting of the -p flag
	    (see 'go help build').

	-run regexp
	    Run only those tests, examples, and fuzz tests matching the regular
	    expression. For tests, the regular expression is split by unbracketed
	    slash (/) characters into a sequence of regular expressions, and each
	    part of a test's identifier must match the corresponding element in
	    the sequence, if any. Note that possible parents of matches are
	    run too, so that -run=X/Y matches and runs and reports the result
	    of all tests matching X, even those without sub-tests matching Y,
	    because it must run them to look for those sub-tests.
	    See also -skip.

	-short
	    Tell long-running tests to shorten their run time.
	    It is off by default but set during all.bash so that installing
	    the Go tree can run a sanity check but not spend time running
	    exhaustive tests.

	-shuffle off,on,N
	    Randomize the execution order of tests and benchmarks.
	    It is off by default. If -shuffle is set to on, then it will seed
	    the randomizer using the system clock. If -shuffle is set to an
	    integer N, then N will be used as the seed value. In both cases,
	    the seed will be reported for reproducibility.

	-skip regexp
	    Run only those tests, examples, fuzz tests, and benchmarks that
	    do not match the regular expression. Like for -run and -bench,
	    for tests and benchmarks, the regular expression is split by unbracketed
	    slash (/) characters into a sequence of regular expressions, and each
	    part of a test's identifier must match the corresponding element in
	    the sequence, if any.

	-timeout d
	    If a test binary runs longer than duration d, panic.
	    If d is 0, the timeout is disabled.
	    The default is 10 minutes (10m).

	-v
	    Verbose output: log all tests as they are run. Also print all
	    text from Log and Logf calls even if the test succeeds.

	-vet list
	    Configure the invocation of "go vet" during "go test"
	    to use the comma-separated list of vet checks.
	    If list is empty, "go test" runs "go vet" with a curated list of
	    checks believed to be always worth addressing.
	    If list is "off", "go test" does not run "go vet" at all.

The following flags are also recognized by 'go test' and can be used to
profile the tests during execution:

	-benchmem
	    Print memory allocation statistics for benchmarks.
	    Allocations made in C or using C.malloc are not counted.

	-blockprofile block.out
	    Write a goroutine blocking profile to the specified file
	    when all tests are complete.
	    Writes test binary as -c would.

	-blockprofilerate n
	    Control the detail provided in goroutine blocking profiles by
	    calling runtime.SetBlockProfileRate with n.
	    See 'go doc runtime.SetBlockProfileRate'.
	    The profiler aims to sample, on average, one blocking event every
	    n nanoseconds the program spends blocked. By default,
	    if -test.blockprofile is set without this flag, all blocking events
	    are recorded, equivalent to -test.blockprofilerate=1.

	-coverprofile cover.out
	    Write a coverage profile to the file after all tests have passed.
	    Sets -cover.

	-cpuprofile cpu.out
	    Write a CPU profile to the specified file before exiting.
	    Writes test binary as -c would.

	-memprofile mem.out
	    Write an allocation profile to the file after all tests have passed.
	    Writes test binary as -c would.

	-memprofilerate n
	    Enable more precise (and expensive) memory allocation profiles by
	    setting runtime.MemProfileRate. See 'go doc runtime.MemProfileRate'.
	    To profile all memory allocations, use -test.memprofilerate=1.

	-mutexprofile mutex.out
	    Write a mutex contention profile to the specified file
	    when all tests are complete.
	    Writes test binary as -c would.

	-mutexprofilefraction n
	    Sample 1 in n stack traces of goroutines holding a
	    contended mutex.

	-outputdir directory
	    Place output files from profiling in the specified directory,
	    by default the directory in which "go test" is running.

	-trace trace.out
	    Write an execution trace to the specified file before exiting.

Each of these flags is also recognized with an optional 'test.' prefix,
as in -test.v. When invoking the generated test binary (the result of
'go test -c') directly, however, the prefix is mandatory.

The 'go test' command rewrites or removes recognized flags,
as appropriate, both before and after the optional package list,
before invoking the test binary.

For instance, the command

	go test -v -myflag testdata -cpuprofile=prof.out -x

will compile the test binary and then run it as

	pkg.test -test.v -myflag testdata -test.cpuprofile=prof.out

(The -x flag is removed because it applies only to the go command's
execution, not to the test itself.)

The test flags that generate profiles (other than for coverage) also
leave the test binary in pkg.test for use when analyzing the profiles.

When 'go test' runs a test binary, it does so from within the
corresponding package's source code directory. Depending on the test,
it may be necessary to do the same when invoking a generated test
binary directly. Because that directory may be located within the
module cache, which may be read-only and is verified by checksums, the
test must not write to it or any other directory within the module
unless explicitly requested by the user (such as with the -fuzz flag,
which writes failures to testdata/fuzz).

The command-line package list, if present, must appear before any
flag not known to the go test command. Continuing the example above,
the package list would have to appear before -myflag, but could appear
on either side of -v.

When 'go test' runs in package list mode, 'go test' caches successful
package test results to avoid unnecessary repeated running of tests. To
disable test caching, use any test flag or argument other than the
cacheable flags. The idiomatic way to disable test caching explicitly
is to use -count=1.

To keep an argument for a test binary from being interpreted as a
known flag or a package name, use -args (see 'go help test') which
passes the remainder of the command line through to the test binary
uninterpreted and unaltered.

For instance, the command

	go test -v -args -x -v

will compile the test binary and then run it as

	pkg.test -test.v -x -v

Similarly,

	go test -args math

will compile the test binary and then run it as

	pkg.test math

In the first example, the -x and the second -v are passed through to the
test binary unchanged and with no effect on the go command itself.
In the second example, the argument math is passed through to the test
binary, instead of being interpreted as the package list.
`,
}

var HelpTestfunc = &base.Command{
	UsageLine: "testfunc",
	Short:     "testing functions",
	Long: `
The 'go test' command expects to find test, benchmark, and example functions
in the "*_test.go" files corresponding to the package under test.

A test function is one named TestXxx (where Xxx does not start with a
lower case letter) and should have the signature,

	func TestXxx(t *testing.T) { ... }

A benchmark function is one named BenchmarkXxx and should have the signature,

	func BenchmarkXxx(b *testing.B) { ... }

A fuzz test is one named FuzzXxx and should have the signature,

	func FuzzXxx(f *testing.F) { ... }

An example function is similar to a test function but, instead of using
*testing.T to report success or failure, prints output to os.Stdout.
If the last comment in the function starts with "Output:" then the output
is compared exactly against the comment (see examples below). If the last
comment begins with "Unordered output:" then the output is compared to the
comment, however the order of the lines is ignored. An example with no such
comment is compiled but not executed. An example with no text after
"Output:" is compiled, executed, and expected to produce no output.

Godoc displays the body of ExampleXxx to demonstrate the use
of the function, constant, or variable Xxx. An example of a method M with
receiver type T or *T is named ExampleT_M. There may be multiple examples
for a given function, constant, or variable, distinguished by a trailing _xxx,
where xxx is a suffix not beginning with an upper case letter.

Here is an example of an example:

	func ExamplePrintln() {
		Println("The output of\nthis example.")
		// Output: The output of
		// this example.
	}

Here is another example where the ordering of the output is ignored:

	func ExamplePerm() {
		for _, value := range Perm(4) {
			fmt.Println(value)
		}

		// Unordered output: 4
		// 2
		// 1
		// 3
		// 0
	}

The entire test file is presented as the example when it contains a single
example function, at least one other function, type, variable, or constant
declaration, and no tests, benchmarks, or fuzz tests.

See the documentation of the testing package for more information.
`,
}
