// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Garbage collector: write barriers.
//
// For the concurrent garbage collector, the Go compiler implements
// updates to pointer-valued fields that may be in heap objects by
// emitting calls to write barriers. The main write barrier for
// individual pointer writes is gcWriteBarrier and is implemented in
// assembly. This file contains write barrier entry points for bulk
// operations. See also mwbbuf.go.

package runtime
