// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// asyncIO implements asynchronous cancelable I/O.
// An asyncIO represents a single asynchronous Read or Write
// operation. The result is returned on the result channel.
// The undergoing I/O system call can either complete or be
// interrupted by a note.

// result is the return value of a Read or Write operation.
