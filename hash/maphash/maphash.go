// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// maphashパッケージはバイト列上のハッシュ関数を提供します。
// これらのハッシュ関数は、任意の文字列やバイト列を符号なし64ビット整数上の均一な分布にマッピングするために使用されることを意図しています。
// ハッシュテーブルやデータ構造の異なるインスタンスごとに個別の [Seed] を使用する必要があります。
//
// これらのハッシュ関数は、暗号的に安全ではありません。
// (暗号的な使用には、crypto/sha256およびcrypto/sha512を参照してください。)
package maphash

// Seedは [Hash] によって計算されるハッシュ関数を選択するランダムな値です。
// もし2つのHashが同じSeedを使う場合、任意の入力に対して同じハッシュ値を計算します。
// もし2つのHashが異なるSeedを使う場合、任意の入力に対して異なるハッシュ値を計算する可能性が非常に高いです。
//
// Seedは [MakeSeed] を呼び出すことで初期化する必要があります。
// ゼロSeedは初期化されておらず、 [Hash] のSetSeedメソッドで使用することはできません。
//
// 各Seedの値は単一のプロセスに対してローカルであり、別のプロセスでシリアライズまたは再作成することはできません。
type Seed struct {
	s uint64
}

// Bytesは与えられたシードでbのハッシュを返します。
//
// Bytesは以下と同等であり、しかもより便利で効率的です：
//
//	var h Hash
//	h.SetSeed(seed)
//	h.Write(b)
//	return h.Sum64()
func Bytes(seed Seed, b []byte) uint64

// Stringは与えられたシードを使ってsのハッシュ値を返します。
//
// Stringは以下のコードと同等ですが、より便利かつ効率的です：
//
//	var h Hash
//	h.SetSeed(seed)
//	h.WriteString(s)
//	return h.Sum64()
func String(seed Seed, s string) uint64

// Hashはバイトシーケンスのシード付きハッシュを計算します。
//
// ゼロのHashは使用する準備ができた有効なHashです。
// ゼロのHashは、Reset、Write、Seed、またはSum64メソッドの最初の呼び出し時に自動でランダムなシードを選択します。
// シードに対する制御には、SetSeedを使用します。
//
// 計算されたハッシュ値は、初期シードとHashオブジェクトに提供されたバイトのシーケンスにのみ依存し、
// バイトの提供方法には依存しません。たとえば、以下の3つのシーケンス
//
//	h.Write([]byte{'f','o','o'})
//	h.WriteByte('f'); h.WriteByte('o'); h.WriteByte('o')
//	h.WriteString("foo")
//
// はすべて同じ効果があります。
//
// Hashは、ハッシュされるバイトシーケンスを制御する敵対者がいる場合でも、衝突耐性が意図されています。
//
// Hashは、複数のゴルーチンによる並行使用には安全ではありませんが、Seedは安全です。
// 複数のゴルーチンが同じシードを使ってハッシュを計算する必要がある場合、
// 各ゴルーチンは独自のHashを宣言し、共通のシードでSetSeedを呼び出すことができます。
type Hash struct {
	_     [0]func()
	seed  Seed
	state Seed
	buf   [bufSize]byte
	n     int
}

// WriteByteは、hによってハッシュされるバイト列にbを追加します。
// 失敗することはありません。エラーの結果は、 [io.ByteWriter] の実装のためです。
func (h *Hash) WriteByte(b byte) error

// Writeはhによってハッシュされたバイトのシーケンスにbを追加します。
// bのすべてを書き込み、失敗することはありません。countとerrorの結果は [io.Writer] を実装するためです。
func (h *Hash) Write(b []byte) (int, error)

// WriteString は文字列 s のバイト列を h によってハッシュ化されたバイト列に追加します。
// いつでも s のすべてを書き込み、失敗することはありません。count と error の結果は [io.StringWriter] の実装のためです。
func (h *Hash) WriteString(s string) (int, error)

// Seedはhのシード値を返します。
func (h *Hash) Seed() Seed

// SetSeedは、hにseedを使用するように設定します。
// seedは [MakeSeed] によって返されたか、別のハッシュのSeedメソッドによって返されたものでなければなりません。
// 同じseedを持つ2つの [Hash] オブジェクトは同じように振る舞います。
// 異なるseedを持つ2つの [Hash] オブジェクトは、非常に異なる振る舞いをする可能性があります。
// この呼び出し前にhに追加されたすべてのバイトは破棄されます。
func (h *Hash) SetSeed(seed Seed)

// Resetはhに追加されたすべてのバイトを破棄します。
// （シードは変わりません。）
func (h *Hash) Reset()

// Sum64はhの現在の64ビット値を返します。これは、hのシードと、hに追加されたバイトのシーケンスに依存します。
// 最後の [Hash.Reset] または [Hash.SetSeed] の呼び出しからの呼び出しを含みます。
// Sum64の結果のすべてのビットはほぼ均等に一様に分布しており、独立しているため、
// ビットマスキング、シフト、またはモジュラ演算で安全に縮小することができます。
func (h *Hash) Sum64() uint64

// MakeSeed は新しいランダムなシードを返す。
func MakeSeed() Seed

// Sumはハッシュの現在の64ビット値をbに追加します。
// [hash.Hash] の実装のために存在します。
// 直接の呼び出しでは、 [Hash.Sum64] を使用する方が効率が良いです。
func (h *Hash) Sum(b []byte) []byte

// Sizeはhのハッシュ値のサイズ、8バイトを返します。
func (h *Hash) Size() int

// BlockSize は h のブロックサイズを返します。
func (h *Hash) BlockSize() int
