// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// res_nsearch, for cgo systems where that's available.

//go:build cgo && !netgo && unix && !(darwin || linux || openbsd)

package net
