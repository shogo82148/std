// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package syscall

// Linux unshare/clone/clone2/clone3 flags, architecture-independent,
// copied from linux/sched.h.
const (
	CLONE_VM             = 0x00000100
	CLONE_FS             = 0x00000200
	CLONE_FILES          = 0x00000400
	CLONE_SIGHAND        = 0x00000800
	CLONE_PIDFD          = 0x00001000
	CLONE_PTRACE         = 0x00002000
	CLONE_VFORK          = 0x00004000
	CLONE_PARENT         = 0x00008000
	CLONE_THREAD         = 0x00010000
	CLONE_NEWNS          = 0x00020000
	CLONE_SYSVSEM        = 0x00040000
	CLONE_SETTLS         = 0x00080000
	CLONE_PARENT_SETTID  = 0x00100000
	CLONE_CHILD_CLEARTID = 0x00200000
	CLONE_DETACHED       = 0x00400000
	CLONE_UNTRACED       = 0x00800000
	CLONE_CHILD_SETTID   = 0x01000000
	CLONE_NEWCGROUP      = 0x02000000
	CLONE_NEWUTS         = 0x04000000
	CLONE_NEWIPC         = 0x08000000
	CLONE_NEWUSER        = 0x10000000
	CLONE_NEWPID         = 0x20000000
	CLONE_NEWNET         = 0x40000000
	CLONE_IO             = 0x80000000

	CLONE_CLEAR_SIGHAND = 0x100000000
	CLONE_INTO_CGROUP   = 0x200000000

	CLONE_NEWTIME = 0x00000080
)

// SysProcIDMapは、Linuxのユーザー名前空間で使用される、コンテナIDからホストIDへの
// マッピングを保持します。user_namespaces(7)を参照してください。
//
// ユーザー名前空間は、多くの一般的なLinuxのバージョンでは（セキュリティ上の問題により）
// 利用できないか、利用可能であってもUbuntu 24.04のようにAppArmorの制限を受けることに
// 注意してください。
type SysProcIDMap struct {
	ContainerID int
	HostID      int
	Size        int
}

type SysProcAttr struct {
	Chroot     string
	Credential *Credential
	// Ptraceは、子にptrace(PTRACE_TRACEME)を呼び出すよう指示します。
	// これを設定してプロセスを開始する前にruntime.LockOSThreadを呼び出し、
	// PtraceSyscallの呼び出しが完了するまでUnlockOSThreadを呼び出さないでください。
	Ptrace bool
	Setsid bool
	// Setpgidは、子のプロセスグループIDをPgidに設定します。
	// Pgid == 0の場合は、新しい子のプロセスIDに設定します。
	Setpgid bool
	// Setcttyは、子の制御端末をファイルディスクリプタCttyに設定します。
	// Cttyは子プロセス内のディスクリプタ番号、つまりProcAttr.Filesへの
	// インデックスでなければなりません。
	// これはSetsidがtrueの場合にのみ意味を持ちます。
	Setctty bool
	Noctty  bool
	Ctty    int
	// Foregroundは、子プロセスグループをフォアグラウンドに配置します。
	// これはSetpgidを暗黙的に含みます。Cttyフィールドには
	// 制御TTYのディスクリプタを設定しなければなりません。
	// Setcttyとは異なり、この場合Cttyは親プロセス内の
	// ディスクリプタ番号でなければなりません。
	Foreground bool
	Pgid       int
	// Pdeathsigがゼロでない場合、作成したスレッドが終了したときにカーネルが
	// 子プロセスに送信するシグナルです。シグナルはスレッドの終了時に送信され、
	// それはプロセスの終了前に発生する可能性があることに注意してください。
	// 詳細は https://go.dev/issue/27505 にあります。
	Pdeathsig    Signal
	Cloneflags   uintptr
	Unshareflags uintptr
	UidMappings  []SysProcIDMap
	GidMappings  []SysProcIDMap
	// GidMappingsEnableSetgroupsはsetgroupsシステムコールを有効にします。
	// falseの場合、子プロセスではsetgroupsシステムコールが無効になります。
	// このパラメータは、GidMappings == nilの場合は何もしません。それ以外の場合、
	// 特権を持たないユーザーでマッピングを機能させるにはfalseに設定するべきです。
	GidMappingsEnableSetgroups bool
	AmbientCaps                []uintptr
	UseCgroupFD                bool
	CgroupFD                   int
	// PidFDがnilでない場合、カーネルがその機能をサポートしていれば子のpidfdを、
	// そうでなければ-1を格納するために使用されます。*PidFDは
	// プロセスが正常に開始した場合にのみ変更されることに注意してください。
	PidFD *int
}
