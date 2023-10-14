// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reflect はランタイムリフレクションを実装し、プログラムが任意の型のオブジェクトを操作できるようにします。典型的な使用方法は、静的型 interface{} の値を取り、TypeOf を呼び出してその動的な型情報を抽出することです。TypeOf は Type を返します。
// ValueOf への呼び出しは、実行時データを表す Value を返します。Zero は Type を受け取り、その型のゼロ値を表す Value を返します。
// Go におけるリフレクションの紹介については、「The Laws of Reflection」を参照してください：
// https://golang.org/doc/articles/laws_of_reflection.html
package reflect

import (
	"github.com/shogo82148/std/internal/abi"
)

// TypeはGoの型の表現です。
//
// すべてのメソッドがすべての種類の型に適用されるわけではありません。制限がある場合は、各メソッドのドキュメントに記載されています。
// kind-specificメソッドを呼び出す前に型の種類を知るためにKindメソッドを使用してください。型の種類に適切でないメソッドを呼び出すと、ランタイムパニックが発生します。
//
// Typeの値は比較可能であり、==演算子などで使用することができます。
// 2つのTypeの値は、同一の型を表している場合に等しいとされます。
type Type interface {

	// Align returns the alignment in bytes of a value of
	// this type when allocated in memory.
	Align() int

	// FieldAlign returns the alignment in bytes of a value of
	// this type when used as a field in a struct.
	FieldAlign() int

	// Method returns the i'th method in the type's method set.
	// It panics if i is not in the range [0, NumMethod()).
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver,
	// and only exported methods are accessible.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	//
	// Methods are sorted in lexicographic order.
	Method(int) Method

	// MethodByName returns the method with that name in the type's
	// method set and a boolean indicating if the method was found.
	//
	// For a non-interface type T or *T, the returned Method's Type and Func
	// fields describe a function whose first argument is the receiver.
	//
	// For an interface type, the returned Method's Type field gives the
	// method signature, without a receiver, and the Func field is nil.
	MethodByName(string) (Method, bool)

	// NumMethod returns the number of methods accessible using Method.
	//
	// For a non-interface type, it returns the number of exported methods.
	//
	// For an interface type, it returns the number of exported and unexported methods.
	NumMethod() int

	// Name returns the type's name within its package for a defined type.
	// For other (non-defined) types it returns the empty string.
	Name() string

	// PkgPath returns a defined type's package path, that is, the import path
	// that uniquely identifies the package, such as "encoding/base64".
	// If the type was predeclared (string, error) or not defined (*T, struct{},
	// []int, or A where A is an alias for a non-defined type), the package path
	// will be the empty string.
	PkgPath() string

	// Size returns the number of bytes needed to store
	// a value of the given type; it is analogous to unsafe.Sizeof.
	Size() uintptr

	// String returns a string representation of the type.
	// The string representation may use shortened package names
	// (e.g., base64 instead of "encoding/base64") and is not
	// guaranteed to be unique among types. To test for type identity,
	// compare the Types directly.
	String() string

	// Kind returns the specific kind of this type.
	Kind() Kind

	// Implements reports whether the type implements the interface type u.
	Implements(u Type) bool

	// AssignableTo reports whether a value of the type is assignable to type u.
	AssignableTo(u Type) bool

	// ConvertibleTo reports whether a value of the type is convertible to type u.
	// Even if ConvertibleTo returns true, the conversion may still panic.
	// For example, a slice of type []T is convertible to *[N]T,
	// but the conversion will panic if its length is less than N.
	ConvertibleTo(u Type) bool

	// Comparable reports whether values of this type are comparable.
	// Even if Comparable returns true, the comparison may still panic.
	// For example, values of interface type are comparable,
	// but the comparison will panic if their dynamic type is not comparable.
	Comparable() bool

	// Bits returns the size of the type in bits.
	// It panics if the type's Kind is not one of the
	// sized or unsized Int, Uint, Float, or Complex kinds.
	Bits() int

	// ChanDir returns a channel type's direction.
	// It panics if the type's Kind is not Chan.
	ChanDir() ChanDir

	// IsVariadic reports whether a function type's final input parameter
	// is a "..." parameter. If so, t.In(t.NumIn() - 1) returns the parameter's
	// implicit actual type []T.
	//
	// For concreteness, if t represents func(x int, y ... float64), then
	//
	//	t.NumIn() == 2
	//	t.In(0) is the reflect.Type for "int"
	//	t.In(1) is the reflect.Type for "[]float64"
	//	t.IsVariadic() == true
	//
	// IsVariadic panics if the type's Kind is not Func.
	IsVariadic() bool

	// Elem returns a type's element type.
	// It panics if the type's Kind is not Array, Chan, Map, Pointer, or Slice.
	Elem() Type

	// Field returns a struct type's i'th field.
	// It panics if the type's Kind is not Struct.
	// It panics if i is not in the range [0, NumField()).
	Field(i int) StructField

	// FieldByIndex returns the nested field corresponding
	// to the index sequence. It is equivalent to calling Field
	// successively for each index i.
	// It panics if the type's Kind is not Struct.
	FieldByIndex(index []int) StructField

	// FieldByName returns the struct field with the given name
	// and a boolean indicating if the field was found.
	FieldByName(name string) (StructField, bool)

	// FieldByNameFunc returns the struct field with a name
	// that satisfies the match function and a boolean indicating if
	// the field was found.
	//
	// FieldByNameFunc considers the fields in the struct itself
	// and then the fields in any embedded structs, in breadth first order,
	// stopping at the shallowest nesting depth containing one or more
	// fields satisfying the match function. If multiple fields at that depth
	// satisfy the match function, they cancel each other
	// and FieldByNameFunc returns no match.
	// This behavior mirrors Go's handling of name lookup in
	// structs containing embedded fields.
	FieldByNameFunc(match func(string) bool) (StructField, bool)

	// In returns the type of a function type's i'th input parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumIn()).
	In(i int) Type

	// Key returns a map type's key type.
	// It panics if the type's Kind is not Map.
	Key() Type

	// Len returns an array type's length.
	// It panics if the type's Kind is not Array.
	Len() int

	// NumField returns a struct type's field count.
	// It panics if the type's Kind is not Struct.
	NumField() int

	// NumIn returns a function type's input parameter count.
	// It panics if the type's Kind is not Func.
	NumIn() int

	// NumOut returns a function type's output parameter count.
	// It panics if the type's Kind is not Func.
	NumOut() int

	// Out returns the type of a function type's i'th output parameter.
	// It panics if the type's Kind is not Func.
	// It panics if i is not in the range [0, NumOut()).
	Out(i int) Type

	common() *abi.Type
	uncommon() *uncommonType
}

