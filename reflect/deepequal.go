// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Deep equality test via reflection

package reflect

// During deepValueEqual, must keep track of checks that are
// in progress.  The comparison algorithm assumes that all
// checks in progress are true when it reencounters them.
// Visited are stored in a map indexed by 17 * a1 + a2;

// DeepEqual tests for deep equality. It uses normal == equality where possible
// but will scan members of arrays, slices, maps, and fields of structs. It correctly
// handles recursive types. Functions are equal only if they are both nil.
func DeepEqual(a1, a2 interface{}) bool
