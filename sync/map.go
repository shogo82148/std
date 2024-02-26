// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sync

import (
	"github.com/shogo82148/std/sync/atomic"
)

// MapはGoのmap[any]anyと似ていますが、追加のロックや調整なしで
// 複数のゴルーチンから同時に使用しても安全です。
// ロード、ストア、削除は平均定数時間で実行されます。
//
// Map型は特殊化されています。ほとんどのコードでは代わりに単純なGoのmapを使用すべきであり、
// 型の安全性を向上させ、他の不変条件とマップのコンテンツを維持しやすくするために別個のロックや調整を行うべきです。
//
// Map型は2つの一般的な使用例に最適化されています：(1)特定のキーのエントリが一度しか書き込まれずに多くの回数読み取られる場合、
// つまり成長のみのキャッシュ、または(2)複数のゴルーチンが交差しないキーのセットの読み取り、書き込み、上書きを行う場合。
// これらの2つの場合において、Mapの使用は、別個のMutexやRWMutexとペアになったGoのマップと比較して、
// ロックの競合を大幅に減らすことができます。
//
// ゼロ値のMapは空で使用準備ができています。Mapは最初の使用後にコピーされてはなりません。
//
// Goメモリモデルの用語では、Mapは書き込み操作が行われたときにそれに続く読み取り操作を "書き込みより前に同期します"。
// ここで、読み取りと書き込み操作は以下のように定義されます。
// Load、LoadAndDelete、LoadOrStore、Swap、CompareAndSwap、CompareAndDeleteは読み取り操作です。
// Delete、LoadAndDelete、Store、Swapは書き込み操作です。
// LoadOrStoreは、loadedがfalseで返された場合に書き込み操作です。
// CompareAndSwapは、swappedがtrueで返された場合に書き込み操作です。
// CompareAndDeleteは、deletedがtrueで返された場合に書き込み操作です。
type Map struct {
	mu Mutex

	// readには、muを保持している場合とそうでない場合の両方で、
	// 同時アクセスが安全なマップの内容の一部が含まれています。
	//
	// readフィールド自体は常に安全に読み込むことができますが、
	// muを保持しているときにのみ保存することができます。
	//
	// readに保存されているエントリは、muなしで同時に更新することができますが、
	// 以前に削除されたエントリの更新には、そのエントリをdirtyマップにコピーし、
	// muを保持してunexpungedする必要があります。
	read atomic.Pointer[readOnly]

	// dirtyには、muの保持が必要なマップの内容の一部が含まれています。dirtyマップがすばやく読み取り用マップに昇格できるようにするため、読み取り用マップの非削除エントリもすべて含まれています。
	// 削除されたエントリはdirtyマップに保存されません。クリーンマップの削除されたエントリは、新しい値を格納する前に未削除にされてdirtyマップに追加する必要があります。
	// dirtyマップがnilの場合、マップへの次の書き込みでは、クリーンマップのステールエントリを省略して浅いコピーを作成して初期化します。
	dirty map[any]*entry

	// misses は、読み込みマップが最後に更新されて以降、キーが存在するかどうかを確認するために mu ロックが必要な読み込み回数を数えます。
	//
	// 十分な数のミスが発生し、汚れたマップのコピーのコストをカバーできるようになると、汚れたマップは読み込みマップに昇格します（修正されていない状態で）。
	misses int
}

// Loadは、キーに対応するマップ内の値を返します。もし値が存在しない場合はnilを返します。
// okの結果は、値がマップ内で見つかったかどうかを示します。
func (m *Map) Load(key any) (value any, ok bool)

// Storeはキーの値を設定します。
func (m *Map) Store(key, value any)

// Clearはすべてのエントリを削除し、結果として空のMapになります。
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
func (m *Map) CompareAndSwap(key, old, new any) bool

// CompareAndDeleteは、keyの値がoldと等しい場合、そのエントリを削除します。
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
