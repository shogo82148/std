// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// debug パッケージには、プログラムが実行中に自己デバッグするための機能が含まれています。
package debug

import (
	"github.com/shogo82148/std/os"
)

// PrintStackはruntime.Stackによって返されたスタックトレースを標準エラー出力に出力します。
func PrintStack()

// Stackはそれを呼び出すgoroutineのフォーマットされたスタックトレースを返します。
// [runtime.Stack] を呼び出して、トレース全体をキャプチャする十分に大きなバッファを使用します。
func Stack() []byte

// SetCrashOutput configures a single additional file where unhandled
// panics and other fatal errors are printed, in addition to standard error.
// There is only one additional file: calling SetCrashOutput again overrides
// any earlier call.
// SetCrashOutput duplicates f's file descriptor, so the caller may safely
// close f as soon as SetCrashOutput returns.
// To disable this additional crash output, call SetCrashOutput(nil).
// If called concurrently with a crash, some in-progress output may be written
// to the old file even after an overriding SetCrashOutput returns.
func SetCrashOutput(f *os.File) error
