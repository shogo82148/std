// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types2

// CommonUnderlying returns the common underlying type of the types in
// the type set implied by t, or nil if there is none.
//
// If t is not a type parameter, the result is t.Underlying().
// (An operand of interface type has the single type t, not the type
// set of the interface.)
//
// If t is a type parameter, the result is the single underlying type U
// shared by every specific type in its constraint's type set. It is nil
// if there are no specific types (e.g. any, comparable) or they do not
// share an underlying type. As a special case, if all are channel types
// with identical element types and directions that do not conflict, the
// result is the most restrictive channel type.
func CommonUnderlying(t Type) Type
