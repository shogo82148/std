// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// This program computes the value of rngCooked in rng.go,
// which is used for seeding all instances of rand.Source.
// a 64bit and a 63bit version of the array is printed to
// the standard output.

package main
