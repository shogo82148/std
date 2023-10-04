// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types
<<<<<<< HEAD

// A _TypeSet represents the type set of an interface.
// Because of existing language restrictions, methods can be "factored out"
// from the terms. The actual type set is the intersection of the type set
// implied by the methods and the type set described by the terms and the
// comparable bit. To test whether a type is included in a type set
// ("implements" relation), the type must implement all methods _and_ be
// an element of the type set described by the terms and the comparable bit.
// If the term list describes the set of all types and comparable is true,
// only comparable types are meant; in all other cases comparable is false.

// topTypeSet may be used as type set for the empty interface.

// byUniqueMethodName method lists can be sorted by their unique method names.

// invalidTypeSet is a singleton type set to signal an invalid type set
// due to an error. It's also a valid empty type set, so consumers of
// type sets may choose to ignore it.
=======
>>>>>>> upstream/master
