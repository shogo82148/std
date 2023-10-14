// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typecheck

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

func AssignConv(n ir.Node, t *types.Type, context string) ir.Node

// LookupNum returns types.LocalPkg.LookupNum(prefix, n).
func LookupNum(prefix string, n int) *types.Sym

// Given funarg struct list, return list of fn args.
func NewFuncParams(origs []*types.Field) []*types.Field

// NodAddr returns a node representing &n at base.Pos.
func NodAddr(n ir.Node) *ir.AddrExpr

// NodAddrAt returns a node representing &n at position pos.
func NodAddrAt(pos src.XPos, n ir.Node) *ir.AddrExpr

// LinksymAddr returns a new expression that evaluates to the address
// of lsym. typ specifies the type of the addressed memory.
func LinksymAddr(pos src.XPos, lsym *obj.LSym, typ *types.Type) *ir.AddrExpr

func NodNil() ir.Node

// AddImplicitDots finds missing fields in obj.field that
// will give the shortest unique addressing and
// modifies the tree with missing field names.
func AddImplicitDots(n *ir.SelectorExpr) *ir.SelectorExpr

// CalcMethods calculates all the methods (including embedding) of a non-interface
// type t.
func CalcMethods(t *types.Type)

// Implements reports whether t implements the interface iface. t can be
// an interface, a type parameter, or a concrete type.
func Implements(t, iface *types.Type) bool

// ImplementsExplain reports whether t implements the interface iface. t can be
// an interface, a type parameter, or a concrete type. If t does not implement
// iface, a non-empty string is returned explaining why.
func ImplementsExplain(t, iface *types.Type) string
