// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package load

var (
	BuildAsmflags   PerPackageFlag
	BuildGcflags    PerPackageFlag
	BuildLdflags    PerPackageFlag
	BuildGccgoflags PerPackageFlag
)

// A PerPackageFlag is a command-line flag implementation (a flag.Value)
// that allows specifying different effective flags for different packages.
// See 'go help build' for more details about per-package flags.
type PerPackageFlag struct {
	raw     string
	present bool
	values  []ppfValue
}

// Set is called each time the flag is encountered on the command line.
func (f *PerPackageFlag) Set(v string) error

func (f *PerPackageFlag) String() string

// Present reports whether the flag appeared on the command line.
func (f *PerPackageFlag) Present() bool

// For returns the flags to use for the given package.
func (f *PerPackageFlag) For(p *Package) []string
