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
