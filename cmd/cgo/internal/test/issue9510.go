// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build cgo && !((ppc64 || ppc64le) && internal)

// Test that we can link together two different cgo packages that both
// use the same libgcc function.

package cgotest
