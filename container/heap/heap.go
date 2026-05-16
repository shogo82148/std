// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージheapは、heap.Interfaceを実装する任意の型に対して
// ヒープ操作を提供します。ヒープは、各ノードがその部分木における
// 最小値ノードであるという性質を持つ木です。
//
// 木の最小要素は、インデックス0のルートです。
//
// ヒープは優先度付きキューを実装する一般的な方法です。優先度付きキューを
// 構築するには、Lessメソッドの順序として（負の）優先度を用いて
// Heapインターフェースを実装します。これにより、Pushは要素を追加し、
// Popはキューから最優先の要素を取り出します。Examplesにはそのような
// 実装が含まれており、example_pq_test.goに完全なソースがあります。
package heap

import "github.com/shogo82148/std/sort"

// Interface型は、このパッケージのルーチンを使用する型に対する
// 要件を記述します。
// これを実装する任意の型は、次の不変条件を満たす
// 最小ヒープとして使用できます（[Init]が呼び出された後、
// またはデータが空もしくはソート済みの場合に成立します）。
//
//	!h.Less(j, i) for 0 <= i < h.Len() and 2*i+1 <= j <= 2*i+2 and j < h.Len()
//
// このインターフェース内の [Push] と [Pop] は、パッケージheapの
// 実装から呼び出すためのものである点に注意してください。
// ヒープへの追加と削除には、[heap.Push] と [heap.Pop] を使用してください。
type Interface interface {
	sort.Interface
	Push(x any)
	Pop() any
}

// Initは、このパッケージ内の他のルーチンが要求するヒープ不変条件を確立します。
// Initはヒープ不変条件に関して冪等であり、
// ヒープ不変条件が崩れた可能性があるときはいつでも呼び出せます。
// 計算量は O(n) です（n = h.Len()）。
func Init(h Interface)

// Pushは要素xをヒープに追加します。
// 計算量は O(log n) です（n = h.Len()）。
func Push(h Interface, x any)

// Popはヒープから最小要素（Lessに従う）を取り出して返します。
// 計算量は O(log n) です（n = h.Len()）。
// Popは [Remove](h, 0) と等価です。
func Pop(h Interface) any

// Removeはヒープのインデックスiにある要素を取り出して返します。
// 計算量は O(log n) です（n = h.Len()）。
func Remove(h Interface, i int) any

// Fixは、インデックスiの要素の値が変更された後にヒープ順序を再確立します。
// インデックスiの要素の値を変更してからFixを呼び出すことは、
// [Remove](h, i) を呼び出して新しい値をPushするのと等価ですが、
// それより低コストです。
// 計算量は O(log n) です（n = h.Len()）。
func Fix(h Interface, i int)
