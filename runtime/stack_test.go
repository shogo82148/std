// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	. "runtime"
	_ "unsafe"
)

type I interface {
	M()
}

// Pass a value to escapeMe to force it to escape.
