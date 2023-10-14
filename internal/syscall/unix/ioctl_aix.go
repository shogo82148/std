// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unix

import (
	"github.com/shogo82148/std/unsafe"
)

func Ioctl(fd int, cmd int, args unsafe.Pointer) (err error)
