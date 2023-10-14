// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !static && !(darwin && internal)

// Excluded in darwin internal linking PIE (which is the default) mode,
// as dynamic export is not supported.

package cgotest

//export IMPIsOpaque
func IMPIsOpaque()

//export IMPInitWithFrame
func IMPInitWithFrame()

//export IMPDrawRect
func IMPDrawRect()

//export IMPWindowResize
func IMPWindowResize()
