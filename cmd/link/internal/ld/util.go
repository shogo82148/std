// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ld

import (
	"github.com/shogo82148/std/cmd/link/internal/loader"
)

func AtExit(f func())

// Exit exits with code after executing all atExitFuncs.
func Exit(code int)

// Exitf logs an error message then calls Exit(2).
func Exitf(format string, a ...any)

// Errorf logs an error message without a specific symbol for context.
// Use ctxt.Errorf when possible.
//
// If more than 20 errors have been printed, exit with an error.
//
// Logging an error means that on exit cmd/link will delete any
// output file and return a non-zero error code.
func Errorf(format string, args ...any)

// Errorf method logs an error message.
//
// If more than 20 errors have been printed, exit with an error.
//
// Logging an error means that on exit cmd/link will delete any
// output file and return a non-zero error code.
func (ctxt *Link) Errorf(s loader.Sym, format string, args ...any)
