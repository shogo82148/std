// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build msan && ((linux && (amd64 || arm64 || loong64)) || (freebsd && amd64))

package msan
