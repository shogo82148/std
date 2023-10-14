// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gover

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/errors"
)

// FromToolchain returns the Go version for the named toolchain,
// derived from the name itself (not by running the toolchain).
// A toolchain is named "goVERSION".
// A suffix after the VERSION introduced by a -, space, or tab is removed.
// Examples:
//
//	FromToolchain("go1.2.3") == "1.2.3"
//	FromToolchain("go1.2.3-bigcorp") == "1.2.3"
//	FromToolchain("invalid") == ""
func FromToolchain(name string) string

// Startup records the information that went into the startup-time version switch.
// It is initialized by switchGoToolchain.
var Startup struct {
	GOTOOLCHAIN   string
	AutoFile      string
	AutoGoVersion string
	AutoToolchain string
}

// A TooNewError explains that a module is too new for this version of Go.
type TooNewError struct {
	What      string
	GoVersion string
	Toolchain string
}

func (e *TooNewError) Error() string

var ErrTooNew = errors.New("module too new")

func (e *TooNewError) Is(err error) bool

// A Switcher provides the ability to switch to a new toolchain in response to TooNewErrors.
// See [cmd/go/internal/toolchain.Switcher] for documentation.
type Switcher interface {
	Error(err error)
	Switch(ctx context.Context)
}
