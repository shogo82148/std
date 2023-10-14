// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build mips64 || mips64le

package cpu

const CacheLinePadSize = 32

// This is initialized by archauxv and should not be changed after it is
// initialized.
var HWCap uint
