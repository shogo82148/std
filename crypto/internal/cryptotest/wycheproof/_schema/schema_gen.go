// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A stand-alone Go module that generates ../schema.go using the
// upstream Wycheproof JSON schema documents.
//
// We maintain this in a separate Go module and vendor the resulting
// generated .go code to avoid the standard library taking a direct
// dependency on the c2sp/wycheproof or atombender/go-jsonschema modules.

package main
