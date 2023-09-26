// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix || linux || darwin || dragonfly || freebsd || openbsd || netbsd || solaris
// +build aix linux darwin dragonfly freebsd openbsd netbsd solaris

package tar

// userMap and groupMap caches UID and GID lookups for performance reasons.
// The downside is that renaming uname or gname by the OS never takes effect.
