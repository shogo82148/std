// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

// MapはGoのmap[any]anyと似ていますが、追加のロックや調整なしで
// 複数のゴルーチンから同時に使用しても安全です。
// ロード、ストア、削除は平均定数時間で実行されます。
//
// Map型は特殊化されています。ほとんどのコードでは代わりに単純なGoのmapを使用すべきであり、
// 型の安全性を向上させ、他の不変条件とマップのコンテンツを維持しやすくするために別個のロックや調整を行うべきです。
//
// Map型は2つの一般的な使用例に最適化されています：(1)特定のキーのエントリーが一度しか書き込まれずに多くの回数読み取られる場合、
// つまり成長のみのキャッシュ、または(2)複数のゴルーチンが交差しないキーのセットの読み取り、書き込み、上書きを行う場合。
// これらの2つの場合において、Mapの使用は、別個の [Mutex] や [RWMutex] とペアになったGoのマップと比較して、
// ロックの競合を大幅に減らすことができます。
//
// ゼロ値のMapは空で使用準備ができています。Mapは最初の使用後にコピーされてはなりません。
//
// [the Go memory model] の用語では、Mapは書き込み操作が行われたときにそれに続く読み取り操作を "書き込みより前に同期します"。
// ここで、読み取りと書き込み操作は以下のように定義されます。
// [Map.Load]、[Map.LoadAndDelete]、[Map.LoadOrStore]、[Map.Swap]、[Map.CompareAndSwap]、[Map.CompareAndDelete] は読み取り操作です。
// [Map.Delete]、[Map.LoadAndDelete]、[Map.Store]、[Map.Swap] は書き込み操作です。
// [Map.LoadOrStore] は、loadedがfalseで返された場合に書き込み操作です。
// [Map.CompareAndSwap] は、swappedがtrueで返された場合に書き込み操作です。
// [Map.CompareAndDelete] は、deletedがtrueで返された場合に書き込み操作です。
//
// [the Go memory model]: https://go.dev/ref/mem
type Map struct {
	_ noCopy

	m isync.HashTrieMap[any, any]
}

// Loadは、キーに対応するマップ内の値を返します。もし値が存在しない場合はnilを返します。
// okの結果は、値がマップ内で見つかったかどうかを示します。
func (m *Map) Load(key any) (value any, ok bool)

// Storeはキーの値を設定します。
func (m *Map) Store(key, value any)

// Clearはすべてのエントリーを削除し、結果として空のMapになります。
func (m *Map) Clear()

// LoadOrStore は、キーが存在する場合は既存の値を返します。
// それ以外の場合は、指定された値を格納して返します。
// 読み込まれた結果が true の場合、値は読み込まれ、false の場合は格納されました。
func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)

// LoadAndDelete はキーに対応する値を削除し、もし値が存在する場合は以前の値を返します。
// 読み込まれた結果はキーが存在するかどうかを報告します。
func (m *Map) LoadAndDelete(key any) (value any, loaded bool)

// Deleteは指定されたキーの値を削除します。
func (m *Map) Delete(key any)

// Swapはキーの値を入れ替え、以前の値（あれば）を返します。
// 読み込まれた結果は、キーが存在するかどうかを報告します。
func (m *Map) Swap(key, value any) (previous any, loaded bool)

// CompareAndSwapは、キーの古い値と新しい値を交換します。
// マップに格納されている値が古い値と等しい場合にのみ行われます。
// 古い値は比較可能な型でなければなりません。
func (m *Map) CompareAndSwap(key, old, new any) (swapped bool)

// CompareAndDeleteは、keyの値がoldと等しい場合、そのエントリーを削除します。
// oldの値は比較可能な型である必要があります。
//
// マップにkeyの現在の値がない場合、CompareAndDeleteはfalseを返します（oldの値がnilインターフェース値であっても）。
func (m *Map) CompareAndDelete(key, old any) (deleted bool)

// Rangeはマップ内の各キーと値に対して順番にfを呼び出します。
// もしfがfalseを返す場合、Rangeは繰り返しを停止します。
//
// RangeはMapの内容に一貫したスナップショットに必ずしも対応していません：
// 各キーは複数回訪れませんが、任意のキーの値が同時に格納または削除される場合（fによっても含まれます）、
// RangeはRange呼び出し中の任意の時点からのそのキーのマッピングを反映することがあります。
// Rangeはレシーバー上の他のメソッドをブロックしません。f自体もmの任意のメソッドを呼び出すことができます。
//
// fが一定回数の呼び出し後にfalseを返す場合でも、Rangeはマップ内の要素数の数に比例するO(N)の可能性があります。
func (m *Map) Range(f func(key, value any) bool)
