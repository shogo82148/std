// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements typechecking of call and selector expressions.

package types

// A getter sets x as the i'th operand, where 0 <= i < n and n is the total
// number of operands (context-specific, and maintained elsewhere). A getter
// type-checks the i'th operand; the details of the actual check are getter-
// specific.
