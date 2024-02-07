// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package httputilは、net/httpパッケージにある一般的なものと補完するHTTPユーティリティ関数を提供します。
package httputil

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/http/internal"
)

<<<<<<< HEAD
// NewChunkedReaderは、rから読み込まれたデータをHTTPの「チャンク」形式から変換して返す新しいchunkedReaderを返します。
// chunkedReaderは、最後の長さ0のチャンクが読み込まれた時にio.EOFを返します。
=======
// NewChunkedReader returns a new chunkedReader that translates the data read from r
// out of HTTP "chunked" format before returning it.
// The chunkedReader returns [io.EOF] when the final 0-length chunk is read.
>>>>>>> upstream/release-branch.go1.22
//
// NewChunkedReaderは通常のアプリケーションでは必要ありません。httpパッケージは、応答ボディを読み込む際に自動的にチャンクをデコードします。
func NewChunkedReader(r io.Reader) io.Reader

// NewChunkedWriter は、w に書き込む前に書き込みを HTTP の "chunked" フォーマットに変換する新しい chunkedWriter を返します。返された chunkedWriter を閉じると、ストリームの終わりを示す最後の長さが 0 のチャンクが送信されますが、トレーラーの後に表示される最後の CRLF は送信されません。トレーラーと最後の CRLF は別個に書き込む必要があります。
// NewChunkedWriter は通常のアプリケーションでは必要ありません。http パッケージは、ハンドラが Content-Length ヘッダーを設定しない場合、自動的にチャンキングを追加します。ハンドラ内で NewChunkedWriter を使用すると、二重チャンキングや Content-Length 長さでのチャンキングなど、間違った結果になります。
func NewChunkedWriter(w io.Writer) io.WriteCloser

// ErrLineTooLong は、行が長すぎる不正なチャンクデータを読み取ると返されます。
var ErrLineTooLong = internal.ErrLineTooLong
