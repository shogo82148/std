// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build wasip1

package syscall

const (
	LOOKUP_SYMLINK_FOLLOW = 0x00000001
)

const (
	OFLAG_CREATE    = 0x0001
	OFLAG_DIRECTORY = 0x0002
	OFLAG_EXCL      = 0x0004
	OFLAG_TRUNC     = 0x0008
)

const (
	FDFLAG_APPEND   = 0x0001
	FDFLAG_DSYNC    = 0x0002
	FDFLAG_NONBLOCK = 0x0004
	FDFLAG_RSYNC    = 0x0008
	FDFLAG_SYNC     = 0x0010
)

const (
	RIGHT_FD_DATASYNC = 1 << iota
	RIGHT_FD_READ
	RIGHT_FD_SEEK
	RIGHT_FDSTAT_SET_FLAGS
	RIGHT_FD_SYNC
	RIGHT_FD_TELL
	RIGHT_FD_WRITE
	RIGHT_FD_ADVISE
	RIGHT_FD_ALLOCATE
	RIGHT_PATH_CREATE_DIRECTORY
	RIGHT_PATH_CREATE_FILE
	RIGHT_PATH_LINK_SOURCE
	RIGHT_PATH_LINK_TARGET
	RIGHT_PATH_OPEN
	RIGHT_FD_READDIR
	RIGHT_PATH_READLINK
	RIGHT_PATH_RENAME_SOURCE
	RIGHT_PATH_RENAME_TARGET
	RIGHT_PATH_FILESTAT_GET
	RIGHT_PATH_FILESTAT_SET_SIZE
	RIGHT_PATH_FILESTAT_SET_TIMES
	RIGHT_FD_FILESTAT_GET
	RIGHT_FD_FILESTAT_SET_SIZE
	RIGHT_FD_FILESTAT_SET_TIMES
	RIGHT_PATH_SYMLINK
	RIGHT_PATH_REMOVE_DIRECTORY
	RIGHT_PATH_UNLINK_FILE
	RIGHT_POLL_FD_READWRITE
	RIGHT_SOCK_SHUTDOWN
	RIGHT_SOCK_ACCEPT
)

const (
	WHENCE_SET = 0
	WHENCE_CUR = 1
	WHENCE_END = 2
)

const (
	FILESTAT_SET_ATIM     = 0x0001
	FILESTAT_SET_ATIM_NOW = 0x0002
	FILESTAT_SET_MTIM     = 0x0004
	FILESTAT_SET_MTIM_NOW = 0x0008
)

func Open(path string, openmode int, perm uint32) (int, error)

func Close(fd int) error

func CloseOnExec(fd int)

func Mkdir(path string, perm uint32) error

func ReadDir(fd int, buf []byte, cookie dircookie) (int, error)

type Stat_t struct {
	Dev      uint64
	Ino      uint64
	Filetype uint8
	Nlink    uint64
	Size     uint64
	Atime    uint64
	Mtime    uint64
	Ctime    uint64

	Mode int

	// Uid and Gid are always zero on wasip1 platforms
	Uid uint32
	Gid uint32
}

func Stat(path string, st *Stat_t) error

func Lstat(path string, st *Stat_t) error

func Fstat(fd int, st *Stat_t) error

func Unlink(path string) error

func Rmdir(path string) error

func Chmod(path string, mode uint32) error

func Fchmod(fd int, mode uint32) error

func Chown(path string, uid, gid int) error

func Fchown(fd int, uid, gid int) error

func Lchown(path string, uid, gid int) error

func UtimesNano(path string, ts []Timespec) error

func Rename(from, to string) error

func Truncate(path string, length int64) error

func Ftruncate(fd int, length int64) error

const ImplementsGetwd = true

func Getwd() (string, error)

func Chdir(path string) error

func Readlink(path string, buf []byte) (n int, err error)

func Link(path, link string) error

func Symlink(path, link string) error

func Fsync(fd int) error

func Read(fd int, b []byte) (int, error)

func Write(fd int, b []byte) (int, error)

func Pread(fd int, b []byte, offset int64) (int, error)

func Pwrite(fd int, b []byte, offset int64) (int, error)

func Seek(fd int, offset int64, whence int) (int64, error)

func Dup(fd int) (int, error)

func Dup2(fd, newfd int) error

func Pipe(fd []int) error

func RandomGet(b []byte) error
