// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || darwin || dragonfly || freebsd || (js && wasm) || linux || nacl || netbsd || openbsd || solaris
// +build aix darwin dragonfly freebsd js,wasm linux nacl netbsd openbsd solaris

// Unix environment variables.

package syscall

func Unsetenv(key string) error

func Getenv(key string) (value string, found bool)

func Setenv(key, value string) error

func Clearenv()

func Environ() []string
