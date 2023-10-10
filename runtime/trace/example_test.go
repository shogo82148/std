// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trace_test

import (
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/runtime/trace"
)

// この例は、trace パッケージを使用して Go プログラムの実行をトレースする方法を示しています。トレースの出力は、ファイル trace.out に書き込まれます。
func Example() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	// ここにあなたのプログラムを書く
	RunMyProgram()
}
