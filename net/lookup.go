// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package net

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/internal/singleflight"
	"github.com/shogo82148/std/net/netip"
	"golang.org/x/sync/singleflight"
)

// DefaultResolverは、パッケージレベルのLookup関数と指定されていないResolverを持つDialersによって使用されるリゾルバです。
var DefaultResolver = &Resolver{}

// Resolverは名前や数値を検索します。
//
// nil *ResolverはゼロのResolverと同等です。
type Resolver struct {

	// PreferGoは、利用可能なプラットフォーム上でGoの組み込みDNSリゾルバーを優先するかどうかを制御します。これはGODEBUG=netdns=goを設定するのと同等ですが、このリゾルバーにのみスコープされます。
	PreferGo bool

	// StrictErrorsは、一時的なエラー（タイムアウト、ソケットエラー、およびSERVFAILを含む）の動作を制御します。Goの組み込みリゾルバを使用する場合、このオプションは複数のサブクエリからなるクエリ（A+AAAAアドレスの検索やDNS検索リストの走査など）に対して、部分的な結果を返す代わりに、エラーが発生した場合にクエリ全体を中止させます。これはデフォルトでは有効にされていませんが、AAAAクエリを正しく処理しないリゾルバとの互換性に影響を与える可能性があるためです。
	StrictErrors bool

	// Dialは、Go言語の組み込みDNSリゾルバがTCPおよびUDP接続を作成するために使用する代替ダイラーをオプションで指定します。アドレスパラメーターのホストは常にリテラルIPアドレスであり、ホスト名ではありません。また、アドレスパラメーターのポートはリテラルポート番号であり、サービス名ではありません。
	// 返されたConnがPacketConnでもある場合、送信および受信されるDNSメッセージはRFC 1035 Section 4.2.1 「UDP使用」に準拠する必要があります。
	Dial func(ctx context.Context, network, address string) (Conn, error)

	// lookupGroupは同じホストのルックアップをまとめてLookupIPAddr呼び出しをマージします。
	// lookupGroupのキーはLookupIPAddr.hostの引数です。
	// 返り値は([]IPAddr, error)です。
	lookupGroup singleflight.Group
}

// LookupHostは、ローカルのリゾルバを使用して指定されたホストを検索します。
// そのホストのアドレスのスライスを返します。
//
<<<<<<< HEAD
// LookupHostは、内部的にcontext.Backgroundを使用します。コンテキストを指定するには、
// Resolver.LookupHostを使用してください。
=======
// LookupHost uses [context.Background] internally; to specify the context, use
// [Resolver.LookupHost].
>>>>>>> upstream/release-branch.go1.22
func LookupHost(host string) (addrs []string, err error)

// LookupHostは、ローカルのリゾルバを使用して指定されたホストを検索します。
// そのホストのアドレスのスライスを返します。
func (r *Resolver) LookupHost(ctx context.Context, host string) (addrs []string, err error)

// LookupIPはローカルリゾルバを使用してホストを検索します。
// それはそのホストのIPv4およびIPv6アドレスのスライスを返します。
func LookupIP(host string) ([]IP, error)

// LookupIPAddrは、ローカルのリゾルバを使用してホストを検索します。
// そのホストのIPv4およびIPv6アドレスのスライスを返します。
func (r *Resolver) LookupIPAddr(ctx context.Context, host string) ([]IPAddr, error)

// LookupIPは、ローカルリゾルバーを使用して指定されたネットワークのホストを検索します。
// networkによって指定されたタイプのホストのIPアドレスのスライスを返します。
// networkは"ip"、"ip4"、または"ip6"のいずれかでなければなりません。
func (r *Resolver) LookupIP(ctx context.Context, network, host string) ([]IP, error)

// LookupNetIPはローカルリゾルバを使用してホストを検索します。
// それは、ネットワークで指定されたタイプのそのホストのIPアドレスのスライスを返します。
// ネットワークは、"ip"、"ip4"、または "ip6"のいずれかでなければなりません。
func (r *Resolver) LookupNetIP(ctx context.Context, network, host string) ([]netip.Addr, error)

var _ context.Context = (*onlyValuesCtx)(nil)

// LookupPortは指定されたネットワークとサービスに対するポートを調べます。
//
<<<<<<< HEAD
// LookupPortは内部でcontext.Backgroundを使用します。コンテキストを指定するには、
// Resolver.LookupPortを使用してください。
=======
// LookupPort uses [context.Background] internally; to specify the context, use
// [Resolver.LookupPort].
>>>>>>> upstream/release-branch.go1.22
func LookupPort(network, service string) (port int, err error)

// LookupPortは、指定されたネットワークとサービスのポートを検索します。
//
// networkは、"tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、または"ip"のいずれかでなければなりません。
func (r *Resolver) LookupPort(ctx context.Context, network, service string) (port int, err error)

