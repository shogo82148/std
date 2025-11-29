// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// these tests rely on inspecting freed memory, so they
// can't be run under any of the memory validating modes.
// TODO: figure out just which test violate which condition
// and split this file out by individual test cases.
// There could be some value to running some of these
// under validation

//go:build goexperiment.runtimesecret && (arm64 || amd64) && linux && !race && !asan && !msan

package secret
