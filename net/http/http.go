// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate bundle -o=h2_bundle.go -prefix=http2 -tags=!nethttpomithttp2 -import=golang.org/x/net/internal/httpcommon=net/http/internal/httpcommon golang.org/x/net/http2

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

// HTTP2Config defines HTTP/2 configuration parameters common to
// both [Transport] and [Server].
type HTTP2Config struct {
	// MaxConcurrentStreams optionally specifies the number of
	// concurrent streams that a peer may have open at a time.
	// If zero, MaxConcurrentStreams defaults to at least 100.
	MaxConcurrentStreams int

	// MaxDecoderHeaderTableSize optionally specifies an upper limit for the
	// size of the header compression table used for decoding headers sent
	// by the peer.
	// A valid value is less than 4MiB.
	// If zero or invalid, a default value is used.
	MaxDecoderHeaderTableSize int

	// MaxEncoderHeaderTableSize optionally specifies an upper limit for the
	// header compression table used for sending headers to the peer.
	// A valid value is less than 4MiB.
	// If zero or invalid, a default value is used.
	MaxEncoderHeaderTableSize int

	// MaxReadFrameSize optionally specifies the largest frame
	// this endpoint is willing to read.
	// A valid value is between 16KiB and 16MiB, inclusive.
	// If zero or invalid, a default value is used.
	MaxReadFrameSize int

	// MaxReceiveBufferPerConnection is the maximum size of the
	// flow control window for data received on a connection.
	// A valid value is at least 64KiB and less than 4MiB.
	// If invalid, a default value is used.
	MaxReceiveBufferPerConnection int

	// MaxReceiveBufferPerStream is the maximum size of
	// the flow control window for data received on a stream (request).
	// A valid value is less than 4MiB.
	// If zero or invalid, a default value is used.
	MaxReceiveBufferPerStream int

	// SendPingTimeout is the timeout after which a health check using a ping
	// frame will be carried out if no frame is received on a connection.
	// If zero, no health check is performed.
	SendPingTimeout time.Duration

	// PingTimeout is the timeout after which a connection will be closed
	// if a response to a ping is not received.
	// If zero, a default of 15 seconds is used.
	PingTimeout time.Duration

	// WriteByteTimeout is the timeout after which a connection will be
	// closed if no data can be written to it. The timeout begins when data is
	// available to write, and is extended whenever any bytes are written.
	WriteByteTimeout time.Duration

	// PermitProhibitedCipherSuites, if true, permits the use of
	// cipher suites prohibited by the HTTP/2 spec.
	PermitProhibitedCipherSuites bool

	// CountError, if non-nil, is called on HTTP/2 errors.
	// It is intended to increment a metric for monitoring.
	// The errType contains only lowercase letters, digits, and underscores
	// (a-z, 0-9, _).
	CountError func(errType string)
}
