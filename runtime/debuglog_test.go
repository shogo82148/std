// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// TODO(austin): All of these tests are skipped if the debuglog build
// tag isn't provided. That means we basically never test debuglog.
// There are two potential ways around this:
//
// 1. Make these tests re-build the runtime test with the debuglog
// build tag and re-invoke themselves.
//
// 2. Always build the whole debuglog infrastructure and depend on
// linker dead-code elimination to drop it. This is easy for dlog()
// since there won't be any calls to it. For printDebugLog, we can
// make panic call a wrapper that is call printDebugLog if the
// debuglog build tag is set, or otherwise do nothing. Then tests
// could call printDebugLog directly. This is the right answer in
// principle, but currently our linker reads in all symbols
// regardless, so this would slow down and bloat all links. If the
// linker gets more efficient about this, we should revisit this
// approach.

package runtime_test
