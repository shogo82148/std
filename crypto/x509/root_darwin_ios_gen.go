// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

// Generates root_darwin_ios.go.
//
// As of iOS 13, there is no API for querying the system trusted X.509 root
// certificates.
//
// Apple publishes the trusted root certificates for iOS and macOS on
// opensource.apple.com so we embed them into the x509 package.
//
// Note that this ignores distrusted and revoked certificates.
package main
