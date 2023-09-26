// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && (386 || amd64 || arm || arm64 || ppc64 || ppc64le)
// +build linux
// +build 386 amd64 arm arm64 ppc64 ppc64le

package runtime_test

import (
	_ "unsafe"
)

//go:linkname vdsoClockgettimeSym runtime.vdsoClockgettimeSym
