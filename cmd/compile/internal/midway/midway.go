// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package midway

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

// SimdMark is added to names of types/functions/variables that have been
// specialized to a particular SIMD vector width, and is followed by the
// width of the vector in bits, or "0" if the specialization is for emulated
// SIMD.
const SimdMark = "@simd"

// SimdInitMark prefixes synthetic package-level proxy variables created by
// midway for top-level variable initializers that depend on SIMD even though
// the variable's type does not (including blank "_" initializers).
//
// The proxy name has the form:
//
//	"@simdinit<vecSize>@<target>"
//
// where <target> is either the original package variable's name (e.g., "N")
// or "_@<id>" for a blank initializer. Because the name starts with "@", it is
// never exported and cannot collide with legal Go identifiers. The unified IR
// writer omits these proxy declarations from pkgDecl and the unified IR reader
// rewrites their initializers to conditionally assign directly to the original
// variable (or blank).
const SimdInitMark = SimdMark + "init"

const SimdPkg = "simd"
const SimdBridgePkg = "simd/internal/bridge"

// IsSpecializedName reports whether name is the name midway gave to a
// size-specialized copy of a package-level declaration, such as "F@simd128".
// "@" can't appear in a Go identifier, so no user declaration has such a name.
func IsSpecializedName(name string) bool

// SimdInitProxyName returns the synthetic init-proxy variable name for target
// at the given SIMD vector bit size (0 for emulation).
func SimdInitProxyName(target string, vecSize int) string

// IsSimdInitProxy reports whether symName is a synthetic SIMD init-proxy
// variable name created by midway.
func IsSimdInitProxy(symName string) bool

// ParseSimdInitProxy parses a synthetic SIMD init-proxy name of the form
// "@simdinit<vecSize>@<target>". If <target> encodes a blank
// initializer ("_" or "_@<id>"), baseName is returned as "_".
func ParseSimdInitProxy(symName string) (baseName string, vecSize int, ok bool)

// ParseSimdSuffix splits a midway-specialized symbol name of the form
// "<baseName>@simd<vecSize>" (for example, "x@simd0", "x@simd256")
// into its original base name and numeric vector bit size (0 for emulation).
// It returns ok == false if symName does not have a valid "@simd<digits>" suffix.
func ParseSimdSuffix(symName string) (baseName string, vecSize int, ok bool)

// InitOrderPlan records the pre-midway Go-spec package variable initialization
// order so it can be restored after midway specializes top-level variable
// declarations and types2 re-checks the rewritten AST.
//
// Splitting a package variable into per-size declarations allows types2 to
// interleave specializations of one variable with unrelated declarations during
// re-check (for example, when a non-SIMD var depends only on the last proxy),
// which can violate Go-spec side-effect ordering under non-default SIMD configs.
type InitOrderPlan struct {
	numSlots   int
	slotByName map[string]int
	slotByPos  map[syntax.Pos]int
}

// Apply filters out pre-rewrite dummy initializers from info.InitOrder and
// reorders all remaining initializers so that every declaration's specialized
// initializers execute atomically at the exact Go-spec position of its original
// declaration.
func (p *InitOrderPlan) Apply(info *types2.Info)
