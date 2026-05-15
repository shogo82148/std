// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http2

import (
	"github.com/shogo82148/std/time"
)

// Config must be kept in sync with net/http.HTTP2Config.
type Config struct {
	MaxConcurrentStreams          int
	StrictMaxConcurrentRequests   bool
	MaxDecoderHeaderTableSize     int
	MaxEncoderHeaderTableSize     int
	MaxReadFrameSize              int
	MaxReceiveBufferPerConnection int
	MaxReceiveBufferPerStream     int
	SendPingTimeout               time.Duration
	PingTimeout                   time.Duration
	WriteByteTimeout              time.Duration
	PermitProhibitedCipherSuites  bool
	CountError                    func(errType string)
}
