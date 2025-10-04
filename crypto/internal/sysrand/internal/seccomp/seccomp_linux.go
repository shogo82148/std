// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package seccomp

// DisableGetrandom makes future calls to getrandom(2) fail with ENOSYS. It
// applies only to the current thread and to any programs executed from it.
// Callers should use [runtime.LockOSThread] in a dedicated goroutine.
func DisableGetrandom() error
