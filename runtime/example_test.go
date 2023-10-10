// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/runtime"
	"github.com/shogo82148/std/strings"
)

func ExampleFrames() {
	c := func() {
		// runtime.Callersを含めて最大10個のPCを要求します。
		pc := make([]uintptr, 10)
		n := runtime.Callers(0, pc)
		if n == 0 {
			// PC（プログラムカウンタ）が利用できません。これは、runtime.Callersの最初の引数が大きい場合に発生する可能性があります。
			//
			// frames.Next以下で返されるはずのゼロのフレームを処理しないため、ここでリターンします。
			return
		}

		pc = pc[:n] // runtime.CallersFramesには有効なプログラムカウンタ（pcs）のみを渡してください。
		frames := runtime.CallersFrames(pc)

		// フレームを取得するためのループ。
		// 固定数のPCが無限のフレームに拡張できます。
		for {
			frame, more := frames.Next()

			// このフレームを処理します。
			//
			// この例の出力を安定させるために
			// テストパッケージに変更があっても
			// runtimeパッケージを抜けるとアンワインドを停止します。
			if !strings.Contains(frame.File, "runtime/") {
				break
			}
			fmt.Printf("- more:%v | %s\n", more, frame.Function)

			// このフレームの処理後にさらにフレームがあるかどうかを確認します。
			if !more {
				break
			}
		}
	}

	b := func() { c() }
	a := func() { b() }

	a()
	// Output:
	// - more:true | runtime.Callers
	// - more:true | runtime_test.ExampleFrames.func1
	// - more:true | runtime_test.ExampleFrames.func2
	// - more:true | runtime_test.ExampleFrames.func3
	// - more:true | runtime_test.ExampleFrames
}
