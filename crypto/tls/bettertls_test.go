// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This test uses Netflix's BetterTLS test suite to test the crypto/x509
// path building and name constraint validation.
//
// The test data in JSON form is around 31MB, so we fetch the BetterTLS
// go module and use it to generate the JSON data on-the-fly in a tmp dir.
//
// For more information, see:
// https://github.com/netflix/bettertls
// https://netflixtechblog.com/bettertls-c9915cd255c0

package tls_test
