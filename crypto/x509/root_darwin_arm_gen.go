// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// Generates root_darwin_armx.go.
//
// As of iOS 8, there is no API for querying the system trusted X.509 root
// certificates. We could use SecTrustEvaluate to verify that a trust chain
// exists for a certificate, but the x509 API requires returning the entire
// chain.
//
// Apple publishes the list of trusted root certificates for iOS on
// support.apple.com. So we parse the list and extract the certificates from
// an OS X machine and embed them into the x509 package.
package main
