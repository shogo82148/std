// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

import "github.com/shogo82148/std/go/constant"

// A Type represents a type of Go.
// All types implement the Type interface.
// (This type originally lived in types2. We moved it here
// so we could depend on it from other packages without
// introducing a circularity.)
type Type interface {
	// Underlying returns the underlying type of a type.
	Underlying() Type

	// String returns a string representation of a type.
	String() string
}

// A TypeAndValue records the type information, constant
// value if known, and various other flags associated with
// an expression.
// This type is similar to types2.TypeAndValue, but exposes
// none of types2's internals.
type TypeAndValue struct {
	Type  Type
	Value constant.Value
	exprFlags
}
