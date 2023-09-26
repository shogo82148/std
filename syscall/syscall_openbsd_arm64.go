// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

func SetKevent(k *Kevent_t, fd, mode, flags int)

func (iov *Iovec) SetLen(length int)

func (msghdr *Msghdr) SetControllen(length int)

func (cmsg *Cmsghdr) SetLen(length int)

// RTM_LOCK only exists in OpenBSD 6.3 and earlier.
const RTM_LOCK = 0x8

// SYS___SYSCTL only exists in OpenBSD 5.8 and earlier, when it was
// was renamed to SYS_SYSCTL.
const SYS___SYSCTL = SYS_SYSCTL
