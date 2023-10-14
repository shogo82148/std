// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkginit

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// Calculate redzone for globals.
func GetRedzoneSizeForGlobal(size int64) int64

// InstrumentGlobalsMap contains only package-local (and unlinknamed from somewhere else)
// globals.
// And the key is the object name. For example, in package p, a global foo would be in this
// map as "foo".
// Consider range over maps is nondeterministic, make a slice to hold all the values in the
// InstrumentGlobalsMap and iterate over the InstrumentGlobalsSlice.
var InstrumentGlobalsMap = make(map[string]ir.Node)
var InstrumentGlobalsSlice = make([]ir.Node, 0, 0)
