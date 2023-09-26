// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// A Type represents a type of Go.
// All types implement the Type interface.
type Type interface {
	Underlying() Type

	String() string
}

// BasicKind describes the kind of basic type.
type BasicKind int

const (
	Invalid BasicKind = iota

	// predeclared types
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

	// types for untyped values
	UntypedBool
	UntypedInt
	UntypedRune
	UntypedFloat
	UntypedComplex
	UntypedString
	UntypedNil

	// aliases
	Byte = Uint8
	Rune = Int32
)

// BasicInfo is a set of flags describing properties of a basic type.
type BasicInfo int

// Properties of basic types.
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

// A Basic represents a basic type.
type Basic struct {
	kind BasicKind
	info BasicInfo
	name string
}

// Kind returns the kind of basic type b.
func (b *Basic) Kind() BasicKind

// Info returns information about properties of basic type b.
func (b *Basic) Info() BasicInfo

// Name returns the name of basic type b.
func (b *Basic) Name() string

// An Array represents an array type.
type Array struct {
	len  int64
	elem Type
}

// NewArray returns a new array type for the given element type and length.
// A negative length indicates an unknown length.
func NewArray(elem Type, len int64) *Array

// Len returns the length of array a.
// A negative result indicates an unknown length.
func (a *Array) Len() int64

// Elem returns element type of array a.
func (a *Array) Elem() Type

// A Slice represents a slice type.
type Slice struct {
	elem Type
}

// NewSlice returns a new slice type for the given element type.
func NewSlice(elem Type) *Slice

// Elem returns the element type of slice s.
func (s *Slice) Elem() Type

// A Struct represents a struct type.
type Struct struct {
	fields []*Var
	tags   []string
}

// NewStruct returns a new struct with the given fields and corresponding field tags.
// If a field with index i has a tag, tags[i] must be that tag, but len(tags) may be
// only as long as required to hold the tag with the largest index i. Consequently,
// if no field has a tag, tags may be nil.
func NewStruct(fields []*Var, tags []string) *Struct

// NumFields returns the number of fields in the struct (including blank and embedded fields).
func (s *Struct) NumFields() int

// Field returns the i'th field for 0 <= i < NumFields().
func (s *Struct) Field(i int) *Var

// Tag returns the i'th field tag for 0 <= i < NumFields().
func (s *Struct) Tag(i int) string

// A Pointer represents a pointer type.
type Pointer struct {
	base Type
}

// NewPointer returns a new pointer type for the given element (base) type.
func NewPointer(elem Type) *Pointer

// Elem returns the element type for the given pointer p.
func (p *Pointer) Elem() Type

// A Tuple represents an ordered list of variables; a nil *Tuple is a valid (empty) tuple.
// Tuples are used as components of signatures and to represent the type of multiple
// assignments; they are not first class types of Go.
type Tuple struct {
	vars []*Var
}

// NewTuple returns a new tuple for the given variables.
func NewTuple(x ...*Var) *Tuple

// Len returns the number variables of tuple t.
func (t *Tuple) Len() int

// At returns the i'th variable of tuple t.
func (t *Tuple) At(i int) *Var

// A Signature represents a (non-builtin) function or method type.
// The receiver is ignored when comparing signatures for identity.
type Signature struct {
	scope    *Scope
	recv     *Var
	params   *Tuple
	results  *Tuple
	variadic bool
}

// NewSignature returns a new function type for the given receiver, parameters,
// and results, either of which may be nil. If variadic is set, the function
// is variadic, it must have at least one parameter, and the last parameter
// must be of unnamed slice type.
func NewSignature(recv *Var, params, results *Tuple, variadic bool) *Signature

// Recv returns the receiver of signature s (if a method), or nil if a
// function. It is ignored when comparing signatures for identity.
//
// For an abstract method, Recv returns the enclosing interface either
// as a *Named or an *Interface. Due to embedding, an interface may
// contain methods whose receiver type is a different interface.
func (s *Signature) Recv() *Var

// Params returns the parameters of signature s, or nil.
func (s *Signature) Params() *Tuple

// Results returns the results of signature s, or nil.
func (s *Signature) Results() *Tuple

// Variadic reports whether the signature s is variadic.
func (s *Signature) Variadic() bool

// An Interface represents an interface type.
type Interface struct {
	methods   []*Func
	embeddeds []Type

	allMethods []*Func
}

// emptyInterface represents the empty (completed) interface

// markComplete is used to mark an empty interface as completely
// set up by setting the allMethods field to a non-nil empty slice.

// NewInterface returns a new (incomplete) interface for the given methods and embedded types.
// Each embedded type must have an underlying type of interface type.
// NewInterface takes ownership of the provided methods and may modify their types by setting
// missing receivers. To compute the method set of the interface, Complete must be called.
//
// Deprecated: Use NewInterfaceType instead which allows any (even non-defined) interface types
// to be embedded. This is necessary for interfaces that embed alias type names referring to
// non-defined (literal) interface types.
func NewInterface(methods []*Func, embeddeds []*Named) *Interface

