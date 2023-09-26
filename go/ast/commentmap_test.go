// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// To avoid a cyclic dependency with go/parser, this file is in a separate package.

package ast_test

import (
	. "go/ast"
)

// res maps a key of the form "line number: node type"
// to the associated comments' text.
//
