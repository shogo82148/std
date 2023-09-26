// Copyright 2010 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ioutil

import (
	"github.com/shogo82148/std/os"
)

// Random number state, accessed without lock; racy but harmless.
// We generate random temporary file names so that there's a good
// chance the file doesn't exist yet - keeps the number of tries in
// TempFile to a minimum.

// TempFile creates a new temporary file in the directory dir
// with a name beginning with prefix, opens the file for reading
// and writing, and returns the resulting *os.File.
// If dir is the empty string, TempFile uses the default directory
// for temporary files (see os.TempDir).
// Multiple programs calling TempFile simultaneously
// will not choose the same file.  The caller can use f.Name()
// to find the name of the file.  It is the caller's responsibility to
// remove the file when no longer needed.
func TempFile(dir, prefix string) (f *os.File, err error)

// TempDir creates a new temporary directory in the directory dir
// with a name beginning with prefix and returns the path of the
// new directory.  If dir is the empty string, TempDir uses the
// default directory for temporary files (see os.TempDir).
// Multiple programs calling TempDir simultaneously
// will not choose the same directory.  It is the caller's responsibility
// to remove the directory when no longer needed.
func TempDir(dir, prefix string) (name string, err error)
