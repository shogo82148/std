// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

// The function is used to test preemption at split stack checks.
// Declaring a var avoids inlining at the call site.

type Matrix [][]float64
