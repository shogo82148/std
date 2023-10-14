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
func Exitf(format string, a ...interface{})

// Errorf logs an error message.
//
// If more than 20 errors have been printed, exit with an error.
//
// Logging an error means that on exit cmd/link will delete any
// output file and return a non-zero error code.
//
// TODO: remove. Use ctxt.Errorf instead.
// All remaining calls use nil as first arg.
func Errorf(dummy *int, format string, args ...interface{})

// Errorf method logs an error message.
//
// If more than 20 errors have been printed, exit with an error.
//
// Logging an error means that on exit cmd/link will delete any
// output file and return a non-zero error code.
func (ctxt *Link) Errorf(s loader.Sym, format string, args ...interface{})
