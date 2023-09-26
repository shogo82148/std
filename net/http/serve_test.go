// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// End-to-end serving tests

package http_test

import (
	. "net/http"
)

// trackLastConnListener tracks the last net.Conn that was accepted.

// repeatReader reads content count times, then EOFs.

// A Response that's just no bigger than 2KB, the buffer-before-chunking threshold.
