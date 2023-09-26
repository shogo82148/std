// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package textproto implements generic support for text-based request/response
// protocols in the style of HTTP, NNTP, and SMTP.
//
// The package provides:
//
// Error, which represents a numeric error response from
// a server.
//
// Pipeline, to manage pipelined requests and responses
// in a client.
//
// Reader, to read numeric response code lines,
// key: value headers, lines wrapped with leading spaces
// on continuation lines, and whole text blocks ending
// with a dot on a line by itself.
//
// Writer, to write dot-encoded text blocks.
//
// Conn, a convenient packaging of Reader, Writer, and Pipeline for use
// with a single network connection.
package textproto

import (
	"github.com/shogo82148/std/io"
)

// An Error represents a numeric error response from a server.
type Error struct {
	Code int
	Msg  string
}

func (e *Error) Error() string

// A ProtocolError describes a protocol violation such
// as an invalid response or a hung-up connection.
type ProtocolError string

func (p ProtocolError) Error() string

// A Conn represents a textual network protocol connection.
// It consists of a Reader and Writer to manage I/O
// and a Pipeline to sequence concurrent requests on the connection.
// These embedded types carry methods with them;
// see the documentation of those types for details.
type Conn struct {
	Reader
	Writer
	Pipeline
	conn io.ReadWriteCloser
}

// NewConn returns a new Conn using conn for I/O.
func NewConn(conn io.ReadWriteCloser) *Conn

// Close closes the connection.
func (c *Conn) Close() error

// Dial connects to the given address on the given network using net.Dial
// and then returns a new Conn for the connection.
func Dial(network, addr string) (*Conn, error)

// Cmd is a convenience method that sends a command after
// waiting its turn in the pipeline. The command text is the
// result of formatting format with args and appending \r\n.
// Cmd returns the id of the command, for use with StartResponse and EndResponse.
//
// For example, a client might run a HELP command that returns a dot-body
// by using:
//
//	id, err := c.Cmd("HELP")
//	if err != nil {
//		return nil, err
//	}
//
//	c.StartResponse(id)
//	defer c.EndResponse(id)
//
//	if _, _, err = c.ReadCodeLine(110); err != nil {
//		return nil, err
//	}
//	text, err := c.ReadDotBytes()
//	if err != nil {
//		return nil, err
//	}
//	return c.ReadCodeLine(250)
func (c *Conn) Cmd(format string, args ...any) (id uint, err error)

// TrimString returns s without leading and trailing ASCII space.
func TrimString(s string) string

// TrimBytes returns b without leading and trailing ASCII space.
func TrimBytes(b []byte) []byte
