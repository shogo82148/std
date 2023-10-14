// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cgotest

import "github.com/shogo82148/std/C"

//export Issue42495A
func Issue42495A(C.T42495A)

//export Issue42495B
func Issue42495B(C.T42495B)
