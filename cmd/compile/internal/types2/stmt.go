// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements typechecking of statements.

package types2

// RangeKeyVal returns the key and value types for a range over typ.
func RangeKeyVal(typ Type) (Type, Type)
