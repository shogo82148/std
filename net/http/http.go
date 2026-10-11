// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

// ProtocolsはHTTPプロトコルのセットです。
// ゼロ値は空のプロトコルセットです。
//
// サポートされているプロトコルは以下の通りです：
//
//   - HTTP1はHTTP/1.0およびHTTP/1.1プロトコルです。
//     HTTP1は非暗号化TCP接続およびTLS接続の両方でサポートされています。
//
//   - HTTP2はTLS接続上のHTTP/2プロトコルです。
//
//   - UnencryptedHTTP2は非暗号化TCP接続上のHTTP/2プロトコルです。
type Protocols struct {
	bits uint8
}

// HTTP1はpにHTTP/1が含まれているかどうかを報告します。
func (p Protocols) HTTP1() bool

// SetHTTP1はpにHTTP/1を追加または削除します。
func (p *Protocols) SetHTTP1(ok bool)

// HTTP2はpにHTTP/2が含まれているかどうかを報告します。
func (p Protocols) HTTP2() bool

// SetHTTP2はpにHTTP/2を追加または削除します。
func (p *Protocols) SetHTTP2(ok bool)

// UnencryptedHTTP2はpに暗号化されていないHTTP/2が含まれているかどうかを報告します。
func (p Protocols) UnencryptedHTTP2() bool

// SetUnencryptedHTTP2はpに暗号化されていないHTTP/2を追加または削除します。
func (p *Protocols) SetUnencryptedHTTP2(ok bool)

func (p Protocols) String() string

// NoBodyはバイト数ゼロの [io.ReadCloser] です。Readは常にEOFを返し、
// Closeは常にnilを返します。リクエストボディがゼロバイトであることを
// 明示的に示すために、送信側のクライアントリクエストで使用できます。
// ただし、単に [Request.Body] をnilに設定することも代替手段です。
var NoBody = noBody{}

var (
	// NoBodyからのio.Copyがバッファを必要としないことを検証する
	_ io.WriterTo   = NoBody
	_ io.ReadCloser = NoBody
)

// PushOptionsは、[Pusher.Push] のオプションを記述します。
type PushOptions struct {

	// Methodは要求されたリクエストのHTTPメソッドを指定します。
	// 設定する場合、"GET"または"HEAD"でなければなりません。空は"GET"を意味します。
	Method string

	// Headerは追加の約束されたリクエストヘッダーを指定します。これには":path"や":scheme"などのHTTP/2疑似ヘッダーフィールドは含めることができませんが、これらは自動的に追加されます。
	Header Header
}

// Pusherは、HTTP/2サーバープッシュをサポートするResponseWritersによって実装されるインターフェースです。
// 詳細については、 https://tools.ietf.org/html/rfc7540#section-8.2 を参照してください。
type Pusher interface {
	Push(target string, opts *PushOptions) error
}

// HTTP2Configは、[Transport] と [Server] の両方に共通する
// HTTP/2の設定パラメータを定義します。
type HTTP2Config struct {
	// MaxConcurrentStreamsは、ピアが同時に開くことができる
	// 同時ストリームの数をオプションで指定します。
	// ゼロの場合、MaxConcurrentStreamsのデフォルトは少なくとも100です。
	MaxConcurrentStreams int

	// MaxDecoderHeaderTableSizeは、ピアから送信されたヘッダーのデコードに使用される
	// ヘッダー圧縮テーブルのサイズの上限をオプションで指定します。
	// 有効な値は4MiB未満です。
	// ゼロまたは無効な場合は、デフォルト値が使用されます。
	MaxDecoderHeaderTableSize int

	// MaxEncoderHeaderTableSizeは、ピアへのヘッダーの送信に使用される
	// ヘッダー圧縮テーブルの上限をオプションで指定します。
	// 有効な値は4MiB未満です。
	// ゼロまたは無効な場合は、デフォルト値が使用されます。
	MaxEncoderHeaderTableSize int

	// MaxReadFrameSizeは、このエンドポイントが読み取る
	// 最大のフレームサイズをオプションで指定します。
	// 有効な値は16KiBから16MiBまで（両端を含む）です。
	// ゼロまたは無効な場合は、デフォルト値が使用されます。
	MaxReadFrameSize int

	// MaxReceiveBufferPerConnectionは、接続で受信するデータの
	// フロー制御ウィンドウの最大サイズです。
	// 有効な値は64KiB以上4MiB未満です。
	// 無効な場合は、デフォルト値が使用されます。
	MaxReceiveBufferPerConnection int

	// MaxReceiveBufferPerStreamは、ストリーム（リクエスト）で受信するデータの
	// フロー制御ウィンドウの最大サイズです。
	// 有効な値は4MiB未満です。
	// ゼロまたは無効な場合は、デフォルト値が使用されます。
	MaxReceiveBufferPerStream int

	// SendPingTimeoutは、接続でフレームが受信されない場合に、
	// pingフレームを使用したヘルスチェックが実行されるまでのタイムアウトです。
	// ゼロの場合、ヘルスチェックは実行されません。
	SendPingTimeout time.Duration

	// PingTimeoutは、pingへの応答が受信されない場合に
	// 接続が閉じられるまでのタイムアウトです。
	// ゼロの場合、デフォルトの15秒が使用されます。
	PingTimeout time.Duration

	// WriteByteTimeoutは、データを書き込めない場合に接続が閉じられるまでの
	// タイムアウトです。タイムアウトは書き込むデータが利用可能になったときに開始し、
	// バイトが書き込まれるたびに延長されます。
	WriteByteTimeout time.Duration

	// PermitProhibitedCipherSuitesがtrueの場合、HTTP/2仕様で
	// 禁止されている暗号スイートの使用を許可します。
	PermitProhibitedCipherSuites bool

	// CountErrorがnilでない場合、HTTP/2エラー時に呼び出されます。
	// これは監視用のメトリクスをインクリメントすることを目的としています。
	// errTypeには小文字、数字、アンダースコア（a-z、0-9、_）のみが含まれます。
	CountError func(errType string)
}
