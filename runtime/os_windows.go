// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// TODO(brainman): should not need those

var (
	// Following syscalls are available on every Windows PC.
	// All these variables are set by the Windows executable
	// loader before the Go program starts.

	_ stdFunction

	// Following syscalls are only available on some Windows PCs.
	// We will load syscalls, if available, before using them.

	_ stdFunction

	// These are from non-kernel32.dll, so we prefer to LoadLibraryEx them.

	_ stdFunction
)

// osRelaxMinNS indicates that sysmon shouldn't osRelax if the next
// timer is less than 60 ms from now. Since osRelaxing may reduce
// timer resolution to 15.6 ms, this keeps timer error under roughly 1
// part in 4.

// haveHighResTimer indicates that the CreateWaitableTimerEx
// CREATE_WAITABLE_TIMER_HIGH_RESOLUTION flag is available.

//go:linkname canUseLongPaths os.canUseLongPaths

// We want this to be large enough to hold the contents of sysDirectory, *plus*
// a slash and another component that itself is greater than MAX_PATH.

// exiting is set to non-zero when the process is exiting.

// suspendLock protects simultaneous SuspendThread operations from
// suspending each other.
