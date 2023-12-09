// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !windows && !plan9

package syslog

import (
	"github.com/shogo82148/std/log"
	"github.com/shogo82148/std/sync"
)

// Priorityは、syslogの施設と重大度の組み合わせです。
// 例えば、LOG_ALERT | LOG_FTPは、FTP施設からのアラート重大度メッセージを送信します。
// デフォルトの重大度はLOG_EMERGで、デフォルトの施設はLOG_KERNです。
type Priority int

const (

	// /usr/include/sys/syslog.hから取得。
	// これらはLinux、BSD、OS Xで同じです。
	LOG_EMERG Priority = iota
	LOG_ALERT
	LOG_CRIT
	LOG_ERR
	LOG_WARNING
	LOG_NOTICE
	LOG_INFO
	LOG_DEBUG
)

const (

	// /usr/include/sys/syslog.hから取得。
	// これらはLinux、BSD、OS XでLOG_FTPまで同じです。
	LOG_KERN Priority = iota << 3
	LOG_USER
	LOG_MAIL
	LOG_DAEMON
	LOG_AUTH
	LOG_SYSLOG
	LOG_LPR
	LOG_NEWS
	LOG_UUCP
	LOG_CRON
	LOG_AUTHPRIV
	LOG_FTP
	_
	_
	_
	_
	LOG_LOCAL0
	LOG_LOCAL1
	LOG_LOCAL2
	LOG_LOCAL3
	LOG_LOCAL4
	LOG_LOCAL5
	LOG_LOCAL6
	LOG_LOCAL7
)

// Writerはsyslogサーバへの接続です。
type Writer struct {
	priority Priority
	tag      string
	hostname string
	network  string
	raddr    string

	mu   sync.Mutex
	conn serverConn
}

// Newはシステムログデーモンへの新しい接続を確立します。戻り値のライターへの
// 各書き込みは、指定された優先度（syslog施設と重大度の組み合わせ）と
// プレフィックスタグを持つログメッセージを送信します。タグが空の場合、os.Args[0]が使用されます。
func New(priority Priority, tag string) (*Writer, error)

// Dialは、指定されたネットワーク上のアドレスraddrに接続することで
// ログデーモンへの接続を確立します。戻り値のライターへの
// 各書き込みは、施設と重大度（priorityから）およびタグを持つログメッセージを送信します。
// タグが空の場合、os.Args[0]が使用されます。
// ネットワークが空の場合、Dialはローカルのsyslogサーバーに接続します。
// それ以外の場合は、ネットワークとraddrの有効な値については、net.Dialのドキュメンテーションを参照してください。
func Dial(network, raddr string, priority Priority, tag string) (*Writer, error)

// Writeはログメッセージをsyslogデーモンに送信します。
func (w *Writer) Write(b []byte) (int, error)

// Closeはsyslogデーモンへの接続を閉じます。
func (w *Writer) Close() error

// Emergは、severity LOG_EMERGのメッセージをログに記録します。Newに渡されたseverityは無視されます。
func (w *Writer) Emerg(m string) error

// Alertは、severity LOG_ALERTのメッセージをログに記録します。Newに渡されたseverityは無視されます。
func (w *Writer) Alert(m string) error

// Critは、severity LOG_CRITのメッセージをログに記録します。Newに渡されたseverityは無視されます。
func (w *Writer) Crit(m string) error

// Errは、severity LOG_ERRのメッセージをログに記録します。Newに渡されたseverityは無視されます。
func (w *Writer) Err(m string) error

// Warningは、severity LOG_WARNINGのメッセージをログに記録します。Newに渡されたseverityは無視されます。
func (w *Writer) Warning(m string) error

// Noticeは、severity LOG_NOTICEのメッセージをログに記録します。Newに渡されたseverityは無視されます。
func (w *Writer) Notice(m string) error

// Infoは、severity LOG_INFOのメッセージをログに記録します。Newに渡されたseverityは無視されます。
func (w *Writer) Info(m string) error

// Debugは、severity LOG_DEBUGのメッセージをログに記録します。Newに渡されたseverityは無視されます。
func (w *Writer) Debug(m string) error

// NewLoggerは、指定された優先度（syslog施設とseverityの組み合わせ）でシステムログサービスに書き込まれるlog.Loggerを作成します。
// logFlag引数は、Loggerを作成するためにlog.Newに渡されるフラグセットです。
func NewLogger(p Priority, logFlag int) (*log.Logger, error)
