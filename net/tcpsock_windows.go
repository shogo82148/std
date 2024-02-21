// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

// SetKeepAliveConfig configures keep-alive messages sent by the operating system.
func (c *TCPConn) SetKeepAliveConfig(config KeepAliveConfig) error
