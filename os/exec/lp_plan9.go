// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package exec

import (
	"github.com/shogo82148/std/errors"
)

// ErrNotFound is the error resulting if a path search failed to find an executable file.
var ErrNotFound = errors.New("executable file not found in $path")

// LookPath searches for an executable named file in the
// directories named by the path environment variable.
// If file begins with "/", "#", "./", or "../", it is tried
// directly and the path is not consulted.
// On success, the result is an absolute path.
//
// In older versions of Go, LookPath could return a path relative to the current directory.
// As of Go 1.19, LookPath will instead return that path along with an error satisfying
// [errors.Is](err, [ErrDot]). See the package documentation for more details.
func LookPath(file string) (string, error)
