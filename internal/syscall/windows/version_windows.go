// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import (
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/unsafe"
)

// SupportFullTCPKeepAlive indicates whether the current Windows version
// supports the full TCP keep-alive configurations.
// The minimal requirement is Windows 10.0.16299.
var SupportFullTCPKeepAlive = sync.OnceValue(func() bool {
	major, _, build := version()
	return major >= 10 && build >= 16299
})

// SupportTCPInitialRTONoSYNRetransmissions indicates whether the current
// Windows version supports the TCP_INITIAL_RTO_NO_SYN_RETRANSMISSIONS.
// The minimal requirement is Windows 10.0.16299.
var SupportTCPInitialRTONoSYNRetransmissions = sync.OnceValue(func() bool {
	major, _, build := version()
	return major >= 10 && build >= 16299
})

// SupportUnixSocket indicates whether the current Windows version supports
// Unix Domain Sockets.
// The minimal requirement is Windows 10.0.17063.
var SupportUnixSocket = sync.OnceValue(func() bool {
	var size uint32

	_, _ = syscall.WSAEnumProtocols(nil, nil, &size)
	n := int32(size) / int32(unsafe.Sizeof(syscall.WSAProtocolInfo{}))

	buf := make([]syscall.WSAProtocolInfo, n)
	n, err := syscall.WSAEnumProtocols(nil, &buf[0], &size)
	if err != nil {
		return false
	}
	for i := int32(0); i < n; i++ {
		if buf[i].AddressFamily == syscall.AF_UNIX {
			return true
		}
	}
	return false
})
