// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux

package unix

func Fchmodat(dirfd int, path string, mode uint32, flags int) error
