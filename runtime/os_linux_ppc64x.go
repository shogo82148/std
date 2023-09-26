// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ppc64 || ppc64le
// +build ppc64 ppc64le

package runtime

// For go:linkname
import _ "github.com/shogo82148/std/unsafe"

//go:linkname cpu_hwcap internal/cpu.ppc64x_hwcap
//go:linkname cpu_hwcap2 internal/cpu.ppc64x_hwcap2
