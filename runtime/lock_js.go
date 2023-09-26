// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package runtime

import (
	_ "github.com/shogo82148/std/unsafe"
)

// events is a stack of calls from JavaScript into Go.

// The timeout event started by beforeIdle.
