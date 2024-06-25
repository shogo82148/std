// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package profilerecord holds internal types used to represent profiling
// records with deep stack traces.
//
// TODO: Consider moving this to internal/runtime, see golang.org/issue/65355.
package profilerecord

type StackRecord struct {
	Stack []uintptr
}

type MemProfileRecord struct {
	AllocBytes, FreeBytes     int64
	AllocObjects, FreeObjects int64
	Stack                     []uintptr
}

func (r *MemProfileRecord) InUseBytes() int64
func (r *MemProfileRecord) InUseObjects() int64

type BlockProfileRecord struct {
	Count  int64
	Cycles int64
	Stack  []uintptr
}
