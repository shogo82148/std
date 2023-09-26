// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Futex is only available on DragonFly BSD, FreeBSD and Linux.
// The race detector emits calls to split stack functions so it breaks
// the test.

//go:build (dragonfly || freebsd || linux) && !race
// +build dragonfly freebsd linux
// +build !race

package runtime_test
