// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm
// +build js,wasm

// Package js gives access to the WebAssembly host environment when using the js/wasm architecture.
// Its API is based on JavaScript semantics.
//
// This package is EXPERIMENTAL. Its current scope is only to allow tests to run, but not yet to provide a
// comprehensive API for users. It is exempt from the Go compatibility promise.
package js

// ref is used to identify a JavaScript value, since the value itself can not be passed to WebAssembly.
// A JavaScript number (64-bit float, except NaN) is represented by its IEEE 754 binary representation.
// All other values are represented as an IEEE 754 binary representation of NaN with bits 0-31 used as
// an ID and bits 32-33 used to differentiate between string, symbol, function and object.

// nanHead are the upper 32 bits of a ref which are set if the value is not a JavaScript number or NaN itself.

// Value represents a JavaScript value.
type Value struct {
	ref ref
}

// Error wraps a JavaScript error.
type Error struct {
	Value
}

// Error implements the error interface.
func (e Error) Error() string

// Undefined returns the JavaScript value "undefined".
func Undefined() Value

// Null returns the JavaScript value "null".
func Null() Value

// Global returns the JavaScript global object, usually "window" or "global".
func Global() Value

// ValueOf returns x as a JavaScript value:
//
//	| Go                     | JavaScript             |
//	| ---------------------- | ---------------------- |
//	| js.Value               | [its value]            |
//	| js.TypedArray          | typed array            |
//	| js.Callback            | function               |
//	| nil                    | null                   |
//	| bool                   | boolean                |
//	| integers and floats    | number                 |
//	| string                 | string                 |
//	| []interface{}          | new array              |
//	| map[string]interface{} | new object             |
func ValueOf(x interface{}) Value

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
func (v Value) Get(p string) Value

// Set sets the JavaScript property p of value v to ValueOf(x).
func (v Value) Set(p string, x interface{})

// Index returns JavaScript index i of value v.
func (v Value) Index(i int) Value

// SetIndex sets the JavaScript index i of value v to ValueOf(x).
func (v Value) SetIndex(i int, x interface{})

// Length returns the JavaScript property "length" of v.
func (v Value) Length() int

// Call does a JavaScript call to the method m of value v with the given arguments.
// It panics if v has no method m.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (v Value) Call(m string, args ...interface{}) Value

// Invoke does a JavaScript call of the value v with the given arguments.
// It panics if v is not a function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (v Value) Invoke(args ...interface{}) Value

// New uses JavaScript's "new" operator with value v as constructor and the given arguments.
// It panics if v is not a function.
// The arguments get mapped to JavaScript values according to the ValueOf function.
func (v Value) New(args ...interface{}) Value

// Float returns the value v as a float64. It panics if v is not a JavaScript number.
func (v Value) Float() float64

// Int returns the value v truncated to an int. It panics if v is not a JavaScript number.
func (v Value) Int() int

// Bool returns the value v as a bool. It panics if v is not a JavaScript boolean.
func (v Value) Bool() bool

// String returns the value v converted to string according to JavaScript type conversions.
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
