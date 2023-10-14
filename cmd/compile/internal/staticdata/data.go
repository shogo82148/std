// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package staticdata

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

// InitAddrOffset writes the static name symbol lsym to n, it does not modify n.
// It's the caller responsibility to make sure lsym is from ONAME/PEXTERN node.
func InitAddrOffset(n *ir.Name, noff int64, lsym *obj.LSym, off int64)

// InitAddr is InitAddrOffset, with offset fixed to 0.
func InitAddr(n *ir.Name, noff int64, lsym *obj.LSym)

// InitSlice writes a static slice symbol {lsym, lencap, lencap} to n+noff, it does not modify n.
// It's the caller responsibility to make sure lsym is from ONAME node.
func InitSlice(n *ir.Name, noff int64, lsym *obj.LSym, lencap int64)

func InitSliceBytes(nam *ir.Name, off int64, s string)

// StringSym returns a symbol containing the string s.
// The symbol contains the string data, not a string header.
func StringSym(pos src.XPos, s string) (data *obj.LSym)

// FuncLinksym returns n·f, the function value symbol for n.
func FuncLinksym(n *ir.Name) *obj.LSym

func GlobalLinksym(n *ir.Name) *obj.LSym

func WriteFuncSyms()

// InitConst writes the static literal c to n.
// Neither n nor c is modified.
func InitConst(n *ir.Name, noff int64, c ir.Node, wid int)
