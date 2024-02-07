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
// ピン留めされたオブジェクトへのポインタは、Cのメモリに直接保存することも、
// Cの関数に渡されるGoのメモリに含めることもできます。もしピン留めされたオブジェクト自体が
// Goのオブジェクトへのポインタを含んでいる場合、それらのオブジェクトはCのコードからアクセスされる場合に
// 別途ピン留めする必要があります。
//
// 引数は任意の型のポインタまたはunsafe.Pointerでなければなりません。
func (p *Pinner) Pin(pointer any)

// UnpinはPinnerのすべてのピン留めされたオブジェクトを解除します。
func (p *Pinner) Unpin()
