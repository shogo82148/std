// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go test cmd/go -v -run=^TestDocsUpToDate$ -fixdocs

package main

import (
	rtrace "runtime/trace"
)

var _ = go11tag
