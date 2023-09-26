// Copyright 2011 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd
// +build darwin dragonfly freebsd netbsd openbsd

// Berkeley packet filter for BSD variants

package syscall

func BpfStmt(code, k int) *BpfInsn

func BpfJump(code, k, jt, jf int) *BpfInsn

func BpfBuflen(fd int) (int, error)

func SetBpfBuflen(fd, l int) (int, error)

func BpfDatalink(fd int) (int, error)

func SetBpfDatalink(fd, t int) (int, error)

func SetBpfPromisc(fd, m int) error

func FlushBpf(fd int) error

func BpfInterface(fd int, name string) (string, error)

func SetBpfInterface(fd int, name string) error

func BpfTimeout(fd int) (*Timeval, error)

func SetBpfTimeout(fd int, tv *Timeval) error

func BpfStats(fd int) (*BpfStat, error)

func SetBpfImmediate(fd, m int) error

func SetBpf(fd int, i []BpfInsn) error

func CheckBpfVersion(fd int) error

func BpfHeadercmpl(fd int) (int, error)

func SetBpfHeadercmpl(fd, f int) error
