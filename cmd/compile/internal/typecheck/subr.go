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
func NewFuncParams(tl *types.Type, mustname bool) []*ir.Field

// NewName returns a new ONAME Node associated with symbol s.
func NewName(s *types.Sym) *ir.Name

// NodAddr returns a node representing &n at base.Pos.
func NodAddr(n ir.Node) *ir.AddrExpr

// NodAddrAt returns a node representing &n at position pos.
func NodAddrAt(pos src.XPos, n ir.Node) *ir.AddrExpr

// If IncrementalAddrtaken is false, we do not compute Addrtaken for an OADDR Node
// when it is built. The Addrtaken bits are set in bulk by computeAddrtaken.
// If IncrementalAddrtaken is true, then when an OADDR Node is built the Addrtaken
// field of its argument is updated immediately.
var IncrementalAddrtaken = false

// If DirtyAddrtaken is true, then there are OADDR whose corresponding arguments
// have not yet been marked as Addrtaken.
var DirtyAddrtaken = false

func ComputeAddrtaken(top []ir.Node)

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

// Is type src assignment compatible to type dst?
// If so, return op code to use in conversion.
// If not, return OXXX. In this case, the string return parameter may
// hold a reason why. In all other cases, it'll be the empty string.
func Assignop(src, dst *types.Type) (ir.Op, string)

func Assignop1(src, dst *types.Type) (ir.Op, string)

// Can we convert a value of type src to a value of type dst?
// If so, return op code to use in conversion (maybe OCONVNOP).
// If not, return OXXX. In this case, the string return parameter may
// hold a reason why. In all other cases, it'll be the empty string.
// srcConstant indicates whether the value of type src is a constant.
func Convertop(srcConstant bool, src, dst *types.Type) (ir.Op, string)

// Implements reports whether t implements the interface iface. t can be
// an interface, a type parameter, or a concrete type.
func Implements(t, iface *types.Type) bool

// ImplementsExplain reports whether t implements the interface iface. t can be
// an interface, a type parameter, or a concrete type. If t does not implement
// iface, a non-empty string is returned explaining why.
func ImplementsExplain(t, iface *types.Type) string
