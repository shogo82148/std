// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect

import (
	"github.com/shogo82148/std/unsafe"
)

// Value is the reflection interface to a Go value.
//
// Not all methods apply to all kinds of values.  Restrictions,
// if any, are noted in the documentation for each method.
// Use the Kind method to find out the kind of value before
// calling kind-specific methods.  Calling a method
// inappropriate to the kind of type causes a run time panic.
//
// The zero Value represents no value.
// Its IsValid method returns false, its Kind method returns Invalid,
// its String method returns "<invalid Value>", and all other methods panic.
// Most functions and methods never return an invalid value.
// If one does, its documentation states the conditions explicitly.
//
// A Value can be used concurrently by multiple goroutines provided that
// the underlying Go value can be used concurrently for the equivalent
// direct operations.
type Value struct {
	typ *commonType

	val unsafe.Pointer

	flag
}

// A ValueError occurs when a Value method is invoked on
// a Value that does not support it.  Such cases are documented
// in the description of each method.
type ValueError struct {
	Method string
	Kind   Kind
}

func (e *ValueError) Error() string

// An iword is the word that would be stored in an
// interface to represent a given value v.  Specifically, if v is
// bigger than a pointer, its word is a pointer to v's data.
// Otherwise, its word holds the data stored
// in its leading bytes (so is not a pointer).
// Because the value sometimes holds a pointer, we use
// unsafe.Pointer to represent it, so that if iword appears
// in a struct, the garbage collector knows that might be
// a pointer.

// emptyInterface is the header for an interface{} value.

// nonEmptyInterface is the header for a interface value with methods.

// Addr returns a pointer value representing the address of v.
// It panics if CanAddr() returns false.
// Addr is typically used to obtain a pointer to a struct field
// or slice element in order to call a method that requires a
// pointer receiver.
func (v Value) Addr() Value

// Bool returns v's underlying value.
// It panics if v's kind is not Bool.
func (v Value) Bool() bool

// Bytes returns v's underlying value.
// It panics if v's underlying value is not a slice of bytes.
func (v Value) Bytes() []byte

// CanAddr returns true if the value's address can be obtained with Addr.
// Such values are called addressable.  A value is addressable if it is
// an element of a slice, an element of an addressable array,
// a field of an addressable struct, or the result of dereferencing a pointer.
// If CanAddr returns false, calling Addr will panic.
func (v Value) CanAddr() bool

// CanSet returns true if the value of v can be changed.
// A Value can be changed only if it is addressable and was not
// obtained by the use of unexported struct fields.
// If CanSet returns false, calling Set or any type-specific
// setter (e.g., SetBool, SetInt64) will panic.
func (v Value) CanSet() bool

// Call calls the function v with the input arguments in.
// For example, if len(in) == 3, v.Call(in) represents the Go call v(in[0], in[1], in[2]).
// Call panics if v's Kind is not Func.
// It returns the output results as Values.
// As in Go, each input argument must be assignable to the
// type of the function's corresponding input parameter.
// If v is a variadic function, Call creates the variadic slice parameter
// itself, copying in the corresponding values.
func (v Value) Call(in []Value) []Value

// CallSlice calls the variadic function v with the input arguments in,
// assigning the slice in[len(in)-1] to v's final variadic argument.
// For example, if len(in) == 3, v.Call(in) represents the Go call v(in[0], in[1], in[2]...).
// Call panics if v's Kind is not Func or if v is not variadic.
// It returns the output results as Values.
// As in Go, each input argument must be assignable to the
// type of the function's corresponding input parameter.
func (v Value) CallSlice(in []Value) []Value

// Cap returns v's capacity.
// It panics if v's Kind is not Array, Chan, or Slice.
func (v Value) Cap() int

// Close closes the channel v.
// It panics if v's Kind is not Chan.
func (v Value) Close()

// Complex returns v's underlying value, as a complex128.
// It panics if v's Kind is not Complex64 or Complex128
func (v Value) Complex() complex128

// Elem returns the value that the interface v contains
// or that the pointer v points to.
// It panics if v's Kind is not Interface or Ptr.
// It returns the zero Value if v is nil.
func (v Value) Elem() Value

// Field returns the i'th field of the struct v.
// It panics if v's Kind is not Struct or i is out of range.
func (v Value) Field(i int) Value

// FieldByIndex returns the nested field corresponding to index.
// It panics if v's Kind is not struct.
func (v Value) FieldByIndex(index []int) Value

// FieldByName returns the struct field with the given name.
// It returns the zero Value if no field was found.
// It panics if v's Kind is not struct.
func (v Value) FieldByName(name string) Value

// FieldByNameFunc returns the struct field with a name
// that satisfies the match function.
// It panics if v's Kind is not struct.
// It returns the zero Value if no field was found.
func (v Value) FieldByNameFunc(match func(string) bool) Value

