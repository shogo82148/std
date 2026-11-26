// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build 386 || amd64

package cpu

const CacheLinePadSize = 64

// Name returns the CPU name given by the vendor.
// If the CPU name can not be determined an
// empty string is returned.
func Name() string

// XSAVEComponent returns the size of XSAVE state component i, and whether it
// must be 64-byte aligned in the compacted format used by XSAVEC. It must only
// be called if X86.HasXSAVEC.
func XSAVEComponent(i uint32) (size uint32, align64 bool)
