// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// When using GOEXPERIMENT=boringcrypto, the test program links in the boringcrypto syso,
// which does not respect GOAMD64, so we skip the test if boringcrypto is enabled.
//go:build !boringcrypto

package amd64_test