// Float returns v's underlying value, as a float64.
// It panics if v's Kind is not Float32 or Float64
func (v Value) Float() float64

// Index returns v's i'th element.
// It panics if v's Kind is not Array or Slice or i is out of range.
func (v Value) Index(i int) Value

// Int returns v's underlying value, as an int64.
// It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64.
func (v Value) Int() int64

// CanInterface returns true if Interface can be used without panicking.
func (v Value) CanInterface() bool

// Interface returns v's current value as an interface{}.
// It is equivalent to:
//
//	var i interface{} = (v's underlying value)
//
// If v is a method obtained by invoking Value.Method
// (as opposed to Type.Method), Interface cannot return an
// interface value, so it panics.
// It also panics if the Value was obtained by accessing
// unexported struct fields.
func (v Value) Interface() (i interface{})

// InterfaceData returns the interface v's value as a uintptr pair.
// It panics if v's Kind is not Interface.
func (v Value) InterfaceData() [2]uintptr

// IsNil returns true if v is a nil value.
// It panics if v's Kind is not Chan, Func, Interface, Map, Ptr, or Slice.
func (v Value) IsNil() bool

// IsValid returns true if v represents a value.
// It returns false if v is the zero Value.
// If IsValid returns false, all other methods except String panic.
// Most functions and methods never return an invalid value.
// If one does, its documentation states the conditions explicitly.
func (v Value) IsValid() bool

// Kind returns v's Kind.
// If v is the zero Value (IsValid returns false), Kind returns Invalid.
func (v Value) Kind() Kind

// Len returns v's length.
// It panics if v's Kind is not Array, Chan, Map, Slice, or String.
func (v Value) Len() int

// MapIndex returns the value associated with key in the map v.
// It panics if v's Kind is not Map.
// It returns the zero Value if key is not found in the map or if v represents a nil map.
// As in Go, the key's value must be assignable to the map's key type.
func (v Value) MapIndex(key Value) Value

// MapKeys returns a slice containing all the keys present in the map,
// in unspecified order.
// It panics if v's Kind is not Map.
// It returns an empty slice if v represents a nil map.
func (v Value) MapKeys() []Value

// Method returns a function value corresponding to v's i'th method.
// The arguments to a Call on the returned function should not include
// a receiver; the returned function will always use v as the receiver.
// Method panics if i is out of range.
func (v Value) Method(i int) Value

// NumMethod returns the number of methods in the value's method set.
func (v Value) NumMethod() int

// MethodByName returns a function value corresponding to the method
// of v with the given name.
// The arguments to a Call on the returned function should not include
// a receiver; the returned function will always use v as the receiver.
// It returns the zero Value if no method was found.
func (v Value) MethodByName(name string) Value

// NumField returns the number of fields in the struct v.
// It panics if v's Kind is not Struct.
func (v Value) NumField() int

// OverflowComplex returns true if the complex128 x cannot be represented by v's type.
// It panics if v's Kind is not Complex64 or Complex128.
func (v Value) OverflowComplex(x complex128) bool

// OverflowFloat returns true if the float64 x cannot be represented by v's type.
// It panics if v's Kind is not Float32 or Float64.
func (v Value) OverflowFloat(x float64) bool

// OverflowInt returns true if the int64 x cannot be represented by v's type.
// It panics if v's Kind is not Int, Int8, int16, Int32, or Int64.
func (v Value) OverflowInt(x int64) bool

// OverflowUint returns true if the uint64 x cannot be represented by v's type.
// It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.
func (v Value) OverflowUint(x uint64) bool

// Pointer returns v's value as a uintptr.
// It returns uintptr instead of unsafe.Pointer so that
// code using reflect cannot obtain unsafe.Pointers
// without importing the unsafe package explicitly.
// It panics if v's Kind is not Chan, Func, Map, Ptr, Slice, or UnsafePointer.
func (v Value) Pointer() uintptr

// Recv receives and returns a value from the channel v.
// It panics if v's Kind is not Chan.
// The receive blocks until a value is ready.
// The boolean value ok is true if the value x corresponds to a send
// on the channel, false if it is a zero value received because the channel is closed.
func (v Value) Recv() (x Value, ok bool)

// Send sends x on the channel v.
// It panics if v's kind is not Chan or if x's type is not the same type as v's element type.
// As in Go, x's value must be assignable to the channel's element type.
func (v Value) Send(x Value)

// Set assigns x to the value v.
// It panics if CanSet returns false.
// As in Go, x's value must be assignable to v's type.
func (v Value) Set(x Value)

// SetBool sets v's underlying value.
// It panics if v's Kind is not Bool or if CanSet() is false.
func (v Value) SetBool(x bool)

// SetBytes sets v's underlying value.
// It panics if v's underlying value is not a slice of bytes.
func (v Value) SetBytes(x []byte)

// SetComplex sets v's underlying value to x.
// It panics if v's Kind is not Complex64 or Complex128, or if CanSet() is false.
func (v Value) SetComplex(x complex128)

