// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// Mkzip creates a FIPS snapshot zip file.
// See GOROOT/lib/fips140/README.md and GOROOT/lib/fips140/Makefile
// for more details about when and why to use this.
//
// Usage:
//
//	cd GOROOT/lib/fips140
//	go run ../../src/cmd/go/internal/fips140/mkzip.go [-b branch] v1.2.3
//
// Mkzip creates a zip file named for the version on the command line
// using the sources in the named branch (default origin/master,
// to avoid accidentally including local commits).
package main
