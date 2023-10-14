// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test to make sure that equality functions (and hash
// functions) don't do unaligned reads on architectures
// that can't do unaligned reads. See issue 46283.

package test
