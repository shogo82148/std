// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Script-driven tests.
// See testdata/script/README for an overview.

//go:generate go test cmd/go -v -run=TestScript/README --fixreadme

package main_test

// testingTBKey is the Context key for a testing.TB.
