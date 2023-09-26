// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// opensslInputEvent enumerates possible inputs that can be sent to an `openssl
// s_client` process.

// opensslOutputSink is an io.Writer that receives the stdout and stderr from an
// `openssl` process and sends a value to handshakeComplete or readKeyUpdate
// when certain messages are seen.

// opensslEndOfHandshake is a message that the “openssl s_server” tool will
// print when a handshake completes if run with “-state”.

// opensslReadKeyUpdate is a message that the “openssl s_server” tool will
// print when a KeyUpdate message is received if run with “-state”.

// clientTest represents a test of the TLS client handshake against a reference
// implementation.

// sctsBase64 contains data from `openssl s_client -serverinfo 18 -connect ritter.vg:443`

// brokenConn wraps a net.Conn and causes all Writes after a certain number to
// fail with brokenConnErr.

// brokenConnErr is the error that brokenConn returns once exhausted.

// writeCountingConn wraps a net.Conn and counts the number of Write calls.
