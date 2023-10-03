// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements typechecking of expressions.

package types

// This is only used for operations that may cause overflow.

// exprKind describes the kind of an expression; the kind
// determines if an expression is valid in 'statement context'.
