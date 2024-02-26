// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/syscall"
	"github.com/shogo82148/std/time"
)

// Dialerはアドレスに接続するためのオプションを含んでいます。
//
// 各フィールドのゼロ値は、そのオプションなしでダイヤルすることと同等です。
// Dialerのゼロ値でダイヤルすることは、単に [Dial] 関数を呼び出すのと同等です。
//
// Dialerのメソッドを同時に呼び出しても安全です。
type Dialer struct {

	// Timeoutはダイヤルが接続の完了を待つ最大時間です。Deadlineも設定されている場合、より早く失敗する可能性があります。
	// デフォルトはタイムアウトなしです。
	// TCPを使用して複数のIPアドレスを持つホスト名にダイヤルする場合、タイムアウトはそれらの間で分割される場合があります。
	// タイムアウトの有無にかかわらず、オペレーティングシステムは独自の早期タイムアウトを課す場合があります。たとえば、TCPのタイムアウトは通常約3分です。
	Timeout time.Duration

	// Deadlineは、ダイヤルが失敗する絶対的な時間です。
	// Timeoutが設定されている場合、それよりも早く失敗することがあります。
	// ゼロは期限がないことを意味し、またはオペレーティングシステムに依存することもあります。
	// (Timeoutオプションと同様に)
	Deadline time.Time

	// LocalAddr is the local address to use when dialing an
	// address. The address must be of a compatible type for the
	// network being dialed.
	// If nil, a local address is automatically chosen.
	LocalAddr Addr

	// DualStackは以前からRFC 6555 Fast Fallback、または「Happy Eyeballs」として知られる機能をサポートしており、IPv6が誤設定されていて正しく動作していない場合にはIPv4がすぐに試されます。
	// 廃止予定：Fast Fallbackはデフォルトで有効になっています。無効にするには、FallbackDelayを負の値に設定してください。
	DualStack bool

	// FallbackDelayは、RFC 6555 Fast Fallback接続を作成する前に待機する時間の長さを指定します。つまり、IPv6が成功するまで待機する時間であり、IPv6の設定が誤っていると仮定し、IPv4に切り替える前に待機する時間です。
	// ゼロの場合、デフォルトの遅延時間は300msです。
	// 負の値はFast Fallbackサポートを無効にします。
	FallbackDelay time.Duration

<<<<<<< HEAD
	// KeepAliveはアクティブなネットワーク接続の間隔を示します。
	// ゼロの場合、keep-aliveプローブはデフォルト値（現在は15秒）で送信されます。
	// プロトコルやオペレーティングシステムがサポートしている場合、ネットワークプロトコルやオペレーティングシステムはkeep-aliveを無視します。
	// ネガティブの場合、keep-aliveプローブは無効になります。
	KeepAlive time.Duration

	// Resolverはオプションで、代替のリゾルバを指定することができます。
=======
	// KeepAlive specifies the interval between keep-alive
	// probes for an active network connection.
	//
	// KeepAlive is ignored if KeepAliveConfig.Enable is true.
	//
	// If zero, keep-alive probes are sent with a default value
	// (currently 15 seconds), if supported by the protocol and operating
	// system. Network protocols or operating systems that do
	// not support keep-alive ignore this field.
	// If negative, keep-alive probes are disabled.
	KeepAlive time.Duration

	// KeepAliveConfig specifies the keep-alive probe configuration
	// for an active network connection, when supported by the
	// protocol and operating system.
	//
	// If KeepAliveConfig.Enable is true, keep-alive probes are enabled.
	// If KeepAliveConfig.Enable is false and KeepAlive is negative,
	// keep-alive probes are disabled.
	KeepAliveConfig KeepAliveConfig

	// Resolver optionally specifies an alternate resolver to use.
>>>>>>> upstream/master
	Resolver *Resolver

	// Cancel is an optional channel whose closure indicates that
	// the dial should be canceled. Not all types of dials support
	// cancellation.
	//
	// Deprecated: Use DialContext instead.
	Cancel <-chan struct{}

	// Controlがnilでない場合、ネットワーク接続の作成後に、実際にダイアルする前に呼び出されます。
	//
	// Control関数に渡されるネットワークとアドレスのパラメータは、必ずしもDialに渡されるものとは限りません。たとえば、Dialに「tcp」を渡すと、Control関数は「tcp4」または「tcp6」で呼び出されます。
	//
	// ControlContextがnilでない場合、Controlは無視されます。
	Control func(network, address string, c syscall.RawConn) error

	// ControlContextがnilでない場合、ネットワークの接続を作成する前に呼び出されます。
	//
	// ControlContext関数に渡されるネットワークおよびアドレスのパラメータは、必ずしもDialに渡されたものではありません。
	// 例えば、Dialに"tcp"を渡すと、ControlContext関数は "tcp4" または "tcp6" とともに呼び出されます。
	//
	// ControlContextがnilでない場合、Controlは無視されます。
	ControlContext func(ctx context.Context, network, address string, c syscall.RawConn) error

	// もしmptcpStatusがMPTCPを許可する値に設定されている場合、"tcp(4|6)"というネットワークを使用するDialの呼び出しは、オペレーティングシステムでサポートされていればMPTCPを使用します。
	mptcpStatus mptcpStatus
}

