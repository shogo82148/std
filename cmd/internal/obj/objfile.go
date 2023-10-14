// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Writing Go object files.

package obj

import (
	"github.com/shogo82148/std/cmd/internal/bio"
)

const UnlinkablePkg = "<unlinkable>"

// Entry point of writing new object file.
func WriteObjFile(ctxt *Link, b *bio.Writer)
