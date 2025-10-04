// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Goexitはそれを呼び出したゴルーチンを終了します。他のゴルーチンには影響を与えません。
// Goexitは終了する前にすべての延期呼び出しを実行します。Goexitはパニックではないため、
// これらの延期された関数内のrecover呼び出しはnilを返します。
//
<<<<<<< HEAD
// メインゴルーチンからGoexitを呼び出すと、そのゴルーチンはfunc mainが戻らない状態で終了します。
// func mainが戻っていないため、プログラムは他のゴルーチンの実行を継続します。
// 他のすべてのゴルーチンが終了すると、プログラムはクラッシュします。
=======
// Calling Goexit from the main goroutine terminates that goroutine
// without func main returning. Since func main has not returned,
// the program continues execution of other goroutines.
// If all other goroutines exit, the program crashes.
//
// It crashes if called from a thread not created by the Go runtime.
>>>>>>> upstream/release-branch.go1.25
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
