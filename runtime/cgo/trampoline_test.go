// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.runtimecgopurego && !msan && !tsan && ((unix && !arm && !mips && !mipsle && !ppc64 && !ppc64le) || (windows && !386))

package cgo
