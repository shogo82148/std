// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The working directory in Plan 9 is effectively per P, so different
// goroutines and even the same goroutine as it's rescheduled on
// different Ps can see different working directories.
//
// Instead, track a Go process-wide intent of the current working directory,
// and switch to it at important points.

package syscall

// このゴルーチンから見えるカレントワーキングディレクトリが、いずれかのゴルーチンで
// 最後に呼び出された[Chdir]と一致することを保証します。これは、相対パス名を使用する
// システムコールを実行する前に内部的に呼び出されます。システムコールが実行される前に
// （異なるワーキングディレクトリを持つ可能性のある）別のスレッドに再スケジュールされるのを
// 防ぐため、ゴルーチンをOSスレッドにロックした状態で呼び出す必要があります。
func Fixwd()

func Getwd() (wd string, err error)

func Chdir(path string) error
