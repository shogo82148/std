// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Getwd returns an absolute path name corresponding to the
// current directory. If the current directory can be
// reached via multiple paths (due to symbolic links),
// Getwd may return any one of them.
//
// On Unix platforms, if the environment variable PWD
// provides an absolute name, and it is a name of the
// current directory, it is returned.
func Getwd() (dir string, err error)