// SetFloat sets v's underlying value to x.
// It panics if v's Kind is not Float32 or Float64, or if CanSet() is false.
func (v Value) SetFloat(x float64)

// SetInt sets v's underlying value to x.
// It panics if v's Kind is not Int, Int8, Int16, Int32, or Int64, or if CanSet() is false.
func (v Value) SetInt(x int64)

// SetLen sets v's length to n.
// It panics if v's Kind is not Slice or if n is negative or
// greater than the capacity of the slice.
func (v Value) SetLen(n int)

// SetMapIndex sets the value associated with key in the map v to val.
// It panics if v's Kind is not Map.
// If val is the zero Value, SetMapIndex deletes the key from the map.
// As in Go, key's value must be assignable to the map's key type,
// and val's value must be assignable to the map's value type.
func (v Value) SetMapIndex(key, val Value)

// SetUint sets v's underlying value to x.
// It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64, or if CanSet() is false.
func (v Value) SetUint(x uint64)

// SetPointer sets the unsafe.Pointer value v to x.
// It panics if v's Kind is not UnsafePointer.
func (v Value) SetPointer(x unsafe.Pointer)

// SetString sets v's underlying value to x.
// It panics if v's Kind is not String or if CanSet() is false.
func (v Value) SetString(x string)

// Slice returns a slice of v.
// It panics if v's Kind is not Array or Slice.
func (v Value) Slice(beg, end int) Value

// String returns the string v's underlying value, as a string.
// String is a special case because of Go's String method convention.
// Unlike the other getters, it does not panic if v's Kind is not String.
// Instead, it returns a string of the form "<T value>" where T is v's type.
func (v Value) String() string

// TryRecv attempts to receive a value from the channel v but will not block.
// It panics if v's Kind is not Chan.
// If the receive cannot finish without blocking, x is the zero Value.
// The boolean ok is true if the value x corresponds to a send
// on the channel, false if it is a zero value received because the channel is closed.
func (v Value) TryRecv() (x Value, ok bool)

// TrySend attempts to send x on the channel v but will not block.
// It panics if v's Kind is not Chan.
// It returns true if the value was sent, false otherwise.
// As in Go, x's value must be assignable to the channel's element type.
func (v Value) TrySend(x Value) bool

// Type returns v's type.
func (v Value) Type() Type

// Uint returns v's underlying value, as a uint64.
// It panics if v's Kind is not Uint, Uintptr, Uint8, Uint16, Uint32, or Uint64.
func (v Value) Uint() uint64

// UnsafeAddr returns a pointer to v's data.
// It is for advanced clients that also import the "unsafe" package.
// It panics if v is not addressable.
func (v Value) UnsafeAddr() uintptr

// StringHeader is the runtime representation of a string.
// It cannot be used safely or portably.
type StringHeader struct {
	Data uintptr
	Len  int
}

// SliceHeader is the runtime representation of a slice.
// It cannot be used safely or portably.
type SliceHeader struct {
	Data uintptr
	Len  int
	Cap  int
}

// Append appends the values x to a slice s and returns the resulting slice.
// As in Go, each x's value must be assignable to the slice's element type.
func Append(s Value, x ...Value) Value

// AppendSlice appends a slice t to a slice s and returns the resulting slice.
// The slices s and t must have the same element type.
func AppendSlice(s, t Value) Value

// Copy copies the contents of src into dst until either
// dst has been filled or src has been exhausted.
// It returns the number of elements copied.
// Dst and src each must have kind Slice or Array, and
// dst and src must have the same element type.
func Copy(dst, src Value) int

// MakeSlice creates a new zero-initialized slice value
// for the specified slice type, length, and capacity.
func MakeSlice(typ Type, len, cap int) Value

// MakeChan creates a new channel with the specified type and buffer size.
func MakeChan(typ Type, buffer int) Value

// MakeMap creates a new map of the specified type.
func MakeMap(typ Type) Value

// Indirect returns the value that v points to.
// If v is a nil pointer, Indirect returns a zero Value.
// If v is not a pointer, Indirect returns v.
func Indirect(v Value) Value

// ValueOf returns a new Value initialized to the concrete value
// stored in the interface i.  ValueOf(nil) returns the zero Value.
func ValueOf(i interface{}) Value

// Zero returns a Value representing the zero value for the specified type.
// The result is different from the zero value of the Value struct,
// which represents no value at all.
// For example, Zero(TypeOf(42)) returns a Value with Kind Int and value 0.
// The returned value is neither addressable nor settable.
func Zero(typ Type) Value

// New returns a Value representing a pointer to a new zero value
// for the specified type.  That is, the returned Value's Type is PtrTo(t).
func New(typ Type) Value

// NewAt returns a Value representing a pointer to a value of the
// specified type, using p as that pointer.
func NewAt(typ Type, p unsafe.Pointer) Value
