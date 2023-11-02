// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// DWARF type information structures.
// The format is heavily biased toward C, but for simplicity
// the String methods use a pseudo-Go syntax.

package dwarf

<<<<<<< HEAD
// A Type conventionally represents a pointer to any of the
// specific Type structures ([CharType], [StructType], etc.).
=======
// Typeは、通常、特定のType構造体（CharType、StructTypeなど）のいずれかへのポインタを表します。
>>>>>>> release-branch.go1.21
type Type interface {
	Common() *CommonType
	String() string
	Size() int64
}

// CommonTypeは、複数の型で共通のフィールドを保持します。
// フィールドが特定の型に対して不明または適用できない場合、
// ゼロ値が使用されます。
type CommonType struct {
	ByteSize int64
	Name     string
}

func (c *CommonType) Common() *CommonType

func (c *CommonType) Size() int64

// BasicTypeは、すべての基本型に共通するフィールドを保持します。
//
<<<<<<< HEAD
// See the documentation for [StructField] for more info on the interpretation of
// the BitSize/BitOffset/DataBitOffset fields.
=======
// BitSize / BitOffset / DataBitOffsetフィールドの解釈に関する詳細については、StructFieldのドキュメントを参照してください。
>>>>>>> release-branch.go1.21
type BasicType struct {
	CommonType
	BitSize       int64
	BitOffset     int64
	DataBitOffset int64
}

func (b *BasicType) Basic() *BasicType

func (t *BasicType) String() string

// CharTypeは、符号付き文字型を表します。
type CharType struct {
	BasicType
}

// UcharTypeは、符号なし文字型を表します。
type UcharType struct {
	BasicType
}

// IntTypeは、符号付き整数型を表します。
type IntType struct {
	BasicType
}

// UintTypeは、符号なし整数型を表します。
type UintType struct {
	BasicType
}

// FloatTypeは、浮動小数点数型を表します。
type FloatType struct {
	BasicType
}

// ComplexTypeは、複素数の浮動小数点数型を表します。
type ComplexType struct {
	BasicType
}

// BoolTypeは、ブール型を表します。
type BoolType struct {
	BasicType
}

// AddrTypeは、マシンアドレス型を表します。
type AddrType struct {
	BasicType
}

// UnspecifiedTypeは、暗黙的な、不明な、曖昧な、または存在しない型を表します。
type UnspecifiedType struct {
	BasicType
}

// QualTypeは、C/C++の「const」、「restrict」、または「volatile」修飾子を持つ型を表します。
type QualType struct {
	CommonType
	Qual string
	Type Type
}

func (t *QualType) String() string

func (t *QualType) Size() int64

// ArrayTypeは、固定サイズの配列型を表します。
type ArrayType struct {
	CommonType
	Type          Type
	StrideBitSize int64
	Count         int64
}

func (t *ArrayType) String() string

func (t *ArrayType) Size() int64

// VoidTypeは、Cのvoid型を表します。
type VoidType struct {
	CommonType
}

func (t *VoidType) String() string

// PtrTypeは、ポインタ型を表します。
type PtrType struct {
	CommonType
	Type Type
}

func (t *PtrType) String() string

// StructTypeは、構造体、共用体、またはC++クラス型を表します。
type StructType struct {
	CommonType
	StructName string
	Kind       string
	Field      []*StructField
	Incomplete bool
}

