// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// If you change the build tags here, see
// internal/syscall/unix/fcntl_linux_32bit.go.

//go:build (linux && 386) || (linux && arm) || (linux && mips) || (linux && mipsle)

package syscall
