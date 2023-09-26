// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// exported value for testing

// intArgRegs is used by the various register assignment
// algorithm implementations in the runtime. These include:.
// - Finalizers (mfinal.go)
// - Windows callbacks (syscall_windows.go)
//
// Both are stripped-down versions of the algorithm since they
// only have to deal with a subset of cases (finalizers only
// take a pointer or interface argument, Go Windows callbacks
// don't support floating point).
//
// It should be modified with care and are generally only
// modified when testing this package.
//
// It should never be set higher than its internal/abi
// constant counterparts, because the system relies on a
// structure that is at least large enough to hold the
// registers the system supports.
//
// Protected by finlock.
