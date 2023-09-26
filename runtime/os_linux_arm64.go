// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64
// +build arm64

package runtime

// For go:linkname
import _ "github.com/shogo82148/std/unsafe"

//go:linkname cpu_hwcap internal/cpu.hwcap

//go:linkname cpu_hwcap2 internal/cpu.hwcap2
