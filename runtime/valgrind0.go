// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Valgrind instrumentation is only available on linux amd64 and arm64.

//go:build !valgrind || !linux || (!amd64 && !arm64)

package runtime
