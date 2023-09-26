// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !linux && !darwin && !dragonfly && !freebsd && !netbsd && (!openbsd || !arm64) && !solaris
// +build !linux
// +build !darwin
// +build !dragonfly
// +build !freebsd
// +build !netbsd
// +build !openbsd !arm64
// +build !solaris

package runtime
