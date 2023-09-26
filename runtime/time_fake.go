// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build faketime && !windows
// +build faketime,!windows

// Faketime isn't currently supported on Windows. This would require:
//
// 1. Shadowing time_now, which is implemented in assembly on Windows.
//    Since that's exported directly to the time package from runtime
//    assembly, this would involve moving it from sys_windows_*.s into
//    its own assembly files build-tagged with !faketime and using the
//    implementation of time_now from timestub.go in faketime mode.
//
// 2. Modifying syscall.Write to call syscall.faketimeWrite,
//    translating the Stdout and Stderr handles into FDs 1 and 2.
//    (See CL 192739 PS 3.)

package runtime

// faketime is the simulated time in nanoseconds since 1970 for the
// playground.
