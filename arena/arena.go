// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.arenas

/*
The arena package provides the ability to allocate memory for a collection
of Go values and free that space manually all at once, safely. The purpose
of this functionality is to improve efficiency: manually freeing memory
before a garbage collection delays that cycle. Less frequent cycles means
the CPU cost of the garbage collector is incurred less frequently.

This functionality in this package is mostly captured in the Arena type.
Arenas allocate large chunks of memory for Go values, so they're likely to
be inefficient for allocating only small amounts of small Go values. They're
best used in bulk, on the order of MiB of memory allocated on each use.

Note that by allowing for this limited form of manual memory allocation
that use-after-free bugs are possible with regular Go values. This package
limits the impact of these use-after-free bugs by preventing reuse of freed
memory regions until the garbage collector is able to determine that it is
safe. Typically, a use-after-free bug will result in a fault and a helpful
error message, but this package reserves the right to not force a fault on
freed memory. That means a valid implementation of this package is to just
allocate all memory the way the runtime normally would, and in fact, it
reserves the right to occasionally do so for some Go values.
*/
package arena

import (
	"github.com/shogo82148/std/unsafe"
)

// Arenaは、まとめて割り当てられ、まとめて解放されるGoの値のコレクションを表します。
// Arenaは手動でランタイムに解放できるため、効率の向上に役立ちますが、
// 解放されたArenaから取得したメモリには、解放後にアクセスしてはいけません。
// Arenaは参照されなくなると自動的に解放されるため、そこから割り当てられたメモリが
// 不要になるまで生存させておく必要があります（runtime.KeepAliveを参照）。
//
// Arenaを複数のゴルーチンから並行して使用してはいけません。
type Arena struct {
	a unsafe.Pointer
}

// NewArenaは新しいArenaを割り当てます。
func NewArena() *Arena

// FreeはArena（およびArenaから割り当てられたすべてのオブジェクト）を解放し、
// Arenaを支えるメモリをガベージコレクションのオーバーヘッドなしにかなり速やかに
// 再利用できるようにします。アプリケーションは、解放後にこのArenaのメソッドを
// 呼び出してはいけません。
func (a *Arena) Free()

// Newは、指定されたArenaに新しい*Tを作成します。*TはArenaが解放された後に
// 使用してはいけません。解放後に値にアクセスするとフォールトが発生する可能性がありますが、
// このフォールトも保証されていません。
func New[T any](a *Arena) *T

// MakeSliceは、指定された容量と長さを持つ新しい[]Tを作成します。[]TはArenaが
// 解放された後に使用してはいけません。解放後にスライスの基礎となるストレージに
// アクセスするとフォールトが発生する可能性がありますが、このフォールトも保証されていません。
func MakeSlice[T any](a *Arena, len, cap int) []T

// Cloneは、入力値の浅いコピーを作成し、そのコピーを返します。コピーは、
// 元の値が割り当てられたArenaにはもはや束縛されません。元の値がArenaから
// 割り当てられていない場合は、そのまま返されます。この関数は、Arenaで割り当てられた値を
// そのArenaよりも長く生存させるのに役立ちます。
// Tはポインタ、スライス、文字列のいずれかでなければならず、そうでない場合この関数はパニックします。
func Clone[T any](s T) T
