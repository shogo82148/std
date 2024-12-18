// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package os

// Stat returns the [FileInfo] structure describing file.
// If there is an error, it will be of type [*PathError].
func (file *File) Stat() (FileInfo, error)
