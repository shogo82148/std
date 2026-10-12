// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// debug パッケージには、プログラムが実行中に自己デバッグするための機能が含まれています。
package debug

import (
	"github.com/shogo82148/std/os"
)

// PrintStack は [runtime.Stack] が返すスタックトレースを標準エラーに出力します。
func PrintStack()

// Stackはそれを呼び出すgoroutineのフォーマットされたスタックトレースを返します。
// [runtime.Stack] を呼び出して、トレース全体をキャプチャする十分に大きなバッファーを使用します。
func Stack() []byte

// CrashOptionsは、致命的なクラッシュメッセージのフォーマットを制御するオプションを提供します。
type CrashOptions struct {
}

// SetCrashOutputは、標準エラーに加えて、処理されないパニックやその他の致命的なエラーが
// 出力される追加のファイルを1つ設定します。
// 追加のファイルは1つだけです。SetCrashOutputを再度呼び出すと、
// 以前の呼び出しが上書きされます。
// SetCrashOutputはfのファイルディスクリプターを複製するため、呼び出し元は
// SetCrashOutputが戻った直後にfを安全に閉じることができます。
// この追加のクラッシュ出力を無効にするには、SetCrashOutput(nil)を呼び出してください。
// クラッシュと並行して呼び出された場合、上書きするSetCrashOutputが戻った後でも、
// 進行中の出力の一部が古いファイルに書き込まれる可能性があります。
func SetCrashOutput(f *os.File, opts CrashOptions) error
