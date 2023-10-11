// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

package os

import (
	"github.com/shogo82148/std/syscall"
)

// すべてのシステムのosパッケージで保証されている信号値は、os.Interrupt（プロセスに割り込みを送信する）とos.Kill（プロセスを強制的に終了する）のみです。Windowsでは、os.Process.Signalに対してos.Interruptを送信することは実装されていません。代わりにエラーが返されます。
var (
	Interrupt Signal = syscall.SIGINT
	Kill      Signal = syscall.SIGKILL
)

// ProcessStateはWaitによって報告されるプロセスの情報を格納します。
type ProcessState struct {
	pid    int
	status syscall.WaitStatus
	rusage *syscall.Rusage
}

// Pidは終了したプロセスのプロセスIDを返します。
func (p *ProcessState) Pid() int

func (p *ProcessState) String() string

// ExitCodeは終了したプロセスの終了コードを返します。プロセスがまだ終了していない場合や、シグナルによって終了した場合は-1を返します。
func (p *ProcessState) ExitCode() int
