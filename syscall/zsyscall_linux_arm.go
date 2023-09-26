// mksyscall.pl -l32 -arm syscall_linux.go syscall_linux_arm.go
// MACHINE GENERATED BY THE COMMAND ABOVE; DO NOT EDIT

package syscall

func Getcwd(buf []byte) (n int, err error)

func Access(path string, mode uint32) (err error)

func Acct(path string) (err error)

func Adjtimex(buf *Timex) (state int, err error)

func Chdir(path string) (err error)

func Chmod(path string, mode uint32) (err error)

func Chroot(path string) (err error)

func Close(fd int) (err error)

func Creat(path string, mode uint32) (fd int, err error)

func Dup(oldfd int) (fd int, err error)

func Dup2(oldfd int, newfd int) (err error)

func EpollCreate(size int) (fd int, err error)

func EpollCreate1(flag int) (fd int, err error)

func EpollCtl(epfd int, op int, fd int, event *EpollEvent) (err error)

func EpollWait(epfd int, events []EpollEvent, msec int) (n int, err error)

func Exit(code int)

func Faccessat(dirfd int, path string, mode uint32, flags int) (err error)

func Fallocate(fd int, mode uint32, off int64, len int64) (err error)

func Fchdir(fd int) (err error)

func Fchmod(fd int, mode uint32) (err error)

func Fchmodat(dirfd int, path string, mode uint32, flags int) (err error)

func Fchownat(dirfd int, path string, uid int, gid int, flags int) (err error)

func Fdatasync(fd int) (err error)

func Flock(fd int, how int) (err error)

func Fsync(fd int) (err error)

func Getdents(fd int, buf []byte) (n int, err error)

func Getpgid(pid int) (pgid int, err error)

func Getpgrp() (pid int)

func Getpid() (pid int)

func Getppid() (ppid int)

func Getrlimit(resource int, rlim *Rlimit) (err error)

func Getrusage(who int, rusage *Rusage) (err error)

func Gettid() (tid int)

func InotifyAddWatch(fd int, pathname string, mask uint32) (watchdesc int, err error)

func InotifyInit() (fd int, err error)

func InotifyInit1(flags int) (fd int, err error)

func InotifyRmWatch(fd int, watchdesc uint32) (success int, err error)

func Kill(pid int, sig Signal) (err error)

func Klogctl(typ int, buf []byte) (n int, err error)

func Link(oldpath string, newpath string) (err error)

func Mkdir(path string, mode uint32) (err error)

func Mkdirat(dirfd int, path string, mode uint32) (err error)

func Mknod(path string, mode uint32, dev int) (err error)

func Mknodat(dirfd int, path string, mode uint32, dev int) (err error)

func Nanosleep(time *Timespec, leftover *Timespec) (err error)

func Pause() (err error)

func PivotRoot(newroot string, putold string) (err error)

func Read(fd int, p []byte) (n int, err error)

func Readlink(path string, buf []byte) (n int, err error)

func Rename(oldpath string, newpath string) (err error)

func Renameat(olddirfd int, oldpath string, newdirfd int, newpath string) (err error)

func Rmdir(path string) (err error)

func Setdomainname(p []byte) (err error)

func Sethostname(p []byte) (err error)

func Setpgid(pid int, pgid int) (err error)

func Setrlimit(resource int, rlim *Rlimit) (err error)

func Setsid() (pid int, err error)

func Settimeofday(tv *Timeval) (err error)

func Setuid(uid int) (err error)

func Symlink(oldpath string, newpath string) (err error)

func Sync()

func Sysinfo(info *Sysinfo_t) (err error)

func Tee(rfd int, wfd int, len int, flags int) (n int64, err error)

func Tgkill(tgid int, tid int, sig Signal) (err error)

func Times(tms *Tms) (ticks uintptr, err error)

func Umask(mask int) (oldmask int)

func Uname(buf *Utsname) (err error)

func Unlink(path string) (err error)

func Unlinkat(dirfd int, path string) (err error)

func Unmount(target string, flags int) (err error)

func Unshare(flags int) (err error)

func Ustat(dev int, ubuf *Ustat_t) (err error)

func Utime(path string, buf *Utimbuf) (err error)

func Write(fd int, p []byte) (n int, err error)

func Madvise(b []byte, advice int) (err error)

func Mprotect(b []byte, prot int) (err error)

func Mlock(b []byte) (err error)

func Munlock(b []byte) (err error)

func Mlockall(flags int) (err error)

func Munlockall() (err error)

func Chown(path string, uid int, gid int) (err error)

func Fchown(fd int, uid int, gid int) (err error)

func Fstat(fd int, stat *Stat_t) (err error)

func Fstatfs(fd int, buf *Statfs_t) (err error)

func Getegid() (egid int)

func Geteuid() (euid int)

func Getgid() (gid int)

func Getuid() (uid int)

func Lchown(path string, uid int, gid int) (err error)

func Listen(s int, n int) (err error)

func Lstat(path string, stat *Stat_t) (err error)

func Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)

func Select(nfd int, r *FdSet, w *FdSet, e *FdSet, timeout *Timeval) (n int, err error)

func Setfsgid(gid int) (err error)

func Setfsuid(uid int) (err error)

func Setgid(gid int) (err error)

func Setregid(rgid int, egid int) (err error)

func Setresgid(rgid int, egid int, sgid int) (err error)

func Setresuid(ruid int, euid int, suid int) (err error)

func Setreuid(ruid int, euid int) (err error)

func Shutdown(fd int, how int) (err error)

func Splice(rfd int, roff *int64, wfd int, woff *int64, len int, flags int) (n int, err error)

func Stat(path string, stat *Stat_t) (err error)

func Statfs(path string, buf *Statfs_t) (err error)

func Gettimeofday(tv *Timeval) (err error)

func Time(t *Time_t) (tt Time_t, err error)

func Pread(fd int, p []byte, offset int64) (n int, err error)

func Pwrite(fd int, p []byte, offset int64) (n int, err error)

func Truncate(path string, length int64) (err error)

func Ftruncate(fd int, length int64) (err error)
