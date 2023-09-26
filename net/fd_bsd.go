// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build freebsd || netbsd || openbsd
// +build freebsd netbsd openbsd

// Waiting for FDs via kqueue/kevent.

package net
