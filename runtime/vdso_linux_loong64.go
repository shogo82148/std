// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build linux && loong64

package runtime

// not currently described in manpages as of May 2022, but will eventually
// appear
// when that happens, see man 7 vdso : loongarch

// initialize to fall back to syscall
