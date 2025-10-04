// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Not all systems have syscall.Mkfifo.
//go:build !aix && !plan9 && !solaris && !wasm && !windows

package wasi_test
