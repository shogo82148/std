// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: This test could be implemented on all (most?) UNIXes if we
// added syscall.Tgkill more widely.

// We skip all of these tests under race mode because our test thread
// spends all of its time in the race runtime, which isn't a safe
// point.

//go:build (amd64 || arm64 || ppc64le) && linux && !race

package runtime_test
