// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements sysSocket for platforms that provide a fast path for
// setting SetNonblock and CloseOnExec, but don't necessarily support it.
// Support for SOCK_* flags as part of the type parameter was added to Oracle
// Solaris in the 11.4 release. Thus, on releases prior to 11.4, we fall back
// to the combination of socket(3c) and fcntl(2).

package net
