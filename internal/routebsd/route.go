// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd

// Package routebsd supports reading interface addresses on BSD systems.
// This is a very stripped down version of x/net/route,
// for use by the net package in the standard library.
package routebsd

// FetchRIBMessages fetches a list of addressing messages for an interface.
// The typ argument is something like syscall.NET_RT_IFLIST.
// The argument is an interface index or 0 for all.
func FetchRIBMessages(typ, arg int) ([]Message, error)
