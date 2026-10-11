// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fork, exec, wait, etc.

package syscall

import (
	"github.com/shogo82148/std/sync"
)

// ForkLockはplan9では使用されません。
var ForkLock sync.RWMutex

// StringSlicePtrは、文字列のスライスをNUL終端されたバイト配列への
// ポインタのスライスに変換します。いずれかの文字列にNULバイトが含まれている場合、
// この関数はエラーを返す代わりにパニックします。
//
// Deprecated: 代わりにSlicePtrFromStringsを使用してください。
func StringSlicePtr(ss []string) []*byte

// SlicePtrFromStringsは、文字列のスライスをNUL終端されたバイト配列への
// ポインタのスライスに変換します。いずれかの文字列にNULバイトが含まれている場合、
// (nil, [EINVAL])を返します。
func SlicePtrFromStrings(ss []string) ([]*byte, error)

// name of the directory containing names and control files for all open file descriptors
var _ = BytePtrFromString("#d")

type ProcAttr struct {
	Dir   string
	Env   []string
	Files []uintptr
	Sys   *SysProcAttr
}

type SysProcAttr struct {
	Rfork int
}

// forkとexecの組み合わせで、スレッドセーフになるよう注意しています。
func ForkExec(argv0 string, argv []string, attr *ProcAttr) (pid int, err error)

// StartProcessは、osパッケージのために [ForkExec] をラップします。
func StartProcess(argv0 string, argv []string, attr *ProcAttr) (pid int, handle uintptr, err error)

// Ordinary exec.
func Exec(argv0 string, argv []string, envv []string) (err error)

// WaitProcessは、実行中のプロセスのpidがwaitメッセージのキューで
// 見つかるまで待機します。これは [ForkExec]/[StartProcess] と組み合わせて、
// 実行中のプロセスの終了を待つために使用されます。
func WaitProcess(pid int, w *Waitmsg) (err error)
