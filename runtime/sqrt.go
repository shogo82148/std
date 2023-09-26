// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copy of math/sqrt.go, here for use by ARM softfloat.
// Modified to not use any floating point arithmetic so
// that we don't clobber any floating-point registers
// while emulating the sqrt instruction.

package runtime
