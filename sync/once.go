// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

<<<<<<< HEAD
import (
	"github.com/shogo82148/std/sync/atomic"
)

// Once is an object that will perform exactly one action.
=======
// Onceは正確に1つのアクションを実行するオブジェクトです。
>>>>>>> release-branch.go1.21
//
// Onceは最初の使用後にコピーしてはいけません。
//
// Goメモリモデルの用語では、fからの戻り値はonce.Do(f)の呼び出しの戻り値よりも前に同期します。
type Once struct {
<<<<<<< HEAD
	// done indicates whether the action has been performed.
	// It is first in the struct because it is used in the hot path.
	// The hot path is inlined at every call site.
	// Placing done first allows more compact instructions on some architectures (amd64/386),
	// and fewer instructions (to calculate offset) on other architectures.
	done atomic.Uint32
=======
	// doneはアクションが実行されたかどうかを示します。
	// これはホットパスで使用されるため、structの最初に配置されています。
	// ホットパスは各呼び出し箇所にインライン展開されます。
	// doneを最初に配置することで、一部のアーキテクチャ（amd64/386）ではよりコンパクトな命令が可能になり、
	// 他のアーキテクチャではオフセットを計算するための命令が少なくなります。
	done uint32
>>>>>>> release-branch.go1.21
	m    Mutex
}

// Doは、Onceのインスタンスで最初にDoが呼び出された場合のみ、関数fを呼び出します。つまり、次のように与えられた場合、
//
//	var once Once
//
// もしこのように複数回once.Do(f)が呼び出された場合、最初の呼び出しのみがfを実行し、
// それぞれの呼び出しでfが異なる値を持っていても、その値に関係なく実行されます。各関数の実行には新しいOnceのインスタンスが必要です。
//
// Doは、一度だけ実行する必要のある初期化に使用されます。fは引数なしの関数ですので、Doによって呼び出される関数の引数を捕捉するために関数リテラルを使用する必要があるかもしれません：
//
//	config.once.Do(func() { config.init(filename) })
//
// Doへの呼び出しがfの返り値が返されるまで戻らないため、fがDoの呼び出しを引き起こすとデッドロックが発生します。
//
// fがパニックした場合、Doはそれを戻ったとみなします。その後のDoの呼び出しはfを呼び出さずに返ります。
func (o *Once) Do(f func())
