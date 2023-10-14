// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements accept for platforms that provide a fast path for
// setting SetNonblock and CloseOnExec, but don't necessarily have accept4.
// This is the code we used for accept in Go 1.17 and earlier.
// On Linux the accept4 system call was introduced in 2.6.28 kernel,
// and our minimum requirement is 2.6.32, so we simplified the function.
// Unfortunately, on ARM accept4 wasn't added until 2.6.36, so for ARM
// only we continue using the older code.

//go:build linux && arm

package poll
