// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !cmd_go_bootstrap

// This code is compiled into the real 'go' binary, but it is not
// compiled into the binary that is built during all.bash, so as
// to avoid needing to build net (and thus use cgo) during the
// bootstrap process.

package web

import (
	"github.com/shogo82148/std/net/http"
)

type Interceptor struct {
	Scheme   string
	FromHost string
	ToHost   string
	Client   *http.Client
}

func EnableTestHooks(interceptors []Interceptor) error

func DisableTestHooks()
