// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 golang.org/x/net/http2

package http

import (
	"github.com/shogo82148/std/io"
)

// NoBodyはバイトを持たないio.ReadCloserです。Readは常にEOFを返し、
// Closeは常にnilを返します。これは、リクエストがゼロバイトであることを
// 明示的に示すために、送信元クライアントのリクエストで使用することができます。
// ただし、代わりにRequest.Bodyをnilに設定することもできます。
var NoBody = noBody{}

var (
	// NoBodyからのio.Copyがバッファを必要としないことを検証する
	_ io.WriterTo   = NoBody
	_ io.ReadCloser = NoBody
)

// PushOptionsは、Pusher.Pushのオプションを記述します。
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
<<<<<<< HEAD
=======
	// Pushは、HTTP/2サーバープッシュを開始します。これは、指定されたターゲットとオプションを使用して合成リクエストを構築し、
	// そのリクエストをPUSH_PROMISEフレームにシリアライズし、そのリクエストをサーバーのリクエストハンドラを使用してディスパッチします。
	// optsがnilの場合、デフォルトのオプションが使用されます。
	//
	// ターゲットは、絶対パス（"/path"のような）または親リクエストと同じスキームと有効なホストを含む絶対URLでなければなりません。
	// ターゲットがパスの場合、親リクエストのスキームとホストを継承します。
	//
	// HTTP/2の仕様では、再帰的なプッシュとクロスオーソリティプッシュが禁止されています。
	// Pushはこれらの無効なプッシュを検出するかもしれませんし、検出しないかもしれません。しかし、無効な
	// プッシュは、準拠しているクライアントによって検出され、キャンセルされます。
	//
	// URL Xをプッシュしたいハンドラは、URL Xのリクエストをトリガーする可能性のあるデータを送信する前にPushを呼び出すべきです。
	// これにより、クライアントがXのPUSH_PROMISEを受信する前にXのリクエストを発行するというレース条件を回避します。
	//
	// Pushは別のゴルーチンで実行されるため、到着の順序は非決定的です。
	// 必要な同期は呼び出し元によって実装する必要があります。
	//
	// Pushは、クライアントがプッシュを無効にした場合、または基本となる接続でプッシュがサポートされていない場合、ErrNotSupportedを返します。
>>>>>>> release-branch.go1.21
	Push(target string, opts *PushOptions) error
}
