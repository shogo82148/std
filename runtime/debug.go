// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

<<<<<<< HEAD
// GOMAXPROCSは同時に実行できる最大CPU数を設定し、前の設定を返します。デフォルトは [runtime.NumCPU] の値です。nが1未満の場合、現在の設定は変更されません。スケジューラの改善が行われると、この呼び出しはなくなります。
func GOMAXPROCS(n int) int

// NumCPUは現在のプロセスで使用可能な論理CPUの数を返します。
=======
// GOMAXPROCS sets the maximum number of CPUs that can be executing
// simultaneously and returns the previous setting. If n < 1, it does not change
// the current setting.
//
// # Default
//
// If the GOMAXPROCS environment variable is set to a positive whole number,
// GOMAXPROCS defaults to that value.
//
// Otherwise, the Go runtime selects an appropriate default value from a combination of
//   - the number of logical CPUs on the machine,
//   - the process’s CPU affinity mask,
//   - and, on Linux, the process’s average CPU throughput limit based on cgroup CPU
//     quota, if any.
//
// If GODEBUG=containermaxprocs=0 is set and GOMAXPROCS is not set by the
// environment variable, then GOMAXPROCS instead defaults to the value of
// [runtime.NumCPU]. Note that GODEBUG=containermaxprocs=0 is [default] for
// language version 1.24 and below.
//
// # Updates
//
// The Go runtime periodically updates the default value based on changes to
// the total logical CPU count, the CPU affinity mask, or cgroup quota. Setting
// a custom value with the GOMAXPROCS environment variable or by calling
// GOMAXPROCS disables automatic updates. The default value and automatic
// updates can be restored by calling [SetDefaultGOMAXPROCS].
//
// If GODEBUG=updatemaxprocs=0 is set, the Go runtime does not perform
// automatic GOMAXPROCS updating. Note that GODEBUG=updatemaxprocs=0 is
// [default] for language version 1.24 and below.
//
// # Compatibility
//
// Note that the default GOMAXPROCS behavior may change as the scheduler
// improves, especially the implementation detail below.
//
// # Implementation details
//
// When computing default GOMAXPROCS via cgroups, the Go runtime computes the
// "average CPU throughput limit" as the cgroup CPU quota / period. In cgroup
// v2, these values come from the cpu.max file. In cgroup v1, they come from
// cpu.cfs_quota_us and cpu.cfs_period_us, respectively. In container runtimes
// that allow configuring CPU limits, this value usually corresponds to the
// "CPU limit" option, not "CPU request".
//
// The Go runtime typically selects the default GOMAXPROCS as the minimum of
// the logical CPU count, the CPU affinity mask count, or the cgroup CPU
// throughput limit. However, it will never set GOMAXPROCS less than 2 unless
// the logical CPU count or CPU affinity mask count are below 2.
//
// If the cgroup CPU throughput limit is not a whole number, the Go runtime
// rounds up to the next whole number.
//
// GOMAXPROCS updates are performed up to once per second, or less if the
// application is idle.
//
// [default]: https://go.dev/doc/godebug#default
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
>>>>>>> upstream/release-branch.go1.25
//
// 利用可能なCPUのセットはプロセスの起動時にオペレーティングシステムによって確認されます。
// プロセスの起動後にオペレーティングシステムのCPU割り当てに変更があっても、それは反映されません。
func NumCPU() int

// NumCgoCall は現在のプロセスによって行われた cgo 呼び出しの数を返します。
func NumCgoCall() int64

// NumGoroutineは現在存在するゴルーチンの数を返します。
func NumGoroutine() int
