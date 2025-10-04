// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<<< HEAD:net/sendfile_unix_test.go
//go:build unix
========
//go:build freebsd || linux || solaris
>>>>>>>> upstream/release-branch.go1.25:os/readfrom_unix_test.go

package net
