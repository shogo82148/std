// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

/*
Input to cgo -godefs
GOARCH=ppc64 go tool cgo -godefs defs_aix.go > defs_aix_ppc64_tmp.go

This is only a helper to create defs_aix_ppc64.go
Go runtime functions require the "linux" name of fields (ss_sp, si_addr, etc)
However, AIX structures don't provide such names and must be modified.

TODO(aix): create a script to automatise defs_aix creation.

Modifications made:
 - sigset replaced by a [4]uint64 array
 - add sigset_all variable
 - siginfo.si_addr uintptr instead of *byte
 - add (*timeval) set_usec
 - stackt.ss_sp uintptr instead of *byte
 - stackt.ss_size uintptr instead of uint64
 - sigcontext.sc_jmpbuf context64 instead of jumbuf
 - ucontext.__extctx is a uintptr because we don't need extctx struct
 - ucontext.uc_mcontext: replace jumbuf structure by context64 structure
 - sigaction.sa_handler represents union field as both are uintptr
 - tstate.* replace *byte by uintptr


*/

package runtime
