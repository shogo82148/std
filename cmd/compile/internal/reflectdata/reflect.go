// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflectdata

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

// TrackSym returns the symbol for tracking use of field/method f, assumed
// to be a member of struct/interface type t.
func TrackSym(t *types.Type, f *types.Field) *obj.LSym

func TypeSymPrefix(prefix string, t *types.Type) *types.Sym

func TypeSym(t *types.Type) *types.Sym

func TypeLinksymPrefix(prefix string, t *types.Type) *obj.LSym

func TypeLinksymLookup(name string) *obj.LSym

func TypeLinksym(t *types.Type) *obj.LSym

// TypePtrAt returns an expression that evaluates to the
// *runtime._type value for t.
func TypePtrAt(pos src.XPos, t *types.Type) *ir.AddrExpr

// ITabLsym returns the LSym representing the itab for concrete type typ implementing
// interface iface. A dummy tab will be created in the unusual case where typ doesn't
// implement iface. Normally, this wouldn't happen, because the typechecker would
// have reported a compile-time error. This situation can only happen when the
// destination type of a type assert or a type in a type switch is parameterized, so
// it may sometimes, but not always, be a type that can't implement the specified
// interface.
func ITabLsym(typ, iface *types.Type) *obj.LSym

// ITabAddrAt returns an expression that evaluates to the
// *runtime.itab value for concrete type typ implementing interface
// iface.
func ITabAddrAt(pos src.XPos, typ, iface *types.Type) *ir.AddrExpr

// InterfaceMethodOffset returns the offset of the i-th method in the interface
// type descriptor, ityp.
func InterfaceMethodOffset(ityp *types.Type, i int64) int64

// NeedRuntimeType ensures that a runtime type descriptor is emitted for t.
func NeedRuntimeType(t *types.Type)

func WriteRuntimeTypes()

func WriteGCSymbols()

func WritePluginTable()

func WriteBasicTypes()

// GCSym returns a data symbol containing GC information for type t, along
// with a boolean reporting whether the UseGCProg bit should be set in the
// type kind, and the ptrdata field to record in the reflect type information.
// GCSym may be called in concurrent backend, so it does not emit the symbol
// content.
func GCSym(t *types.Type) (lsym *obj.LSym, useGCProg bool, ptrdata int64)

// ZeroAddr returns the address of a symbol with at least
// size bytes of zeros.
func ZeroAddr(size int64) ir.Node

// NeedEmit reports whether typ is a type that we need to emit code
// for (e.g., runtime type descriptors, method wrappers).
func NeedEmit(typ *types.Type) bool

var ZeroSize int64

// MarkTypeUsedInInterface marks that type t is converted to an interface.
// This information is used in the linker in dead method elimination.
func MarkTypeUsedInInterface(t *types.Type, from *obj.LSym)

func MarkTypeSymUsedInInterface(tsym *obj.LSym, from *obj.LSym)

// MarkUsedIfaceMethod marks that an interface method is used in the current
// function. n is OCALLINTER node.
func MarkUsedIfaceMethod(n *ir.CallExpr)
