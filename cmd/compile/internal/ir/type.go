// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ir

import (
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

// TypeNode returns the Node representing the type t.
func TypeNode(t *types.Type) Node

// A DynamicType represents a type expression whose exact type must be
// computed dynamically.
type DynamicType struct {
	miniExpr

	// RType is an expression that yields a *runtime._type value
	// representing the asserted type.
	//
	// BUG(mdempsky): If ITab is non-nil, RType may be nil.
	RType Node

	// ITab is an expression that yields a *runtime.itab value
	// representing the asserted type within the assertee expression's
	// original interface type.
	//
	// ITab is only used for assertions (including type switches) from
	// non-empty interface type to a concrete (i.e., non-interface)
	// type. For all other assertions, ITab is nil.
	ITab Node
}

func NewDynamicType(pos src.XPos, rtype Node) *DynamicType

// ToStatic returns static type of dt if it is actually static.
func (dt *DynamicType) ToStatic() Node
