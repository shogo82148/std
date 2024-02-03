// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package debug_test

// ExampleSetCrashOutput_monitor shows an example of using
// [debug.SetCrashOutput] to direct crashes to a "monitor" process,
// for automated crash reporting. The monitor is the same executable,
// invoked in a special mode indicated by an environment variable.
func ExampleSetCrashOutput_monitor() {
	appmain()

	// This Example doesn't actually run as a test because its
	// purpose is to crash, so it has no "Output:" comment
	// within the function body.
	//
	// To observe the monitor in action, replace the entire text
	// of this comment with "Output:" and run this command:
	//
	//    $ go test -run=ExampleSetCrashOutput_monitor runtime/debug
	//    panic: oops
	//    ...stack...
	//    monitor: saved crash report at /tmp/10804884239807998216.crash
}