// MultipathTCPはMPTCPを使用するかどうかを報告します。
//
// このメソッドは、オペレーティングシステムがMPTCPをサポートしているかどうかをチェックしません。
func (d *Dialer) MultipathTCP() bool

// SetMultipathTCPは、オペレーティングシステムでサポートされている場合、[Dial] メソッドがMPTCPを使用するかどうかを指示します。
// このメソッドは、システムのデフォルトとGODEBUG=multipathtcp=...の設定を上書きします。
//
// ホストでMPTCPが利用できない場合やサーバーでサポートされていない場合、DialメソッドはTCPにフォールバックします。
func (d *Dialer) SetMultipathTCP(use bool)

// Dialは指定されたネットワークのアドレスに接続します。
//
// 知られているネットワークは "tcp", "tcp4" (IPv4のみ), "tcp6" (IPv6のみ),
// "udp", "udp4" (IPv4のみ), "udp6" (IPv6のみ), "ip", "ip4"
// (IPv4のみ), "ip6" (IPv6のみ), "unix", "unixgram" および
// "unixpacket" です。
//
// TCPとUDPのネットワークの場合、アドレスは "ホスト:ポート" の形式で指定します。
// ホストはリテラルのIPアドレスであるか、IPアドレスに解決できるホスト名である必要があります。
// ポートはリテラルのポート番号またはサービス名である必要があります。
// ホストがリテラルのIPv6アドレスの場合、"[2001:db8::1]:80" または "[fe80::1%zone]:80" のように角括弧で囲む必要があります。
// ゾーンは、RFC 4007で定義されているリテラルのIPv6アドレスのスコープを指定します。
// 関数 [JoinHostPort] と [SplitHostPort] は、この形式のホストとポートのペアを操作します。
// TCPを使用し、ホストが複数のIPアドレスに解決される場合、Dialは順番に各IPアドレスを試し、成功したものを使用します。
//
// 例:
//
// Dial("tcp", "golang.org:http")
// Dial("tcp", "192.0.2.1:http")
// Dial("tcp", "198.51.100.1:80")
// Dial("udp", "[2001:db8::1]:domain")
// Dial("udp", "[fe80::1%lo0]:53")
// Dial("tcp", ":80")
//
// IPネットワークの場合、ネットワークは "ip", "ip4" または "ip6" の後にコロンとリテラルのプロトコル番号またはプロトコル名が続き、
// アドレスは "ホスト" の形式となります。ホストはリテラルのIPアドレスまたはゾーン付きのリテラルのIPv6アドレスである必要があります。
// "0" や "255" などの広く知られていないプロトコル番号の場合、各オペレーティングシステムによって動作が異なることによります。
//
// 例:
//
// Dial("ip4:1", "192.0.2.1")
// Dial("ip6:ipv6-icmp", "2001:db8::1")
// Dial("ip6:58", "fe80::1%lo0")
//
// TCP、UDP、およびIPネットワークの場合、ホストが空白またはリテラルの未指定IPアドレスの場合、
// すなわち ":80", "0.0.0.0:80" または "[::]:80" などの場合、TCPおよびUDPでは、
// ""、"0.0.0.0" または "::" などの場合、IPでは、ローカルシステムが仮定されます。
//
// UNIXネットワークの場合、アドレスはファイルシステムのパスである必要があります。
func Dial(network, address string) (Conn, error)