// NewInterfaceType returns a new (incomplete) interface for the given methods and embedded types.
// Each embedded type must have an underlying type of interface type (this property is not
// verified for defined types, which may be in the process of being set up and which don't
// have a valid underlying type yet).
// NewInterfaceType takes ownership of the provided methods and may modify their types by setting
// missing receivers. To compute the method set of the interface, Complete must be called.
func NewInterfaceType(methods []*Func, embeddeds []Type) *Interface

// NumExplicitMethods returns the number of explicitly declared methods of interface t.
func (t *Interface) NumExplicitMethods() int

// ExplicitMethod returns the i'th explicitly declared method of interface t for 0 <= i < t.NumExplicitMethods().
// The methods are ordered by their unique Id.
func (t *Interface) ExplicitMethod(i int) *Func

// NumEmbeddeds returns the number of embedded types in interface t.
func (t *Interface) NumEmbeddeds() int

// Embedded returns the i'th embedded defined (*Named) type of interface t for 0 <= i < t.NumEmbeddeds().
// The result is nil if the i'th embedded type is not a defined type.
//
// Deprecated: Use EmbeddedType which is not restricted to defined (*Named) types.
func (t *Interface) Embedded(i int) *Named

// EmbeddedType returns the i'th embedded type of interface t for 0 <= i < t.NumEmbeddeds().
func (t *Interface) EmbeddedType(i int) Type

// NumMethods returns the total number of methods of interface t.
func (t *Interface) NumMethods() int

// Method returns the i'th method of interface t for 0 <= i < t.NumMethods().
// The methods are ordered by their unique Id.
func (t *Interface) Method(i int) *Func

// Empty returns true if t is the empty interface.
func (t *Interface) Empty() bool

// Complete computes the interface's method set. It must be called by users of
// NewInterfaceType and NewInterface after the interface's embedded types are
// fully defined and before using the interface type in any way other than to
// form other types. Complete returns the receiver.
func (t *Interface) Complete() *Interface

// A Map represents a map type.
type Map struct {
	key, elem Type
}

// NewMap returns a new map for the given key and element types.
func NewMap(key, elem Type) *Map

// Key returns the key type of map m.
func (m *Map) Key() Type

// Elem returns the element type of map m.
func (m *Map) Elem() Type

// A Chan represents a channel type.
type Chan struct {
	dir  ChanDir
	elem Type
}

// A ChanDir value indicates a channel direction.
type ChanDir int

// The direction of a channel is indicated by one of these constants.
const (
	SendRecv ChanDir = iota
	SendOnly
	RecvOnly
)

// NewChan returns a new channel type for the given direction and element type.
func NewChan(dir ChanDir, elem Type) *Chan

// Dir returns the direction of channel c.
func (c *Chan) Dir() ChanDir

// Elem returns the element type of channel c.
func (c *Chan) Elem() Type

// A Named represents a named type.
type Named struct {
	obj        *TypeName
	underlying Type
	methods    []*Func
}

// NewNamed returns a new named type for the given type name, underlying type, and associated methods.
// If the given type name obj doesn't have a type yet, its type is set to the returned named type.
// The underlying type must not be a *Named.
func NewNamed(obj *TypeName, underlying Type, methods []*Func) *Named

// Obj returns the type name for the named type t.
func (t *Named) Obj() *TypeName

// NumMethods returns the number of explicit methods whose receiver is named type t.
func (t *Named) NumMethods() int

// Method returns the i'th method of named type t for 0 <= i < t.NumMethods().
func (t *Named) Method(i int) *Func

// SetUnderlying sets the underlying type and marks t as complete.
func (t *Named) SetUnderlying(underlying Type)

// AddMethod adds method m unless it is already in the method list.
func (t *Named) AddMethod(m *Func)

func (b *Basic) Underlying() Type
func (a *Array) Underlying() Type
func (s *Slice) Underlying() Type
func (s *Struct) Underlying() Type
func (p *Pointer) Underlying() Type
func (t *Tuple) Underlying() Type
func (s *Signature) Underlying() Type
func (t *Interface) Underlying() Type
func (m *Map) Underlying() Type
func (c *Chan) Underlying() Type
func (t *Named) Underlying() Type

func (b *Basic) String() string
func (a *Array) String() string
func (s *Slice) String() string
func (s *Struct) String() string
func (p *Pointer) String() string
func (t *Tuple) String() string
func (s *Signature) String() string
func (t *Interface) String() string
func (m *Map) String() string
func (c *Chan) String() string
func (t *Named) String() string
