// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || freebsd || linux || netbsd || openbsd || solaris

package socktest

import "github.com/shogo82148/std/syscall"

// Accept4 wraps syscall.Accept4.
func (sw *Switch) Accept4(s, flags int) (ns int, sa syscall.Sockaddr, err error)
