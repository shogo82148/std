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

// When loading DLLs, we prefer to use LoadLibraryEx with
// LOAD_LIBRARY_SEARCH_* flags, if available. LoadLibraryEx is not
// available on old Windows, though, and the LOAD_LIBRARY_SEARCH_*
// flags are not available on some versions of Windows without a
// security patch.
//
// https://msdn.microsoft.com/en-us/library/ms684179(v=vs.85).aspx says:
// "Windows 7, Windows Server 2008 R2, Windows Vista, and Windows
// Server 2008: The LOAD_LIBRARY_SEARCH_* flags are available on
// systems that have KB2533623 installed. To determine whether the
// flags are available, use GetProcAddress to get the address of the
// AddDllDirectory, RemoveDllDirectory, or SetDefaultDllDirectories
// function. If GetProcAddress succeeds, the LOAD_LIBRARY_SEARCH_*
// flags can be used with LoadLibraryEx."

// osRelaxMinNS indicates that sysmon shouldn't osRelax if the next
// timer is less than 60 ms from now. Since osRelaxing may reduce
// timer resolution to 15.6 ms, this keeps timer error under roughly 1
// part in 4.

// useQPCTime controls whether time.now and nanotime use QueryPerformanceCounter.
// This is only set to 1 when running under Wine.

// exiting is set to non-zero when the process is exiting.

// suspendLock protects simultaneous SuspendThread operations from
// suspending each other.
