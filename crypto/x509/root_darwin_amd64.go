// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !ios
// +build !ios

package x509

import (
	macOS "crypto/x509/internal/macos"
)

// loadSystemRootsWithCgo is set in root_cgo_darwin_amd64.go when cgo is
// available, and is only used for testing.
