// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//specgen:category Basics

package spec

// Len returns the number of elements in x.
func Len[E EltOrMask, W Width](x Vec[E, W]) int

// String returns x in the form "{x[0],x[1],...}".
func String[E Nums, W Width](x Vec[E, W]) string

// StringMask returns x in the form "{x[0],x[1],...}", where 0 means false and 1
// means true.
//
//specgen:name String
func StringMask[E MaskElt, W Width](x Vec[E, W]) string
