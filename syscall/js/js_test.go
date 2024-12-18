// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

// To run these tests:
//
// - Install Node
// - Add /path/to/go/lib/wasm to your $PATH (so that "go test" can find
//   "go_js_wasm_exec").
// - GOOS=js GOARCH=wasm go test
//
// See -exec in "go help test", and "go help run" for details.

package js_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/syscall/js"
)

func ExampleFuncOf() {
	var cb js.Func
	cb = js.FuncOf(func(this js.Value, args []js.Value) any {
		fmt.Println("button clicked")
		cb.Release() // release the function if the button will not be clicked again
		return nil
	})
	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)
}
