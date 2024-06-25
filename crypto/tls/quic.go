// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tls

import (
	"github.com/shogo82148/std/context"
)

// QUICEncryptionLevelは、ハンドシェイクメッセージを送信するために使用されるQUIC暗号化レベルを表します。
type QUICEncryptionLevel int

const (
	QUICEncryptionLevelInitial = QUICEncryptionLevel(iota)
	QUICEncryptionLevelEarly
	QUICEncryptionLevelHandshake
	QUICEncryptionLevelApplication
)

func (l QUICEncryptionLevel) String() string

// QUICConnは、RFC 9001で説明されているように、基礎となるトランスポートとしてQUICの実装を使用する接続を表します。
//
// QUICConnのメソッドは同時に使用することはできません。
type QUICConn struct {
	conn *Conn

	sessionTicketSent bool
}

// QUICConfigは [QUICConn] を設定します。
type QUICConfig struct {
	TLSConfig *Config

	// EnableStoreSessionEvent may be set to true to enable the
	// [QUICStoreSession] event for client connections.
	// When this event is enabled, sessions are not automatically
	// stored in the client session cache.
	// The application should use [QUICConn.StoreSession] to store sessions.
	EnableStoreSessionEvent bool
}

// QUICEventKindはQUIC接続上での操作の種類です。
type QUICEventKind int

const (
	// QUICNoEventは利用可能なイベントが存在しないことを示します。
	QUICNoEvent QUICEventKind = iota

	// QUICSetReadSecretとQUICSetWriteSecretは、特定の暗号化レベルの読み取りと書き込みの秘密情報を提供します。
	// QUICEvent.Level、QUICEvent.Data、QUICEvent.Suiteが設定されます。
	//
	// Initial暗号化レベルの秘密情報は、最初の宛先接続IDから派生され、QUICConnによって提供されません。
	QUICSetReadSecret
	QUICSetWriteSecret

	// QUICWriteDataはCRYPTOフレームでピアに送信するデータを提供します。
	// QUICEvent.Dataが設定されています。
	QUICWriteData

	// QUICTransportParametersは相手のQUICトランスポートパラメータを提供します。
	// QUICEvent.Dataが設定されています。
	QUICTransportParameters

	// QUICTransportParametersRequiredは、呼び出し元がピアに送信するためのQUICトランスポートパラメータを提供する必要があることを示します。呼び出し元は、QUICConn.SetTransportParametersを使用してトランスポートパラメータを設定し、QUICConn.NextEventを再度呼び出す必要があります。
	// QUICConn.Startを呼び出す前にトランスポートパラメータが設定されている場合、接続は決してQUICTransportParametersRequiredイベントを生成しません。
	QUICTransportParametersRequired

	// QUICRejectedEarlyDataは、サーバーが私たちが提供したものであっても、0-RTTデータを拒否したことを示しています。これは、QUICEncryptionLevelApplicationのキーが返される前に返されます。
	// このイベントはクライアント接続でのみ発生します。
	QUICRejectedEarlyData

	// QUICHandshakeDone は、TLS ハンドシェイクが完了したことを示します。
	QUICHandshakeDone

	// QUICResumeSession indicates that a client is attempting to resume a previous session.
	// [QUICEvent.SessionState] is set.
	//
	// For client connections, this event occurs when the session ticket is selected.
	// For server connections, this event occurs when receiving the client's session ticket.
	//
	// The application may set [QUICEvent.SessionState.EarlyData] to false before the
	// next call to [QUICConn.NextEvent] to decline 0-RTT even if the session supports it.
	QUICResumeSession

	// QUICStoreSession indicates that the server has provided state permitting
	// the client to resume the session.
	// [QUICEvent.SessionState] is set.
	// The application should use [QUICConn.Store] session to store the [SessionState].
	// The application may modify the [SessionState] before storing it.
	// This event only occurs on client connections.
	QUICStoreSession
)

// QUICEventはQUIC接続で発生するイベントです。
//
// イベントの種類はKindフィールドで指定されます。
// 他のフィールドの内容は、種別によって異なります。
type QUICEvent struct {
	Kind QUICEventKind

	// QUICSetReadSecret、QUICSetWriteSecret、およびQUICWriteDataに対する設定。
	Level QUICEncryptionLevel

	// QUICTransportParameters、QUICSetReadSecret、QUICSetWriteSecret、およびQUICWriteDataに設定します。
	// この内容はcrypto/tlsによって所有され、次のNextEvent呼び出しまで有効です。
	Data []byte

	// QUICSetReadSecretおよびQUICSetWriteSecretに設定します。
	Suite uint16

	// Set for QUICResumeSession and QUICStoreSession.
	SessionState *SessionState
}

// QUICClientは、QUICTransportを基礎とした新しいTLSクライアント側接続を返します。設定はnilであってはなりません。
//
// 設定のMinVersionは、少なくともTLS 1.3である必要があります。
func QUICClient(config *QUICConfig) *QUICConn

// QUICServerは、下層トランスポートとしてQUICTransportを使用した新しいTLSサーバーサイド接続を返します。設定はnilにできません。
//
// 設定のMinVersionは、少なくともTLS 1.3である必要があります。
func QUICServer(config *QUICConfig) *QUICConn

// Startはクライアントまたはサーバーのハンドシェイクプロトコルを開始します。
// 接続イベントを生成する場合があり、 [QUICConn.NextEvent] で読み取ることができます。
//
// Startは1度以上呼び出すことはできません。
func (q *QUICConn) Start(ctx context.Context) error

// NextEventは接続で発生する次のイベントを返します。
// イベントが利用できない場合は、Kindが [QUICNoEvent] のイベントを返します。
func (q *QUICConn) NextEvent() QUICEvent

// Closeは接続を閉じ、進行中のハンドシェイクを停止します。
func (q *QUICConn) Close() error

// HandleDataはピアから受信したハンドシェイクバイトを処理します。
// 接続イベントを生成することがあり、 [QUICConn.NextEvent] で読み取ることができます。
func (q *QUICConn) HandleData(level QUICEncryptionLevel, data []byte) error

type QUICSessionTicketOptions struct {
	// EarlyDataは0-RTTで使用できるかどうかを指定します。
	EarlyData bool
	Extra     [][]byte
}

// SendSessionTicketはクライアントにセッションチケットを送信します。
// これにより、接続イベントが生成され、 [QUICConn.NextEvent] で読み取ることができます。
// 現在、一度しか呼び出すことはできません。
func (q *QUICConn) SendSessionTicket(opts QUICSessionTicketOptions) error

// StoreSessionは、QUICStoreSessionイベントで以前受け取ったセッションをClientSessionCacheに保存します。
// アプリケーションは、セッションを保存する前に追加のイベントを処理したり、SessionStateを変更したりすることができます。
func (q *QUICConn) StoreSession(session *SessionState) error

// ConnectionStateは接続に関する基本的なTLSの詳細を返します。
func (q *QUICConn) ConnectionState() ConnectionState

// SetTransportParametersはピアに送信するためのトランスポートパラメータを設定します。
//
// サーバ接続では、クライアントのトランスポートパラメータを受信した後にトランスポートパラメータを設定することができます。 [QUICTransportParametersRequired] を参照してください。
func (q *QUICConn) SetTransportParameters(params []byte)
