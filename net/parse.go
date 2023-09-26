// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Simple file i/o and string manipulation, to avoid
// depending on strconv and bufio and strings.

package net

import (
	_ "github.com/shogo82148/std/unsafe"
)

// Bigger than we need, not too big to worry about overflow
