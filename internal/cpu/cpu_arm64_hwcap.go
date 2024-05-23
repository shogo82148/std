// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build arm64 && linux

package cpu

// HWCap may be initialized by archauxv and
// should not be changed after it was initialized.
//
// Other widely used packages
// access HWCap using linkname as well, most notably:
//   - github.com/klauspost/cpuid/v2
//
// Do not remove or change the type signature.
// See go.dev/issue/67401.
//
//go:linkname HWCap
var HWCap uint
