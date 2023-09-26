// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Auxiliary information if the File describes a directory

// allowReadDirFileID indicates whether File.readdir should try to use FILE_ID_BOTH_DIR_INFO
// if the underlying file system supports it.
// Useful for testing purposes.
