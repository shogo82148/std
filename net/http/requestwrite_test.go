// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

// delegateReader is a reader that delegates to another reader,
// once it arrives on a channel.

// dumpConn is a net.Conn that writes to Writer and reads from Reader.
