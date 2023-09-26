// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// Interface for all IO operations.

// IO completion result parameters.

// anOp implements functionality common to all IO operations.

// bufOp is used by IO operations that read / write
// data from / to client buffer.

// resultSrv will retrieve all IO completion results from
// iocp and send them to the correspondent waiting client
// goroutine via channel supplied in the request.

// ioSrv executes net IO requests.

// Start helper goroutines.

// Network file descriptor.
