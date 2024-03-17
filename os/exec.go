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

<<<<<<< HEAD
// ErrProcessDone は、プロセスが終了したことを示します。
var ErrProcessDone = errors.New("os: process already finished")

// ProcessはStartProcessによって作成されたプロセスに関する情報を格納します。
=======
// ErrProcessDone indicates a [Process] has finished.
var ErrProcessDone = errors.New("os: process already finished")

// Process stores the information about a process created by [StartProcess].
>>>>>>> upstream/master
type Process struct {
	Pid    int
	handle atomic.Uintptr
	isdone atomic.Bool
	sigMu  sync.RWMutex
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
<<<<<<< HEAD
// 返されるProcessは、基礎となるオペレーティングシステムのプロセスに関する情報を取得するために使用できます。
=======
// The [Process] it returns can be used to obtain information
// about the underlying operating system process.
>>>>>>> upstream/master
//
// Unixシステムでは、FindProcessは常に成功し、プロセスが存在するかどうかに関わらず、指定されたpidのProcessを返します。
// 実際にプロセスが存在するかどうかをテストするには、p.Signal(syscall.Signal(0))がエラーを報告するかどうかを確認してください。
func FindProcess(pid int) (*Process, error)

<<<<<<< HEAD
// StartProcessは、name、argv、attrで指定されたプログラム、引数、属性で新しいプロセスを開始します。
// argvスライスは新しいプロセスでos.Argsになるため、通常はプログラム名で始まります。
//
// 呼び出し元のgoroutineがruntime.LockOSThreadでオペレーティングシステムスレッドをロックし、継承可能なOSレベルのスレッド状態を変更した場合（例えば、LinuxやPlan 9の名前空間）、新しいプロセスは呼び出し元のスレッド状態を継承します。
//
// StartProcessは低レベルなインターフェースです。os/execパッケージはより高レベルなインターフェースを提供します。
//
// エラーが発生した場合、*PathError型となります。
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

// Releaseはプロセスpに関連付けられたリソースを解放し、将来使用できなくします。
// Waitが呼び出されない場合にのみReleaseを呼び出す必要があります。
func (p *Process) Release() error

// Killはプロセスを直ちに終了させます。Killはプロセスが実際に終了するのを待ちません。これによってプロセス自体のみを終了させるため、プロセスが起動した他のプロセスには影響を与えません。
func (p *Process) Kill() error

// Waitはプロセスの終了を待ち、その後、状態とエラー（あれば）を示すProcessStateを返します。
// Waitはプロセスに関連するリソースを解放します。
// ほとんどのオペレーティングシステムでは、プロセスは現在のプロセスの子であるか、エラーが返されます。
func (p *Process) Wait() (*ProcessState, error)

// SignalはProcessにシグナルを送信します。
// Windowsでは中断を送信することは実装されていません。
=======
// StartProcess starts a new process with the program, arguments and attributes
// specified by name, argv and attr. The argv slice will become [os.Args] in the
// new process, so it normally starts with the program name.
//
// If the calling goroutine has locked the operating system thread
// with [runtime.LockOSThread] and modified any inheritable OS-level
// thread state (for example, Linux or Plan 9 name spaces), the new
// process will inherit the caller's thread state.
//
// StartProcess is a low-level interface. The [os/exec] package provides
// higher-level interfaces.
//
// If there is an error, it will be of type [*PathError].
func StartProcess(name string, argv []string, attr *ProcAttr) (*Process, error)

// Release releases any resources associated with the [Process] p,
// rendering it unusable in the future.
// Release only needs to be called if [Process.Wait] is not.
func (p *Process) Release() error

// Kill causes the [Process] to exit immediately. Kill does not wait until
// the Process has actually exited. This only kills the Process itself,
// not any other processes it may have started.
func (p *Process) Kill() error

// Wait waits for the [Process] to exit, and then returns a
// ProcessState describing its status and an error, if any.
// Wait releases any resources associated with the Process.
// On most operating systems, the Process must be a child
// of the current process or an error will be returned.
func (p *Process) Wait() (*ProcessState, error)

// Signal sends a signal to the [Process].
// Sending [Interrupt] on Windows is not implemented.
>>>>>>> upstream/master
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

<<<<<<< HEAD
// Sysはプロセスに関するシステム依存の終了情報を返します。
// それを適切な基礎となる型に変換して、その内容にアクセスします。
// 例：Unixの場合、syscall.WaitStatusとして変換します。
func (p *ProcessState) Sys() any

// SysUsageは終了したプロセスのシステム依存のリソース使用状況情報を返します。それを適切な基に変換してください、例えばUnixでは*syscall.Rusage型など、その内容にアクセスするために。 (Unixでは、*syscall.Rusageはgetrusage(2)マニュアルページで定義されているstruct rusageに一致します。)
=======
// Sys returns system-dependent exit information about
// the process. Convert it to the appropriate underlying
// type, such as [syscall.WaitStatus] on Unix, to access its contents.
func (p *ProcessState) Sys() any

// SysUsage returns system-dependent resource usage information about
// the exited process. Convert it to the appropriate underlying
// type, such as [*syscall.Rusage] on Unix, to access its contents.
// (On Unix, *syscall.Rusage matches struct rusage as defined in the
// getrusage(2) manual page.)
>>>>>>> upstream/master
func (p *ProcessState) SysUsage() any
