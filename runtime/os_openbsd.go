// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// From OpenBSD's <sys/sysctl.h>

// sigPerThreadSyscall is only used on linux, so we assign a bogus signal
// number.
