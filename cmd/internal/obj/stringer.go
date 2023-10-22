// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

// This is a mini version of the stringer tool customized for the Anames table
// in the architecture support for obj.
// This version just generates the slice of strings, not the String method.

package main

import (
	"github.com/shogo82148/std/regexp"
)

var Are = regexp.MustCompile(`^\tA([A-Za-z0-9]+)`)
