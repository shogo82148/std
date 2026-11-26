// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Exclude -race because Go's race_darwin_*.syso exports a 2-arg
// __tsan_acquire(ThreadState*, void*) that would collide with C TSAN's
// 1-arg __tsan_acquire(void*).
//
//go:build goexperiment.runtimecgopurego && darwin && !race && !msan

package cgo
