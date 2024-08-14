// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ascii85パッケージはbtoaツールやAdobeのPostScriptおよびPDFドキュメント形式で使用されているascii85データエンコーディングを実装しています。
package ascii85

import (
	"github.com/shogo82148/std/io"
)

// Encode encodes src into at most [MaxEncodedLen](len(src))
// bytes of dst, returning the actual number of bytes written.
//
// The encoding handles 4-byte chunks, using a special encoding
// for the last fragment, so Encode is not appropriate for use on
// individual blocks of a large data stream. Use [NewEncoder] instead.
//
// しばしば、ascii85でエンコードされたデータは<~と~>の記号で囲まれていますが、Encodeはこれを追加しません。
func Encode(dst, src []byte) int

// MaxEncodedLenは、n個のソースバイトのエンコーディングの最大長を返します。
func MaxEncodedLen(n int) int

// NewEncoderは新しいascii85ストリームエンコーダーを返します。返されたライターに書き込まれたデータはエンコードされ、wに書き込まれます。Ascii85エンコーディングは32ビットのブロックで動作します。書き込みが終了したら、呼び出し元は残りの部分ブロックをフラッシュするために返されたエンコーダーを閉じる必要があります。
func NewEncoder(w io.Writer) io.WriteCloser

type CorruptInputError int64

func (e CorruptInputError) Error() string

// Decode decodes src into dst, returning both the number
// of bytes written to dst and the number consumed from src.
// If src contains invalid ascii85 data, Decode will return the
// number of bytes successfully written and a [CorruptInputError].
// Decode ignores space and control characters in src.
// Often, ascii85-encoded data is wrapped in <~ and ~> symbols.
// Decode expects these to have been stripped by the caller.
//
// flushがtrueの場合、Decodeはsrcが入力ストリームの終わりを表し、別の32ビットブロックの完了を待つのではなく、完全に処理すると想定します。
//
// [NewDecoder] wraps an [io.Reader] interface around Decode.
func Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error)

// NewDecoder は新しい ascii85 ストリームデコーダを構築します。
func NewDecoder(r io.Reader) io.Reader
