// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// recordingConn is a net.Conn that records the traffic that passes through it.
// WriteTo can be used to produce Go code that contains the recorded traffic.

// Script of interaction with gnutls implementation.
// The values for this test are obtained by building and running in server mode:
//   % go test -test.run "TestRunServer" -serve
// The recorded bytes are written to stdout.
