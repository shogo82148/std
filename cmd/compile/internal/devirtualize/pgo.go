// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package devirtualize

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/pgo"
)

// CallStat summarizes a single call site.
//
// This is used only for debug logging.
type CallStat struct {
	Pkg string
	Pos string

	Caller string

	// Direct or indirect call.
	Direct bool

	// For indirect calls, interface call or other indirect function call.
	Interface bool

	// Total edge weight from this call site.
	Weight int64

	// Hottest callee from this call site, regardless of type
	// compatibility.
	Hottest       string
	HottestWeight int64

	// Devirtualized callee if != "".
	//
	// Note that this may be different than Hottest because we apply
	// type-check restrictions, which helps distinguish multiple calls on
	// the same line.
	Devirtualized       string
	DevirtualizedWeight int64
}

// ProfileGuided performs call devirtualization of indirect calls based on
// profile information.
//
// Specifically, it performs conditional devirtualization of interface calls
// for the hottest callee. That is, it performs a transformation like:
//
//	type Iface interface {
//		Foo()
//	}
//
//	type Concrete struct{}
//
//	func (Concrete) Foo() {}
//
//	func foo(i Iface) {
//		i.Foo()
//	}
//
// to:
//
//	func foo(i Iface) {
//		if c, ok := i.(Concrete); ok {
//			c.Foo()
//		} else {
//			i.Foo()
//		}
//	}
//
// The primary benefit of this transformation is enabling inlining of the
// direct call.
func ProfileGuided(fn *ir.Func, p *pgo.Profile)
