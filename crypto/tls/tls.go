// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tls partially implements TLS 1.0, as specified in RFC 2246.
package tls

import (
	"github.com/shogo82148/std/net"
)

// Server returns a new TLS server side connection
// using conn as the underlying transport.
// The configuration config must be non-nil and must have
// at least one certificate.
func Server(conn net.Conn, config *Config) *Conn

// Client returns a new TLS client side connection
// using conn as the underlying transport.
// Client interprets a nil configuration as equivalent to
// the zero configuration; see the documentation of Config
// for the defaults.
func Client(conn net.Conn, config *Config) *Conn

// A listener implements a network listener (net.Listener) for TLS connections.

// NewListener creates a Listener which accepts connections from an inner
// Listener and wraps each connection with Server.
// The configuration config must be non-nil and must have
// at least one certificate.
func NewListener(inner net.Listener, config *Config) net.Listener

// Listen creates a TLS listener accepting connections on the
// given network address using net.Listen.
// The configuration config must be non-nil and must have
// at least one certificate.
func Listen(network, laddr string, config *Config) (net.Listener, error)

// Dial connects to the given network address using net.Dial
// and then initiates a TLS handshake, returning the resulting
// TLS connection.
// Dial interprets a nil configuration as equivalent to
// the zero configuration; see the documentation of Config
// for the defaults.
func Dial(network, addr string, config *Config) (*Conn, error)

// LoadX509KeyPair reads and parses a public/private key pair from a pair of
// files. The files must contain PEM encoded data.
func LoadX509KeyPair(certFile, keyFile string) (cert Certificate, err error)

// X509KeyPair parses a public/private key pair from a pair of
// PEM encoded data.
func X509KeyPair(certPEMBlock, keyPEMBlock []byte) (cert Certificate, err error)
