// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package reflect implements run-time reflection, allowing a program to
// manipulate objects with arbitrary types.  The typical use is to take a value
// with static type interface{} and extract its dynamic type information by
// calling TypeOf, which returns a Type.
//
// A call to ValueOf returns a Value representing the run-time data.
// Zero takes a Type and returns a Value representing a zero value
// for that type.
//
// See "The Laws of Reflection" for an introduction to reflection in Go:
// http://golang.org/doc/articles/laws_of_reflection.html
package reflect

// Type is the representation of a Go type.
//
// Not all methods apply to all kinds of types.  Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of type before
// calling kind-specific methods.  Calling a method
// inappropriate to the kind of type causes a run-time panic.
type Type interface {
	Align() int

	FieldAlign() int

	Method(int) Method

	MethodByName(string) (Method, bool)

	NumMethod() int

	Name() string

	PkgPath() string

	Size() uintptr

	String() string

	Kind() Kind

	Implements(u Type) bool

	AssignableTo(u Type) bool

	Bits() int

	ChanDir() ChanDir

	IsVariadic() bool

	Elem() Type

	Field(i int) StructField

	FieldByIndex(index []int) StructField

	FieldByName(name string) (StructField, bool)

	FieldByNameFunc(match func(string) bool) (StructField, bool)

	In(i int) Type

	Key() Type

	Len() int

	NumField() int

	NumIn() int

	NumOut() int

	Out(i int) Type

	runtimeType() *runtimeType
	common() *commonType
	uncommon() *uncommonType
}

// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
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
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

// The compiler can only construct empty interface values at
// compile time; non-empty interface values get created
// during initialization.  Type is an empty interface
// so that the compiler can lay out references as data.
// The underlying type is *reflect.ArrayType and so on.

// commonType is the common implementation of most values.
// It is embedded in other, public struct types, but always
// with a unique tag like `reflect:"array"` or `reflect:"ptr"`
// so that code cannot convert from, say, *arrayType to *ptrType.

// Method on non-interface type

// uncommonType is present only for types with names or methods
// (if T is a named type, the uncommonTypes for T and *T have methods).
// Using a pointer to this struct reduces the overall size required
// to describe an unnamed type with no methods.

// ChanDir represents a channel type's direction.
type ChanDir int

const (
	RecvDir ChanDir             = 1 << iota
	SendDir
	BothDir = RecvDir | SendDir
)

// arrayType represents a fixed array type.

// chanType represents a channel type.

// funcType represents a function type.

// imethod represents a method on an interface type

// interfaceType represents an interface type.

// mapType represents a map type.

// ptrType represents a pointer type.

// sliceType represents a slice type.

// Struct field

// structType represents a struct type.

// Method represents a single method.
type Method struct {
	Name    string
	PkgPath string

	Type  Type
	Func  Value
	Index int
}

// High bit says whether type has
// embedded pointers,to help garbage collector.

func (k Kind) String() string

func (d ChanDir) String() string

// A StructField describes a single field in a struct.
type StructField struct {
	Name    string
	PkgPath string

	Type      Type
	Tag       StructTag
	Offset    uintptr
	Index     []int
	Anonymous bool
}

// A StructTag is the tag string in a struct field.
//
// By convention, tag strings are a concatenation of
// optionally space-separated key:"value" pairs.
// Each key is a non-empty string consisting of non-control
// characters other than space (U+0020 ' '), quote (U+0022 '"'),
// and colon (U+003A ':').  Each value is quoted using U+0022 '"'
// characters and Go string literal syntax.
type StructTag string

// Get returns the value associated with key in the tag string.
// If there is no such key in the tag, Get returns the empty string.
// If the tag does not have the conventional format, the value
// returned by Get is unspecified.
func (tag StructTag) Get(key string) string

// TypeOf returns the reflection Type of the value in the interface{}.
// TypeOf(nil) returns nil.
func TypeOf(i interface{}) Type

// ptrMap is the cache for PtrTo.

// PtrTo returns the pointer type with element t.
// For example, if t represents type Foo, PtrTo(t) represents *Foo.
func PtrTo(t Type) Type
