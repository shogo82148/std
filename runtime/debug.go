// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

<<<<<<< HEAD
// GOMAXPROCSは同時に実行できる最大CPU数を設定し、前の設定を返します。デフォルトはruntime.NumCPUの値です。nが1未満の場合、現在の設定は変更されません。スケジューラの改善が行われると、この呼び出しはなくなります。
=======
// GOMAXPROCS sets the maximum number of CPUs that can be executing
// simultaneously and returns the previous setting. It defaults to
// the value of [runtime.NumCPU]. If n < 1, it does not change the current setting.
// This call will go away when the scheduler improves.
>>>>>>> upstream/master
func GOMAXPROCS(n int) int

// NumCPUは現在のプロセスで使用可能な論理CPUの数を返します。
//
// 利用可能なCPUのセットはプロセスの起動時にオペレーティングシステムによって確認されます。
// プロセスの起動後にオペレーティングシステムのCPU割り当てに変更があっても、それは反映されません。
func NumCPU() int

// NumCgoCall は現在のプロセスによって行われた cgo 呼び出しの数を返します。
func NumCgoCall() int64

// NumGoroutineは現在存在するゴルーチンの数を返します。
func NumGoroutine() int
