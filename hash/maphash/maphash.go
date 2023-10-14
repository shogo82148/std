// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<< HEAD
// パッケージmaphashはバイト列上のハッシュ関数を提供します。
// これらのハッシュ関数は、任意の文字列やバイト列を符号なし64ビット整数上の均一な分布にマッピングするために使用されることを意図しています。
// ハッシュテーブルやデータ構造の異なるインスタンスごとに個別のSeedを使用する必要があります。
=======
// Package maphash provides hash functions on byte sequences.
// These hash functions are intended to be used to implement hash tables or
// other data structures that need to map arbitrary strings or byte
// sequences to a uniform distribution on unsigned 64-bit integers.
// Each different instance of a hash table or data structure should use its own [Seed].
>>>>>>> upstream/master
//
// これらのハッシュ関数は、暗号的に安全ではありません。
// (暗号的な使用には、crypto/sha256およびcrypto/sha512を参照してください。)
package maphash

<<<<<<< HEAD
// Seedは特定のハッシュ関数を選択するランダムな値です。
// もし2つのHashが同じSeedを使う場合、任意の入力に対して同じハッシュ値を計算します。
// もし2つのHashが異なるSeedを使う場合、任意の入力に対して異なるハッシュ値を計算する可能性が非常に高いです。
//
// SeedはMakeSeedを呼び出すことで初期化する必要があります。
// ゼロSeedは初期化されておらず、HashのSetSeedメソッドで使用することはできません。
=======
// A Seed is a random value that selects the specific hash function
// computed by a [Hash]. If two Hashes use the same Seeds, they
// will compute the same hash values for any given input.
// If two Hashes use different Seeds, they are very likely to compute
// distinct hash values for any given input.
//
// A Seed must be initialized by calling [MakeSeed].
// The zero seed is uninitialized and not valid for use with [Hash]'s SetSeed method.
>>>>>>> upstream/master
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
//  h.Write([]byte{'f','o','o'})
//  h.WriteByte('f'); h.WriteByte('o'); h.WriteByte('o')
//  h.WriteString("foo")
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

<<<<<<< HEAD
// WriteByteは、hによってハッシュされるバイト列にbを追加します。
// 失敗することはありません。エラーの結果は、io.ByteWriterの実装のためです。
func (h *Hash) WriteByte(b byte) error

// Writeはhによってハッシュされたバイトのシーケンスにbを追加します。
// bのすべてを書き込み、失敗することはありません。countとerrorの結果はio.Writerを実装するためです。
func (h *Hash) Write(b []byte) (int, error)

// WriteString は文字列 s のバイト列を h によってハッシュ化されたバイト列に追加します。
// いつでも s のすべてを書き込み、失敗することはありません。count と error の結果は io.StringWriter の実装のためです。
=======
// WriteByte adds b to the sequence of bytes hashed by h.
// It never fails; the error result is for implementing [io.ByteWriter].
func (h *Hash) WriteByte(b byte) error

// Write adds b to the sequence of bytes hashed by h.
// It always writes all of b and never fails; the count and error result are for implementing [io.Writer].
func (h *Hash) Write(b []byte) (int, error)

// WriteString adds the bytes of s to the sequence of bytes hashed by h.
// It always writes all of s and never fails; the count and error result are for implementing [io.StringWriter].
>>>>>>> upstream/master
func (h *Hash) WriteString(s string) (int, error)

// Seedはhのシード値を返します。
func (h *Hash) Seed() Seed

<<<<<<< HEAD
// SetSeedは、hにseedを使用するように設定します。
// seedはMakeSeedによって返されたか、別のハッシュのSeedメソッドによって返されたものでなければなりません。
// 同じseedを持つ2つのハッシュオブジェクトは同じように振る舞います。
// 異なるseedを持つ2つのハッシュオブジェクトは、非常に異なる振る舞いをする可能性があります。
// この呼び出し前にhに追加されたすべてのバイトは破棄されます。
=======
// SetSeed sets h to use seed, which must have been returned by [MakeSeed]
// or by another [Hash.Seed] method.
// Two [Hash] objects with the same seed behave identically.
// Two [Hash] objects with different seeds will very likely behave differently.
// Any bytes added to h before this call will be discarded.
>>>>>>> upstream/master
func (h *Hash) SetSeed(seed Seed)

// Resetはhに追加されたすべてのバイトを破棄します。
// （シードは変わりません。）
func (h *Hash) Reset()

<<<<<<< HEAD
// Sum64はhの現在の64ビット値を返します。これは、hのシードと、hに追加されたバイトのシーケンスに依存します。
// 最後のResetまたはSetSeedの呼び出しからの呼び出しを含みます。
// Sum64の結果のすべてのビットはほぼ均等に一様に分布しており、独立しているため、
// ビットマスキング、シフト、またはモジュラ演算で安全に縮小することができます。
=======
// Sum64 returns h's current 64-bit value, which depends on
// h's seed and the sequence of bytes added to h since the
// last call to [Hash.Reset] or [Hash.SetSeed].
//
// All bits of the Sum64 result are close to uniformly and
// independently distributed, so it can be safely reduced
// by using bit masking, shifting, or modular arithmetic.
>>>>>>> upstream/master
func (h *Hash) Sum64() uint64

// MakeSeed は新しいランダムなシードを返す。
func MakeSeed() Seed

<<<<<<< HEAD
// Sumはハッシュの現在の64ビット値をbに追加します。
// hash.Hashの実装のために存在します。
// 直接の呼び出しでは、Sum64を使用する方が効率が良いです。
=======
// Sum appends the hash's current 64-bit value to b.
// It exists for implementing [hash.Hash].
// For direct calls, it is more efficient to use [Hash.Sum64].
>>>>>>> upstream/master
func (h *Hash) Sum(b []byte) []byte

// Sizeはhのハッシュ値のサイズ、8バイトを返します。
func (h *Hash) Size() int

// BlockSize は h のブロックサイズを返します。
func (h *Hash) BlockSize() int
