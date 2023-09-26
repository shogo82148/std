// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http_test

import (
	. "net/http"
)

// delayedEOFReader never returns (n > 0, io.EOF), instead putting
// off the io.EOF until a subsequent Read call.

// infiniteReader satisfies Read requests as if the contents of buf
// loop indefinitely.
