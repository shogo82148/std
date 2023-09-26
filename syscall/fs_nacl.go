// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// A simulated Unix-like file system for use within NaCl.
//
// The simulation is not particularly tied to NaCl other than the reuse
// of NaCl's definition for the Stat_t structure.
//
// The file system need never be written to disk, so it is represented as
// in-memory Go data structures, never in a serialized form.
//
// TODO: Perhaps support symlinks, although they muck everything up.

package syscall

// An fsys is a file system.
// Since there is no I/O (everything is in memory),
// the global lock mu protects the whole file system state,
// and that's okay.

// A devFile is the implementation required of device files
// like /dev/null or /dev/random.

// An inode is a (possibly special) file in the file system.

// A dirent describes a single directory entry.

// An fsysFile is the fileImpl implementation backed by the file system.

func ReadDirent(fd int, buf []byte) (int, error)

func Open(path string, openmode int, perm uint32) (fd int, err error)

func Mkdir(path string, perm uint32) error

func Getcwd(buf []byte) (n int, err error)

func Stat(path string, st *Stat_t) error

func Lstat(path string, st *Stat_t) error

func Unlink(path string) error

func Rmdir(path string) error

func Chmod(path string, mode uint32) error

func Fchmod(fd int, mode uint32) error

func Chown(path string, uid, gid int) error

func Fchown(fd int, uid, gid int) error

func Lchown(path string, uid, gid int) error

func UtimesNano(path string, ts []Timespec) error

func Link(path, link string) error

func Rename(from, to string) error

func Truncate(path string, length int64) error

func Ftruncate(fd int, length int64) error

func Chdir(path string) error

func Fchdir(fd int) error

func Readlink(path string, buf []byte) (n int, err error)

func Symlink(path, link string) error

func Fsync(fd int) error
