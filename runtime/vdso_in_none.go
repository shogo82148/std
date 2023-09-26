// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (linux && !386 && !amd64 && !arm && !arm64 && !ppc64 && !ppc64le) || !linux
// +build linux,!386,!amd64,!arm,!arm64,!ppc64,!ppc64le !linux

package runtime
