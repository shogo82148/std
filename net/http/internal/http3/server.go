// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http3

import (
	"github.com/shogo82148/std/net/http"

	"golang.org/x/net/quic"
)

type ServerOpts struct {
	// QUICConfig is the QUIC configuration used by the server.
	// QUICConfig may be nil and should not be modified after calling
	// RegisterServer.
	// If QUICConfig.TLSConfig is nil, the TLSConfig of the net/http Server
	// given to RegisterServer will be used.
	QUICConfig *quic.Config
}

// RegisterServer adds HTTP/3 support to a net/http Server.
//
// RegisterServer must be called before s begins serving, and only affects
// s.ListenAndServeTLS.
func RegisterServer(s *http.Server, opts ServerOpts) error
