// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm
// +build js,wasm

package js

// TypedArray represents a JavaScript typed array.
type TypedArray struct {
	Value
}

// Release frees up resources allocated for the typed array.
// The typed array and its buffer must not be accessed after calling Release.
func (a TypedArray) Release()

// TypedArrayOf returns a JavaScript typed array backed by the slice's underlying array.
//
// The supported types are []int8, []int16, []int32, []uint8, []uint16, []uint32, []float32 and []float64.
// Passing an unsupported value causes a panic.
//
// TypedArray.Release must be called to free up resources when the typed array will not be used any more.
func TypedArrayOf(slice interface{}) TypedArray
