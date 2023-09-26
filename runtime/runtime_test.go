// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime_test

import (
	. "runtime"
)

// flagQuick is set by the -quick option to skip some relatively slow tests.
// This is used by the cmd/dist test runtime:cpu124.
// The cmd/dist test passes both -test.short and -quick;
// there are tests that only check testing.Short, and those tests will
// not be skipped if only -quick is used.

var One = []int64{1}
