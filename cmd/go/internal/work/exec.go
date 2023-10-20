// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Action graph execution.

package work

import (
	"github.com/shogo82148/std/context"

	"github.com/shogo82148/std/cmd/go/internal/load"
)

// Do runs the action graph rooted at root.
func (b *Builder) Do(ctx context.Context, root *Action)

// VetTool is the path to an alternate vet tool binary.
// The caller is expected to set it (if needed) before executing any vet actions.
var VetTool string

// VetFlags are the default flags to pass to vet.
// The caller is expected to set them before executing any vet actions.
var VetFlags []string

// VetExplicit records whether the vet flags were set explicitly on the command line.
var VetExplicit bool

// PkgconfigCmd returns a pkg-config binary name
// defaultPkgConfig is defined in zdefaultcc.go, written by cmd/dist.
func (b *Builder) PkgconfigCmd() string

// BuildInstallFunc is the action for installing a single package or executable.
func BuildInstallFunc(b *Builder, ctx context.Context, a *Action) (err error)

// AllowInstall returns a non-nil error if this invocation of the go command is
// allowed to install a.Target.
//
// The build of cmd/go running under its own test is forbidden from installing
// to its original GOROOT. The var is exported so it can be set by TestMain.
var AllowInstall = func(*Action) error { return nil }

// GccCmd returns a gcc command line prefix
// defaultCC is defined in zdefaultcc.go, written by cmd/dist.
func (b *Builder) GccCmd(incdir, workdir string) []string

// GxxCmd returns a g++ command line prefix
// defaultCXX is defined in zdefaultcc.go, written by cmd/dist.
func (b *Builder) GxxCmd(incdir, workdir string) []string

// CFlags returns the flags to use when invoking the C, C++ or Fortran compilers, or cgo.
func (b *Builder) CFlags(p *load.Package) (cppflags, cflags, cxxflags, fflags, ldflags []string, err error)