// Kindは、Typeが表す特定の種類の型を表します。
// ゼロのKindは有効な種類ではありません。
type Kind uint

const (
	Invalid Kind = iota
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
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

// PtrはPointer種別の旧名称です。
const Ptr = Pointer

// ChanDirはチャネルの方向を表します。
type ChanDir int

const (
	RecvDir ChanDir = 1 << iota
	SendDir
	BothDir = RecvDir | SendDir
)

// Methodは単一のメソッドを表します。
type Method struct {
	// Nameはメソッド名です。
	Name string

	// PkgPathは、小文字（エクスポートされていない）のメソッド名を修飾するパッケージパスです。大文字（エクスポートされた）のメソッド名の場合は空です。
	// PkgPathとNameの組み合わせは、メソッドセット内のメソッドを一意に識別します。
	// https://golang.org/ref/spec#Uniqueness_of_identifiers を参照してください。
	PkgPath string

	Type  Type
	Func  Value
	Index int
}

// IsExportedはメソッドがエクスポートされているかどうかを報告します。
func (m Method) IsExported() bool

// String は k の名前を返します。
func (k Kind) String() string

func (d ChanDir) String() string

// StructFieldはstruct内のフィールドを1つ記述します。
type StructField struct {
	// Nameはフィールド名です。
	Name string

	// PkgPathは小文字（エクスポートされていない）のフィールド名を修飾するパッケージパスです。大文字（エクスポートされた）のフィールド名には空です。
	// 詳細はhttps://golang.org/ref/spec#Uniqueness_of_identifiersを参照してください。
	PkgPath string

	Type      Type
	Tag       StructTag
	Offset    uintptr
	Index     []int
	Anonymous bool
}

// IsExportedはフィールドがエクスポートされているかどうかを報告します。
func (f StructField) IsExported() bool

// StructTagは、structフィールドのタグ文字列です。
//
// 慣例として、タグ文字列はオプションでスペースで区切られたkey:"value"の連結です。
// 各キーは、スペース（U+0020 ' '）、引用符（U+0022 '"'）、
// コロン（U+003A ':'）以外の非制御文字で構成される空でない文字列です。
// 各値は、U+0022 '"'文字とGoの文字列リテラル構文を使用して引用されます。
type StructTag string

// Getはtag文字列内のキーに関連付けられた値を返します。
// もしtag内にそのようなキーが存在しない場合、Getは空の文字列を返します。
// もしtagが従来の形式を持っていない場合、Getが返す値は不特定です。
// タグが明示的に空の文字列に設定されているかどうかを判断するには、Lookupを使用してください。
func (tag StructTag) Get(key string) string

// Lookupは、タグ文字列内のキーに関連する値を返します。
// キーがタグ内に存在する場合、その値（空かもしれません）が返されます。
// キーがタグに明示的に設定されていない場合、返される値は空の文字列になります。
// okの返り値は、値がタグ文字列に明示的に設定されているかどうかを報告します。
// タグに通常の形式がない場合、Lookupによって返される値は指定されていません。
func (tag StructTag) Lookup(key string) (value string, ok bool)

// TypeOfは、iの動的な型を表す反射Typeを返します。
// もしiがnilのインターフェース値である場合、TypeOfはnilを返します。
func TypeOf(i any) Type

// PtrToは、要素tを持つポインタ型を返します。
// 例えば、もしtがFoo型を表すなら、PtrTo(t)は*Fooを表します。
//
// PtrToはPointerToの古い綴りです。
// これらの2つの関数は同じように動作します。
func PtrTo(t Type) Type

// PointerToは要素tを持つポインタ型を返します。
// 例えば、もしtがFoo型を表すなら、PointerTo(t)は*Fooを表します。
func PointerTo(t Type) Type

// ChanOfは指定された方向と要素の型を持つチャネル型を返します。
// たとえば、tがintを表す場合、ChanOf(RecvDir, t)は<-chan intを表します。
//
// gcのランタイムは、チャネルの要素型に64 kBの制限を課しています。
// もしtのサイズがこの制限以上である場合、ChanOfはパニックを発生させます。
func ChanOf(dir ChanDir, t Type) Type

// MapOfは与えられたキーと要素の型のマップタイプを返します。
// 例えば、もしkがintを表し、eがstringを表すならば、MapOf(k, e)はmap[int]stringを表します。
//
// もしキーの型が有効なマップキータイプではない場合（つまり、Goの==演算子を実装していない場合）、MapOfはパニックを起こします。
func MapOf(key, elem Type) Type

// FuncOfは与えられた引数と戻り値の型を持つ関数型を返します。
// 例えば、kがintを表し、eがstringを表す場合、
// FuncOf([]Type{k}, []Type{e}, false)はfunc(int) stringを表します。
//
// 可変引数の引数は、関数が可変引数かどうかを制御します。FuncOfは、
// variadicがtrueであり、in[len(in)-1]がスライスを表していない場合にパニックを起こします。
func FuncOf(in, out []Type, variadic bool) Type

// SliceOfは要素の型がtのスライス型を返します。
// 例えば、tがintを表す場合、SliceOf(t)は[]intを表します。
func SliceOf(t Type) Type

// StructOfはフィールドを含む構造体の型を返します。
// OffsetとIndexのフィールドは無視され、コンパイラによって計算されます。
//
// StructOfは現在、埋め込まれたフィールドに対してラッパーメソッドを生成せず、
// 非公開のStructFieldsが渡された場合はパニックします。
// これらの制限は将来のバージョンで解除される可能性があります。
func StructOf(fields []StructField) Type

// ArrayOfは、与えられた長さと要素の型を持つ配列型を返します。
// 例えば、tがintを表す場合、ArrayOf(5, t)は[5]intを表します。
//
// もし結果の型が利用可能なアドレススペースよりも大きくなる場合、ArrayOfはパニックを発生させます。
func ArrayOf(length int, elem Type) Type
