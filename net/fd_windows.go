// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// Interface for all io operations.

// IO completion result parameters.

// anOp implements functionality common to all io operations.

// bufOp is used by io operations that read / write
// data from / to client buffer.

// resultSrv will retrieve all io completion results from
// iocp and send them to the correspondent waiting client
// goroutine via channel supplied in the request.

// ioSrv executes net io requests.

// Start helper goroutines.

// Network file descriptor.
