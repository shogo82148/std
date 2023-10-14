// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !static && (!darwin || (!internal_pie && !arm64))

// Excluded in darwin internal linking PIE mode, as dynamic export is not
// supported.
// Excluded in internal linking mode on darwin/arm64, as it is always PIE.

package cgotest

//export IMPIsOpaque
func IMPIsOpaque()

//export IMPInitWithFrame
func IMPInitWithFrame()

//export IMPDrawRect
func IMPDrawRect()

//export IMPWindowResize
func IMPWindowResize()
