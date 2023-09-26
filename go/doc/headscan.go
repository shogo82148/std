// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

/*
The headscan command extracts comment headings from package files;
it is used to detect false positives which may require an adjustment
to the comment formatting heuristics in comment.go.

Usage: headscan [-root root_directory]

By default, the $GOROOT/src directory is scanned.
*/
package main

// ToHTML in comment.go assigns a (possibly blank) ID to each heading
