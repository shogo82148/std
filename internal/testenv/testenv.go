// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package testenv provides information about what functionality
// is available in different testing environments run by the Go team.
//
// It is an internal package because these details are specific
// to the Go team's test setup (on build.golang.org) and not
// fundamental to tests in general.
package testenv

import (
	"github.com/shogo82148/std/testing"
)

// Builder reports the name of the builder running this test
// (for example, "linux-amd64" or "windows-386-gce").
// If the test is not running on the build infrastructure,
// Builder returns the empty string.
func Builder() string

// HasGoBuild reports whether the current system can build programs with “go build”
// and then run them with os.StartProcess or exec.Command.
func HasGoBuild() bool

// MustHaveGoBuild checks that the current system can build programs with “go build”
// and then run them with os.StartProcess or exec.Command.
// If not, MustHaveGoBuild calls t.Skip with an explanation.
func MustHaveGoBuild(t testing.TB)

// HasGoRun reports whether the current system can run programs with “go run”.
func HasGoRun() bool

// MustHaveGoRun checks that the current system can run programs with “go run”.
// If not, MustHaveGoRun calls t.Skip with an explanation.
func MustHaveGoRun(t testing.TB)

// HasParallelism reports whether the current system can execute multiple
// threads in parallel.
// There is a copy of this function in cmd/dist/test.go.
func HasParallelism() bool

// MustHaveParallelism checks that the current system can execute multiple
// threads in parallel. If not, MustHaveParallelism calls t.Skip with an explanation.
func MustHaveParallelism(t testing.TB)

// GoToolPath reports the path to the Go tool.
// It is a convenience wrapper around GoTool.
// If the tool is unavailable GoToolPath calls t.Skip.
// If the tool should be available and isn't, GoToolPath calls t.Fatal.
func GoToolPath(t testing.TB) string

// GOROOT reports the path to the directory containing the root of the Go
// project source tree. This is normally equivalent to runtime.GOROOT, but
// works even if the test binary was built with -trimpath and cannot exec
// 'go env GOROOT'.
//
// If GOROOT cannot be found, GOROOT skips t if t is non-nil,
// or panics otherwise.
func GOROOT(t testing.TB) string

// GoTool reports the path to the Go tool.
func GoTool() (string, error)

// MustHaveSource checks that the entire source tree is available under GOROOT.
// If not, it calls t.Skip with an explanation.
func MustHaveSource(t testing.TB)

// HasExternalNetwork reports whether the current system can use
// external (non-localhost) networks.
func HasExternalNetwork() bool

// MustHaveExternalNetwork checks that the current system can use
// external (non-localhost) networks.
// If not, MustHaveExternalNetwork calls t.Skip with an explanation.
func MustHaveExternalNetwork(t testing.TB)

// HasCGO reports whether the current system can use cgo.
func HasCGO() bool

// MustHaveCGO calls t.Skip if cgo is not available.
func MustHaveCGO(t testing.TB)

// CanInternalLink reports whether the current system can link programs with
// internal linking.
func CanInternalLink(withCgo bool) bool

// MustInternalLink checks that the current system can link programs with internal
// linking.
// If not, MustInternalLink calls t.Skip with an explanation.
func MustInternalLink(t testing.TB, withCgo bool)

// MustInternalLinkPIE checks whether the current system can link PIE binary using
// internal linking.
// If not, MustInternalLinkPIE calls t.Skip with an explanation.
func MustInternalLinkPIE(t testing.TB)

// MustHaveBuildMode reports whether the current system can build programs in
// the given build mode.
// If not, MustHaveBuildMode calls t.Skip with an explanation.
func MustHaveBuildMode(t testing.TB, buildmode string)

// HasSymlink reports whether the current system can use os.Symlink.
func HasSymlink() bool

// MustHaveSymlink reports whether the current system can use os.Symlink.
// If not, MustHaveSymlink calls t.Skip with an explanation.
func MustHaveSymlink(t testing.TB)

// HasLink reports whether the current system can use os.Link.
func HasLink() bool

// MustHaveLink reports whether the current system can use os.Link.
// If not, MustHaveLink calls t.Skip with an explanation.
func MustHaveLink(t testing.TB)

func SkipFlaky(t testing.TB, issue int)

func SkipFlakyNet(t testing.TB)

// CPUIsSlow reports whether the CPU running the test is suspected to be slow.
func CPUIsSlow() bool

// SkipIfShortAndSlow skips t if -short is set and the CPU running the test is
// suspected to be slow.
//
// (This is useful for CPU-intensive tests that otherwise complete quickly.)
func SkipIfShortAndSlow(t testing.TB)

// SkipIfOptimizationOff skips t if optimization is disabled.
func SkipIfOptimizationOff(t testing.TB)

// WriteImportcfg writes an importcfg file used by the compiler or linker to
// dstPath containing entries for the file mappings in packageFiles, as well
// as for the packages transitively imported by the package(s) in pkgs.
//
// pkgs may include any package pattern that is valid to pass to 'go list',
// so it may also be a list of Go source files all in the same directory.
func WriteImportcfg(t testing.TB, dstPath string, packageFiles map[string]string, pkgs ...string)

// SyscallIsNotSupported reports whether err may indicate that a system call is
// not supported by the current platform or execution environment.
func SyscallIsNotSupported(err error) bool

// ParallelOn64Bit calls t.Parallel() unless there is a case that cannot be parallel.
// This function should be used when it is necessary to avoid t.Parallel on
// 32-bit machines, typically because the test uses lots of memory.
func ParallelOn64Bit(t *testing.T)
