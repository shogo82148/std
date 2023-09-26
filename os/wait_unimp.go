// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// aix, darwin, js/wasm, openbsd, solaris and wasip1/wasm don't implement
// waitid/wait6.

//go:build aix || darwin || (js && wasm) || openbsd || solaris || wasip1

package os
