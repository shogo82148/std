// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build asan

package runtime

import (
	"github.com/shogo82148/std/unsafe"
)

// Public address sanitizer API.
func ASanRead(addr unsafe.Pointer, len int)

func ASanWrite(addr unsafe.Pointer, len int)

// Private interface for the runtime.
