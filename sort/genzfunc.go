// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This program is run via "go generate" (via a directive in sort.go)
// to generate zfuncversion.go.
//
// It copies sort.go to zfuncversion.go, only retaining funcs which
// take a "data Interface" parameter, and renaming each to have a
// "_func" suffix and taking a "data lessSwap" instead. It then rewrites
// each internal function call to the appropriate _func variants.

package main
