// 「go test -run=Generate -write=all」によって生成されたコードです。編集しないでください。

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// BasicKindは基本型の種類を表します。
type BasicKind int

const (
	Invalid BasicKind = iota

	// 事前に宣言された型
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	String
	UnsafePointer

	// 型のない値に対する型
	UntypedBool
	UntypedInt
	UntypedRune
	UntypedFloat
	UntypedComplex
	UntypedString
	UntypedNil

	// エイリアス
	Byte = Uint8
	Rune = Int32
)

// BasicInfoは基本型のプロパティを記述するフラグの集合です。
type BasicInfo int

// 基本型のプロパティ。
const (
	IsBoolean BasicInfo = 1 << iota
	IsInteger
	IsUnsigned
	IsFloat
	IsComplex
	IsString
	IsUntyped

	IsOrdered   = IsInteger | IsFloat | IsString
	IsNumeric   = IsInteger | IsFloat | IsComplex
	IsConstType = IsBoolean | IsNumeric | IsString
)

// Basicは基本型を表します。
type Basic struct {
	kind BasicKind
	info BasicInfo
	name string
}

// Kindは基本型bの種類を返します。
func (b *Basic) Kind() BasicKind

// Infoは基本型bのプロパティに関する情報を返します。
func (b *Basic) Info() BasicInfo

// Nameは基本型bの名前を返します。
func (b *Basic) Name() string

func (b *Basic) Underlying() Type
func (b *Basic) String() string
