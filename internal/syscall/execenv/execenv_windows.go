// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build windows

package execenv

import (
	"github.com/shogo82148/std/syscall"
)

// Default will return the default environment
// variables based on the process attributes
// provided.
//
// If the process attributes contain a token, then
// the environment variables will be sourced from
// the defaults for that user token, otherwise they
// will be sourced from syscall.Environ().
func Default(sys *syscall.SysProcAttr) (env []string, err error)
