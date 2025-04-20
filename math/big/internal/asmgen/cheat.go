// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// This program can be compiled with -S to produce a “cheat sheet”
// for filling out a new Arch: the compiler will show you how to implement
// the various operations.
//
// Usage (replace TARGET with your target architecture):
//
//	GOOS=linux GOARCH=TARGET go build -gcflags='-p=cheat -S' cheat.go

package p
