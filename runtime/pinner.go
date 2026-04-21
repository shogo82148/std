// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

<<<<<<< HEAD
// Pinnerは、メモリ内の固定された場所に各Goオブジェクトが固定されたセットです。
// [Pinner.Pin]メソッドは1つのオブジェクトを固定し、[Pinner.Unpin]メソッドはすべての固定されたオブジェクトを解除します。
// 詳細については、それぞれのコメントを参照してください。
=======
// A Pinner is a set of Go objects each pinned to a fixed location in memory. The
// [Pinner.Pin] method pins one object, while [Pinner.Unpin] unpins all pinned
// objects.
//
// The purpose of a Pinner is two-fold.
// First, it allows C code to safely use Go pointers that have not been passed
// explicitly to the C code via a cgo call.
// For example, for safely interacting with a pointer stored inside of a struct
// whose pointer is passed to a C function.
// Second, it allows C memory to safely retain that Go pointer even after the
// cgo call returns, provided the object remains pinned.
//
// A Pinner arranges for its objects to be automatically unpinned some time after
// it becomes unreachable, so its referents will not leak. However, this means the
// Pinner itself must be kept alive across a cgo call, or as long as C retains a
// reference to the pinned Go pointers.
//
// Reusing a Pinner is safe, and in fact encouraged, to avoid the cost of
// initializing new Pinners on first use.
//
// The zero value of Pinner is ready to use.
>>>>>>> upstream/release-branch.go1.26
type Pinner struct {
	*pinner
}

// PinはGoオブジェクトをピン留めし、[Pinner.Unpin] メソッドが呼び出されるまで、
// ガベージコレクタによって移動または解放されるのを防ぎます。
//
// 固定されたオブジェクトへのポインタは、Cメモリに直接格納されるか、C関数に渡されるGoメモリに含まれることができます。
// 固定されたオブジェクト自体がGoオブジェクトへのポインタを含む場合、これらのオブジェクトがCコードからアクセスされる場合は、別途固定する必要があります。
//
<<<<<<< HEAD
// 引数は、任意の型のポインタまたは [unsafe.Pointer] である必要があります。
// 非Goポインタに対してPinを呼び出すことは安全であり、その場合、Pinは何もしません。
func (p *Pinner) Pin(pointer any)

// Unpinは [Pinner] のすべてのピン留めされたオブジェクトを解除します。
=======
// The argument must be a pointer of any type or an [unsafe.Pointer].
//
// It's safe to call Pin on non-Go pointers, in which case Pin will do nothing.
func (p *Pinner) Pin(pointer any)

// Unpin unpins all pinned objects of the [Pinner].
// It's safe and encouraged to reuse a Pinner after calling Unpin.
>>>>>>> upstream/release-branch.go1.26
func (p *Pinner) Unpin()
