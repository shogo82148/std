// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dwarf_test

import (
	. "debug/dwarf"
)

// As Apple converts gcc to a clang-based front end
// they keep breaking the DWARF output. This map lists the
// conversion from real answer to Apple answer.
