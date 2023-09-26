// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

// Strings and slices that don't escape and fit into tmpBuf are stack allocated,
// which defeats using AllocsPerRun to test other optimizations.
