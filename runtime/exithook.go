// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// exitHook stores a function to be run on program exit, registered
// by the utility runtime.addExitHook.

// exitHooks stores state related to hook functions registered to
// run when program execution terminates.
