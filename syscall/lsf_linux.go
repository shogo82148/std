// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Linux socket filter

package syscall

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func LsfStmt(code, k int) *SockFilter

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func LsfJump(code, k, jt, jf int) *SockFilter

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func LsfSocket(ifindex, proto int) (int, error)

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func SetLsfPromisc(name string, m bool) error

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func AttachLsf(fd int, i []SockFilter) error

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func DetachLsf(fd int) error
