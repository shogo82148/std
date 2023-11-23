// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

// CommonTypeはすべての型の要素を保持します。
// これは歴史的な遺物であり、バイナリ互換性を保つため、および型記述子のエンコーディングのために
// パッケージが利用するためだけに保持されています。クライアントによる直接的な使用は意図されていません。
type CommonType struct {
	Name string
	Id   typeId
}

// GobEncoderは、GobDecoderに送信するための値のエンコーディング表現を提供するデータを
// 描写するインターフェースです。GobEncoderとGobDecoderを実装する型は、そのデータの表現に
// 完全な制御を持つため、通常はgobストリームで送信できないプライベートフィールド、チャネル、
// 関数などを含むことができます。
//
// 注意: gobsは永続的に保存できるため、ソフトウェアが進化するにつれてGobEncoderによって
// 使用されるエンコーディングが安定していることを保証することは良い設計です。例えば、GobEncodeが
// エンコーディングにバージョン番号を含めることは理にかなっているかもしれません。
type GobEncoder interface {
	// GobEncodeは、通常は同じ具体的な型のGobDecoderに送信するための
	// 受信者のエンコーディングを表すバイトスライスを返します。
	GobEncode() ([]byte, error)
}

// GobDecoderは、GobEncoderによって送信された値のデコーディングルーチンを提供するデータを
// 描写するインターフェースです。
type GobDecoder interface {
	// GobDecodeは、受信者（ポインタでなければならない）を、
	// バイトスライスによって表される値で上書きします。このバイトスライスは、
	// 通常は同じ具体的な型のためにGobEncodeによって書き込まれます。
	GobDecode([]byte) error
}

// RegisterNameはRegisterと同様ですが、型のデフォルトではなく提供された名前を使用します。
func RegisterName(name string, value any)

// Registerは、その型の値によって識別される型を、
// 内部型名の下に記録します。その名前は、インターフェース変数として送受信される値の
// 具体的な型を識別します。インターフェース値の実装として転送される型のみを登録する必要があります。
// 初期化時にのみ使用されることを期待しており、型と名前の間のマッピングが全単射でない場合はパニックを引き起こします。
func Register(value any)
