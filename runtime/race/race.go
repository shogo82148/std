// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race && ((linux && (amd64 || arm64 || loong64 || ppc64le || riscv64 || s390x)) || ((freebsd || netbsd || openbsd || windows) && amd64))

package race
