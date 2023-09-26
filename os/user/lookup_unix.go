// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build (darwin || dragonfly || freebsd || (js && wasm) || (!android && linux) || nacl || netbsd || openbsd || solaris) && (!cgo || osusergo)
// +build darwin dragonfly freebsd js,wasm !android,linux nacl netbsd openbsd solaris
// +build !cgo osusergo

package user

// lineFunc returns a value, an error, or (nil, nil) to skip the row.
