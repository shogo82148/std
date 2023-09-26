// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

// keyPEM is the same as rsaKeyPEM, but declares itself as just
// "PRIVATE KEY", not "RSA PRIVATE KEY".  https://golang.org/issue/4477

// changeImplConn is a net.Conn which can change its Write and Close
// methods.
