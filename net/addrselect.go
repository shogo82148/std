// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Minimal RFC 6724 address selection.

package net

// RFC 6724 section 2.1.

// byMaskLength sorts policyTableEntry by the size of their Prefix.Mask.Size,
// from smallest mask, to largest.

// RFC 6724 section 3.1.
