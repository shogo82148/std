// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syscall

// InterfaceMulticastAddrMessageは、ネットワークインターフェースのアドレスエントリーを
// 含むルーティングメッセージを表します。
//
// Deprecated: 代わりにgolang.org/x/net/routeを使用してください。
type InterfaceMulticastAddrMessage struct {
	Header IfmaMsghdr2
	Data   []byte
}
