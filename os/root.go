// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

import (
	"github.com/shogo82148/std/io/fs"
)

// OpenInRoot opens the file name in the directory dir.
// It is equivalent to OpenRoot(dir) followed by opening the file in the root.
//
// OpenInRoot returns an error if any component of the name
// references a location outside of dir.
//
// See [Root] for details and limitations.
func OpenInRoot(dir, name string) (*File, error)

// Root may be used to only access files within a single directory tree.
//
// Methods on Root can only access files and directories beneath a root directory.
// If any component of a file name passed to a method of Root references a location
// outside the root, the method returns an error.
// File names may reference the directory itself (.).
//
// Methods on Root will follow symbolic links, but symbolic links may not
// reference a location outside the root.
// Symbolic links must not be absolute.
//
// Methods on Root do not prohibit traversal of filesystem boundaries,
// Linux bind mounts, /proc special files, or access to Unix device files.
//
// Methods on Root are safe to be used from multiple goroutines simultaneously.
//
// On most platforms, creating a Root opens a file descriptor or handle referencing
// the directory. If the directory is moved, methods on Root reference the original
// directory in its new location.
//
// Root's behavior differs on some platforms:
//
//   - When GOOS=windows, file names may not reference Windows reserved device names
//     such as NUL and COM1.
//   - When GOOS=js, Root is vulnerable to TOCTOU (time-of-check-time-of-use)
//     attacks in symlink validation, and cannot ensure that operations will not
//     escape the root.
//   - When GOOS=plan9 or GOOS=js, Root does not track directories across renames.
//     On these platforms, a Root references a directory name, not a file descriptor.
type Root struct {
	root root
}

// OpenRoot opens the named directory.
// If there is an error, it will be of type *PathError.
func OpenRoot(name string) (*Root, error)

// Name returns the name of the directory presented to OpenRoot.
//
// It is safe to call Name after [Close].
func (r *Root) Name() string

// Close closes the Root.
// After Close is called, methods on Root return errors.
func (r *Root) Close() error

// Open opens the named file in the root for reading.
// See [Open] for more details.
func (r *Root) Open(name string) (*File, error)

// Create creates or truncates the named file in the root.
// See [Create] for more details.
func (r *Root) Create(name string) (*File, error)

// OpenFile opens the named file in the root.
// See [OpenFile] for more details.
//
// If perm contains bits other than the nine least-significant bits (0o777),
// OpenFile returns an error.
func (r *Root) OpenFile(name string, flag int, perm FileMode) (*File, error)

// OpenRoot opens the named directory in the root.
// If there is an error, it will be of type *PathError.
func (r *Root) OpenRoot(name string) (*Root, error)

// Mkdir creates a new directory in the root
// with the specified name and permission bits (before umask).
// See [Mkdir] for more details.
//
// If perm contains bits other than the nine least-significant bits (0o777),
// OpenFile returns an error.
func (r *Root) Mkdir(name string, perm FileMode) error

// Remove removes the named file or (empty) directory in the root.
// See [Remove] for more details.
func (r *Root) Remove(name string) error

// Stat returns a [FileInfo] describing the named file in the root.
// See [Stat] for more details.
func (r *Root) Stat(name string) (FileInfo, error)

// Lstat returns a [FileInfo] describing the named file in the root.
// If the file is a symbolic link, the returned FileInfo
// describes the symbolic link.
// See [Lstat] for more details.
func (r *Root) Lstat(name string) (FileInfo, error)

// FS returns a file system (an fs.FS) for the tree of files in the root.
//
// The result implements [io/fs.StatFS], [io/fs.ReadFileFS] and
// [io/fs.ReadDirFS].
func (r *Root) FS() fs.FS
