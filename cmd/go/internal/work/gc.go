// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package work

import (
	"github.com/shogo82148/std/runtime"
)

// Tests can override this by setting $TESTGO_TOOLCHAIN_VERSION.
var ToolchainVersion = runtime.Version()
