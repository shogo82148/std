// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

// FilterNil packs non-nil (non-zero) values in bufp together
// at the beginning of bufp, returning the length of the
// packed buffer. It treats bufp as an array of size n.
//
// TODO(mknyszek): Add a faster SIMD-based implementation.
func FilterNil(bufp *uintptr, n int32) int32
