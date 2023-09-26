// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Clone, the Linux rfork.

// gsignalInitQuirk, if non-nil, is called for every allocated gsignal G.
//
// TODO(austin): Remove this after Go 1.15 when we remove the
// mlockGsignal workaround.

// touchStackBeforeSignal stores an errno value. If non-zero, it means
// that we should touch the signal stack before sending a signal.
// This is used on systems that have a bug when the signal stack must
// be faulted in.  See #35777 and #37436.
//
// This is accessed atomically as it is set and read in different threads.
//
// TODO(austin): Remove this after Go 1.15 when we remove the
// mlockGsignal workaround.