// DialTimeoutは、タイムアウトを設定して [Dial] と同様の動作をします。
//
// 必要に応じて名前解決も含まれたタイムアウト処理が行われます。
// TCPを使用している場合、アドレスパラメータのホストが複数のIPアドレスに解決される場合は、
// タイムアウトは各連続したダイヤルに均等に分散され、それぞれが適切な時間の一部を接続に割り当てます。
//
// ネットワークとアドレスパラメータの詳細については、func Dialを参照してください。
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)

// Dialは指定されたネットワーク上のアドレスに接続します。
//
// ネットワークとアドレスの詳細は、func Dialの説明を参照してください。
//
// Dialは内部的に [context.Background] を使用します。コンテキストを指定するには、[Dialer.DialContext] を使用してください。
func (d *Dialer) Dial(network, address string) (Conn, error)

// DialContextは、指定されたコンテキストを使用して、指定されたネットワーク上のアドレスに接続します。
//
// 提供されたコンテキストは、nilでない必要があります。接続が完了する前にコンテキストが期限切れになると、エラーが返されます。接続が成功した後、コンテキストの期限切れは接続に影響しません。
//
// TCPを使用し、アドレスパラメータのホストが複数のネットワークアドレスに解決される場合、ダイヤルタイムアウト（d.Timeoutまたはctxから）は、各連続したダイヤルに均等に分散されます。それぞれのダイヤルには、適切な接続時間の割合が与えられます。
// 例えば、ホストが4つのIPアドレスを持ち、タイムアウトが1分の場合、次のアドレスを試す前に、各単一のアドレスへの接続には15秒の時間が与えられます。
//
<<<<<<< HEAD
// ネットワークやアドレスパラメータの説明については、[Dial] 関数を参照してください。
=======
// ネットワークやアドレスパラメータの説明については、func [Dial] を参照してください。
>>>>>>> release-branch.go1.22
func (d *Dialer) DialContext(ctx context.Context, network, address string) (Conn, error)

// ListenConfig はアドレスのリッスンに関するオプションを含んでいます。
type ListenConfig struct {

	// Controlがnilでない場合、ネットワーク接続を作成した後、
	// オペレーティングシステムにバインドする前に呼び出されます。
	//
	// Controlメソッドに渡されるネットワークとアドレスのパラメータは、
	// 必ずしもListenに渡されるものとは限りません。例えば、"tcp"を
	// Listenに渡すと、Control関数へは"tcp4"または"tcp6"が渡されます。
	Control func(network, address string, c syscall.RawConn) error

<<<<<<< HEAD
	// KeepAliveは、このリスナーによって受け入れられたネットワーク接続のキープアライブ期間を指定します。
	// ゼロの場合、プロトコルとオペレーティングシステムがサポートしている場合にキープアライブが有効になります。
	// キープアライブをサポートしていないネットワークプロトコルやオペレーティングシステムは、このフィールドを無視します。
	// マイナスの値の場合、キープアライブは無効になります。
	KeepAlive time.Duration

	// もしmptcpStatusがMultipath TCP（MPTCP）を許可する値に設定されている場合、ネットワークとして"tcp(4|6)"でListenを呼び出すと、オペレーティングシステムがサポートしている場合にはMPTCPが使用されます。
=======
	// KeepAlive specifies the keep-alive period for network
	// connections accepted by this listener.
	//
	// KeepAlive is ignored if KeepAliveConfig.Enable is true.
	//
	// If zero, keep-alive are enabled if supported by the protocol
	// and operating system. Network protocols or operating systems
	// that do not support keep-alive ignore this field.
	// If negative, keep-alive are disabled.
	KeepAlive time.Duration

	// KeepAliveConfig specifies the keep-alive probe configuration
	// for an active network connection, when supported by the
	// protocol and operating system.
	//
	// If KeepAliveConfig.Enable is true, keep-alive probes are enabled.
	// If KeepAliveConfig.Enable is false and KeepAlive is negative,
	// keep-alive probes are disabled.
	KeepAliveConfig KeepAliveConfig

	// If mptcpStatus is set to a value allowing Multipath TCP (MPTCP) to be
	// used, any call to Listen with "tcp(4|6)" as network will use MPTCP if
	// supported by the operating system.
>>>>>>> upstream/master
	mptcpStatus mptcpStatus
}

