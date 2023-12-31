// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Issue 8148.  A typedef of an unnamed struct didn't work when used
// with an exported Go function.  No runtime test; just make sure it
// compiles.

package cgotest

func Issue8148() int
