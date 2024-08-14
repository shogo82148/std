// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file provides an internal debug logging facility. The debug
// log is a lightweight, in-memory, per-M ring buffer. By default, the
// runtime prints the debug log on panic.
//
// To print something to the debug log, call dlog to obtain a dlogger
// and use the methods on that to add values. The values will be
// space-separated in the output (much like println).
//
// This facility can be enabled by passing -tags debuglog when
// building. Without this tag, dlog calls compile to nothing.
//
// Implementation notes
//
// There are two implementations of the dlog interface: dloggerImpl and
// dloggerFake. dloggerFake is a no-op implementation. dlogger is type-aliased
// to one or the other depending on the debuglog build tag. However, both types
// always exist and are always built. This helps ensure we compile as much of
// the implementation as possible in the default build configuration, while also
// enabling us to achieve good test coverage of the real debuglog implementation
// even when the debuglog build tag is not set.

package runtime
