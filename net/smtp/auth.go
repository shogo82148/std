// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package smtp

// Auth is implemented by an SMTP authentication mechanism.
type Auth interface {
	Start(server *ServerInfo) (proto string, toServer []byte, err error)

	Next(fromServer []byte, more bool) (toServer []byte, err error)
}

// ServerInfo records information about an SMTP server.
type ServerInfo struct {
	Name string
	TLS  bool
	Auth []string
}

// PlainAuth returns an Auth that implements the PLAIN authentication
// mechanism as defined in RFC 4616.
// The returned Auth uses the given username and password to authenticate
// on TLS connections to host and act as identity. Usually identity will be
// left blank to act as username.
func PlainAuth(identity, username, password, host string) Auth

// CRAMMD5Auth returns an Auth that implements the CRAM-MD5 authentication
// mechanism as defined in RFC 2195.
// The returned Auth uses the given username and secret to authenticate
// to the server using the challenge-response mechanism.
func CRAMMD5Auth(username, secret string) Auth
