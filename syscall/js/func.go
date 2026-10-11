// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package js

import (
	"github.com/shogo82148/std/internal/synctest"
)

// Funcは、JavaScriptから呼び出されるようにラップされたGoの関数です。
type Func struct {
	Value
	bubble *synctest.Bubble
	id     uint32
}

// FuncOfは、JavaScriptで使用される関数を返します。
//
// Goの関数fnは、JavaScriptの"this"キーワードの値と呼び出しの引数で呼び出されます。
// 呼び出しの戻り値は、ValueOfに従ってJavaScriptにマッピングし直された
// Goの関数の結果です。
//
// JavaScriptからラップされたGoの関数を呼び出すと、イベントループが一時停止し、
// 新しいゴルーチンが生成されます。
// GoからJavaScriptへの呼び出し中にトリガーされた他のラップされた関数は、
// 同じゴルーチン上で実行されます。
//
// その結果、ラップされた関数の1つがブロックすると、JavaScriptのイベントループは
// その関数が戻るまでブロックされます。したがって、fetch（http.Client）のような
// イベントループを必要とする非同期のJavaScript APIを呼び出すと、即座にデッドロックが
// 発生します。そのため、ブロックする関数は明示的に新しいゴルーチンを開始するべきです。
//
// 関数がもう呼び出されなくなったら、リソースを解放するためにFunc.Releaseを呼び出す必要があります。
func FuncOf(fn func(this Value, args []Value) any) Func

// Releaseは、関数のために割り当てられたリソースを解放します。
// Releaseを呼び出した後に関数を呼び出してはいけません。
// 関数がまだ実行中にReleaseを呼び出すことは許可されています。
func (c Func) Release()
