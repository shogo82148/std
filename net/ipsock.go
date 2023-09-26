// Copyright 2009 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// IP sockets

package net

type InvalidAddrError string

func (e InvalidAddrError) Error() string
func (e InvalidAddrError) Timeout() bool
func (e InvalidAddrError) Temporary() bool

// SplitHostPort splits a network address of the form
// "host:port" or "[host]:port" into host and port.
// The latter form must be used when host contains a colon.
func SplitHostPort(hostport string) (host, port string, err error)

// JoinHostPort combines host and port into a network address
// of the form "host:port" or, if host contains a colon, "[host]:port".
func JoinHostPort(host, port string) string
