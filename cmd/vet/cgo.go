// Copyright 2015 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Check for invalid cgo pointer passing.
// This looks for code that uses cgo to call C code passing values
// whose types are almost always invalid according to the cgo pointer
// sharing rules.
// Specifically, it warns about attempts to pass a Go chan, map, func,
// or slice to C, either directly, or via a pointer, array, or struct.

package main
