// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements commonly used type predicates.

package types

// IsInterface reports whether typ is an interface type.
func IsInterface(typ Type) bool

// Comparable reports whether values of type T are comparable.
func Comparable(T Type) bool

// Identical reports whether x and y are identical.
func Identical(x, y Type) bool

// An ifacePair is a node in a stack of interface type pairs compared for identity.
