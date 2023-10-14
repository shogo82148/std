// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

func (f *Func) String() string

func StmtString(p src.XPos) string
