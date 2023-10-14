// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that -static works when not using cgo.  This test is in
// misc/cgo to take advantage of the testing framework support for
// when -static is expected to work.

package nocgo

func NoCgo() int
