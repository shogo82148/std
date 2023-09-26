// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fcgi implements the FastCGI protocol.
//
// See https://fast-cgi.github.io/ for an unofficial mirror of the
// original documentation.
//
// Currently only the responder role is supported.
package fcgi

// recType is a record type, as defined by
// https://web.archive.org/web/20150420080736/http://www.fastcgi.com/drupal/node/6?q=node/22#S8

// keep the connection between web-server and responder open after request

// for padding so we don't have to allocate all the time
// not synchronized because we don't care what the contents are

// conn sends records over rwc

// bufWriter encapsulates bufio.Writer but also closes the underlying stream when
// Closed.

// streamWriter abstracts out the separation of a stream into discrete records.
// It only writes maxWrite bytes at a time.
