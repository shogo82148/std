// Copyright 2012 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !cmd_go_bootstrap
// +build !cmd_go_bootstrap

// This code is compiled into the real 'go' binary, but it is not
// compiled into the binary that is built during all.bash, so as
// to avoid needing to build net (and thus use cgo) during the
// bootstrap process.

package main

// httpClient is the default HTTP client, but a variable so it can be
// changed by tests, without modifying http.DefaultClient.

// impatientInsecureHTTPClient is used in -insecure mode,
// when we're connecting to https servers that might not be there
// or might be using self-signed certificates.
