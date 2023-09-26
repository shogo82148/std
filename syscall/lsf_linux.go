// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Linux socket filter

package syscall

func LsfStmt(code, k int) *SockFilter

func LsfJump(code, k, jt, jf int) *SockFilter

func LsfSocket(ifindex, proto int) (int, error)

func SetLsfPromisc(name string, m bool) error

func AttachLsf(fd int, i []SockFilter) error

func DetachLsf(fd int) error
