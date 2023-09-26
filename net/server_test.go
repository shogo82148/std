// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// Do not test empty datagrams by default.
// It causes unexplained timeouts on some systems,
// including Snow Leopard.  I think that the kernel
// doesn't quite expect them.
