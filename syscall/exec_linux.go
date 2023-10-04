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

// SysProcIDMap holds Container ID to Host ID mappings used for User Namespaces in Linux.
// See user_namespaces(7).
type SysProcIDMap struct {
	ContainerID int
	HostID      int
	Size        int
}

type SysProcAttr struct {
	Chroot     string
	Credential *Credential
	// Ptrace tells the child to call ptrace(PTRACE_TRACEME).
	// Call runtime.LockOSThread before starting a process with this set,
	// and don't call UnlockOSThread until done with PtraceSyscall calls.
	Ptrace bool
	Setsid bool
	// Setpgid sets the process group ID of the child to Pgid,
	// or, if Pgid == 0, to the new child's process ID.
	Setpgid bool
	// Setctty sets the controlling terminal of the child to
	// file descriptor Ctty. Ctty must be a descriptor number
	// in the child process: an index into ProcAttr.Files.
	// This is only meaningful if Setsid is true.
	Setctty bool
	Noctty  bool
	Ctty    int
	// Foreground places the child process group in the foreground.
	// This implies Setpgid. The Ctty field must be set to
	// the descriptor of the controlling TTY.
	// Unlike Setctty, in this case Ctty must be a descriptor
	// number in the parent process.
	Foreground bool
	Pgid       int
	// Pdeathsig, if non-zero, is a signal that the kernel will send to
	// the child process when the creating thread dies. Note that the signal
	// is sent on thread termination, which may happen before process termination.
	// There are more details at https://go.dev/issue/27505.
	Pdeathsig    Signal
	Cloneflags   uintptr
	Unshareflags uintptr
	UidMappings  []SysProcIDMap
	GidMappings  []SysProcIDMap
	// GidMappingsEnableSetgroups enabling setgroups syscall.
	// If false, then setgroups syscall will be disabled for the child process.
	// This parameter is no-op if GidMappings == nil. Otherwise for unprivileged
	// users this should be set to false for mappings work.
	GidMappingsEnableSetgroups bool
	AmbientCaps                []uintptr
	UseCgroupFD                bool
	CgroupFD                   int
}
