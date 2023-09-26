// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepath

// IsAbs returns true if the path is absolute.
func IsAbs(path string) (b bool)

// VolumeName returns leading volume name.
// Given "C:\foo\bar" it returns "C:" under windows.
// Given "\\host\share\foo" it returns "\\host\share".
// On other platforms it returns "".
func VolumeName(path string) (v string)

// HasPrefix exists for historical compatibility and should not be used.
func HasPrefix(p, prefix string) bool
