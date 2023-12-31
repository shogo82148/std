// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

// Package js gives access to the WebAssembly host environment when using the js/wasm architecture.
// Its API is based on JavaScript semantics.
//
// This package is EXPERIMENTAL. Its current scope is only to allow tests to run, but not yet to provide a
// comprehensive API for users. It is exempt from the Go compatibility promise.
package js

// Value represents a JavaScript value. The zero value is the JavaScript value "undefined".
// Values can be checked for equality with the Equal method.
type Value struct {
	_     [0]func()
	ref   ref
	gcPtr *ref
}

// Error wraps a JavaScript error.
type Error struct {
	// Value is the underlying JavaScript error value.
	Value
}

// Error implements the error interface.
func (e Error) Error() string

// Equal reports whether v and w are equal according to JavaScript's === operator.
func (v Value) Equal(w Value) bool

// Undefined returns the JavaScript value "undefined".
func Undefined() Value

// IsUndefined reports whether v is the JavaScript value "undefined".
func (v Value) IsUndefined() bool

// Null returns the JavaScript value "null".
func Null() Value

// IsNull reports whether v is the JavaScript value "null".
func (v Value) IsNull() bool

// IsNaN reports whether v is the JavaScript value "NaN".
func (v Value) IsNaN() bool

// Global returns the JavaScript global object, usually "window" or "global".
func Global() Value

// ValueOf returns x as a JavaScript value:
//
//	| Go                     | JavaScript             |
//	| ---------------------- | ---------------------- |
//	| js.Value               | [its value]            |
//	| js.Func                | function               |
//	| nil                    | null                   |
//	| bool                   | boolean                |
//	| integers and floats    | number                 |
//	| string                 | string                 |
//	| []interface{}          | new array              |
//	| map[string]interface{} | new object             |
//
// Panics if x is not one of the expected types.
func ValueOf(x any) Value

// Type represents the JavaScript type of a Value.
type Type int

const (
	TypeUndefined Type = iota
	TypeNull
	TypeBoolean
	TypeNumber
	TypeString
	TypeSymbol
	TypeObject
	TypeFunction
)

func (t Type) String() string

// Type returns the JavaScript type of the value v. It is similar to JavaScript's typeof operator,
// except that it returns TypeNull instead of TypeObject for null.
func (v Value) Type() Type

// Get returns the JavaScript property p of value v.
// It panics if v is not a JavaScript object.
func (v Value) Get(p string) Value

// Set sets the JavaScript property p of value v to ValueOf(x).
// It panics if v is not a JavaScript object.
func (v Value) Set(p string, x any)

// Delete deletes the JavaScript property p of value v.
// It panics if v is not a JavaScript object.
func (v Value) Delete(p string)

// Index returns JavaScript index i of value v.
// It panics if v is not a JavaScript object.
func (v Value) Index(i int) Value

// SetIndex sets the JavaScript index i of value v to ValueOf(x).
// It panics if v is not a JavaScript object.
func (v Value) SetIndex(i int, x any)

// Length returns the JavaScript property "length" of v.
// It panics if v is not a JavaScript object.
func (v Value) Length() int

// Call does a JavaScript call to the method m of value v with the given arguments.
// It panics if v has no method m.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (v Value) Call(m string, args ...any) Value

// Invoke does a JavaScript call of the value v with the given arguments.
// It panics if v is not a JavaScript function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (v Value) Invoke(args ...any) Value

// New uses JavaScript's "new" operator with value v as constructor and the given arguments.
// It panics if v is not a JavaScript function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (v Value) New(args ...any) Value

// Float returns the value v as a float64.
// It panics if v is not a JavaScript number.
func (v Value) Float() float64

// Int returns the value v truncated to an int.
// It panics if v is not a JavaScript number.
func (v Value) Int() int

// Bool returns the value v as a bool.
// It panics if v is not a JavaScript boolean.
func (v Value) Bool() bool

// Truthy returns the JavaScript "truthiness" of the value v. In JavaScript,
// false, 0, "", null, undefined, and NaN are "falsy", and everything else is
// "truthy". See https://developer.mozilla.org/en-US/docs/Glossary/Truthy.
func (v Value) Truthy() bool

// String returns the value v as a string.
// String is a special case because of Go's String method convention. Unlike the other getters,
// it does not panic if v's Type is not TypeString. Instead, it returns a string of the form "<T>"
// or "<T: V>" where T is v's type and V is a string representation of v's value.
func (v Value) String() string

// InstanceOf reports whether v is an instance of type t according to JavaScript's instanceof operator.
func (v Value) InstanceOf(t Value) bool

// A ValueError occurs when a Value method is invoked on
// a Value that does not support it. Such cases are documented
// in the description of each method.
type ValueError struct {
	Method string
	Type   Type
}

func (e *ValueError) Error() string

// CopyBytesToGo copies bytes from src to dst.
// It panics if src is not a Uint8Array or Uint8ClampedArray.
// It returns the number of bytes copied, which will be the minimum of the lengths of src and dst.
func CopyBytesToGo(dst []byte, src Value) int

// CopyBytesToJS copies bytes from src to dst.
// It panics if dst is not a Uint8Array or Uint8ClampedArray.
// It returns the number of bytes copied, which will be the minimum of the lengths of src and dst.
func CopyBytesToJS(dst Value, src []byte) int
