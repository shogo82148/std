// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// syslogパッケージは、システムログサービスへのシンプルなインターフェースを提供します。
// UNIXドメインソケット、UDP、またはTCPを使用してsyslogデーモンにメッセージを送信することができます。
//
// Dialへの呼び出しは一度だけ必要です。書き込み失敗時、
// syslogクライアントはサーバーへの再接続を試み、再度書き込みを行います。
//
// syslogパッケージは凍結されており、新たな機能は受け入れていません。
// 一部の外部パッケージがより多くの機能を提供しています。参照：
//
//	https://godoc.org/?q=syslog
package syslog
