// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fs

// A SubFS is a file system with a Sub method.
type SubFS interface {
	FS

	Sub(dir string) (FS, error)
}

// Sub returns an FS corresponding to the subtree rooted at fsys's dir.
//
// If fs implements SubFS, Sub calls returns fsys.Sub(dir).
// Otherwise, if dir is ".", Sub returns fsys unchanged.
// Otherwise, Sub returns a new FS implementation sub that,
// in effect, implements sub.Open(dir) as fsys.Open(path.Join(dir, name)).
// The implementation also translates calls to ReadDir, ReadFile, and Glob appropriately.
//
// Note that Sub(os.DirFS("/"), "prefix") is equivalent to os.DirFS("/prefix")
// and that neither of them guarantees to avoid operating system
// accesses outside "/prefix", because the implementation of os.DirFS
// does not check for symbolic links inside "/prefix" that point to
// other directories. That is, os.DirFS is not a general substitute for a
// chroot-style security mechanism, and Sub does not change that fact.
func Sub(fsys FS, dir string) (FS, error)
