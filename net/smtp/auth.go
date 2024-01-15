// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package smtp

// AuthはSMTP認証メカニズムによって実装されます。
type Auth interface {
	Start(server *ServerInfo) (proto string, toServer []byte, err error)

	Next(fromServer []byte, more bool) (toServer []byte, err error)
}

// ServerInfoはSMTPサーバーの情報を記録します。
type ServerInfo struct {
	Name string
	TLS  bool
	Auth []string
}

<<<<<<< HEAD
// PlainAuthは、RFC 4616で定義されているPLAIN認証メカニズムを実装するAuthを返します。返されたAuthは、指定したユーザー名とパスワードを使用してホストに認証し、アイデンティティとして動作します。通常、アイデンティティは空の文字列であるべきです。これにより、ユーザー名として機能します。
// PlainAuthは、接続がTLSを使用しているか、またはlocalhostに接続されている場合にのみ認証情報を送信します。それ以外の場合は、認証に失敗し、認証情報は送信されません。
func PlainAuth(identity, username, password, host string) Auth

// CRAMMD5AuthはRFC 2195で定義されたCRAM-MD5認証メカニズムを実装するAuthを返します。
// 返されたAuthは、ユーザー名と秘密情報を使用して、チャレンジ-レスポンスメカニズムを使ってサーバーに認証します。
=======
// PlainAuth returns an [Auth] that implements the PLAIN authentication
// mechanism as defined in RFC 4616. The returned Auth uses the given
// username and password to authenticate to host and act as identity.
// Usually identity should be the empty string, to act as username.
//
// PlainAuth will only send the credentials if the connection is using TLS
// or is connected to localhost. Otherwise authentication will fail with an
// error, without sending the credentials.
func PlainAuth(identity, username, password, host string) Auth

// CRAMMD5Auth returns an [Auth] that implements the CRAM-MD5 authentication
// mechanism as defined in RFC 2195.
// The returned Auth uses the given username and secret to authenticate
// to the server using the challenge-response mechanism.
>>>>>>> upstream/master
func CRAMMD5Auth(username, secret string) Auth
