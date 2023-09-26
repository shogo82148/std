// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// SHA1 hash algorithm.  See RFC 3174.

package sha1

import (
	"fmt"
	"io"
)

func ExampleNew() {
	h := New()
	io.WriteString(h, "His money is twice tainted: 'taint yours and 'taint mine.")
	fmt.Printf("% x", h.Sum(nil))
	// Output: 59 7f 6a 54 00 10 f9 4c 15 d7 18 06 a9 9a 2c 87 10 e7 47 bd
}
