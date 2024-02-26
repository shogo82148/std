// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Pinnerは、メモリ内の固定された場所に各Goオブジェクトが固定されたセットです。
// [Pinner.Pin]メソッドは1つのオブジェクトを固定し、[Pinner.Unpin]メソッドはすべての固定されたオブジェクトを解除します。
// 詳細については、それぞれのコメントを参照してください。
type Pinner struct {
	*pinner
}

// PinはGoオブジェクトをピン留めし、[Pinner.Unpin] メソッドが呼び出されるまで、
// ガベージコレクタによって移動または解放されるのを防ぎます。
//
// 固定されたオブジェクトへのポインタは、Cメモリに直接格納されるか、C関数に渡されるGoメモリに含まれることができます。
// 固定されたオブジェクト自体がGoオブジェクトへのポインタを含む場合、これらのオブジェクトがCコードからアクセスされる場合は、別途固定する必要があります。
//
// 引数は、任意の型のポインタまたは [unsafe.Pointer] である必要があります。
// 非Goポインタに対してPinを呼び出すことは安全であり、その場合、Pinは何もしません。
func (p *Pinner) Pin(pointer any)

// Unpinは [Pinner] のすべてのピン留めされたオブジェクトを解除します。
func (p *Pinner) Unpin()
