// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Declarations for operating systems implementing time.now directly in assembly.
// Those systems are also expected to have nanotime subtract startNano,
// so that time.now and nanotime return the same monotonic clock readings.

//go:build (darwin && amd64) || (darwin && 386) || windows
// +build darwin,amd64 darwin,386 windows

package runtime

import _ "github.com/shogo82148/std/unsafe"