<<<<<<< HEAD
// LookupCNAMEは指定されたホストの正式な名前（カノニカル名）を返します。
// カノニカル名に関心がない場合は、LookupHostまたはLookupIPを直接呼び出すことができます。
// どちらも、ルックアップの一部としてカノニカル名の解決を行います。
=======
// LookupCNAME returns the canonical name for the given host.
// Callers that do not care about the canonical name can call
// [LookupHost] or [LookupIP] directly; both take care of resolving
// the canonical name as part of the lookup.
>>>>>>> upstream/release-branch.go1.22
//
// カノニカル名は、ゼロまたは複数のCNAMEレコードを辿った後の最終的な名前です。
// hostにDNSの"CNAME"レコードが含まれていない場合でも、hostがアドレスレコードに解決されている限り、LookupCNAMEはエラーを返しません。
//
// 返されるカノニカル名は、正しくフォーマットされたプレゼンテーション形式のドメイン名であることが検証されます。
//
<<<<<<< HEAD
// LookupCNAMEは内部的にcontext.Backgroundを使用します。コンテキストを指定するには、Resolver.LookupCNAMEを使用してください。
func LookupCNAME(host string) (cname string, err error)

// LookupCNAMEは指定されたホストの正規名を返します。
// 正規名に関心を持たない呼び出し元は、
// LookupHostまたはLookupIPを直接呼び出すことができます。
// 両者は名前解決の一環として正規名を処理します。
=======
// LookupCNAME uses [context.Background] internally; to specify the context, use
// [Resolver.LookupCNAME].
func LookupCNAME(host string) (cname string, err error)

// LookupCNAME returns the canonical name for the given host.
// Callers that do not care about the canonical name can call
// [LookupHost] or [LookupIP] directly; both take care of resolving
// the canonical name as part of the lookup.
>>>>>>> upstream/release-branch.go1.22
//
// 正規名は、ゼロ個以上のCNAMEレコードをたどった後の最終名です。
// LookupCNAMEは、ホストがDNSの"CNAME"レコードを含まない場合でも、
// ホストがアドレスレコードに解決されている限り、エラーを返しません。
//
// 返される正規名は、適切な形式のドメイン名であることが検証されます。
func (r *Resolver) LookupCNAME(ctx context.Context, host string) (string, error)

<<<<<<< HEAD
// LookupSRVは、指定されたサービス、プロトコル、およびドメイン名のSRVクエリを解決しようとします。
// protoは「tcp」または「udp」です。
// 返されるレコードは優先度に従ってソートされ、各優先度内で重みによってランダムになります。
=======
// LookupSRV tries to resolve an [SRV] query of the given service,
// protocol, and domain name. The proto is "tcp" or "udp".
// The returned records are sorted by priority and randomized
// by weight within a priority.
>>>>>>> upstream/release-branch.go1.22
//
// LookupSRVはRFC 2782に従って調べるDNS名を構築します。
// つまり、_service._proto.nameを検索します。非標準の名前でSRVレコードを公開するサービスに対応するために、
// serviceとprotoの両方が空の文字列の場合、LookupSRVは直接nameを検索します。
//
// 返されたサービス名は、適切な形式のプレゼンテーション形式のドメイン名であることが検証されます。
// 応答に無効な名前が含まれている場合、これらのレコードはフィルタリングされ、エラーが返されます。
// 残りの結果がある場合は、これらのエラーと一緒に返されます。
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)

<<<<<<< HEAD
// LookupSRVは、指定されたサービス、プロトコル、ドメイン名のSRVクエリを解決しようとします。
// プロトコルは「tcp」または「udp」です。
// 返されるレコードは優先度でソートされ、優先度内でのウェイトによってランダムになります。
=======
// LookupSRV tries to resolve an [SRV] query of the given service,
// protocol, and domain name. The proto is "tcp" or "udp".
// The returned records are sorted by priority and randomized
// by weight within a priority.
>>>>>>> upstream/release-branch.go1.22
//
// LookupSRVは、RFC 2782に従ってルックアップするためのDNS名を構築します。
// つまり、_service._proto.nameをルックアップします。非標準の名前の下にSRVレコードを公開するサービスを収容するために、
// serviceとprotoの両方が空の文字列の場合、LookupSRVは直接nameをルックアップします。
//
// 返されるサービス名は、正しくフォーマットされたプレゼンテーション形式のドメイン名であることが検証されます。
// レスポンスに無効な名前が含まれている場合、それらのレコードはフィルタリングされ、エラーが返されます。
// 残りの結果がある場合、それらと一緒にエラーが返されます。
func (r *Resolver) LookupSRV(ctx context.Context, service, proto, name string) (string, []*SRV, error)

