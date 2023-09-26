// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package mail implements parsing of mail messages.

For the most part, this package follows the syntax as specified by RFC 5322.
Notable divergences:
  - Obsolete address formats are not parsed, including addresses with
    embedded route information.
  - Group addresses are not parsed.
  - The full range of spacing (the CFWS syntax element) is not supported,
    such as breaking addresses across lines.
*/
package mail

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/time"
)

// A Message represents a parsed mail message.
type Message struct {
	Header Header
	Body   io.Reader
}

// ReadMessage reads a message from r.
// The headers are parsed, and the body of the message will be available
// for reading from r.
func ReadMessage(r io.Reader) (msg *Message, err error)

// Layouts suitable for passing to time.Parse.
// These are tried in order.

// A Header represents the key-value pairs in a mail message header.
type Header map[string][]string

// Get gets the first value associated with the given key.
// If there are no values associated with the key, Get returns "".
func (h Header) Get(key string) string

var ErrHeaderNotPresent = errors.New("mail: header not in message")

// Date parses the Date header field.
func (h Header) Date() (time.Time, error)

// AddressList parses the named header field as a list of addresses.
func (h Header) AddressList(key string) ([]*Address, error)

// Address represents a single mail address.
// An address such as "Barry Gibbs <bg@example.com>" is represented
// as Address{Name: "Barry Gibbs", Address: "bg@example.com"}.
type Address struct {
	Name    string
	Address string
}

// String formats the address as a valid RFC 5322 address.
// If the address's name contains non-ASCII characters
// the name will be rendered according to RFC 2047.
func (a *Address) String() string
