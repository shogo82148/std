// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This file is run by the x509 tests to ensure that a program with minimal
// imports can sign certificates without errors resulting from missing hash
// functions.
package main
