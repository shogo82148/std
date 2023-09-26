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
// https://golang.org/doc/articles/laws_of_reflection.html
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

	ConvertibleTo(u Type) bool

	Comparable() bool

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

	common() *rtype
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

// rtype is the common implementation of most values.
// It is embedded in other, public struct types, but always
// with a unique tag like `reflect:"array"` or `reflect:"ptr"`
// so that code cannot convert from, say, *arrayType to *ptrType.

// a copy of runtime.typeAlg

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

func (k Kind) String() string

func (d ChanDir) String() string

// A StructField describes a single field in a struct.
type StructField struct {
	Name string

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

// A fieldScan represents an item on the fieldByNameFunc scan work list.

// TypeOf returns the reflection Type that represents the dynamic type of i.
// If i is a nil interface value, TypeOf returns nil.
func TypeOf(i interface{}) Type

// ptrMap is the cache for PtrTo.

// PtrTo returns the pointer type with element t.
// For example, if t represents type Foo, PtrTo(t) represents *Foo.
func PtrTo(t Type) Type

// The lookupCache caches ChanOf, MapOf, and SliceOf lookups.

// A cacheKey is the key for use in the lookupCache.
// Four values describe any of the types we are looking for:
// type kind, one or two subtypes, and an extra integer.

// The funcLookupCache caches FuncOf lookups.
// FuncOf does not share the common lookupCache since cacheKey is not
// sufficient to represent functions unambiguously.

// ChanOf returns the channel type with the given direction and element type.
// For example, if t represents int, ChanOf(RecvDir, t) represents <-chan int.
//
// The gc runtime imposes a limit of 64 kB on channel element types.
// If t's size is equal to or exceeds this limit, ChanOf panics.
func ChanOf(dir ChanDir, t Type) Type

// MapOf returns the map type with the given key and element types.
// For example, if k represents int and e represents string,
// MapOf(k, e) represents map[int]string.
//
// If the key type is not a valid map key type (that is, if it does
// not implement Go's == operator), MapOf panics.
func MapOf(key, elem Type) Type

// FuncOf returns the function type with the given argument and result types.
// For example if k represents int and e represents string,
// FuncOf([]Type{k}, []Type{e}, false) represents func(int) string.
//
// The variadic argument controls whether the function is variadic. FuncOf
// panics if the in[len(in)-1] does not represent a slice and variadic is
// true.
func FuncOf(in, out []Type, variadic bool) Type

// Make sure these routines stay in sync with ../../runtime/hashmap.go!
// These types exist only for GC, so we only fill out GC relevant info.
// Currently, that's just size and the GC program.  We also fill in string
// for possible debugging use.

// SliceOf returns the slice type with element type t.
// For example, if t represents int, SliceOf(t) represents []int.
func SliceOf(t Type) Type

// See cmd/compile/internal/gc/reflect.go for derivation of constant.

// ArrayOf returns the array type with the given count and element type.
// For example, if t represents int, ArrayOf(5, t) represents [5]int.
//
// If the resulting type would be larger than the available address space,
// ArrayOf panics.
func ArrayOf(count int, elem Type) Type

// Layout matches runtime.BitVector (well enough).
