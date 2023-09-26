// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// This program generates zipdata.go from $GOROOT/lib/time/zoneinfo.zip.
package main

// header is put at the start of the generated file.
// The string addition avoids this file (generate_zipdata.go) from
// matching the "generated file" regexp.
