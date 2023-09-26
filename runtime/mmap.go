// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !plan9 && !solaris && !windows && (!linux || !amd64) && (!linux || !arm64) && !js && !darwin && !aix
// +build !plan9
// +build !solaris
// +build !windows
// +build !linux !amd64
// +build !linux !arm64
// +build !js
// +build !darwin
// +build !aix

package runtime
