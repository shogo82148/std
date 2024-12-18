// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// An AlertError is a TLS alert.
//
// When using a QUIC transport, QUICConn methods will return an error
// which wraps AlertError rather than sending a TLS alert.
type AlertError uint8

func (e AlertError) Error() string
