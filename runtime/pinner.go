// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Pinnerは固定されたGoオブジェクトのセットです。オブジェクトはPinメソッドで固定でき、Pinnerのすべての固定されたオブジェクトはUnpinメソッドで固定を解除できます。
type Pinner struct {
	*pinner
}

// PinはGoのオブジェクトをピン留めし、ガベージコレクターによる移動や解放を防止します。
// Unpinメソッドが呼び出されるまで、ピン留めされたオブジェクトへのポインタは、
// Cメモリに直接格納するか、Goメモリに含めてC関数に渡すことができます。
// ピン留めされたオブジェクト自体がGoオブジェクトへのポインタを持っている場合、
// Cコードからアクセスするためにはそれらのオブジェクトも別途ピン留めする必要があります。
//
<<<<<<< HEAD
// 引数は任意の型のポインタまたはunsafe.Pointerである必要があります。
// newの呼び出し結果、複合リテラルのアドレス、またはローカル変数のアドレスを取る必要があります。
// 上記の条件のいずれかが満たされない場合、Pinはパニックを引き起こします。
=======
// A pointer to a pinned
// object can be directly stored in C memory or can be contained in Go memory
// passed to C functions. If the pinned object itself contains pointers to Go
// objects, these objects must be pinned separately if they are going to be
// accessed from C code.
//
// The argument must be a pointer of any type or an unsafe.Pointer.
>>>>>>> upstream/release-branch.go1.21
func (p *Pinner) Pin(pointer any)

// UnpinはPinnerのすべてのピン留めされたオブジェクトを解除します。
func (p *Pinner) Unpin()
