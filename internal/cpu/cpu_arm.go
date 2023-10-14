// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cpu

const CacheLinePadSize = 32

// arm doesn't have a 'cpuid' equivalent, so we rely on HWCAP/HWCAP2.
// These are initialized by archauxv() and should not be changed after they are
// initialized.
var HWCap uint
var HWCap2 uint
