// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modfetch

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var HelpGoproxy = &base.Command{
	UsageLine: "goproxy",
	Short:     "module proxy protocol",
	Long: `
A Go module proxy is any web server that can respond to GET requests for
URLs of a specified form. The requests have no query parameters, so even
a site serving from a fixed file system (including a file:/// URL)
can be a module proxy.

For details on the GOPROXY protocol, see
https://golang.org/ref/mod#goproxy-protocol.
`,
}

// TryProxies iterates f over each configured proxy (including "noproxy" and
// "direct" if applicable) until f returns no error or until f returns an
// error that is not equivalent to fs.ErrNotExist on a proxy configured
// not to fall back on errors.
//
// TryProxies then returns that final error.
//
// If GOPROXY is set to "off", TryProxies invokes f once with the argument
// "off".
func TryProxies(f func(proxy string) error) error
