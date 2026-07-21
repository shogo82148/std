// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package parser

import (
	"github.com/shogo82148/std/go/ast"
)

// ResolveFile applies the parser's deprecated
// [ast.Ident]-to-[ast.Object] resolution to the specified file's
// syntax tree. This is the same operation that is skipped on files
// parsed with the [SkipObjectResolution] mode flag.
//
// ResolveFile is idempotent and concurrency safe.
// Once ResolveFile returns, the file is resolved.
//
// Resolution error messages are discarded.
//
// Deprecated: [ast.Object] should not be used in new designs; use the
// [go/types] package instead. This function is provided to ease
// migration in applications that have disabled legacy object
// resolution by default but still need it in some circumstances.
func ResolveFile(file *ast.File)
