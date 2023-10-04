// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements various error reporters.

package types
<<<<<<< HEAD

// An error_ represents a type-checking error.
// To report an error_, call Checker.report.

// An errorDesc describes part of a type-checking error.

// The positioner interface is used to extract the position of type-checker
// errors.

// posSpan holds a position range along with a highlighted position within that
// range. This is used for positioning errors, with pos by convention being the
// first position in the source where the error is known to exist, and start
// and end defining the full span of syntax being considered when the error was
// detected. Invariant: start <= pos < end || start == pos == end.

// atPos wraps a token.Pos to implement the positioner interface.
=======
>>>>>>> upstream/release-branch.go1.21
