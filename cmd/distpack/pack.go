// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Distpack creates the tgz and zip files for a Go distribution.
// It writes into GOROOT/pkg/distpack:
//
//   - a binary distribution (tgz or zip) for the current GOOS and GOARCH
//   - a source distribution that is independent of GOOS/GOARCH
//   - the module mod, info, and zip files for a distribution in module form
//     (as used by GOTOOLCHAIN support in the go command).
//
// Distpack is typically invoked by the -distpack flag to make.bash.
// A cross-compiled distribution for goos/goarch can be built using:
//
//	GOOS=goos GOARCH=goarch ./make.bash -distpack
//
// To test that the module downloads are usable with the go command:
//
//	./make.bash -distpack
//	mkdir -p /tmp/goproxy/golang.org/toolchain/
//	ln -sf $(pwd)/../pkg/distpack /tmp/goproxy/golang.org/toolchain/@v
//	GOPROXY=file:///tmp/goproxy GOTOOLCHAIN=$(sed 1q ../VERSION) gotip version
//
// gotip can be replaced with an older released Go version once there is one.
// It just can't be the one make.bash built, because it knows it is already that
// version and will skip the download.
package main
