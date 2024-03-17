// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

<<<<<<< HEAD
// Executableは、現在のプロセスを開始した実行可能ファイルのパス名を返します。
// パスがまだ正しい実行可能ファイルを指しているとは限りません。
// プロセスの開始にシンボリックリンクが使用された場合、オペレーティングシステムによって結果は
// シンボリックリンクまたはそれが指していたパスになる可能性があります。
// 安定した結果が必要な場合は、path/filepath.EvalSymlinksが役立ちます。
=======
// Executable returns the path name for the executable that started
// the current process. There is no guarantee that the path is still
// pointing to the correct executable. If a symlink was used to start
// the process, depending on the operating system, the result might
// be the symlink or the path it pointed to. If a stable result is
// needed, [path/filepath.EvalSymlinks] might help.
>>>>>>> upstream/master
//
// Executableは、エラーが発生しない限り、絶対パスを返します。
//
// 主な利用ケースは、実行可能ファイルに対して相対的に配置されたリソースを見つけることです。
func Executable() (string, error)
