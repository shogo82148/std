// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// encgen writes the helper functions for encoding. Intended to be
// used with go generate; see the invocation in encode.go.

// TODO: We could do more by being unsafe. Add a -unsafe flag?

package main

type Type struct {
	lower   string
	upper   string
	zero    string
	encoder string
}
