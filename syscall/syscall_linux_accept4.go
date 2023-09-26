// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides the Accept function used on all systems
// other than arm. See syscall_linux_accept.go for why.

//go:build linux && !arm

package syscall

func Accept(fd int) (nfd int, sa Sockaddr, err error)
