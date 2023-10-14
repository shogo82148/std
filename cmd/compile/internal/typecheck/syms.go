// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typecheck

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
)

// LookupRuntime returns a function or variable declared in
// _builtin/runtime.go. If types_ is non-empty, successive occurrences
// of the "any" placeholder type will be substituted.
func LookupRuntime(name string, types_ ...*types.Type) *ir.Name

// AutoLabel generates a new Name node for use with
// an automatically generated label.
// prefix is a short mnemonic (e.g. ".s" for switch)
// to help with debugging.
// It should begin with "." to avoid conflicts with
// user labels.
func AutoLabel(prefix string) *types.Sym

func Lookup(name string) *types.Sym

// InitRuntime loads the definitions for the low-level runtime functions,
// so that the compiler can generate calls to them,
// but does not make them visible to user code.
func InitRuntime()

// LookupRuntimeFunc looks up Go function name in package runtime. This function
// must follow the internal calling convention.
func LookupRuntimeFunc(name string) *obj.LSym

// LookupRuntimeVar looks up a variable (or assembly function) name in package
// runtime. If this is a function, it may have a special calling
// convention.
func LookupRuntimeVar(name string) *obj.LSym

// LookupRuntimeABI looks up a name in package runtime using the given ABI.
func LookupRuntimeABI(name string, abi obj.ABI) *obj.LSym

// InitCoverage loads the definitions for routines called
// by code coverage instrumentation (similar to InitRuntime above).
func InitCoverage()

// LookupCoverage looks up the Go function 'name' in package
// runtime/coverage. This function must follow the internal calling
// convention.
func LookupCoverage(name string) *ir.Name
