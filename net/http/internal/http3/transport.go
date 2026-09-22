// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http3

import (
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/http"

	"golang.org/x/net/quic"
)

type TransportOpts struct {
	// ListenQUIC determines how the transport will open a QUIC endpoint.
	// By default, quic.Listen("udp", addr, config) is used.
	// ListenQUIC might be called multiple times.
	ListenQUIC func(addr string, config *quic.Config) (*quic.Endpoint, error)

	// ListenPacket specifies the function for creating a UDP listener.
	// If ListenPacket is nil, then the transport listens using net.ListenPacket.
	//
	// If ListenQUIC and ListenPacket are both set, ListenQUIC takes priority.
	ListenPacket func(network, addr string) (net.PacketConn, error)

	// QUICConfig is the QUIC configuration used by the transport.
	// QUICConfig may be nil and should not be modified after calling
	// RegisterTransport.
	//
	// The QUICConfig's TLSConfig is not used.
	// Set the TLSConfig on the net/http Transport instead.
	QUICConfig *quic.Config
}

// RegisterTransport configures a net/http HTTP/1 Transport to use HTTP/3.
func RegisterTransport(tr *http.Transport, opts TransportOpts) error
