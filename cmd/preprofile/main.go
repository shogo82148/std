// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Preprofile handles pprof files.
//
// Usage:
//
//	go tool preprofile [-v] [-o output] [-i (pprof)input]
//
//

package main

type NodeMapKey struct {
	CallerName     string
	CalleeName     string
	CallSiteOffset int
}
