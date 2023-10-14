// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Test that SIGSETXID runs on signal stack, since it's likely to
// overflow if it runs on the Go stack.

package cgotest
