// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go test . -run=^TestGenerated$ -fix

package platform

// An OSArch is a pair of GOOS and GOARCH values indicating a platform.
type OSArch struct {
	GOOS, GOARCH string
}

func (p OSArch) String() string

// RaceDetectorSupported reports whether goos/goarch supports the race
// detector. There is a copy of this function in cmd/dist/test.go.
// Race detector only supports 48-bit VMA on arm64. But it will always
// return true for arm64, because we don't have VMA size information during
// the compile time.
func RaceDetectorSupported(goos, goarch string) bool

// MSanSupported reports whether goos/goarch supports the memory
// sanitizer option.
func MSanSupported(goos, goarch string) bool

// ASanSupported reports whether goos/goarch supports the address
// sanitizer option.
func ASanSupported(goos, goarch string) bool

// FuzzSupported reports whether goos/goarch supports fuzzing
// ('go test -fuzz=.').
func FuzzSupported(goos, goarch string) bool

// FuzzInstrumented reports whether fuzzing on goos/goarch uses coverage
// instrumentation. (FuzzInstrumented implies FuzzSupported.)
func FuzzInstrumented(goos, goarch string) bool

// MustLinkExternal reports whether goos/goarch requires external linking
// with or without cgo dependencies.
func MustLinkExternal(goos, goarch string, withCgo bool) bool

// BuildModeSupported reports whether goos/goarch supports the given build mode
// using the given compiler.
// There is a copy of this function in cmd/dist/test.go.
func BuildModeSupported(compiler, buildmode, goos, goarch string) bool

func InternalLinkPIESupported(goos, goarch string) bool

// DefaultPIE reports whether goos/goarch produces a PIE binary when using the
// "default" buildmode. On Windows this is affected by -race,
// so force the caller to pass that in to centralize that choice.
func DefaultPIE(goos, goarch string, isRace bool) bool

// ExecutableHasDWARF reports whether the linked executable includes DWARF
// symbols on goos/goarch.
func ExecutableHasDWARF(goos, goarch string) bool

// CgoSupported reports whether goos/goarch supports cgo.
func CgoSupported(goos, goarch string) bool

// FirstClass reports whether goos/goarch is considered a “first class” port.
// (See https://go.dev/wiki/PortingPolicy#first-class-ports.)
func FirstClass(goos, goarch string) bool

// Broken reports whether goos/goarch is considered a broken port.
// (See https://go.dev/wiki/PortingPolicy#broken-ports.)
func Broken(goos, goarch string) bool
