// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (linux && !amd64 && !arm64) || (freebsd && !amd64)
// +build linux,!amd64,!arm64 freebsd,!amd64

package runtime
