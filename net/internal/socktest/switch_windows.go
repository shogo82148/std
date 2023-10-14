// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package socktest

import "github.com/shogo82148/std/syscall"

// Sockets maps a socket descriptor to the status of socket.
type Sockets map[syscall.Handle]Status
