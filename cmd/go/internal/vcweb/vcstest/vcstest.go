// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package vcstest serves the repository scripts in cmd/go/testdata/vcstest
// using the [vcweb] script engine.
package vcstest

import (
	"github.com/shogo82148/std/cmd/go/internal/vcweb"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/net/http/httptest"
)

var Hosts = []string{
	"vcs-test.golang.org",
}

type Server struct {
	vcweb   *vcweb.Server
	workDir string
	HTTP    *httptest.Server
	HTTPS   *httptest.Server
}

// NewServer returns a new test-local vcweb server that serves VCS requests
// for modules with paths that begin with "vcs-test.golang.org" using the
// scripts in cmd/go/testdata/vcstest.
func NewServer() (srv *Server, err error)

func (srv *Server) Close() error

func (srv *Server) WriteCertificateFile() (string, error)

// TLSClient returns an http.Client that can talk to the httptest.Server
// whose certificate is written to the given file path.
func TLSClient(certFile string) (*http.Client, error)
