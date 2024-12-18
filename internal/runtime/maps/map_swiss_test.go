// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Tests of map internals that need to use the builtin map type, and thus must
// be built with GOEXPERIMENT=swissmap.

//go:build goexperiment.swissmap

package maps_test
