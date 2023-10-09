// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/crypto/x509"
)

// SessionStateは再開可能なセッションです。
type SessionState struct {

	// Extraはcrypto/tlsによって無視されますが、[SessionState.Bytes]によってエンコードされ、[ParseSessionState]によって解析されます。
	//
	// これにより、[Config.UnwrapSession]/[Config.WrapSession]や[ClientSessionCache]の実装が、このセッションとともに追加のデータを格納および取得できるようになります。
	//
	// プロトコルスタックの異なるレイヤーがこのフィールドを共有するために、アプリケーションは追加するだけでなく、置き換えることなく、順不同でも認識できるエントリを使用する必要があります（例：idとバージョンのプレフィックスを付けるなど）。
	Extra [][]byte

	// EarlyDataは、QUIC接続で0-RTTに使用できるかを示します。
	// サポートされていても、アプリケーションが0-RTTを提供しないことを拒否する場合は、これをfalseに設定することができます。
	EarlyData bool

	version     uint16
	isClient    bool
	cipherSuite uint16

	// createdAtはサーバー上でのシークレットの生成時間（TLS 1.0-1.2の場合、現在のセッションよりも前かもしれません）およびクライアントでのチケット受信時間です。
	createdAt         uint64
	secret            []byte
	extMasterSecret   bool
	peerCertificates  []*x509.Certificate
	activeCertHandles []*activeCert
	ocspResponse      []byte
	scts              [][]byte
	verifiedChains    [][]*x509.Certificate
	alpnProtocol      string

	// クライアント側のTLS 1.3専用フィールド。
	useBy  uint64
	ageAdd uint32
}

// Bytesはセッションをエンコードし、[ParseSessionState]によって解析できるようにします。エンコードには、将来のセッションのセキュリティに重要な秘密値が含まれている可能性があります。
//
// 具体的なエンコーディングは不透明と見なされ、Goのバージョン間で非互換性がある可能性があります。
func (s *SessionState) Bytes() ([]byte, error)

// ParseSessionStateは[SessionState.Bytes]でエンコードされた[SessionState]を解析します。
func ParseSessionState(data []byte) (*SessionState, error)

// EncryptTicketは、Configで設定された（またはデフォルトの）セッションチケットキーを使用してチケットを暗号化します。これは[Config.WrapSession]の実装として使用できます。
func (c *Config) EncryptTicket(cs ConnectionState, ss *SessionState) ([]byte, error)

// DecryptTicketは[Config.EncryptTicket]によって暗号化されたチケットを復号化します。[Config.UnwrapSession]の実装として使用することができます。
//
// もしチケットを復号化または解析できない場合、DecryptTicketは(nil, nil)を返します。
func (c *Config) DecryptTicket(identity []byte, cs ConnectionState) (*SessionState, error)

// ClientSessionState は、クライアントが前のTLSセッションを再開するために必要な状態を含んでいます。
type ClientSessionState struct {
	ticket  []byte
	session *SessionState
}

// ResumptionStateは、サーバーが送信したセッションチケット（セッションの識別子としても知られている）と、このセッションを再開するために必要な状態を返します。
//
// [ClientSessionCache.Put]によって呼び出され、セッションをシリアル化（[SessionState.Bytes]を使用）して保存することができます。
func (cs *ClientSessionState) ResumptionState() (ticket []byte, state *SessionState, err error)

// NewResumptionStateは、前のセッションを再開するために[ClientSessionCache.Get]によって返されるstate値を返します。
//
// stateは[ParseSessionState]によって返される必要があり、ticketとsessionの状態は[ClientSessionState.ResumptionState]によって返される必要があります。
func NewResumptionState(ticket []byte, state *SessionState) (*ClientSessionState, error)
