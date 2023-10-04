// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements typechecking of statements.

package types
<<<<<<< HEAD
<<<<<<< HEAD

// stmtContext is a bitset describing which
// control-flow statements are permissible,
// and provides additional context information
// for better error messages.

// A valueMap maps a case value (of a basic Go type) to a list of positions
// where the same case value appeared, together with the corresponding case
// types.
// Since two case values may have the same "underlying" value but different
// types we need to also check the value's types (e.g., byte(1) vs myByte(1))
// when the switch expression is of interface type.
=======
>>>>>>> upstream/master
=======
>>>>>>> release-branch.go1.21
