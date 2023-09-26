// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1 || windows

package net

// SetUnlinkOnClose sets whether the underlying socket file should be removed
// from the file system when the listener is closed.
//
// The default behavior is to unlink the socket file only when package net created it.
// That is, when the listener and the underlying socket file were created by a call to
// Listen or ListenUnix, then by default closing the listener will remove the socket file.
// but if the listener was created by a call to FileListener to use an already existing
// socket file, then by default closing the listener will not remove the socket file.
func (l *UnixListener) SetUnlinkOnClose(unlink bool)
