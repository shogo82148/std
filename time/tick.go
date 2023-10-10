// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// Tickerは、時計の間隔で "ticks" を配信するチャネルを保持します。
type Ticker struct {
	C <-chan Time
	r runtimeTimer
}

// NewTickerは、各チック後にチャネルに現在の時刻を送信するチッカーを返します。チックの期間はduration引数で指定されます。チッカーは、遅い受信者のために時間間隔を調整するか、チックを削除します。 dの期間はゼロより大きくする必要があります。そうでない場合、NewTickerはパニックを起こします。チッカーを停止して関連するリソースを解放します。
func NewTicker(d Duration) *Ticker

// Stopはティッカーを停止します。Stop後は、もうティックが送信されません。
// Stopはチャネルを閉じないため、チャネルから読み取り中の並行ゴルーチンが誤った「ティック」を見ることを防ぎます。
func (t *Ticker) Stop()

// Resetはタイマーを停止し、指定された期間でタイマーをリセットします。
// 新しい期間が経過すると、次のティックが到着します。期間dは0より大きくなければなりません。もしそうでない場合、Resetはパニックを引き起こします。
func (t *Ticker) Reset(d Duration)

// Tickは、ティッキングチャネルにのみアクセスを提供するためのNewTickerの便利なラッパーです。Tickerをシャットダウンする必要のないクライアントに便利ですが、シャットダウンの方法がないため、ガベージコレクタによって元のTickerは回収されません。"リーク"します。
// NewTickerとは異なり、d <= 0の場合にはnilが返されます。
func Tick(d Duration) <-chan Time
