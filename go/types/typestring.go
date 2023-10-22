// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはタイプの出力を実装しています。

package types

import (
	"github.com/shogo82148/std/bytes"
)

<<<<<<< HEAD
// Qualifierは、TypeString、ObjectString、およびSelectionStringの呼び出しで名前付きのパッケージレベルのオブジェクトが出力される方法を制御します。
=======
// A Qualifier controls how named package-level objects are printed in
// calls to [TypeString], [ObjectString], and [SelectionString].
>>>>>>> upstream/master
//
// これらの3つのフォーマットルーチンは、Qualifierを各パッケージレベルのオブジェクトOに対して呼び出し、Qualifierが空ではない文字列pを返すと、オブジェクトはp.Oの形式で表示されます。
// フォーマットルーチンが空文字列を返す場合、オブジェクト名Oのみが表示されます。
//
<<<<<<< HEAD
// nilのQualifierを使用することは、(*Package).Pathを使用することと同じです：オブジェクトはインポートパスで修飾されます。例えば、"encoding/json.Marshal"となります。
type Qualifier func(*Package) string

// RelativeToは、pkg以外のすべてのパッケージのメンバーを完全に修飾するQualifierを返します。
func RelativeTo(pkg *Package) Qualifier

// TypeStringはtypの文字列表現を返します。
// Qualifierはパッケージレベルのオブジェクトの印刷を制御し、nilである可能性があります。
func TypeString(typ Type, qf Qualifier) string

// WriteTypeはtypの文字列表現をbufに書き込みます。
// Qualifierはpackageレベルのオブジェクトの表示を制御し、nilであることもあります。
func WriteType(buf *bytes.Buffer, typ Type, qf Qualifier)

// WriteSignatureは、関数キーワードを付けずにシグネチャsigの表現をbufに書き込みます。
// Qualifierはパッケージレベルのオブジェクトの印刷を制御し、nilである可能性があります。
=======
// Using a nil Qualifier is equivalent to using (*[Package]).Path: the
// object is qualified by the import path, e.g., "encoding/json.Marshal".
type Qualifier func(*Package) string

// RelativeTo returns a [Qualifier] that fully qualifies members of
// all packages other than pkg.
func RelativeTo(pkg *Package) Qualifier

// TypeString returns the string representation of typ.
// The [Qualifier] controls the printing of
// package-level objects, and may be nil.
func TypeString(typ Type, qf Qualifier) string

// WriteType writes the string representation of typ to buf.
// The [Qualifier] controls the printing of
// package-level objects, and may be nil.
func WriteType(buf *bytes.Buffer, typ Type, qf Qualifier)

// WriteSignature writes the representation of the signature sig to buf,
// without a leading "func" keyword. The [Qualifier] controls the printing
// of package-level objects, and may be nil.
>>>>>>> upstream/master
func WriteSignature(buf *bytes.Buffer, sig *Signature, qf Qualifier)
