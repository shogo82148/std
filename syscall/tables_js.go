// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build js && wasm

package syscall

// These were originally used by Nacl, then later also used by
// js/wasm. Now that they're only used by js/wasm, these numbers are
// just arbitrary.
//
// TODO: delete? replace with something meaningful?

// TODO: Auto-generate some day. (Hard-coded in binaries so not likely to change.)
const (
	// native_client/src/trusted/service_runtime/include/sys/errno.h
	// The errors are mainly copied from Linux.
	EPERM           Errno = 1
	ENOENT          Errno = 2
	ESRCH           Errno = 3
	EINTR           Errno = 4
	EIO             Errno = 5
	ENXIO           Errno = 6
	E2BIG           Errno = 7
	ENOEXEC         Errno = 8
	EBADF           Errno = 9
	ECHILD          Errno = 10
	EAGAIN          Errno = 11
	ENOMEM          Errno = 12
	EACCES          Errno = 13
	EFAULT          Errno = 14
	EBUSY           Errno = 16
	EEXIST          Errno = 17
	EXDEV           Errno = 18
	ENODEV          Errno = 19
	ENOTDIR         Errno = 20
	EISDIR          Errno = 21
	EINVAL          Errno = 22
	ENFILE          Errno = 23
	EMFILE          Errno = 24
	ENOTTY          Errno = 25
	EFBIG           Errno = 27
	ENOSPC          Errno = 28
	ESPIPE          Errno = 29
	EROFS           Errno = 30
	EMLINK          Errno = 31
	EPIPE           Errno = 32
	ENAMETOOLONG    Errno = 36
	ENOSYS          Errno = 38
	EDQUOT          Errno = 122
	EDOM            Errno = 33
	ERANGE          Errno = 34
	EDEADLK         Errno = 35
	ENOLCK          Errno = 37
	ENOTEMPTY       Errno = 39
	ELOOP           Errno = 40
	ENOMSG          Errno = 42
	EIDRM           Errno = 43
	ECHRNG          Errno = 44
	EL2NSYNC        Errno = 45
	EL3HLT          Errno = 46
	EL3RST          Errno = 47
	ELNRNG          Errno = 48
	EUNATCH         Errno = 49
	ENOCSI          Errno = 50
	EL2HLT          Errno = 51
	EBADE           Errno = 52
	EBADR           Errno = 53
	EXFULL          Errno = 54
	ENOANO          Errno = 55
	EBADRQC         Errno = 56
	EBADSLT         Errno = 57
	EDEADLOCK       Errno = EDEADLK
	EBFONT          Errno = 59
	ENOSTR          Errno = 60
	ENODATA         Errno = 61
	ETIME           Errno = 62
	ENOSR           Errno = 63
	ENONET          Errno = 64
	ENOPKG          Errno = 65
	EREMOTE         Errno = 66
	ENOLINK         Errno = 67
	EADV            Errno = 68
	ESRMNT          Errno = 69
	ECOMM           Errno = 70
	EPROTO          Errno = 71
	EMULTIHOP       Errno = 72
	EDOTDOT         Errno = 73
	EBADMSG         Errno = 74
	EOVERFLOW       Errno = 75
	ENOTUNIQ        Errno = 76
	EBADFD          Errno = 77
	EREMCHG         Errno = 78
	ELIBACC         Errno = 79
	ELIBBAD         Errno = 80
	ELIBSCN         Errno = 81
	ELIBMAX         Errno = 82
	ELIBEXEC        Errno = 83
	EILSEQ          Errno = 84
	EUSERS          Errno = 87
	ENOTSOCK        Errno = 88
	EDESTADDRREQ    Errno = 89
	EMSGSIZE        Errno = 90
	EPROTOTYPE      Errno = 91
	ENOPROTOOPT     Errno = 92
	EPROTONOSUPPORT Errno = 93
	ESOCKTNOSUPPORT Errno = 94
	EOPNOTSUPP      Errno = 95
	EPFNOSUPPORT    Errno = 96
	EAFNOSUPPORT    Errno = 97
	EADDRINUSE      Errno = 98
	EADDRNOTAVAIL   Errno = 99
	ENETDOWN        Errno = 100
	ENETUNREACH     Errno = 101
	ENETRESET       Errno = 102
	ECONNABORTED    Errno = 103
	ECONNRESET      Errno = 104
	ENOBUFS         Errno = 105
	EISCONN         Errno = 106
	ENOTCONN        Errno = 107
	ESHUTDOWN       Errno = 108
	ETOOMANYREFS    Errno = 109
	ETIMEDOUT       Errno = 110
	ECONNREFUSED    Errno = 111
	EHOSTDOWN       Errno = 112
	EHOSTUNREACH    Errno = 113
	EALREADY        Errno = 114
	EINPROGRESS     Errno = 115
	ESTALE          Errno = 116
	ENOTSUP         Errno = EOPNOTSUPP
	ENOMEDIUM       Errno = 123
	ECANCELED       Errno = 125
	ELBIN           Errno = 2048
	EFTYPE          Errno = 2049
	ENMFILE         Errno = 2050
	EPROCLIM        Errno = 2051
	ENOSHARE        Errno = 2052
	ECASECLASH      Errno = 2053
	EWOULDBLOCK     Errno = EAGAIN
)

// TODO: Auto-generate some day. (Hard-coded in binaries so not likely to change.)

// Do the interface allocations only once for common
// Errno values.
