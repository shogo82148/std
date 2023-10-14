// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x86

import (
	"github.com/shogo82148/std/cmd/internal/obj"
)

// EncodeRegisterRange packs [reg0-reg1] list into 64-bit value that
// is intended to be stored inside obj.Addr.Offset with TYPE_REGLIST.
func EncodeRegisterRange(reg0, reg1 int16) int64

// ParseSuffix handles the special suffix for the 386/AMD64.
// Suffix bits are stored into p.Scond.
//
// Leading "." in cond is ignored.
func ParseSuffix(p *obj.Prog, cond string) error
