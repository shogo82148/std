// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build dragonfly || freebsd || netbsd || (openbsd && mips64)

package unix

func Eaccess(path string, mode uint32) error
