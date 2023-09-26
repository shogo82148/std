// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// The read and write file descriptors used by the sigNote functions.

//go:linkname executablePath os.executablePath

// sigPerThreadSyscall is only used on linux, so we assign a bogus signal
// number.
