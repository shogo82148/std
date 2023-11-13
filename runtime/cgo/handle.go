// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgo

// Handleは、Goで割り当てられたメモリのポインタ（Goのポインタ）を含む値を、
// cgoのポインタ渡し規則を破ることなく、GoとCの間でやり取りするための手段を提供します。
// Handleは、任意のGoの値を表すことができる整数値です。
// Handleは、Cを介してGoに渡すことも、Goのコードで元のGoの値を取得するためにHandleを使用することもできます。
//
// Handleの基になる型は、ポインタのビットパターンを保持できるだけの十分に大きな整数型に収まることが保証されています。
// Handleのゼロ値は無効であり、CのAPIでセンチネルとして安全に使用することができます。
//
// たとえば、Goの側では：
//
//	package main
//
//	/*
//	#include <stdint.h> // for uintptr_t
//
//	extern void MyGoPrint(uintptr_t handle);
//	void myprint(uintptr_t handle);
//	*/
//	import "C"
//	import "runtime/cgo"
//
//	//export MyGoPrint
//	func MyGoPrint(handle C.uintptr_t) {
//		h := cgo.Handle(handle)
//		val := h.Value().(string)
//		println(val)
//		h.Delete()
//	}
//
//	func main() {
//		val := "hello Go"
//		C.myprint(C.uintptr_t(cgo.NewHandle(val)))
//		// Output: hello Go
//	}
//
// Cの側では：
//
//	#include <stdint.h> // for uintptr_t
//
//	// A Go function
//	extern void MyGoPrint(uintptr_t handle);
//
//	// A C function
//	void myprint(uintptr_t handle) {
//	    MyGoPrint(handle);
//	}
//
<<<<<<< HEAD
// 特定のCの関数は、呼び出し元が提供した任意のデータの値を指すvoid*引数を受け入れます。
// cgo.Handle（整数）をGoのunsafe.Pointerに強制変換することは安全ではありませんが、
// 代わりにcgo.Handleのアドレスをvoid*パラメータに渡すことができます。次に示す前の例のバリアントでは、このようにします。
=======
// Some C functions accept a void* argument that points to an arbitrary
// data value supplied by the caller. It is not safe to coerce a [cgo.Handle]
// (an integer) to a Go [unsafe.Pointer], but instead we can pass the address
// of the cgo.Handle to the void* parameter, as in this variant of the
// previous example:
>>>>>>> upstream/master
//
//	package main
//
//	/*
//	extern void MyGoPrint(void *context);
//	static inline void myprint(void *context) {
//	    MyGoPrint(context);
//	}
//	*/
//	import "C"
//	import (
//		"runtime/cgo"
//		"unsafe"
//	)
//
//	//export MyGoPrint
//	func MyGoPrint(context unsafe.Pointer) {
//		h := *(*cgo.Handle)(context)
//		val := h.Value().(string)
//		println(val)
//		h.Delete()
//	}
//
//	func main() {
//		val := "hello Go"
//		h := cgo.NewHandle(val)
//		C.myprint(unsafe.Pointer(&h))
//		// Output: hello Go
//	}
type Handle uintptr

// NewHandleは指定された値のハンドルを返します。
//
// このハンドルはprogramがそれに対してDeleteを呼び出すまで有効です。ハンドルはリソースを使用し、
// このパッケージではCコードがハンドルを保持している可能性があるため、プログラムはハンドルが不要になったら
// 明示的にDeleteを呼び出す必要があります。
//
// この関数の意図された使用方法は、返されたハンドルをCコードに渡し、
// CコードがそれをGoに戻し、GoがValueを呼び出すことです。
func NewHandle(v any) Handle

// Valueは有効なハンドルに関連付けられたGoの値を返します。
//
// ハンドルが無効な場合、このメソッドはパニックを発生させます。
func (h Handle) Value() any

// Deleteはハンドルを無効にします。このメソッドは、プログラムがもはやCにハンドルを渡す必要がなくなり、Cのコードがハンドルの値のコピーを持っていない場合にのみ呼び出すべきです。
//
// ハンドルが無効な場合、このメソッドはパニックを引き起こします。
func (h Handle) Delete()
