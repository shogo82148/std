// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !goexperiment.swissmap

// This file allows non-GOEXPERIMENT=swissmap builds (i.e., old map builds) to
// construct a swissmap table for running the tests in this package.

package maps
