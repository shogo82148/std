// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http_test

import (
	. "net/http"
)

// infiniteReader satisfies Read requests as if the contents of buf
// loop indefinitely.
