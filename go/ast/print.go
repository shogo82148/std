// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルにはASTの出力サポートが含まれています。

package ast

import (
	"github.com/shogo82148/std/go/token"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/reflect"
)

// 出力を制御するために、FprintにFieldFilterを指定することができます。
type FieldFilter func(name string, value reflect.Value) bool

// NotNilFilterは、nilでないフィールド値に対してtrueを返します。
// それ以外の場合はfalseを返します。
func NotNilFilter(_ string, v reflect.Value) bool

// FprintはASTノードxから始まる(サブ)ツリーをwに出力します。
// もしfset != nilなら、位置情報はそのファイルセットに対して相対的に解釈されます。
// それ以外の場合は位置は整数値（ファイルセット固有のオフセット）として表示されます。
//
// 非nilのFieldFilter fが提供された場合、出力を制御するために使用されます：
// f(fieldname, fieldvalue)がtrueを返す構造体フィールドだけが出力されます。
// それ以外のものは出力からフィルタリングされます。エクスポートされていない構造体フィールドは常に出力されません。
func Fprint(w io.Writer, fset *token.FileSet, x any, f FieldFilter) error

// Print関数は、nilのフィールドをスキップしてxを標準出力に出力します。
// Print(fset, x)は、Fprint(os.Stdout, fset, x, NotNilFilter)と同じです。
func Print(fset *token.FileSet, x any) error
