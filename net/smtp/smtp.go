// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package smtp implements the Simple Mail Transfer Protocol as defined in RFC 5321.
// It also implements the following extensions:
//
//	8BITMIME  RFC 1652
//	AUTH      RFC 2554
//	STARTTLS  RFC 3207
//
// Additional extensions may be handled by clients.
package smtp

import (
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/textproto"
)

// A Client represents a client connection to an SMTP server.
type Client struct {
	Text *textproto.Conn

	conn net.Conn

	tls        bool
	serverName string

	ext map[string]string

	auth []string
}

// Dial returns a new Client connected to an SMTP server at addr.
func Dial(addr string) (*Client, error)

// NewClient returns a new Client using an existing connection and host as a
// server name to be used when authenticating.
func NewClient(conn net.Conn, host string) (*Client, error)

// StartTLS sends the STARTTLS command and encrypts all further communication.
// Only servers that advertise the STARTTLS extension support this function.
func (c *Client) StartTLS(config *tls.Config) error

// Verify checks the validity of an email address on the server.
// If Verify returns nil, the address is valid. A non-nil return
// does not necessarily indicate an invalid address. Many servers
// will not verify addresses for security reasons.
func (c *Client) Verify(addr string) error

// Auth authenticates a client using the provided authentication mechanism.
// A failed authentication closes the connection.
// Only servers that advertise the AUTH extension support this function.
func (c *Client) Auth(a Auth) error

// Mail issues a MAIL command to the server using the provided email address.
// If the server supports the 8BITMIME extension, Mail adds the BODY=8BITMIME
// parameter.
// This initiates a mail transaction and is followed by one or more Rcpt calls.
func (c *Client) Mail(from string) error

// Rcpt issues a RCPT command to the server using the provided email address.
// A call to Rcpt must be preceded by a call to Mail and may be followed by
// a Data call or another Rcpt call.
func (c *Client) Rcpt(to string) error

// Data issues a DATA command to the server and returns a writer that
// can be used to write the data. The caller should close the writer
// before calling any more methods on c.
// A call to Data must be preceded by one or more calls to Rcpt.
func (c *Client) Data() (io.WriteCloser, error)

// SendMail connects to the server at addr, switches to TLS if possible,
// authenticates with mechanism a if possible, and then sends an email from
// address from, to addresses to, with message msg.
func SendMail(addr string, a Auth, from string, to []string, msg []byte) error

// Extension reports whether an extension is support by the server.
// The extension name is case-insensitive. If the extension is supported,
// Extension also returns a string that contains any parameters the
// server specifies for the extension.
func (c *Client) Extension(ext string) (bool, string)

// Reset sends the RSET command to the server, aborting the current mail
// transaction.
func (c *Client) Reset() error

// Quit sends the QUIT command and closes the connection to the server.
func (c *Client) Quit() error
