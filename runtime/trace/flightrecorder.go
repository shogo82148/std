// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/time"
)

// FlightRecorderは、Go実行トレースの単一のコンシューマを表します。
// ランタイムによって生成される実行トレースの移動ウィンドウを追跡し、
// 常に最新のトレースデータを含みます。
//
// 一度にアクティブにできるフライトレコーダーは最大1つですが、
// フライトレコーディングはtrace.Startを使用するトレースコンシューマと
// 同時にアクティブにすることができます。
// フライトレコーダーが1つだけというこの制限は、将来削除される可能性があります。
type FlightRecorder struct {
	err error

	// State specific to the recorder.
	header [16]byte
	active rawGeneration
	ringMu sync.Mutex
	ring   []rawGeneration
	freq   frequency

	// Externally-set options.
	targetSize   uint64
	targetPeriod time.Duration

	enabled bool
	writing sync.Mutex

	// The values of targetSize and targetPeriod we've committed to since the last Start.
	wantSize uint64
	wantDur  time.Duration
}

// NewFlightRecorderは、指定された設定から新しいフライトレコーダーを作成します。
func NewFlightRecorder(cfg FlightRecorderConfig) *FlightRecorder

// Startはフライトレコーダーをアクティブにし、トレースデータの記録を開始します。
// 一度にアクティブにできるtrace.Startの呼び出しは1つだけです。
// さらに、現在プログラム内でアクティブにできるフライトレコーダーは1つだけです。
// フライトレコーダーを開始できない場合、またはすでに開始されている場合はエラーを返します。
func (fr *FlightRecorder) Start() error

// Stopはトレースデータの記録を終了します。並行するWriteTo呼び出しが完了するまでブロックします。
func (fr *FlightRecorder) Stop()

// Enabledは、フライトレコーダーがアクティブな場合にtrueを返します。
// 具体的には、Startがエラーを返さず、Stopがまだ呼び出されていない場合にtrueを返します。
// 複数のゴルーチンから同時に呼び出しても安全です。
func (fr *FlightRecorder) Enabled() bool

// WriteToは、フライトレコーダーが追跡している移動ウィンドウのスナップショットを取得します。
// スナップショットには、WriteToが呼び出された時点で最新のデータが含まれることが期待されますが、
// これは厳密な保証ではありません。
// 一度にWriteToを実行できるゴルーチンは1つだけです。
// wへの書き込みに失敗した場合、別のWriteTo呼び出しがすでに進行中の場合、
// またはフライトレコーダーが非アクティブな場合はエラーが返されます。
func (fr *FlightRecorder) WriteTo(w io.Writer) (n int64, err error)

type FlightRecorderConfig struct {
	// MinAgeは、フライトレコーダーのウィンドウ内のイベントの経過時間の下限です。
	//
	// フライトレコーダーは最小経過時間より古いイベントを速やかに破棄するよう努めますが、
	// 古いイベントがウィンドウのスナップショットに現れる場合があります。経過時間の設定は
	// 常にMaxBytesによって上書きされます。
	//
	// これが0の場合、最小経過時間は実装定義ですが、秒単位程度であると想定できます。
	MinAge time.Duration

	// MaxBytesは、ウィンドウのサイズのバイト単位の上限です。
	//
	// この設定はMinAgeより優先されます。
	// ただし、WriteToが書き込むデータのサイズについては何も保証せず、
	// メモリのオーバーヘッドが常にMaxBytes未満に保たれることも保証しません。
	// ヒントとして扱ってください。
	//
	// これが0の場合、最大サイズは実装定義です。
	MaxBytes uint64
}
