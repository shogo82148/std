// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Netlink sockets and messages

package syscall

// NetlinkRouteRequestは、カーネルからルーティングとリンクの状態を受信するための
// リクエストメッセージを表します。
type NetlinkRouteRequest struct {
	Header NlMsghdr
	Data   RtGenmsg
}

// NetlinkRIBは、ネットワーク機能の情報、状態、パラメーターで構成される
// ルーティング情報ベース（RIBとも呼ばれる）を返します。
func NetlinkRIB(proto, family int) ([]byte, error)

// NetlinkMessageは、netlinkメッセージを表します。
type NetlinkMessage struct {
	Header NlMsghdr
	Data   []byte
}

// ParseNetlinkMessageは、bをnetlinkメッセージの配列としてパースし、
// NetlinkMessage構造体を含むスライスを返します。
func ParseNetlinkMessage(b []byte) ([]NetlinkMessage, error)

// NetlinkRouteAttrは、netlinkのルート属性を表します。
type NetlinkRouteAttr struct {
	Attr  RtAttr
	Value []byte
}

// ParseNetlinkRouteAttrは、mのペイロードをnetlinkのルート属性の配列として
// パースし、NetlinkRouteAttr構造体を含むスライスを返します。
func ParseNetlinkRouteAttr(m *NetlinkMessage) ([]NetlinkRouteAttr, error)
