// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unique

// Handleは型Tの値に対するグローバルに一意な識別子です。
//
// 2つのハンドルが等しいと比較されるのは、それらのハンドルを生成した2つの値が
// 等しいと比較される場合に限ります。ハンドル同士の比較は単純で、
// 通常、元の値同士を比較するよりもはるかに効率的です。
type Handle[T comparable] struct {
	value *T
}

// Valueは、Handleを生成したT値の浅いコピーを返します。
// Valueは複数のゴルーチンから同時に安全に使用できます。
func (h Handle[T]) Value() T

// Makeは型Tの値に対するグローバルに一意なハンドルを返します。
// ハンドル同士が等しいのは、それらを生成した値が等しい場合のみです。
// Makeは複数のゴルーチンから同時に安全に使用できます。
func Make[T comparable](value T) Handle[T]
