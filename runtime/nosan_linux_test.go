// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The file contains tests that cannot run under race detector (or asan or msan) for some reason.
//
//go:build !race && !asan && !msan

package runtime_test
