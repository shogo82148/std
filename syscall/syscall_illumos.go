// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build illumos

package syscall

// F_DUP2FD_CLOEXEC has different values on Solaris and Illumos.
const F_DUP2FD_CLOEXEC = 0x24

func Flock(fd int, how int) error
