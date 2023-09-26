// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build boringcrypto

package tls

// A self-signed test certificate with an RSA key of size 2048, for testing
// RSA-PSS with SHA512. SAN of example.golang.
