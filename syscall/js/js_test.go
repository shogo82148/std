// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm
// +build js,wasm

package js_test

import (
	"fmt"
	"syscall/js"
)

func ExampleNewCallback() {
	var cb js.Callback
	cb = js.NewCallback(func(args []js.Value) {
		fmt.Println("button clicked")
		cb.Release() // release the callback if the button will not be clicked again
	})
	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)
}
