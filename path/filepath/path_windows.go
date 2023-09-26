// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepath

// IsAbs reports whether the path is absolute.
func IsAbs(path string) (b bool)

// HasPrefix exists for historical compatibility and should not be used.
func HasPrefix(p, prefix string) bool
