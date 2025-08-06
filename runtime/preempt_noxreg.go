// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !amd64

// This provides common support for architectures that DO NOT use extended
// register state in asynchronous preemption.

package runtime
