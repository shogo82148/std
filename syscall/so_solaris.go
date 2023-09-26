// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// soError describes reasons for shared library load failures.

// A so implements access to a single shared library object.

// A proc implements access to a procedure inside a shared library.

// A lazySO implements access to a single shared library.  It will delay
// the load of the shared library until the first call to its Handle method
// or to one of its lazyProc's Addr method.

// A lazyProc implements access to a procedure inside a lazySO.
// It delays the lookup until the Addr method is called.
