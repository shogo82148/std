// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// GOMAXPROCS sets the maximum number of CPUs that can be executing
// simultaneously and returns the previous setting. If n < 1, it does not change
// the current setting.
//
// If the GOMAXPROCS environment variable is set to a positive whole number,
// GOMAXPROCS defaults to that value.
//
// Otherwise, the Go runtime selects an appropriate default value based on the
// number of logical CPUs on the machine, the process’s CPU affinity mask, and,
// on Linux, the process’s average CPU throughput limit based on cgroup CPU
// quota, if any.
//
// The Go runtime periodically updates the default value based on changes to
// the total logical CPU count, the CPU affinity mask, or cgroup quota. Setting
// a custom value with the GOMAXPROCS environment variable or by calling
// GOMAXPROCS disables automatic updates. The default value and automatic
// updates can be restored by calling [SetDefaultGOMAXPROCS].
//
// If GODEBUG=containermaxprocs=0 is set, GOMAXPROCS defaults to the value of
// [runtime.NumCPU]. If GODEBUG=updatemaxprocs=0 is set, the Go runtime does
// not perform automatic GOMAXPROCS updating.
//
// The default GOMAXPROCS behavior may change as the scheduler improves.
func GOMAXPROCS(n int) int

// SetDefaultGOMAXPROCS updates the GOMAXPROCS setting to the runtime
// default, as described by [GOMAXPROCS], ignoring the GOMAXPROCS
// environment variable.
//
// SetDefaultGOMAXPROCS can be used to enable the default automatic updating
// GOMAXPROCS behavior if it has been disabled by the GOMAXPROCS
// environment variable or a prior call to [GOMAXPROCS], or to force an immediate
// update if the caller is aware of a change to the total logical CPU count, CPU
// affinity mask or cgroup quota.
func SetDefaultGOMAXPROCS()

// NumCPU returns the number of logical CPUs usable by the current process.
//
// The set of available CPUs is checked by querying the operating system
// at process startup. Changes to operating system CPU allocation after
// process startup are not reflected.
func NumCPU() int

// NumCgoCall returns the number of cgo calls made by the current process.
func NumCgoCall() int64

// NumGoroutine returns the number of goroutines that currently exist.
func NumGoroutine() int
