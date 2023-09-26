// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build illumos

// Illumos system calls not present on Solaris.

package syscall

func Accept4(fd int, flags int) (int, Sockaddr, error)

func Flock(fd int, how int) error
