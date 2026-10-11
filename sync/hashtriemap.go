// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.synchashtriemap

package sync

// Mapは、Goのmap[any]anyと似ていますが、追加のロックや調整なしに
// 複数のゴルーチンから並行して使用しても安全です。
// 読み込み、格納、削除は償却定数時間で実行されます。
//
// Map型は特殊化されています。ほとんどのコードは、型安全性を高め、
// マップの内容とともに他の不変条件を維持しやすくするために、
// 代わりに通常のGoのマップを個別のロックや調整とともに使用するべきです。
//
// Map型は2つの一般的なユースケースに最適化されています: (1) 特定のキーのエントリが
// 一度だけ書き込まれ、何度も読み取られる場合（例えば、増加するだけのキャッシュ）、
// または (2) 複数のゴルーチンが互いに素なキーの集合のエントリを読み取り、書き込み、
// 上書きする場合です。これら2つのケースでは、Mapを使用すると、個別の[Mutex]や[RWMutex]と
// 組み合わせたGoのマップと比べて、ロックの競合を大幅に減らすことができます。
//
// ゼロ値のMapは空で、すぐに使用できます。Mapは最初の使用後にコピーしてはいけません。
//
// [the Go memory model]の用語では、Mapは、書き込み操作が、その書き込みの効果を観測する
// 任意の読み取り操作よりも「前に同期する」ように調整します。ここで、読み取り操作と
// 書き込み操作は次のように定義されます。
// [Map.Load]、[Map.LoadAndDelete]、[Map.LoadOrStore]、[Map.Swap]、[Map.CompareAndSwap]、
// [Map.CompareAndDelete]は読み取り操作です。
// [Map.Delete]、[Map.LoadAndDelete]、[Map.Store]、[Map.Swap]は書き込み操作です。
// [Map.LoadOrStore]は、loadedがfalseに設定されて返される場合に書き込み操作です。
// [Map.CompareAndSwap]は、swappedがtrueに設定されて返される場合に書き込み操作です。
// [Map.CompareAndDelete]は、deletedがtrueに設定されて返される場合に書き込み操作です。
//
// [the Go memory model]: https://go.dev/ref/mem
type Map struct {
	_ noCopy

	m isync.HashTrieMap[any, any]
}

// Loadは、キーに対してマップに格納されている値を返します。値が存在しない場合はnilを返します。
// 結果のokは、値がマップ内で見つかったかどうかを示します。
func (m *Map) Load(key any) (value any, ok bool)

// Storeはキーに対する値を設定します。
func (m *Map) Store(key, value any)

// Clearはすべてのエントリを削除し、空のMapにします。
func (m *Map) Clear()

// LoadOrStoreは、キーが存在する場合はその既存の値を返します。
// そうでない場合は、指定された値を格納して返します。
// 結果のloadedは、値が読み込まれた場合はtrue、格納された場合はfalseです。
func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)

// LoadAndDeleteはキーに対する値を削除し、以前の値があればそれを返します。
// 結果のloadedは、キーが存在したかどうかを報告します。
func (m *Map) LoadAndDelete(key any) (value any, loaded bool)

// Deleteはキーに対する値を削除します。
// キーがマップに存在しない場合、Deleteは何もしません。
func (m *Map) Delete(key any)

// Swapはキーに対する値を交換し、以前の値があればそれを返します。
// 結果のloadedは、キーが存在したかどうかを報告します。
func (m *Map) Swap(key, value any) (previous any, loaded bool)

// CompareAndSwapは、マップに格納されている値がoldと等しい場合に、
// keyの古い値と新しい値を交換します。
// 古い値は比較可能な型でなければなりません。
func (m *Map) CompareAndSwap(key, old, new any) (swapped bool)

// CompareAndDeleteは、keyの値がoldと等しい場合にそのエントリを削除します。
// 古い値は比較可能な型でなければなりません。
//
// マップにkeyの現在の値がない場合、CompareAndDeleteは
// （古い値がnilインターフェース値であっても）falseを返します。
func (m *Map) CompareAndDelete(key, old any) (deleted bool)

// Rangeは、マップに存在する各キーと値に対して順番にfを呼び出します。
// fがfalseを返すと、rangeは反復を停止します。
//
// Rangeは、必ずしもMapの内容の一貫したスナップショットに対応するわけではありません。
// どのキーも2回以上訪問されることはありませんが、いずれかのキーの値が
// （fによるものも含めて）並行して格納または削除された場合、Rangeは
// Rangeの呼び出し中の任意の時点のそのキーのマッピングを反映する可能性があります。
// Rangeはレシーバの他のメソッドをブロックしません。f自体がmの任意のメソッドを呼び出すこともできます。
//
// fが一定回数の呼び出し後にfalseを返したとしても、Rangeはマップ内の要素数に対して
// O(N)になる可能性があります。
func (m *Map) Range(f func(key, value any) bool)