// LookupMXは指定されたドメイン名のDNS MXレコードを優先度に従ってソートして返します。
//
// 返されるメールサーバー名は、正しくフォーマットされた表示形式のドメイン名であることが検証されます。
// レスポンスに無効な名前が含まれている場合、それらのレコードはフィルタリングされ、エラーと共に残りの結果が返されます（もしあれば）。
//
<<<<<<< HEAD
// LookupMXは内部的にcontext.Backgroundを使用します。コンテキストを指定するには、Resolver.LookupMXを使用してください。
=======
// LookupMX uses [context.Background] internally; to specify the context, use
// [Resolver.LookupMX].
>>>>>>> upstream/release-branch.go1.22
func LookupMX(name string) ([]*MX, error)

// LookupMXは、指定されたドメイン名のDNS MXレコードを優先度に基づいてソートして返します。
// 返されるメールサーバー名は正しくフォーマットされたプレゼンテーション形式のドメイン名であることが検証されます。
// レスポンスに無効な名前が含まれている場合、それらのレコードはフィルタリングされ、エラーが返されます。
// 残りの結果がある場合、それらとともにエラーが返されます。
func (r *Resolver) LookupMX(ctx context.Context, name string) ([]*MX, error)

// LookupNSは指定されたドメイン名のDNS NSレコードを返します。
//
// 返されるネームサーバ名は、正しくフォーマットされた表示形式のドメイン名であることが検証されます。
// 応答に無効な名前が含まれている場合、これらのレコードはフィルタリングされ、エラーが残りの結果と共に返されます。
//
<<<<<<< HEAD
// LookupNSは内部的にcontext.Backgroundを使用します。コンテキストを指定するには、Resolver.LookupNSを使用します。
=======
// LookupNS uses [context.Background] internally; to specify the context, use
// [Resolver.LookupNS].
>>>>>>> upstream/release-branch.go1.22
func LookupNS(name string) ([]*NS, error)

// LookupNSは指定されたドメイン名のDNS NSレコードを返します。
//
// 返されたネームサーバの名前は、正しくフォーマットされた
// プレゼンテーション形式のドメイン名であることが検証されます。
// もしレスポンスに無効な名前が含まれている場合、それらのレコードは
// フィルタリングされ、エラーが返されます。
// 残りの結果がある場合、それらとともにエラーが返されます。
func (r *Resolver) LookupNS(ctx context.Context, name string) ([]*NS, error)

// LookupTXTは指定されたドメイン名のDNS TXTレコードを返します。
//
<<<<<<< HEAD
// LookupTXTは内部でcontext.Backgroundを使用します。コンテキストを指定するには、
// Resolver.LookupTXTを使用してください。
=======
// LookupTXT uses [context.Background] internally; to specify the context, use
// [Resolver.LookupTXT].
>>>>>>> upstream/release-branch.go1.22
func LookupTXT(name string) ([]string, error)

// LookupTXTは指定されたドメイン名のDNSのTXTレコードを返します。
func (r *Resolver) LookupTXT(ctx context.Context, name string) ([]string, error)

<<<<<<< HEAD
// LookupAddrは与えられたアドレスに対して逆引きを行い、そのアドレスにマッピングされる名前のリストを返します。
// 返された名前は適切にフォーマットされたプレゼンテーション形式のドメイン名であることが検証されます。応答に無効な名前が含まれている場合、それらのレコードはフィルタリングされ、エラーと一緒に残りの結果（ある場合）が返されます。
// ホストCライブラリリゾルバを使用する場合、最大で1つの結果が返されます。ホストリゾルバをバイパスするには、カスタムリゾルバを使用してください。
// LookupAddrは内部でcontext.Backgroundを使用します。コンテキストを指定するには、Resolver.LookupAddrを使用してください。
=======
// LookupAddr performs a reverse lookup for the given address, returning a list
// of names mapping to that address.
//
// The returned names are validated to be properly formatted presentation-format
// domain names. If the response contains invalid names, those records are filtered
// out and an error will be returned alongside the remaining results, if any.
//
// When using the host C library resolver, at most one result will be
// returned. To bypass the host resolver, use a custom [Resolver].
//
// LookupAddr uses [context.Background] internally; to specify the context, use
// [Resolver.LookupAddr].
>>>>>>> upstream/release-branch.go1.22
func LookupAddr(addr string) (names []string, err error)

// LookupAddrは指定されたアドレスの逆引きを行い、そのアドレスにマッピングされる名前のリストを返します。
//
// 返された名前は適切なフォーマットのプレゼンテーション形式のドメイン名であることが検証されます。
// もし回答に無効な名前が含まれている場合、それらのレコードはフィルタリングされ、
// 残りの結果がある場合はエラーが返されます。
func (r *Resolver) LookupAddr(ctx context.Context, addr string) ([]string, error)
