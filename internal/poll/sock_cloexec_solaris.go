// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements accept for platforms that provide a fast path for
// setting SetNonblock and CloseOnExec, but don't necessarily have accept4.
// The accept4(3c) function was added to Oracle Solaris in the Solaris 11.4.0
// release. Thus, on releases prior to 11.4, we fall back to the combination
// of accept(3c) and fcntl(2).

package poll
