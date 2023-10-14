// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements accept for platforms that do not provide a fast path for
// setting SetNonblock and CloseOnExec.

//go:build aix || darwin || (js && wasm) || wasip1

package poll