// StructFieldは、構造体、共用体、またはC++クラス型のフィールドを表します。
//
// # ビットフィールド
//
// BitSize、BitOffset、DataBitOffsetフィールドは、C/C++のstruct/union/class型でビットフィールドとして宣言されたデータメンバーのビットサイズとオフセットを記述します。
//
// BitSizeは、ビットフィールドのビット数です。
//
// DataBitOffsetが0以外の場合、それは、
// ビットフィールドの開始位置から囲むエンティティ（例：含まれるstruct/class/union）の開始位置までのビット数です。
// これは、DWARF 4で導入されたDW_AT_data_bit_offset DWARF属性に対応します。
//
// BitOffsetが0以外の場合、それは、
// ビットフィールドを保持するストレージユニットの最上位ビットからビットフィールドの最上位ビットまでのビット数です。
// ここで「ストレージユニット」とは、ビットフィールドの前の型名です（フィールド「unsigned x：17」の場合、ストレージユニットは「unsigned」です）。
// BitOffsetの値は、システムのエンディアンによって異なる場合があります。
// BitOffsetは、DWARF 4で廃止され、DWARF 5で削除されたDW_AT_bit_offset DWARF属性に対応します。
//
// DataBitOffsetとBitOffsetのうち、最大1つだけが非ゼロになります。
// DataBitOffset/BitOffsetは、BitSizeが非ゼロの場合にのみ非ゼロになります。
// Cコンパイラがどちらを使用するかは、コンパイラのバージョンとコマンドラインオプションに依存します。
//
// 以下は、C/C++のビットフィールドの使用例と、DWARFビットオフセット情報の期待値の例です。
// このコードを考えてみましょう。
//
//	struct S {
//		int q;
//		int j:5;
//		int k:6;
//		int m:5;
//		int n:8;
//	} s;
//
// 上記のコードに対して、GCC 8を使用している場合、DW_AT_bit_offsetの値は次のようになることが期待されます。
//
//	       Little   |     Big
//	       Endian   |    Endian
//	                |
//	"j":     27     |     0
//	"k":     21     |     5
//	"m":     16     |     11
//	"n":     8      |     16
//
// 上記のように、j/k/m/nの含まれるストレージユニットに対するオフセットのみが考慮されます。これらの値は、含まれる構造体の前のデータメンバーのサイズに基づいて変化することはありません。
//
// コンパイラがDW_AT_data_bit_offsetを出力する場合、期待される値は次のようになります。
//
//	"j":     32
//	"k":     37
//	"m":     43
//	"n":     48
//
// ここで、"j"の値32は、ビットフィールドが他のデータメンバーに続くことを反映しています（DW_AT_data_bit_offset値は、含まれる構造体の開始位置を基準としています）。
// したがって、DW_AT_data_bit_offset値は、多数のフィールドを持つ構造体ではかなり大きくなる可能性があります。
//
// DWARFは、ビットサイズとビットオフセットがゼロでない基本型の可能性も許容しているため、これらの情報は基本型に対してもキャプチャされます。
// ただし、主流の言語を使用してこの動作をトリガーすることはできないことに注意する価値があります。
type StructField struct {
	Name          string
	Type          Type
	ByteOffset    int64
	ByteSize      int64
	BitOffset     int64
	DataBitOffset int64
	BitSize       int64
}

func (t *StructType) String() string

func (t *StructType) Defn() string

<<<<<<< HEAD
// An EnumType represents an enumerated type.
// The only indication of its native integer type is its ByteSize
// (inside [CommonType]).
=======
// EnumTypeは、列挙型を表します。
// ネイティブの整数型の唯一の指標は、CommonType内のByteSizeです。
>>>>>>> release-branch.go1.21
type EnumType struct {
	CommonType
	EnumName string
	Val      []*EnumValue
}

// EnumValueは、単一の列挙値を表します。
type EnumValue struct {
	Name string
	Val  int64
}

func (t *EnumType) String() string

// FuncTypeは、関数の型を表します。
type FuncType struct {
	CommonType
	ReturnType Type
	ParamType  []Type
}

func (t *FuncType) String() string

// DotDotDotTypeは、可変長の...関数パラメータを表します。
type DotDotDotType struct {
	CommonType
}

func (t *DotDotDotType) String() string

// TypedefTypeは、名前付き型を表します。
type TypedefType struct {
	CommonType
	Type Type
}

func (t *TypedefType) String() string

func (t *TypedefType) Size() int64

// UnsupportedTypeは、サポートされていない型に遭遇した場合に返されるプレースホルダーです。
type UnsupportedType struct {
	CommonType
	Tag Tag
}

func (t *UnsupportedType) String() string

// Typeは、DWARF「info」セクションのoffにある型を読み取ります。
func (d *Data) Type(off Offset) (Type, error)
