// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements the host side of CGI (being the webserver
// parent process).

// Package cgi implements CGI (Common Gateway Interface) as specified
// in RFC 3875.
//
// Note that using CGI means starting a new process to handle each
// request, which is typically less efficient than using a
// long-running server.  This package is intended primarily for
// compatibility with existing systems.
package cgi

import (
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/net/http"
)

// Handler runs an executable in a subprocess with a CGI environment.
type Handler struct {
	Path string
	Root string

	Dir string

	Env        []string
	InheritEnv []string
	Logger     *log.Logger
	Args       []string

	PathLocationHandler http.Handler
}

func (h *Handler) ServeHTTP(rw http.ResponseWriter, req *http.Request)
