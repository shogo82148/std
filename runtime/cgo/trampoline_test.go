// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.runtimecgopurego && !msan && !tsan && (((amd64 || arm64) && (darwin || (linux && !android))) || (windows && amd64))

package cgo
