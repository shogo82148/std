// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package runtime

// GOARCH=wasm currently has 64 bits pointers, but the WebAssembly host expects
// pointers to be 32 bits so we use this type alias to represent pointers in
// structs and arrays passed as arguments to WASI functions.
//
// Note that the use of an integer type prevents the compiler from tracking
// pointers passed to WASI functions, so we must use KeepAlive to explicitly
// retain the objects that could otherwise be reclaimed by the GC.

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-size-u32

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-errno-variant

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-filesize-u64

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-timestamp-u64

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-clockid-variant

// https://github.com/WebAssembly/WASI/blob/a2b96e81c0586125cc4dc79a5be0b78d9a059925/legacy/preview1/docs.md#-iovec-record

// The go:wasmimport directive currently does not accept values of type uint16
// in arguments or returns of the function signature. Most WASI imports return
// an errno value, which we have to define as uint32 because of that limitation.
// However, the WASI errno type is intended to be a 16 bits integer, and in the
// event struct the error field should be of type errno. If we used the errno
// type for the error field it would result in a mismatching field alignment and
// struct size because errno is declared as a 32 bits type, so we declare the
// error field as a plain uint16.
