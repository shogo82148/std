// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package encodingは、データをバイトレベルやテキスト表現に変換する他のパッケージで共有されるインターフェースを定義します。
// これらのインターフェースをチェックするパッケージにはencoding/gob、encoding/json、encoding/xmlなどがあります。
// そのため、一度インターフェースを実装することで、1つの型が複数のエンコーディングで有用になることがあります。
// これらのインターフェースを実装する標準の型にはtime.Timeやnet.IPがあります。
// これらのインターフェースは、エンコードデータの生成と消費を行うペアとして提供されます。
// 既存の型にエンコード/デコードのメソッドを追加することは、破壊的な変更となる可能性があるため、注意が必要です。
// なぜなら、これらのメソッドは異なるライブラリバージョンで書かれたプログラムとの通信において
// シリアライズに使用されるからです。
// Goプロジェクトによって管理されるパッケージのポリシーは、既存の適切なマーシャリングが存在しない場合にのみ、
// マーシャリング関数の追加を許可することです。
package encoding

// BinaryMarshalerは、自身をバイナリ形式に変換できるオブジェクトによって実装されるインターフェースです。
//
// MarshalBinaryは、レシーバをバイナリ形式にエンコードし、その結果を返します。
type BinaryMarshaler interface {
	MarshalBinary() (data []byte, err error)
}

// BinaryUnmarshalerは、自身のバイナリ表現をアンマーシャルできるオブジェクトによって実装されるインターフェースです。
//
// UnmarshalBinaryは、MarshalBinaryによって生成された形式をデコードできる必要があります。
// UnmarshalBinaryは、データを保持したい場合はデータをコピーする必要があります。
//処理を終えた後のデータを残したい場合は、データをコピーする必要があります。
type BinaryUnmarshaler interface {
	UnmarshalBinary(data []byte) error
}

// TextMarshalerは、自身をテキスト形式にマーシャリングできるオブジェクトによって実装されるインターフェースです。
//
// MarshalTextは、レシーバをUTF-8でエンコードされたテキストに変換し、結果を返します。
type TextMarshaler interface {
	MarshalText() (text []byte, err error)
}

// TextUnmarshalerは、自身のテキスト表現をUnmarshalできるオブジェクトが実装するインターフェースです。
//
// UnmarshalTextは、MarshalTextによって生成された形式をデコードできる必要があります。
// UnmarshalTextは、戻り値の後にテキストを保持する場合は、テキストをコピーする必要があります。
type TextUnmarshaler interface {
	UnmarshalText(text []byte) error
}
