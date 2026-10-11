// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package syscall

// RouteRIBは、ネットワーク機能の情報、状態、パラメータで構成される
// ルーティング情報ベース（RIBとも呼ばれる）を返します。
//
// Deprecated: 代わりにgolang.org/x/net/routeを使用してください。
func RouteRIB(facility, param int) ([]byte, error)

// RoutingMessageは、ルーティングメッセージを表します。
//
// Deprecated: 代わりにgolang.org/x/net/routeを使用してください。
type RoutingMessage interface {
	sockaddr() ([]Sockaddr, error)
}

// RouteMessageは、ルーティングエントリを含むルーティングメッセージを表します。
//
// Deprecated: 代わりにgolang.org/x/net/routeを使用してください。
type RouteMessage struct {
	Header RtMsghdr
	Data   []byte
}

// InterfaceMessageは、ネットワークインターフェースのエントリを含む
// ルーティングメッセージを表します。
//
// Deprecated: 代わりにgolang.org/x/net/routeを使用してください。
type InterfaceMessage struct {
	Header IfMsghdr
	Data   []byte
}

// InterfaceAddrMessageは、ネットワークインターフェースのアドレスエントリを含む
// ルーティングメッセージを表します。
//
// Deprecated: 代わりにgolang.org/x/net/routeを使用してください。
type InterfaceAddrMessage struct {
	Header IfaMsghdr
	Data   []byte
}

// ParseRoutingMessageは、bをルーティングメッセージとしてパースし、
// [RoutingMessage]インターフェースを含むスライスを返します。
//
// Deprecated: 代わりにgolang.org/x/net/routeを使用してください。
func ParseRoutingMessage(b []byte) (msgs []RoutingMessage, err error)

// ParseRoutingSockaddrは、msgのペイロードを生のsockaddrとしてパースし、
// [Sockaddr]インターフェースを含むスライスを返します。
//
// Deprecated: 代わりにgolang.org/x/net/routeを使用してください。
func ParseRoutingSockaddr(msg RoutingMessage) ([]Sockaddr, error)
