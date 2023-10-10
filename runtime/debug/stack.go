// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// debug パッケージには、プログラムが実行中に自己デバッグするための機能が含まれています。
package debug

// PrintStackはruntime.Stackによって返されたスタックトレースを標準エラー出力に出力します。
func PrintStack()

// Stackはそれを呼び出すgoroutineのフォーマットされたスタックトレースを返します。
// runtime.Stackを呼び出して、トレース全体をキャプチャする十分に大きなバッファを使用します。
func Stack() []byte
