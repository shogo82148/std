// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// smtp パッケージは RFC 5321 で定義されている Simple Mail Transfer Protocol を実装しています。
// さらに、以下の拡張も実装しています:
//
//	8BITMIME  RFC 1652
//	AUTH      RFC 2554
//	STARTTLS  RFC 3207
//
// クライアント側で追加の拡張も扱うことができます。
//
// smtp パッケージは凍結されており、新しい機能の追加は受け付けていません。
// いくつかの外部パッケージがより多機能を提供しています。以下を参照してください:
//
//	https://godoc.org/?q=smtp
package smtp

import (
	"github.com/shogo82148/std/crypto/tls"
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net"
	"github.com/shogo82148/std/net/textproto"
)

// ClientはSMTPサーバーへのクライアント接続を表します。
type Client struct {

	// TextはClientによって使用されるtextproto.Connです。拡張機能を追加できるように、公開されています。
	Text *textproto.Conn

	// 後でTLS接続を作成するために接続への参照を保持する
	conn net.Conn
	// クライアントがTLSを使用しているかどうか
	tls        bool
	serverName string
	// サポートされる拡張子のマップ
	ext map[string]string
	// サポートされている認証メカニズム
	auth       []string
	localName  string
	didHello   bool
	helloError error
}

<<<<<<< HEAD
// Dialはaddrに指定されたポート付きのSMTPサーバーに接続された新しいClientを返します。
// addrは"mail.example.com:smtp"のような形式である必要があります。
func Dial(addr string) (*Client, error)

// NewClient は既存の接続とホストを使用して新しいクライアントを返します。認証時に使用するサーバー名です。
=======
// Dial returns a new [Client] connected to an SMTP server at addr.
// The addr must include a port, as in "mail.example.com:smtp".
func Dial(addr string) (*Client, error)

// NewClient returns a new [Client] using an existing connection and host as a
// server name to be used when authenticating.
>>>>>>> upstream/master
func NewClient(conn net.Conn, host string) (*Client, error)

// Closeは接続をクローズします。
func (c *Client) Close() error

// Helloメソッドは、指定されたホスト名としてサーバーにHELOまたはEHLOを送信します。
// クライアントが使用するホスト名を制御する必要がある場合にのみ、このメソッドを呼び出す必要があります。
// それ以外の場合は、クライアントは自動的に「localhost」として自己紹介します。
// Helloメソッドを呼び出す場合は、他のメソッドのいずれかを呼び出す前に呼び出す必要があります。
func (c *Client) Hello(localName string) error

// StartTLSはSTARTTLSコマンドを送信し、以降のすべての通信を暗号化します。
// この機能をサポートするのは、STARTTLS拡張機能を広告するサーバーのみです。
func (c *Client) StartTLS(config *tls.Config) error

<<<<<<< HEAD
// TLSConnectionState はクライアントのTLS接続状態を返します。
// StartTLS が成功しなかった場合、返り値はゼロ値になります。
=======
// TLSConnectionState returns the client's TLS connection state.
// The return values are their zero values if [Client.StartTLS] did
// not succeed.
>>>>>>> upstream/master
func (c *Client) TLSConnectionState() (state tls.ConnectionState, ok bool)

// Verifyはサーバー上でメールアドレスの妥当性をチェックします。
// Verifyがnilを返す場合、アドレスは有効です。非nilの返り値は
// 必ずしも無効なアドレスを示すわけではありません。セキュリティ上の理由から、
// 多くのサーバーはアドレスの検証を行わない場合があります。
func (c *Client) Verify(addr string) error

// Authは提供された認証メカニズムを使用してクライアントを認証します。
// 認証に失敗した場合、接続は閉じられます。
// この機能は、AUTH拡張機能をサポートしているサーバーのみが広告しています。
func (c *Client) Auth(a Auth) error

<<<<<<< HEAD
// Mailは提供されたメールアドレスを使用してサーバーにMAILコマンドを発行します。
// サーバーが8BITMIME拡張をサポートしている場合、MailはBODY=8BITMIMEパラメータを追加します。
// サーバーがSMTPUTF8拡張をサポートしている場合、MailはSMTPUTF8パラメータを追加します。
// これにより、メールのトランザクションが開始され、その後に1つ以上のRcpt呼び出しが続きます。
func (c *Client) Mail(from string) error

// Rcptは提供されたメールアドレスを使用してサーバーにRCPTコマンドを発行します。
// Rcptの呼び出しは、Mailの呼び出しの前に行われなければならず、Dataの呼び出しまたは別のRcptの呼び出しの後に続く場合があります。
func (c *Client) Rcpt(to string) error

// DataはサーバーにDATAコマンドを送信し、メールのヘッダーと本文を書き込むために使用できるライターを返します。呼び出し元は、cの他のメソッドを呼び出す前にライターを閉じる必要があります。Dataの呼び出しは、一つ以上のRcptの呼び出しに先行する必要があります。
=======
// Mail issues a MAIL command to the server using the provided email address.
// If the server supports the 8BITMIME extension, Mail adds the BODY=8BITMIME
// parameter. If the server supports the SMTPUTF8 extension, Mail adds the
// SMTPUTF8 parameter.
// This initiates a mail transaction and is followed by one or more [Client.Rcpt] calls.
func (c *Client) Mail(from string) error

// Rcpt issues a RCPT command to the server using the provided email address.
// A call to Rcpt must be preceded by a call to [Client.Mail] and may be followed by
// a [Client.Data] call or another Rcpt call.
func (c *Client) Rcpt(to string) error

// Data issues a DATA command to the server and returns a writer that
// can be used to write the mail headers and body. The caller should
// close the writer before calling any more methods on c. A call to
// Data must be preceded by one or more calls to [Client.Rcpt].
>>>>>>> upstream/master
func (c *Client) Data() (io.WriteCloser, error)

// SendMailはaddrで指定されたサーバに接続し、可能な場合はTLSに切り替え、必要に応じてオプションのメカニズムaで認証し、fromからのアドレス、toへのアドレス、メッセージmsgを送信します。
// addrにはポートを含める必要があります。例："mail.example.com:smtp"
//
// toパラメータのアドレスは、SMTPのRCPTアドレスです。
//
// msgパラメータは、ヘッダ、空行、メッセージ本文の順になったRFC 822スタイルの電子メールである必要があります。msgの各行はCRLFで終端する必要があります。msgのヘッダには通常、"From"、"To"、"Subject"、"Cc"などのフィールドが含まれるべきです。"Bcc"メッセージを送信するには、toパラメータにメールアドレスを含め、msgのヘッダには含めません。
//
// SendMail関数とnet/smtpパッケージは低レベルのメカニズムであり、DKIM署名、MIME添付ファイル（mime/multipartパッケージを参照）、その他のメール機能をサポートしていません。高レベルのパッケージは標準ライブラリの外部に存在します。
func SendMail(addr string, a Auth, from string, to []string, msg []byte) error

// Extensionはサーバーが対応している拡張機能かどうかを報告します。
// 拡張機能名は大文字小文字を区別しません。もし拡張機能が対応されている場合、
// Extensionは拡張機能に対してサーバーが指定する任意のパラメータを含む文字列も返します。
func (c *Client) Extension(ext string) (bool, string)

// Resetは、現在のメールトランザクションを中止し、サーバーにRSETコマンドを送信します。
func (c *Client) Reset() error

// NoopはサーバーにNOOPコマンドを送信します。これによってサーバーとの接続が正常であることを確認します。
func (c *Client) Noop() error

// QuitはQUITコマンドを送信し、サーバーへの接続を閉じます。
func (c *Client) Quit() error
