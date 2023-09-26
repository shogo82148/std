// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Linux socket filter

package syscall

// Deprecated: Use golang.org/x/net/bpf instead.
func LsfStmt(code, k int) *SockFilter

// Deprecated: Use golang.org/x/net/bpf instead.
func LsfJump(code, k, jt, jf int) *SockFilter

// Deprecated: Use golang.org/x/net/bpf instead.
func LsfSocket(ifindex, proto int) (int, error)

// Deprecated: Use golang.org/x/net/bpf instead.
func SetLsfPromisc(name string, m bool) error

// Deprecated: Use golang.org/x/net/bpf instead.
func AttachLsf(fd int, i []SockFilter) error

// Deprecated: Use golang.org/x/net/bpf instead.
func DetachLsf(fd int) error
