// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Fork, exec, wait, etc.

package syscall

import (
	"github.com/shogo82148/std/sync"
)

// ForkLockはWindowsでは使用されません。
var ForkLock sync.RWMutex

// EscapeArgは、https://msdn.microsoft.com/en-us/library/ms880421 で
// 規定されているとおりにコマンドライン引数sを書き換えます。
// sが空の場合、この関数は""（2つのダブルクォート）を返します。
// それ以外の場合、次の変換が行われます:
//   - すべてのバックスラッシュ（\）は、直後にダブルクォート（"）が続く場合にのみ
//     2つに重ねられます。
//   - すべてのダブルクォート（"）はバックスラッシュ（\）でエスケープされます。
//   - 最後に、s内にスペースまたはタブがある場合にのみ、sはダブルクォートで
//     囲まれます（arg -> "arg"）。
func EscapeArg(s string) string

func CloseOnExec(fd Handle)

func SetNonblock(fd Handle, nonblocking bool) (err error)

// FullPathは、指定されたファイルのフルパスを取得します。
func FullPath(name string) (path string, err error)

type ProcAttr struct {
	Dir   string
	Env   []string
	Files []uintptr
	Sys   *SysProcAttr
}

type SysProcAttr struct {
	HideWindow                 bool
	CmdLine                    string
	CreationFlags              uint32
	Token                      Token
	ProcessAttributes          *SecurityAttributes
	ThreadAttributes           *SecurityAttributes
	NoInheritHandles           bool
	AdditionalInheritedHandles []Handle
	ParentProcess              Handle
}

func StartProcess(argv0 string, argv []string, attr *ProcAttr) (pid int, handle uintptr, err error)

func Exec(argv0 string, argv []string, envv []string) (err error)
