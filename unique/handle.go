// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package unique

// Handle is a globally unique identity for some value of type T.
//
// Two handles compare equal exactly if the two values used to create the handles
// would have also compared equal. The comparison of two handles is trivial and
// typically much more efficient than comparing the values used to create them.
type Handle[T comparable] struct {
	value *T
}

// Value returns a shallow copy of the T value that produced the Handle.
// Value is safe for concurrent use by multiple goroutines.
func (h Handle[T]) Value() T

// Make returns a globally unique handle for a value of type T. Handles
// are equal if and only if the values used to produce them are equal.
// Make is safe for concurrent use by multiple goroutines.
func Make[T comparable](value T) Handle[T]
