// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix

package syscall

// origRlimitNofile, if not {0, 0}, is the original soft RLIMIT_NOFILE.
// When we can assume that we are bootstrapping with Go 1.19,
// this can be atomic.Pointer[Rlimit].

func Setrlimit(resource int, rlim *Rlimit) error
