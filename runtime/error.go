// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Errorはpanicで使用されるランタイムエラーを識別します。
//
// Goランタイムは、Go言語仕様で説明されている様々なケース、例えばスライスや配列の範囲外アクセス、nilチャネルのクローズ、型アサーションの失敗などでpanicを発生させます。
//
// これらのケースが発生した場合、GoランタイムはErrorを実装するエラーでpanicします。これは、panicからのリカバリー時に、カスタムアプリケーションのpanicと基本的なランタイムpanicを区別するのに役立ちます。
//
// Go標準ライブラリ以外のパッケージでErrorを実装すべきではありません。
type Error interface {
	error

	RuntimeError()
}

// TypeAssertionErrorは、型アサーションの失敗を説明します。
type TypeAssertionError struct {
	_interface    *_type
	concrete      *_type
	asserted      *_type
	missingMethod string
}

func (*TypeAssertionError) RuntimeError()

func (e *TypeAssertionError) Error() string
