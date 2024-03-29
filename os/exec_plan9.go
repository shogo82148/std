// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/syscall"
)

// osパッケージですべてのシステムで存在が保証されている信号値は、Interrupt（プロセスに割り込みを送信する）とKill（プロセスを強制終了する）です。InterruptはWindowsでは実装されていません。[os.Process.Signal] で使用するとエラーが返されます。
var (
	Interrupt Signal = syscall.Note("interrupt")
	Kill      Signal = syscall.Note("kill")
)

// ProcessStateは、Waitによって報告されるプロセスに関する情報を格納します。
type ProcessState struct {
	pid    int
	status *syscall.Waitmsg
}

// Pidは終了したプロセスのプロセスIDを返します。
func (p *ProcessState) Pid() int

func (p *ProcessState) String() string

// ExitCodeは終了したプロセスの終了コードを返します。もしプロセスが終了していないか、シグナルで終了された場合は-1が返ります。
func (p *ProcessState) ExitCode() int
