// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains an driver.UI implementation
// that provides the readline functionality if possible.

//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris || windows) && !appengine && !android
// +build darwin dragonfly freebsd linux netbsd openbsd solaris windows
// +build !appengine
// +build !android

package main

// readlineUI implements driver.UI interface using the
// golang.org/x/crypto/ssh/terminal package.
// The upstream pprof command implements the same functionality
// using the github.com/chzyer/readline package.
