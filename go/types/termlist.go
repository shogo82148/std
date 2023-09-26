// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// A termlist represents the type set represented by the union
// t1 ∪ y2 ∪ ... tn of the type sets of the terms t1 to tn.
// A termlist is in normal form if all terms are disjoint.
// termlist operations don't require the operands to be in
// normal form.

// allTermlist represents the set of all types.
// It is in normal form.