// MultipathTCPはMPTCPが使用されるかどうかを報告します。
//
// このメソッドはオペレーティングシステムがMPTCPをサポートしているかどうかを確認しません。
func (lc *ListenConfig) MultipathTCP() bool

// SetMultipathTCPは、オペレーティングシステムがサポートしている場合、[Listen] メソッドがMPTCPを使用するかどうかを指示します。
// このメソッドは、システムのデフォルトおよびGODEBUG=multipathtcp=...の設定を上書きします。
//
// ホスト上でMPTCPが利用できない場合、またはクライアントがサポートしていない場合、
// ListenメソッドはTCPにフォールバックします。
func (lc *ListenConfig) SetMultipathTCP(use bool)

// Listenはローカルネットワークアドレスでアナウンスします。
//
// ネットワークおよびアドレスの詳細については、func Listenを参照してください。
func (lc *ListenConfig) Listen(ctx context.Context, network, address string) (Listener, error)

// ListenPacketはローカルネットワークアドレスでアナウンスします。
//
// ネットワークとアドレスのパラメーターの説明については、func ListenPacketを参照してください。
func (lc *ListenConfig) ListenPacket(ctx context.Context, network, address string) (PacketConn, error)

// Listenはローカルネットワークアドレスでアナウンスします。
//
// ネットワークは"tcp"、"tcp4"、"tcp6"、"unix"、または"unixpacket"である必要があります。
//
// TCPネットワークの場合、アドレスパラメータのホストが空または明示的に指定されていないIPアドレスの場合、Listenは利用可能なすべてのユニキャストおよびエニーキャストIPアドレスでリッスンします。
// IPv4のみを使用する場合は、ネットワークに"tcp4"を使用します。
// アドレスにはホスト名を使用できますが、これは推奨されないため、ホストのIPアドレスの最大1つのリスナーが作成されます。
// アドレスパラメータのポートが空または"0"の場合、例えば"127.0.0.1:"や"[::1]:0"のように、ポート番号が自動的に選択されます。
// [Listener] の [Addr] メソッドを使用して、選択されたポートを取得できます。
//
<<<<<<< HEAD
// ネットワークおよびアドレスパラメータの説明については、[Dial] 関数を参照してください。
=======
// ネットワークおよびアドレスパラメータの説明については、func [Dial] を参照してください。
>>>>>>> release-branch.go1.22
//
// Listenは内部的にcontext.Backgroundを使用します。コンテキストを指定するには、[ListenConfig.Listen] を使用してください。
func Listen(network, address string) (Listener, error)

// ListenPacketはローカルネットワークアドレスでの通知を行います。
//
// ネットワークは「udp」「udp4」「udp6」「unixgram」またはIPトランスポートである必要があります。
// IPトランスポートは、次の形式で「ip」「ip4」、「ip6」のいずれかの後に「:」とリテラルのプロトコル番号またはプロトコル名が続きます。
// 例：「ip:1」または「ip:icmp」。
//
// UDPとIPネットワークの場合、アドレスパラメータのホストが空白またはリテラルの未指定のIPアドレスの場合、
// ListenPacketはマルチキャストIPアドレス以外のすべての利用可能なローカルシステムのIPアドレスでリスンします。
// IPv4のみを使用する場合は、ネットワークに「udp4」または「ip4:proto」を使用します。
// アドレスはホスト名を使用することもできますが、これは推奨されません。
// なぜなら、それによってホストのIPアドレスのうちの最大で1つのリスナが作成されるからです。
// アドレスパラメータのポートが空または「0」の場合、「127.0.0.1:」や「[::1]:0」といった形式で、ポート番号は自動的に選択されます。
// [PacketConn] のLocalAddrメソッドを使用して選択されたポートを特定することができます。
//
<<<<<<< HEAD
// ネットワークおよびアドレスパラメータの説明については、[Dial] 関数を参照してください。
=======
// ネットワークおよびアドレスパラメータの説明については、func [Dial] を参照してください。
>>>>>>> release-branch.go1.22
//
// ListenPacketは内部的にcontext.Backgroundを使用します。コンテキストを指定するには、
// [ListenConfig.ListenPacket] を使用してください。
func ListenPacket(network, address string) (PacketConn, error)
