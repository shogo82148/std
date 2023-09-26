// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types_test

import (
	"flag"

	. "go/types"
)

var (
	H = flag.Int("H", 5, "Hilbert matrix size")
)
