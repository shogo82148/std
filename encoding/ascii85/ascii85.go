// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ascii85パッケージはbtoaツールやAdobeのPostScriptおよびPDFドキュメント形式で使用されているascii85データエンコーディングを実装しています。
package ascii85

import (
	"github.com/shogo82148/std/io"
)

// Encodeは、srcを最大 [MaxEncodedLen](len(src)) バイトのdstにエンコードし、
// 実際に書き込まれたバイト数を返します。
//
// エンコーディングは4バイトのチャンク単位で処理され、最後の断片には特別な
// エンコーディングを使用するため、Encodeは大きなデータストリームの個々のブロックに
// 使用するのには適していません。代わりに [NewEncoder] を使用してください。
//
// しばしば、ascii85でエンコードされたデータは<~と~>の記号で囲まれていますが、Encodeはこれを追加しません。
func Encode(dst, src []byte) int

// MaxEncodedLenは、n個のソースバイトのエンコーディングの最大長を返します。
func MaxEncodedLen(n int) int

// NewEncoderは新しいascii85ストリームエンコーダーを返します。返されたライターに書き込まれたデータはエンコードされ、wに書き込まれます。Ascii85エンコーディングは32ビットのブロックで動作します。書き込みが終了したら、呼び出し元は残りの部分ブロックをフラッシュするために返されたエンコーダーを閉じる必要があります。
func NewEncoder(w io.Writer) io.WriteCloser

type CorruptInputError int64

func (e CorruptInputError) Error() string

// Decodeは、srcをdstにデコードし、dstに書き込まれたバイト数と
// srcから消費されたバイト数の両方を返します。
// srcに不正なascii85データが含まれている場合、Decodeは正常に書き込まれた
// バイト数と [CorruptInputError] を返します。
// Decodeはsrc内の空白と制御文字を無視します。
// ascii85でエンコードされたデータは、しばしば<~と~>の記号で囲まれています。
// Decodeは、これらが呼び出し元によって取り除かれていることを想定しています。
//
// flushがtrueの場合、Decodeはsrcが入力ストリームの終わりを表し、別の32ビットブロックの完了を待つのではなく、完全に処理すると想定します。
//
// [NewDecoder] は、Decodeを [io.Reader] インターフェースでラップします。
func Decode(dst, src []byte, flush bool) (ndst, nsrc int, err error)

// NewDecoder は新しい ascii85 ストリームデコーダを構築します。
func NewDecoder(r io.Reader) io.Reader
