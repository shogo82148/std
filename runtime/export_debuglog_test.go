// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Export debuglog guts for testing.

package runtime

const DlogEnabled = dlogEnabled

const DebugLogBytes = debugLogBytes

const DebugLogStringLimit = debugLogStringLimit

var Dlog = dlog
