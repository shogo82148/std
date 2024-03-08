// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package windows

import (
	"github.com/shogo82148/std/sync"
)

// Version retrieves the major, minor, and build version numbers
// of the current Windows OS from the RtlGetNtVersionNumbers API
// and parse the results properly.
func Version() (major, minor, build uint32)

// SupportFullTCPKeepAlive indicates whether the current Windows version
// supports the full TCP keep-alive configurations, the minimal requirement
// is Windows 10, version 1709.
var SupportFullTCPKeepAlive = sync.OnceValue(func() bool {
	major, _, build := Version()
	return major >= 10 && build >= 16299
})

// SupportTCPInitialRTONoSYNRetransmissions indicates whether the current
// Windows version supports the TCP_INITIAL_RTO_NO_SYN_RETRANSMISSIONS, the
// minimal requirement is Windows 10.0.16299.
var SupportTCPInitialRTONoSYNRetransmissions = sync.OnceValue(func() bool {
	major, _, build := Version()
	return major >= 10 && build >= 16299
})

// SupportUnixSocket indicates whether the current Windows version supports
// Unix Domain Sockets, the minimal requirement is Windows 10, build 17063.
var SupportUnixSocket = sync.OnceValue(func() bool {
	major, _, build := Version()
	return major >= 10 && build >= 17063
})
