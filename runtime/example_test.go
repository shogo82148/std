// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/os"
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

<<<<<<< HEAD
			// このフレームを処理します。
			//
			// この例の出力を安定させるために
			// テストパッケージに変更があっても
			// runtimeパッケージを抜けるとアンワインドを停止します。
			if !strings.Contains(frame.File, "runtime/") {
=======
			// Canonicalize function name and skip callers of this function
			// for predictable example output.
			// You probably don't need this in your own code.
			function := strings.ReplaceAll(frame.Function, "main.main", "runtime_test.ExampleFrames")
			fmt.Printf("- more:%v | %s\n", more, function)
			if function == "runtime_test.ExampleFrames" {
>>>>>>> upstream/release-branch.go1.25
				break
			}

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

func ExampleAddCleanup() {
	tempFile, err := os.CreateTemp(os.TempDir(), "file.*")
	if err != nil {
		fmt.Println("failed to create temp file:", err)
		return
	}

	ch := make(chan struct{})

	// Attach a cleanup function to the file object.
	runtime.AddCleanup(&tempFile, func(fileName string) {
		if err := os.Remove(fileName); err == nil {
			fmt.Println("temp file has been removed")
		}
		ch <- struct{}{}
	}, tempFile.Name())

	if err := tempFile.Close(); err != nil {
		fmt.Println("failed to close temp file:", err)
		return
	}

	// Run the garbage collector to reclaim unreachable objects
	// and enqueue their cleanup functions.
	runtime.GC()

	// Wait until cleanup function is done.
	<-ch

	// Output:
	// temp file has been removed
}
