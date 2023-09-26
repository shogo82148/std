// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Linux system call wrappers that provide POSIX semantics through the
// corresponding cgo->libc (nptl) wrappers for various system calls.

//go:build linux
// +build linux

package cgo

//go:cgo_import_static _cgo_libc_setegid
//go:linkname _cgo_libc_setegid _cgo_libc_setegid
//go:linkname cgo_libc_setegid syscall.cgo_libc_setegid

//go:cgo_import_static _cgo_libc_seteuid
//go:linkname _cgo_libc_seteuid _cgo_libc_seteuid
//go:linkname cgo_libc_seteuid syscall.cgo_libc_seteuid

//go:cgo_import_static _cgo_libc_setregid
//go:linkname _cgo_libc_setregid _cgo_libc_setregid
//go:linkname cgo_libc_setregid syscall.cgo_libc_setregid

//go:cgo_import_static _cgo_libc_setresgid
//go:linkname _cgo_libc_setresgid _cgo_libc_setresgid
//go:linkname cgo_libc_setresgid syscall.cgo_libc_setresgid

//go:cgo_import_static _cgo_libc_setresuid
//go:linkname _cgo_libc_setresuid _cgo_libc_setresuid
//go:linkname cgo_libc_setresuid syscall.cgo_libc_setresuid

//go:cgo_import_static _cgo_libc_setreuid
//go:linkname _cgo_libc_setreuid _cgo_libc_setreuid
//go:linkname cgo_libc_setreuid syscall.cgo_libc_setreuid

//go:cgo_import_static _cgo_libc_setgroups
//go:linkname _cgo_libc_setgroups _cgo_libc_setgroups
//go:linkname cgo_libc_setgroups syscall.cgo_libc_setgroups

//go:cgo_import_static _cgo_libc_setgid
//go:linkname _cgo_libc_setgid _cgo_libc_setgid
//go:linkname cgo_libc_setgid syscall.cgo_libc_setgid

//go:cgo_import_static _cgo_libc_setuid
//go:linkname _cgo_libc_setuid _cgo_libc_setuid
//go:linkname cgo_libc_setuid syscall.cgo_libc_setuid
