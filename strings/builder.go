// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

// Builderは、 [Builder.Write] メソッドを使用して効率的に文字列を構築するために使用されます。
// メモリのコピーを最小限に抑えます。ゼロ値はすぐに使用できます。
// 非ゼロのBuilderをコピーしないでください。
type Builder struct {
	addr *Builder
	buf  []byte
}

// Stringは、蓄積された文字列を返します。
func (b *Builder) String() string

// Lenは、蓄積されたバイト数を返します。b.Len() == len(b.String())です。
func (b *Builder) Len() int

// Capは、ビルダーの基礎となるバイトスライスの容量を返します。
// 構築中の文字列に割り当てられた総スペースを含み、すでに書き込まれたバイトも含みます。
func (b *Builder) Cap() int

// Resetは、 [Builder] を空にリセットします。
func (b *Builder) Reset()

// Growは、必要に応じてbの容量を拡張し、別のnバイトのスペースを保証します。
// Grow(n)の後、少なくともnバイトを別の割り当てなしでbに書き込むことができます。
// nが負の場合、Growはパニックを引き起こします。
func (b *Builder) Grow(n int)

// Writeは、pの内容をbのバッファに追加します。
// Writeは常にlen(p)、nilを返します。
func (b *Builder) Write(p []byte) (int, error)

// WriteByteは、バイトcをbのバッファに追加します。
// 返されるエラーは常にnilです。
func (b *Builder) WriteByte(c byte) error

// WriteRuneは、UnicodeコードポイントrのUTF-8エンコーディングをbのバッファに追加します。
// rの長さとnilエラーを返します。
func (b *Builder) WriteRune(r rune) (int, error)

// WriteStringは、sの内容をbのバッファに追加します。
// sの長さとnilエラーを返します。
func (b *Builder) WriteString(s string) (int, error)
