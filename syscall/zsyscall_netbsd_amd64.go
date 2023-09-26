// mksyscall.pl -netbsd -tags netbsd,amd64 syscall_bsd.go syscall_netbsd.go syscall_netbsd_amd64.go
// Code generated by the command above; DO NOT EDIT.

//go:build netbsd && amd64
// +build netbsd,amd64

package syscall

func Shutdown(s int, how int) (err error)

func Access(path string, mode uint32) (err error)

func Adjtime(delta *Timeval, olddelta *Timeval) (err error)

func Chdir(path string) (err error)

func Chflags(path string, flags int) (err error)

func Chmod(path string, mode uint32) (err error)

func Chown(path string, uid int, gid int) (err error)

func Chroot(path string) (err error)

func Close(fd int) (err error)

func Dup(fd int) (nfd int, err error)

func Dup2(from int, to int) (err error)

func Fchdir(fd int) (err error)

func Fchflags(fd int, flags int) (err error)

func Fchmod(fd int, mode uint32) (err error)

func Fchown(fd int, uid int, gid int) (err error)

func Flock(fd int, how int) (err error)

func Fpathconf(fd int, name int) (val int, err error)

func Fstat(fd int, stat *Stat_t) (err error)

func Fsync(fd int) (err error)

func Ftruncate(fd int, length int64) (err error)

func Getegid() (egid int)

func Geteuid() (uid int)

func Getgid() (gid int)

func Getpgid(pid int) (pgid int, err error)

func Getpgrp() (pgrp int)

func Getpid() (pid int)

func Getppid() (ppid int)

func Getpriority(which int, who int) (prio int, err error)

func Getrlimit(which int, lim *Rlimit) (err error)

func Getrusage(who int, rusage *Rusage) (err error)

func Getsid(pid int) (sid int, err error)

func Gettimeofday(tv *Timeval) (err error)

func Getuid() (uid int)

func Issetugid() (tainted bool)

func Kill(pid int, signum Signal) (err error)

func Kqueue() (fd int, err error)

func Lchown(path string, uid int, gid int) (err error)

func Link(path string, link string) (err error)

func Listen(s int, backlog int) (err error)

func Lstat(path string, stat *Stat_t) (err error)

func Mkdir(path string, mode uint32) (err error)

func Mkfifo(path string, mode uint32) (err error)

func Mknod(path string, mode uint32, dev int) (err error)

func Nanosleep(time *Timespec, leftover *Timespec) (err error)

func Open(path string, mode int, perm uint32) (fd int, err error)

func Pathconf(path string, name int) (val int, err error)

func Pread(fd int, p []byte, offset int64) (n int, err error)

func Pwrite(fd int, p []byte, offset int64) (n int, err error)

func Readlink(path string, buf []byte) (n int, err error)

func Rename(from string, to string) (err error)

func Revoke(path string) (err error)

func Rmdir(path string) (err error)

func Seek(fd int, offset int64, whence int) (newoffset int64, err error)

func Select(n int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (err error)

func Setegid(egid int) (err error)

func Seteuid(euid int) (err error)

func Setgid(gid int) (err error)

func Setpgid(pid int, pgid int) (err error)

func Setpriority(which int, who int, prio int) (err error)

func Setregid(rgid int, egid int) (err error)

func Setreuid(ruid int, euid int) (err error)

func Setrlimit(which int, lim *Rlimit) (err error)

func Setsid() (pid int, err error)

func Settimeofday(tp *Timeval) (err error)

func Setuid(uid int) (err error)

func Stat(path string, stat *Stat_t) (err error)

func Symlink(path string, link string) (err error)

func Sync() (err error)

func Truncate(path string, length int64) (err error)

func Umask(newmask int) (oldmask int)

func Unlink(path string) (err error)

func Unmount(path string, flags int) (err error)
