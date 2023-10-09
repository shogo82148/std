// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ast

import (
	"github.com/shogo82148/std/go/token"
)

// SortImportsはfのimportブロック内の連続したimport行をソートします。
// データの損失なしに重複するimportを削除することも可能です。
func SortImports(fset *token.FileSet, f *File)
