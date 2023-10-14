// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package metrics

// Float64Histogramはfloat64値の分布を表します。
type Float64Histogram struct {

	// Countsには、各ヒストグラムバケットの重みが格納されています。
	//
	// バケット数Nが与えられた場合、Count[n]は範囲[bucket[n]、bucket[n+1])の重みです。
	// ただし、0 <= n < N。
	Counts []uint64

	// Buckets はヒストグラムのバケットの境界を含み、増加する順序で格納します。
	//
	// Buckets[0] は最小バケットの包括的な下限であり、
	// Buckets[len(Buckets)-1] は最大バケットの排他的な上限です。
	// したがって、len(Buckets)-1 個のカウントがあります。さらに、len(Buckets) != 1 は常に成り立ちます。
	// なぜならば、少なくとも 2 つの境界が必要で、0 個の境界を使用して 0 個のバケットを記述するためです。
	//
	// Buckets[0] には値 -Inf を許可し、Buckets[len(Buckets)-1] には値 Inf を許可します。
	//
	// 特定のメトリック名に対して、Buckets の値はプログラムの終了まで呼び出しの間変わらないことが保証されています。
	//
	// このスライスの値は、他の Float64Histograms の Buckets フィールドとエイリアスを持つことが許可されているため、
	// 内部の値は読み取ることしかできません。変更する必要がある場合は、ユーザーがコピーを作成する必要があります。
	Buckets []float64
}
