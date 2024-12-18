// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements CGI from the perspective of a child
// process.

package cgi

import (
	"github.com/shogo82148/std/net/http"
)

// Request returns the HTTP request as represented in the current
// environment. This assumes the current program is being run
// by a web server in a CGI environment.
// The returned Request's Body is populated, if applicable.
func Request() (*http.Request, error)

// RequestFromMap creates an [http.Request] from CGI variables.
// The returned Request's Body field is not populated.
func RequestFromMap(params map[string]string) (*http.Request, error)

// Serve executes the provided [Handler] on the currently active CGI
// request, if any. If there's no current CGI environment
// an error is returned. The provided handler may be nil to use
// [http.DefaultServeMux].
func Serve(handler http.Handler) error
