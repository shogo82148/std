// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build msan

package runtime

import (
	"github.com/shogo82148/std/unsafe"
)

func MSanRead(addr unsafe.Pointer, len int)

func MSanWrite(addr unsafe.Pointer, len int)
