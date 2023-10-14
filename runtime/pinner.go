// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Pinnerは、メモリ内の固定された場所に各Goオブジェクトが固定されたセットです。
// [Pin]メソッドは1つのオブジェクトを固定し、[Unpin]メソッドはすべての固定されたオブジェクトを解除します。
// 詳細については、それぞれのコメントを参照してください。
type Pinner struct {
	*pinner
}

// PinはGoのオブジェクトをピン留めし、ガベージコレクターによる移動や解放を防止します。
// Unpinメソッドが呼び出されるまで、ピン留めされたオブジェクトへのポインタは、
// Cメモリに直接格納するか、Goメモリに含めてC関数に渡すことができます。
// ピン留めされたオブジェクト自体がGoオブジェクトへのポインタを持っている場合、
// Cコードからアクセスするためにはそれらのオブジェクトも別途ピン留めする必要があります。
//
// 固定されたオブジェクトへのポインタは、Cメモリに直接格納されるか、C関数に渡されるGoメモリに含まれることができます。
// 固定されたオブジェクト自体がGoオブジェクトへのポインタを含む場合、これらのオブジェクトがCコードからアクセスされる場合は、別途固定する必要があります。
//
// 引数は、任意の型のポインタまたはunsafe.Pointerである必要があります。
// 非Goポインタに対してPinを呼び出すことは安全であり、その場合、Pinは何もしません。
func (p *Pinner) Pin(pointer any)

// UnpinはPinnerのすべてのピン留めされたオブジェクトを解除します。
func (p *Pinner) Unpin()
