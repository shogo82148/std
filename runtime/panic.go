// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Goexitはそれを呼び出したゴルーチンを終了します。他のゴルーチンには影響を与えません。
// Goexitは終了する前にすべての延期呼び出しを実行します。Goexitはパニックではないため、
// これらの延期された関数内のrecover呼び出しはnilを返します。
//
// メインゴルーチンからGoexitを呼び出すと、そのゴルーチンはfunc mainがreturnせずに終了します。
// func mainがreturnしていないため、他のゴルーチンの実行は継続されます。
// 他のすべてのゴルーチンが終了すると、プログラムはクラッシュします。
//
// Goランタイムによって作成されていないスレッドから呼び出すとクラッシュします。
func Goexit()

// PanicNilErrorは、コードがpanic(nil)を呼び出したときに発生します。
//
// Go 1.21より前のバージョンでは、panic(nil)を呼び出すプログラムでは、recoverがnilを返すことが観察されました。
// Go 1.21以降、panic(nil)を呼び出すプログラムでは、recoverが*PanicNilErrorを返すことが観察されます。
// プログラムは、GODEBUG=panicnil=1を設定することで古い動作に戻すことができます。
type PanicNilError struct {

	// このフィールドによって、PanicNilErrorはこのパッケージの他の構造体とは異なる構造を持ちます。_は、他のパッケージの構造体とも異なります。
	// これにより、この構造体と同じフィールドを共有する他の構造体との間で誤って変換が可能になることを防ぎます。go.dev/issue/56603で発生したような事故を回避します。
	_ [0]*PanicNilError
}

func (*PanicNilError) Error() string
func (*PanicNilError) RuntimeError()
