// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains a driver.UI implementation
// that provides the readline functionality if possible.

//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows) && !appengine && !android

package main
