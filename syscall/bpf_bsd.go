// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd

// Berkeley packet filter for BSD variants

package syscall

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func BpfStmt(code, k int) *BpfInsn

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func BpfJump(code, k, jt, jf int) *BpfInsn

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func BpfBuflen(fd int) (int, error)

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func SetBpfBuflen(fd, l int) (int, error)

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func BpfDatalink(fd int) (int, error)

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func SetBpfDatalink(fd, t int) (int, error)

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func SetBpfPromisc(fd, m int) error

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func FlushBpf(fd int) error

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func BpfInterface(fd int, name string) (string, error)

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func SetBpfInterface(fd int, name string) error

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func BpfTimeout(fd int) (*Timeval, error)

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func SetBpfTimeout(fd int, tv *Timeval) error

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func BpfStats(fd int) (*BpfStat, error)

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func SetBpfImmediate(fd, m int) error

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func SetBpf(fd int, i []BpfInsn) error

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func CheckBpfVersion(fd int) error

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func BpfHeadercmpl(fd int) (int, error)

// Deprecated: 代わりにgolang.org/x/net/bpfを使用してください。
func SetBpfHeadercmpl(fd, f int) error
