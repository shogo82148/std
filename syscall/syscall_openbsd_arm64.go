// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func SetKevent(k *Kevent_t, fd, mode, flags int)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)

// RTM_LOCKはOpenBSD 6.3以前にのみ存在します。
const RTM_LOCK = 0x8

// SYS___SYSCTLはOpenBSD 5.8以前にのみ存在し、その後SYS_SYSCTLに
// 名前が変更されました。
const SYS___SYSCTL = SYS_SYSCTL
