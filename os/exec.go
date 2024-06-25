// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/sync"
	"github.com/shogo82148/std/sync/atomic"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// ErrProcessDone は、[Process] が終了したことを示します。
var ErrProcessDone = errors.New("os: process already finished")

// Processは [StartProcess] によって作成されたプロセスに関する情報を格納します。
type Process struct {
	Pid int

	mode processMode

	// State contains the atomic process state.
	//
	// In modePID, this consists only of the processStatus fields, which
	// indicate if the process is done/released.
	//
	// In modeHandle, the lower bits also contain a reference count for the
	// handle field.
	//
	// The Process itself initially holds 1 persistent reference. Any
	// operation that uses the handle with a system call temporarily holds
	// an additional transient reference. This prevents the handle from
	// being closed prematurely, which could result in the OS allocating a
	// different handle with the same value, leading to Process' methods
	// operating on the wrong process.
	//
	// Release and Wait both drop the Process' persistent reference, but
	// other concurrent references may delay actually closing the handle
	// because they hold a transient reference.
	//
	// Regardless, we want new method calls to immediately treat the handle
	// as unavailable after Release or Wait to avoid extending this delay.
	// This is achieved by setting either processStatus flag when the
	// Process' persistent reference is dropped. The only difference in the
	// flags is the reason the handle is unavailable, which affects the
	// errors returned by concurrent calls.
	state atomic.Uint64

	// Used only in modePID.
	sigMu sync.RWMutex

	// handle is the OS handle for process actions, used only in
	// modeHandle.
	//
	// handle must be accessed only via the handleTransientAcquire method
	// (or during closeHandle), not directly! handle is immutable.
	//
	// On Windows, it is a handle from OpenProcess.
	// On Linux, it is a pidfd.
	// It is unused on other GOOSes.
	handle uintptr
}

// ProcAttrはStartProcessによって開始される新しいプロセスに適用される属性を保持します。
type ProcAttr struct {

	// Dir が空でない場合、子プロセスを作成する前にディレクトリに変更します。
	Dir string

	// もしEnvがnilでない場合、それは新しいプロセスの環境変数をEnvironによって返される形式で指定します。
	// もしEnvがnilである場合、Environの結果が使用されます。
	Env []string

	// Filesは新しいプロセスに引き継がれるオープンファイルを指定します。最初の3つのエントリは標準入力、標準出力、標準エラーに対応します。実装は、基になるオペレーティングシステムに応じて、追加のエントリをサポートすることがあります。nilのエントリは、そのファイルがプロセスの開始時に閉じられることを意味します。
	// Unixシステムでは、StartProcessはこれらのFile値をブロッキングモードに変更します。つまり、SetDeadlineは動作しなくなり、Closeを呼び出してもReadまたはWriteが中断されません。
	Files []*File

	// オペレーティングシステム固有のプロセス作成属性です。
	// このフィールドを設定すると、プログラムが正常に実行されない場合や、
	// 一部のオペレーティングシステムではコンパイルすらできないことがあります。
	Sys *syscall.SysProcAttr
}

// Signalはオペレーティングシステムのシグナルを表します。
// 通常、基礎となる実装はオペレーティングシステムに依存します：
// Unixではsyscall.Signalです。
type Signal interface {
	String() string
	Signal()
}

// Getpidは呼び出し元のプロセスIDを返します。
func Getpid() int

// Getppidは呼び出し元の親プロセスのプロセスIDを返します。
func Getppid() int

// FindProcessは、pidによって実行中のプロセスを検索します。
//
// 返される [Process] は、基礎となるオペレーティングシステムのプロセスに関する情報を取得するために使用できます。
//
// Unixシステムでは、FindProcessは常に成功し、プロセスが存在するかどうかに関わらず、指定されたpidのProcessを返します。
// 実際にプロセスが存在するかどうかをテストするには、p.Signal(syscall.Signal(0))がエラーを報告するかどうかを確認してください。
func FindProcess(pid int) (*Process, error)

// StartProcessは、name、argv、attrで指定されたプログラム、引数、属性で新しいプロセスを開始します。
// argvスライスは新しいプロセスで [os.Args] になるため、通常はプログラム名で始まります。
//
// 呼び出し元のgoroutineが [runtime.LockOSThread] でオペレーティングシステムスレッドをロックし、継承可能なOSレベルのスレッド状態を変更した場合（例えば、LinuxやPlan 9の名前空間）、新しいプロセスは呼び出し元のスレッド状態を継承します。
//
// StartProcessは低レベルなインターフェースです。[os/exec] パッケージはより高レベルなインターフェースを提供します。
//
// エラーが発生した場合、[*PathError] 型となります。
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

// Releaseは [Process] pに関連付けられたリソースを解放し、将来使用できなくします。
// [Process.Wait] が呼び出されない場合にのみReleaseを呼び出す必要があります。
func (p *Process) Release() error

// Killは [Process] を直ちに終了させます。Killはプロセスが実際に終了するのを待ちません。これによってプロセス自体のみを終了させるため、プロセスが起動した他のプロセスには影響を与えません。
func (p *Process) Kill() error

// Waitは [Process] の終了を待ち、その後、状態とエラー（あれば）を示すProcessStateを返します。
// Waitはプロセスに関連するリソースを解放します。
// ほとんどのオペレーティングシステムでは、プロセスは現在のプロセスの子であるか、エラーが返されます。
func (p *Process) Wait() (*ProcessState, error)

// Signalは [Process] にシグナルを送信します。
// Windowsでは [Interrupt] を送信することは実装されていません。
func (p *Process) Signal(sig Signal) error

// UserTimeは終了したプロセスおよびその子プロセスのユーザーCPU時間を返します。
func (p *ProcessState) UserTime() time.Duration

// SystemTimeは終了したプロセスとその子プロセスのシステムCPU時間を返します。
func (p *ProcessState) SystemTime() time.Duration

// Exited はプログラムが終了したかどうかを報告します。
// Unixシステムでは、この関数はプログラムが exit を呼び出して終了した場合に true を返し、
// シグナルによってプログラムが終了した場合には false を返します。
func (p *ProcessState) Exited() bool

// Successは、プログラムが正常に終了したかどうかを報告します。
// たとえば、Unixでは終了ステータス0で終了した場合などです。
func (p *ProcessState) Success() bool

// Sysはプロセスに関するシステム依存の終了情報を返します。
// それを適切な基礎となる型に変換して、その内容にアクセスします。
// 例：Unixの場合、[syscall.WaitStatus] として変換します。
func (p *ProcessState) Sys() any

// SysUsageは終了したプロセスのシステム依存のリソース使用状況情報を返します。それを適切な基に変換してください、例えばUnixでは [*syscall.Rusage] 型など、その内容にアクセスするために。 (Unixでは、*syscall.Rusageはgetrusage(2)マニュアルページで定義されているstruct rusageに一致します。)
func (p *ProcessState) SysUsage() any
