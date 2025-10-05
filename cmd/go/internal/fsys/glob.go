// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package fsys

// Glob is like filepath.Glob but uses the overlay file system.
func Glob(pattern string) (matches []string, err error)
