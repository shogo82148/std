// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはタイプの出力を実装しています。

package types

import (
	"github.com/shogo82148/std/bytes"
)

// Qualifierは、[TypeString]、[ObjectString]、および [SelectionString] の呼び出しで名前付きのパッケージレベルのオブジェクトが出力される方法を制御します。
//
// これらの3つのフォーマットルーチンは、Qualifierを各パッケージレベルのオブジェクトOに対して呼び出し、Qualifierが空ではない文字列pを返すと、オブジェクトはp.Oの形式で表示されます。
// フォーマットルーチンが空文字列を返す場合、オブジェクト名Oのみが表示されます。
//
// nilのQualifierを使用することは、(*[Package]).Pathを使用することと同じです：オブジェクトはインポートパスで修飾されます。例えば、"encoding/json.Marshal"となります。
type Qualifier func(*Package) string

// RelativeToは、pkg以外のすべてのパッケージのメンバーを完全に修飾する [Qualifier] を返します。
func RelativeTo(pkg *Package) Qualifier

// TypeStringはtypの文字列表現を返します。
// [Qualifier] はパッケージレベルのオブジェクトの印刷を制御し、nilである可能性があります。
func TypeString(typ Type, qf Qualifier) string

// WriteTypeはtypの文字列表現をbufに書き込みます。
// [Qualifier] はpackageレベルのオブジェクトの表示を制御し、nilであることもあります。
func WriteType(buf *bytes.Buffer, typ Type, qf Qualifier)

// WriteSignatureは、関数キーワードを付けずにシグネチャsigの表現をbufに書き込みます。
// [Qualifier] はパッケージレベルのオブジェクトの印刷を制御し、nilである可能性があります。
func WriteSignature(buf *bytes.Buffer, sig *Signature, qf Qualifier)
