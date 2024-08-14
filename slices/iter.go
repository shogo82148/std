// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slices

import (
	"github.com/shogo82148/std/cmp"
	"github.com/shogo82148/std/iter"
)

// Allは、スライス内のインデックスと値のペアを通常の順序で反復するイテレータを返します。
func All[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]

// Backwardは、スライス内のインデックスと値のペアを逆順で反復するイテレータを返します。
// インデックスを降順でたどります。
func Backward[Slice ~[]E, E any](s Slice) iter.Seq2[int, E]

// Valuesは、スライス要素を順番に生成するイテレータを返します。
func Values[Slice ~[]E, E any](s Slice) iter.Seq[E]

// AppendSeqは、seqからスライスに値を追加し、拡張されたスライスを返します。
func AppendSeq[Slice ~[]E, E any](s Slice, seq iter.Seq[E]) Slice

// Collectは、seqから値を収集して新しいスライスに格納し、それを返します。
func Collect[E any](seq iter.Seq[E]) []E

// Sortedは、seqから値を収集して新しいスライスに格納し、スライスをソートして返します。
func Sorted[E cmp.Ordered](seq iter.Seq[E]) []E

// SortedFuncは、seqから値を収集して新しいスライスに格納し、
// 比較関数を使用してスライスをソートし、それを返します。
func SortedFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E

// SortedStableFuncは、seqから値を収集して新しいスライスに格納します。
// 次に、比較関数を使用して要素を比較しながら、元の順序を保持してスライスをソートします。
// 新しいスライスを返します。
func SortedStableFunc[E any](seq iter.Seq[E], cmp func(E, E) int) []E

// Chunkは、sの最大n要素の連続する部分スライスを反復するイテレータを返します。
// 最後の部分スライス以外はすべてサイズnになります。
// すべての部分スライスは、長さを超える容量を持たないようにクリップされます。
// sが空の場合、シーケンスも空です：シーケンスに空のスライスはありません。
// nが1未満の場合、Chunkはパニックを起こします。
func Chunk[Slice ~[]E, E any](s Slice, n int) iter.Seq[Slice]
