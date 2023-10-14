// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package toolchain

import (
	"github.com/shogo82148/std/context"

	"github.com/shogo82148/std/cmd/go/internal/gover"
)

// A Switcher collects errors to be reported and then decides
// between reporting the errors or switching to a new toolchain
// to resolve them.
//
// The client calls [Switcher.Error] repeatedly with errors encountered
// and then calls [Switcher.Switch]. If the errors included any
// *gover.TooNewErrors (potentially wrapped) and switching is
// permitted by GOTOOLCHAIN, Switch switches to a new toolchain.
// Otherwise Switch prints all the errors using base.Error.
//
// See https://go.dev/doc/toolchain#switch.
type Switcher struct {
	TooNew *gover.TooNewError
	Errors []error
}

// Error reports the error to the Switcher,
// which saves it for processing during Switch.
func (s *Switcher) Error(err error)

// NeedSwitch reports whether Switch would attempt to switch toolchains.
func (s *Switcher) NeedSwitch() bool

// Switch decides whether to switch to a newer toolchain
// to resolve any of the saved errors.
// It switches if toolchain switches are permitted and there is at least one TooNewError.
//
// If Switch decides not to switch toolchains, it prints the errors using base.Error and returns.
//
// If Switch decides to switch toolchains but cannot identify a toolchain to use.
// it prints the errors along with one more about not being able to find the toolchain
// and returns.
//
// Otherwise, Switch prints an informational message giving a reason for the
// switch and the toolchain being invoked and then switches toolchains.
// This operation never returns.
func (s *Switcher) Switch(ctx context.Context)

// SwitchOrFatal attempts a toolchain switch based on the information in err
// and otherwise falls back to base.Fatal(err).
func SwitchOrFatal(ctx context.Context, err error)

// NewerToolchain returns the name of the toolchain to use when we need
// to switch to a newer toolchain that must support at least the given Go version.
// See https://go.dev/doc/toolchain#switch.
//
// If the latest major release is 1.N.0, we use the latest patch release of 1.(N-1) if that's >= version.
// Otherwise we use the latest 1.N if that's allowed.
// Otherwise we use the latest release.
func NewerToolchain(ctx context.Context, version string) (string, error)

// HasAuto reports whether the GOTOOLCHAIN setting allows "auto" upgrades.
func HasAuto() bool

// HasPath reports whether the GOTOOLCHAIN setting allows "path" upgrades.
func HasPath() bool
