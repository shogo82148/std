// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pprof

// A profileBuilder writes a profile incrementally from a
// stream of profile samples delivered by the runtime.

// symbolizeFlag keeps track of symbolization result.
//   0                  : no symbol lookup was performed
//   1<<0 (lookupTried) : symbol lookup was performed
//   1<<1 (lookupFailed): symbol lookup was performed but failed
