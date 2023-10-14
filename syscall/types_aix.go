// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

/*
Input to cgo -godefs.  See also mkerrors.sh and mkall.sh
*/

// +godefs map struct_in_addr [4]byte /* in_addr */
// +godefs map struct_in6_addr [16]byte /* in6_addr */

package syscall

import "C"

const (
	PathMax = C.PATH_MAX
)

type Timespec C.struct_timespec

type Timeval C.struct_timeval

type Timeval32 C.struct_timeval32

type Timezone C.struct_timezone

type Rusage C.struct_rusage

type Rlimit C.struct_rlimit

type Flock_t C.struct_flock

type Stat_t C.struct_stat

type Statfs_t C.struct_statfs

type Fsid64_t C.fsid64_t

type StTimespec_t C.st_timespec_t

type Dirent C.struct_dirent

type RawSockaddrInet4 C.struct_sockaddr_in

type RawSockaddrInet6 C.struct_sockaddr_in6

type RawSockaddrUnix C.struct_sockaddr_un

type RawSockaddrDatalink C.struct_sockaddr_dl

type RawSockaddr C.struct_sockaddr

type RawSockaddrAny C.struct_sockaddr_any

type Cmsghdr C.struct_cmsghdr

type ICMPv6Filter C.struct_icmp6_filter

type Iovec C.struct_iovec

type IPMreq C.struct_ip_mreq

type IPv6Mreq C.struct_ipv6_mreq

type Linger C.struct_linger

type Msghdr C.struct_msghdr

const (
	SizeofSockaddrInet4    = C.sizeof_struct_sockaddr_in
	SizeofSockaddrInet6    = C.sizeof_struct_sockaddr_in6
	SizeofSockaddrAny      = C.sizeof_struct_sockaddr_any
	SizeofSockaddrUnix     = C.sizeof_struct_sockaddr_un
	SizeofSockaddrDatalink = C.sizeof_struct_sockaddr_dl
	SizeofLinger           = C.sizeof_struct_linger
	SizeofIPMreq           = C.sizeof_struct_ip_mreq
	SizeofIPv6Mreq         = C.sizeof_struct_ipv6_mreq
	SizeofMsghdr           = C.sizeof_struct_msghdr
	SizeofCmsghdr          = C.sizeof_struct_cmsghdr
	SizeofICMPv6Filter     = C.sizeof_struct_icmp6_filter
)

const (
	PTRACE_TRACEME = C.PT_TRACE_ME
	PTRACE_CONT    = C.PT_CONTINUE
	PTRACE_KILL    = C.PT_KILL
)

const (
	SizeofIfMsghdr = C.sizeof_struct_if_msghdr
)

type IfMsgHdr C.struct_if_msghdr

type Utsname C.struct_utsname

type Termios C.struct_termios
