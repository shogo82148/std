// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package time

// Tickerは、時計の間隔で "ticks" を配信するチャネルを保持します。
type Ticker struct {
	C    <-chan Time
	self *Ticker
}

// NewTickerは、各ティック後にチャネル上の現在の時間を送信するチャネルを含む新しい [Ticker] を返します。
// ティックの周期は、duration引数で指定されます。Tickerは、受信速度が遅い場合に時間間隔を調整したり、
// ティックをドロップしたりします。duration dはゼロより大きくなければならず、そうでない場合、
// NewTickerはパニックを引き起こします。
//
// Go 1.23より前では、ガベージコレクターはまだ期限切れになっていないか停止していない
// tickerを回収しなかったため、コードはしばしばNewTickerを呼び出した直後にt.Stopを即時に遅延させ、
// tickerが不要になったときに回収可能にしました。
// Go 1.23以降では、ガベージコレクターは参照されていないtickerを回収できます、
// たとえそれらが停止していなくても。
// Stopメソッドはもはやガベージコレクターを助けるためには必要ありません。
// （もちろん、コードは他の理由でtickerを停止させるためにStopを呼び出すことを望むかもしれません。）
func NewTicker(d Duration) *Ticker

// Stopはtickerを停止します。Stop後、これ以上のティックは送信されません。
// Stopはチャネルを閉じません。これにより [Ticker.Reset] を呼び出すことができ、
// チャネルから読み取っている並行ゴルーチンが
// 誤ったティックを見ないようにするためです。
func (t *Ticker) Stop()

// Resetはタイマーを停止し、指定された期間でタイマーをリセットします。
// 新しい期間が経過すると、次のティックが到着します。期間dは0より大きくなければなりません。もしそうでない場合、Resetはパニックを引き起こします。
func (t *Ticker) Reset(d Duration)

// Tickは、ティックするチャネルへのアクセスのみを提供する [NewTicker] の便利なラッパーです。
// NewTickerとは異なり、d <= 0の場合、Tickはnilを返します。
//
// Go 1.23より前では、このドキュメントは、基礎となる [Ticker] はガベージコレクターによって
// 決して回収されないこと、そして効率が問題になる場合はコードは代わりにNewTickerを使用し、
// ティッカーが不要になったら [Ticker.Stop] を呼び出すべきであることを警告していました。
// Go 1.23以降、ガベージコレクターは、停止されていなくても参照されていないティッカーを
// 回収できます。
// ガベージコレクターを助けるためにStopメソッドを呼び出す必要はもうありません。
// Tickで十分な場合にNewTickerを優先する理由はもうありません。
func Tick(d Duration) <-chan Time
