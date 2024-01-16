// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/net/netip"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// TCPAddrはTCPエンドポイントのアドレスを表します。
type TCPAddr struct {
	IP   IP
	Port int
	Zone string
}

// AddrPortは [TCPAddr] aを [netip.AddrPort] として返します。
//
// もしa.Portがuint16に収まらない場合、静かに切り捨てられます。
//
// もしaがnilの場合、ゼロ値が返されます。
func (a *TCPAddr) AddrPort() netip.AddrPort

// Networkはアドレスのネットワーク名「tcp」を返します。
func (a *TCPAddr) Network() string

func (a *TCPAddr) String() string

// ResolveTCPAddrはTCPエンドポイントのアドレスを返します。
//
// ネットワークはTCPのネットワーク名である必要があります。
//
// アドレスパラメータのホストがリテラルIPアドレスでない場合や、
// ポートがリテラルのポート番号でない場合、ResolveTCPAddrは
// TCPエンドポイントのアドレスに解決します。
// そうでなければ、アドレスをリテラルのIPアドレスとポート番号のペアとして解析します。
// アドレスパラメータはホスト名を使用することもできますが、
// ホスト名のIPアドレスの一つを最大で返すため、推奨されていません。
//
// ネットワークとアドレスパラメータの詳細については、
// [Dial] 関数の説明を参照してください。
func ResolveTCPAddr(network, address string) (*TCPAddr, error)

// TCPAddrFromAddrPortはaddrを [TCPAddr] として返します。もしaddrがIsValid()がfalseである場合、
// 返されるTCPAddrにはnilのIPフィールドが含まれ、アドレスファミリーに依存しない未指定のアドレスを示します。
func TCPAddrFromAddrPort(addr netip.AddrPort) *TCPAddr

// TCPConnはTCPネットワーク接続の [Conn] インターフェースの実装です。
type TCPConn struct {
	conn
}

// SyscallConnは生のネットワーク接続を返します。
// これは [syscall.Conn] インターフェースを実装しています。
func (c *TCPConn) SyscallConn() (syscall.RawConn, error)

// ReadFrom は [io.ReaderFrom] の ReadFrom メソッドを実装します。
func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)

// WriteToは、io.WriterToのWriteToメソッドを実装します。
func (c *TCPConn) WriteTo(w io.Writer) (int64, error)

// CloseReadはTCP接続の読み込み側をシャットダウンします。
// ほとんどの呼び出し元は、単にCloseを使用するだけで十分です。
func (c *TCPConn) CloseRead() error

// CloseWrite は TCP 接続の書き込み側をシャットダウンします。
// ほとんどの呼び出し元は Close を使用すべきです。
func (c *TCPConn) CloseWrite() error

// SetLingerは、まだ送信または確認待ちのデータがある接続に対してCloseの振る舞いを設定します。
// sec < 0（デフォルト）の場合、オペレーティングシステムはバックグラウンドでデータの送信を完了します。
// sec == 0の場合、オペレーティングシステムは未送信または確認待ちのデータを破棄します。
// sec > 0の場合、データはsec < 0と同様にバックグラウンドで送信されます。
// Linuxを含む一部のオペレーティングシステムでは、これによりCloseが全てのデータの送信または破棄が完了するまでブロックする場合があります。
// sec秒経過後、未送信のデータは破棄される可能性があります。
func (c *TCPConn) SetLinger(sec int) error

// SetKeepAliveは、オペレーティングシステムが接続に対して
// keep-aliveメッセージを送信するかどうかを設定します。
func (c *TCPConn) SetKeepAlive(keepalive bool) error

// SetKeepAlivePeriodは、Keep-Alive間の期間を設定します。
func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error

// SetNoDelayは、パケットの送信を遅延させるかどうかを制御します。これにより、より少ないパケットで送信することが期待されます（Nagleのアルゴリズム）。デフォルト値はtrue（遅延なし）であり、Writeの後で可能な限りすぐにデータが送信されます。
func (c *TCPConn) SetNoDelay(noDelay bool) error

// MultipathTCPは、現在の接続がMPTCPを使用しているかどうかを報告します。
//
// ホスト、他のピア、またはその間にあるデバイスによってMultipath TCPがサポートされていない場合、
// 意図的に/意図せずにフィルタリングされた場合、TCPへのフォールバックが行われます。
// このメソッドは、MPTCPが使用されているかどうかを確認するために最善を尽くします。
//
// Linuxでは、カーネルのバージョンがv5.16以上の場合、さらに条件が検証され、結果が改善されます。
func (c *TCPConn) MultipathTCP() (bool, error)

// DialTCPはTCPネットワークのための [Dial] のように振る舞います。
//
// ネットワークはTCPネットワーク名でなければなりません。詳細についてはfunc Dialを参照してください。
//
// laddrがnilの場合、自動的にローカルアドレスが選択されます。
// raddrのIPフィールドがnilまたは未指定のIPアドレスの場合、ローカルシステムが使用されます。
func DialTCP(network string, laddr, raddr *TCPAddr) (*TCPConn, error)

// TCPListenerはTCPネットワークリスナーです。クライアントは通常、TCPを仮定する代わりに [Listener] 型の変数を使用するべきです。
type TCPListener struct {
	fd *netFD
	lc ListenConfig
}

// SyscallConn は生のネットワーク接続を返します。
// これは [syscall.Conn] インターフェースを実装しています。
//
// 返された RawConn は Control の呼び出しのみをサポートします。
// Read と Write はエラーを返します。
func (l *TCPListener) SyscallConn() (syscall.RawConn, error)

// AcceptTCPは次の着信呼び出しを受け入れ、新しい接続を返します。
func (l *TCPListener) AcceptTCP() (*TCPConn, error)

// Accept implements the Accept method in the [Listener] interface; it
// waits for the next call and returns a generic [Conn].
func (l *TCPListener) Accept() (Conn, error)

// Close は TCP アドレスのリスニングを停止します。
// 既に受け入れられた接続は閉じられません。
func (l *TCPListener) Close() error

// Addrはリスナーのネットワークアドレス、[*TCPAddr] を返します。
// 返されるAddrはAddrのすべての呼び出しで共有されるため、
// 変更しないでください。
func (l *TCPListener) Addr() Addr

// SetDeadlineはリスナーに関連付けられた締め切りを設定します。
// ゼロの時刻値は締め切りを無効にします。
func (l *TCPListener) SetDeadline(t time.Time) error

// File は元の [os.File] のコピーを返します。
// 終了した後、f を閉じる責任は呼び出し元にあります。
// l を閉じても f には影響を与えませんし、f を閉じても l には影響を与えません。
//
// 返された os.File のファイルディスクリプタは、接続のものとは異なります。
// この複製を使用して元のもののプロパティを変更しようとすると、
// 望ましい効果が現れるかどうかは不明です。
func (l *TCPListener) File() (f *os.File, err error)

// ListenTCPはTCPネットワーク用の [Listen] のように機能します。
//
// ネットワークはTCPネットワーク名でなければなりません。詳細はfunc Dialを参照してください。
//
// laddrのIPフィールドがnilまたは未指定のIPアドレスの場合、ListenTCPはローカルシステムの利用可能なユニキャストおよびエニーキャストIPアドレスすべてでリスンします。
// laddrのPortフィールドが0の場合、ポート番号は自動的に選択されます。
func ListenTCP(network string, laddr *TCPAddr) (*TCPListener, error)
