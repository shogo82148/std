// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Stat returns a FileInfo structure describing the named file.
// If there is an error, it will be of type *PathError.
func Stat(name string) (FileInfo, error)

// Lstat returns the FileInfo structure describing the named file.
// If the file is a symbolic link (though Plan 9 does not have symbolic links),
// the returned FileInfo describes the symbolic link.  Lstat makes no attempt to follow the link.
// If there is an error, it will be of type *PathError.
func Lstat(name string) (FileInfo, error)
