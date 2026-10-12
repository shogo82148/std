// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sgutil

type PathList struct {
	paths     []string
	fromShell bool
}

func FlagPathList(name string, usage string, value ...string) *PathList

func (l *PathList) String() string

func (l *PathList) Set(val string) error

// Find returns the first element of l containing a file by the given name. If
// file is not found in the path list, it returns a descriptive error.
func (l *PathList) Find(file string) (string, error)

// FlagXEDPath registers and returns a global -xedPath flag.
func FlagXEDPath(genRoot string) *PathList

func ResolveXEDPath(flag *PathList) (string, error)

// FlagARM64Path registers and returns a global -arm64Path flag.
func FlagARM64Path(genRoot string) *PathList

func ResolveARM64Path(flag *PathList) (string, error)
