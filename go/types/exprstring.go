// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは式の出力を実装しています。

package types

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/go/ast"
)

// ExprStringはxの（短縮された可能性のある）文字列表現を返します。
// 短縮された表現はユーザーインターフェースに適していますが、Goの構文に必ずしも従っているわけではありません。
func ExprString(x ast.Expr) string

// WriteExprは、xの（短縮されたかもしれない）文字列表現をbufに書き込みます。
// 短縮表示はユーザーインターフェースに適していますが、必ずしもGoの構文に従うとは限りません。
func WriteExpr(buf *bytes.Buffer, x ast.Expr)
