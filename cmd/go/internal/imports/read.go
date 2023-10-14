// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copied from Go distribution src/go/build/read.go.

package imports

import (
	"github.com/shogo82148/std/io"
)

// ReadComments is like io.ReadAll, except that it only reads the leading
// block of comments in the file.
func ReadComments(f io.Reader) ([]byte, error)

// ReadImports is like io.ReadAll, except that it expects a Go file as input
// and stops reading the input once the imports have completed.
func ReadImports(f io.Reader, reportSyntaxError bool, imports *[]string) ([]byte, error)
