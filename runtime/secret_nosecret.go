// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

<<<<<<<< HEAD:runtime/defs_windows.go
// Windows architecture-independent definitions.
========
//go:build !(amd64 || arm64) || !linux
>>>>>>>> upstream/release-branch.go1.26:runtime/secret_nosecret.go

package runtime
