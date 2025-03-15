// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO: expand the set of supported platforms, with testing. Nothing about
// the instrumentation is OS specific, but only amd64 and arm64 are
// supported in the runtime. See src/runtime/libfuzzer*.
//
// If you update this constraint, also update internal/platform.FuzzInstrumented.
//
//go:build !((darwin || linux || windows || freebsd || openbsd) && (amd64 || arm64))

package fuzz
