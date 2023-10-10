// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

// Descriptionはランタイムメトリクスを説明します。
type Description struct {

	// Nameは単位を含むメトリクスのフルネームです。
	//
	// メトリクスの形式は以下の正規表現で表されます。
	//
	// 	^(?P<name>/[^:]+):(?P<unit>[^:*/]+(?:[*/][^:*/]+)*)$
	//
	// この形式は、名前をコロンで区切られた2つのコンポーネントに分割します。1つは常に/で始まるパスであり、もう1つは機械で解釈可能な単位です。名前は/文字の間に任意の有効なUnicodeコードポイントを含むことができますが、慣例として小文字の英字とハイフンを使用することが推奨されています。このようなパスの例としては、"/memory/heap/free"があります。
	//
	// 単位は、*または/で区切られた接頭辞のない小文字の英語の単位名（単数形または複数形）の連続です。単位名には、区切り文字でない有効なUnicodeコードポイントを含めることができます。単位の例には、"seconds"、"bytes"、"bytes/second"、"cpu-seconds"、"byte*cpu-seconds"、および"bytes/second/second"があります。
	//
	// ヒストグラムの場合、複数の単位が適用される場合があります。たとえば、バケットの単位とカウントの単位です。ヒストグラムの場合、カウントの単位は常に"samples"であり、サンプルのタイプはメトリクスの名前で明示されますが、名前の単位はバケットの単位を指定します。
	//
	// 完全な名前の例としては、"/memory/heap/free:bytes"があります。
	Name string

	// Descriptionはメトリックを説明する英文です。
	Description string

	// Kindはこのメトリックの値の種類です。
	//
	// このフィールドの目的は、アプリケーションが理解できない型の値を持つメトリックをフィルタリングできるようにすることです。
	Kind ValueKind

	// Cumulativeはメトリクスが累積的かどうかを示します。累積メトリクスが単一の数値である場合、それは単調に増加します。メトリクスが分布である場合、各バケットのカウントが単調に増加します。
	//
	// したがって、このフラグはこの値からレートを計算することが有用かどうかを示します。
	Cumulative bool
}

// Allは、サポートされているすべてのメトリックの説明を含むスライスを返します。
func All() []Description
