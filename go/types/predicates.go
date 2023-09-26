// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements commonly used type predicates.

package types

// IsInterface reports whether typ is an interface type.
func IsInterface(typ Type) bool

// Comparable reports whether values of type T are comparable.
func Comparable(T Type) bool

// An ifacePair is a node in a stack of interface type pairs compared for identity.

// Default returns the default "typed" type for an "untyped" type;
// it returns the incoming type for all other types. The default type
// for untyped nil is untyped nil.
func Default(typ Type) Type
