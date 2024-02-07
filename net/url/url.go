// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package urlはURLを解析し、クエリのエスケープを実装します。
package url

// Errorはエラーと、それが発生した操作とURLを報告します。
type Error struct {
	Op  string
	URL string
	Err error
}

func (e *Error) Unwrap() error
func (e *Error) Error() string

func (e *Error) Timeout() bool

func (e *Error) Temporary() bool

type EscapeError string

func (e EscapeError) Error() string

type InvalidHostError string

func (e InvalidHostError) Error() string

<<<<<<< HEAD
// QueryUnescapeはQueryEscapeの逆変換を行います
// "%AB"のような形式の3バイトエンコードされた部分文字列を
// 16進数でデコードされたバイト0xABに変換します
// もし%の後に2桁の16進数が続かない場合、エラーが返されます。
func QueryUnescape(s string) (string, error)

// PathUnescapeはPathEscapeの逆の変換を行います。形式が"%AB"の各3バイトエンコードされた部分文字列をhexデコードされたバイト0xABに変換します。もし%の後に2桁の16進数が続かない場合、エラーが返されます。
//
// PathUnescapeはQueryUnescapeと同じですが、'+'を' '（スペース）に変換しない点が異なります。
func PathUnescape(s string) (string, error)

func QueryEscape(s string) string

// PathEscapeは、文字列を安全にURLパスセグメント内に配置できるようにエスケープします。
// 必要に応じて特殊文字（/を含む）を%XXシーケンスで置き換えます。
=======
// QueryUnescape does the inverse transformation of [QueryEscape],
// converting each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB.
// It returns an error if any % is not followed by two hexadecimal
// digits.
func QueryUnescape(s string) (string, error)

// PathUnescape does the inverse transformation of [PathEscape],
// converting each 3-byte encoded substring of the form "%AB" into the
// hex-decoded byte 0xAB. It returns an error if any % is not followed
// by two hexadecimal digits.
//
// PathUnescape is identical to [QueryUnescape] except that it does not
// unescape '+' to ' ' (space).
func PathUnescape(s string) (string, error)

// QueryEscape escapes the string so it can be safely placed
// inside a [URL] query.
func QueryEscape(s string) string

// PathEscape escapes the string so it can be safely placed inside a [URL] path segment,
// replacing special characters (including /) with %XX sequences as needed.
>>>>>>> upstream/release-branch.go1.22
func PathEscape(s string) string

// URLは解析されたURL（厳密にはURIリファレンス）を表します。
//
// 全般的な形式は次のようになります：
//
//	[スキーム:][//[ユーザー情報@]ホスト][/パス][?クエリ][#フラグメント]
//
// スキームの後にスラッシュで始まらないURLは次のように解釈されます：
//
//	スキーム:透明部分[?クエリ][#フラグメント]
//
// Hostフィールドには、URLのホストとポートのサブコンポーネントが含まれます。
// ポートが存在する場合、コロンでホストから分離されます。
// ホストがIPv6アドレスの場合、角括弧で囲む必要があります: "[fe80::1]:80"。
// [net.JoinHostPort] 関数は、必要に応じてホストに角括弧を追加して、ホストとポートを文字列に結合します。
//
<<<<<<< HEAD
// Pathフィールドは、デコードされた形式で保存されます：/%47%6f%2fは/Go/になります。
// 結果として、Path内のどのスラッシュが生のURL内のスラッシュであり、どのスラッシュが%2fであるかを区別することはできません。
// この区別はほとんど重要ではありませんが、重要な場合は、コードはEscapedPathメソッドを使用する必要があります。
// このメソッドは、Pathの元のエンコーディングを保持します。
=======
// Note that the Path field is stored in decoded form: /%47%6f%2f becomes /Go/.
// A consequence is that it is impossible to tell which slashes in the Path were
// slashes in the raw URL and which were %2f. This distinction is rarely important,
// but when it is, the code should use the [URL.EscapedPath] method, which preserves
// the original encoding of Path.
>>>>>>> upstream/release-branch.go1.22
//
// RawPathフィールドは、デフォルトのパスのエンコードがエスケープされたパスと異なる場合にのみ設定されるオプションのフィールドです。
// 詳細については、EscapedPathメソッドを参照してください。
//
// URLのStringメソッドは、パスを取得するためにEscapedPathメソッドを使用します。
type URL struct {
	Scheme      string
	Opaque      string
	User        *Userinfo
	Host        string
	Path        string
	RawPath     string
	OmitHost    bool
	ForceQuery  bool
	RawQuery    string
	Fragment    string
	RawFragment string
}

<<<<<<< HEAD
// Userは、提供されたユーザー名を含むUserinfoを返します
// パスワードは設定されていません。
func User(username string) *Userinfo

// UserPasswordは提供されたユーザー名とパスワードを含むUserinfoを返します。
// この機能は、レガシーウェブサイトでのみ使用するべきです。
// RFC 2396は、この方法でUserinfoを解釈することを「推奨されない」と警告しています。
// 「URIなどで平文で認証情報を渡すことは、ほとんどの場合セキュリティリスクとなっている」と述べています。
func UserPassword(username, password string) *Userinfo

// Userinfo型は、URLのユーザー名とパスワードの詳細を不変なカプセル化します。既存のUserinfo値には、ユーザー名が設定されていることが保証されています（RFC 2396で許可されているように、空にすることも可能です）、また、オプションでパスワードも持つことができます。
=======
// User returns a [Userinfo] containing the provided username
// and no password set.
func User(username string) *Userinfo

// UserPassword returns a [Userinfo] containing the provided username
// and password.
//
// This functionality should only be used with legacy web sites.
// RFC 2396 warns that interpreting Userinfo this way
// “is NOT RECOMMENDED, because the passing of authentication
// information in clear text (such as URI) has proven to be a
// security risk in almost every case where it has been used.”
func UserPassword(username, password string) *Userinfo

// The Userinfo type is an immutable encapsulation of username and
// password details for a [URL]. An existing Userinfo value is guaranteed
// to have a username set (potentially empty, as allowed by RFC 2396),
// and optionally a password.
>>>>>>> upstream/release-branch.go1.22
type Userinfo struct {
	username    string
	password    string
	passwordSet bool
}

// Usernameはユーザー名を返します。
func (u *Userinfo) Username() string

// Passwordは設定されている場合はパスワードを返し、設定されているかどうかも返します。
func (u *Userinfo) Password() (string, bool)

// Stringは「username [: password]」の標準形式でエンコードされたユーザー情報を返します。
func (u *Userinfo) String() string

<<<<<<< HEAD
// Parseは生のURLをURL構造に解析します。
=======
// Parse parses a raw url into a [URL] structure.
>>>>>>> upstream/release-branch.go1.22
//
// URLは相対的なもの（ホストなしのパス）または絶対的なもの（スキームで始まる）である可能性があります。
// スキームなしでホスト名とパスを解析しようとすることは無効ですが、解析の曖昧さにより、
// エラーを返さない場合があります。
func Parse(rawURL string) (*URL, error)

<<<<<<< HEAD
// ParseRequestURIは生のURLをURL構造体に解析します。これは、URLがHTTPリクエストで受け取られたものであることを前提としており、urlは絶対URIまたは絶対パスとしてのみ解釈されます。
// 文字列urlには#fragmentの接尾辞がないことが前提とされています。
// （ウェブブラウザはURLをウェブサーバーに送信する前に#fragmentを取り除きます。）
func ParseRequestURI(rawURL string) (*URL, error)

// EscapedPathはu.Pathのエスケープされた形式を返します。
// 一般的には、任意のパスには複数のエスケープされた形式が存在します。
// EscapedPathはu.RawPathがu.Pathの有効なエスケープである場合にはu.RawPathを返します。
// そうでない場合、EscapedPathはu.RawPathを無視し、独自のエスケープ形式を計算します。
// StringメソッドとRequestURIメソッドは、それぞれの結果を構築するためにEscapedPathを使用します。
// 一般的に、コードはu.RawPathを直接読むのではなく、EscapedPathを呼び出すべきです。
func (u *URL) EscapedPath() string

// EscapedFragmentはu.Fragmentのエスケープ形式を返します。
// 一般的には、任意のフラグメントには複数のエスケープ形式が存在します。
// u.Fragmentが有効なエスケープである場合、EscapedFragmentはu.RawFragmentを返します。
// そうでない場合、EscapedFragmentはu.RawFragmentを無視し、独自のエスケープ形式を計算します。
// Stringメソッドは、その結果を構築するためにEscapedFragmentを使用します。
// 一般的には、コードはu.RawFragmentを直接読む代わりにEscapedFragmentを呼び出すべきです。
func (u *URL) EscapedFragment() string

// StringはURLを有効なURL文字列に再構築します。
// 結果の一般的な形式は次のいずれかです:
=======
// ParseRequestURI parses a raw url into a [URL] structure. It assumes that
// url was received in an HTTP request, so the url is interpreted
// only as an absolute URI or an absolute path.
// The string url is assumed not to have a #fragment suffix.
// (Web browsers strip #fragment before sending the URL to a web server.)
func ParseRequestURI(rawURL string) (*URL, error)

// EscapedPath returns the escaped form of u.Path.
// In general there are multiple possible escaped forms of any path.
// EscapedPath returns u.RawPath when it is a valid escaping of u.Path.
// Otherwise EscapedPath ignores u.RawPath and computes an escaped
// form on its own.
// The [URL.String] and [URL.RequestURI] methods use EscapedPath to construct
// their results.
// In general, code should call EscapedPath instead of
// reading u.RawPath directly.
func (u *URL) EscapedPath() string

// EscapedFragment returns the escaped form of u.Fragment.
// In general there are multiple possible escaped forms of any fragment.
// EscapedFragment returns u.RawFragment when it is a valid escaping of u.Fragment.
// Otherwise EscapedFragment ignores u.RawFragment and computes an escaped
// form on its own.
// The [URL.String] method uses EscapedFragment to construct its result.
// In general, code should call EscapedFragment instead of
// reading u.RawFragment directly.
func (u *URL) EscapedFragment() string

// String reassembles the [URL] into a valid URL string.
// The general form of the result is one of:
>>>>>>> upstream/release-branch.go1.22
//
// scheme:opaque?query#fragment
// scheme://userinfo@host/path?query#fragment
//
// もしu.Opaqueが空でない場合、Stringは最初の形式を使用します。
// そうでなければ、2番目の形式を使用します。
// ホスト内の非ASCII文字はエスケープされます。
// パスを取得するために、Stringはu.EscapedPath()を使用します。
//
// 2番目の形式では、以下のルールが適用されます:
//   - もしu.Schemeが空なら、scheme:は省略されます。
//   - もしu.Userがnilなら、userinfo@は省略されます。
//   - もしu.Hostが空なら、host/は省略されます。
//   - もしu.Schemeとu.Hostが空であり、u.Userがnilなら、
//     scheme://userinfo@host/全体が省略されます。
//   - もしu.Hostが空でなく、u.Pathが/で始まるなら、
//     host/pathの形式は独自の/を追加しません。
//   - もしu.RawQueryが空なら、?queryは省略されます。
//   - もしu.Fragmentが空なら、#fragmentは省略されます。
func (u *URL) String() string

<<<<<<< HEAD
// RedactedはStringと似ていますが、パスワードを「xxxxx」で置き換えます。
// u.User内のパスワードのみが伏せられます。
=======
// Redacted is like [URL.String] but replaces any password with "xxxxx".
// Only the password in u.User is redacted.
>>>>>>> upstream/release-branch.go1.22
func (u *URL) Redacted() string

// Valuesは文字列のキーを値のリストにマップします。
// 通常、クエリパラメータやフォームの値に使用されます。
// http.Headerマップとは異なり、Valuesマップのキーは大文字小文字を区別します。
type Values map[string][]string

// Getは指定したキーに関連付けられた最初の値を返します。
// キーに関連付けられた値がない場合、Getは空の文字列を返します。
// 複数の値にアクセスするには、直接マップを使用してください。
func (v Values) Get(key string) string

// Setはキーを値にセットします。既存の値を置き換えます。
func (v Values) Set(key, value string)

// Addはkeyに値を追加します。keyと関連付けられた既存の値に追加します。
func (v Values) Add(key, value string)

// Delはキーに関連付けられた値を削除します。
func (v Values) Del(key string)

// Hasは与えられたキーが設定されているかどうかを確認します。
func (v Values) Has(key string) bool

// ParseQueryはURLエンコードされたクエリ文字列を解析して、
// 各キーに指定された値をリストしたマップを返します。
// ParseQueryは常に、最初にエンコードできないエラーが見つかった場合を示すnon-nilのマップを返します。エラーの詳細はerrに記載されます。
//
// クエリはアンパサンドで区切られたキー=値のリストとして期待されます。
// イコール記号がない設定は、空の値に設定されたキーとして解釈されます。
// URLエンコードされていないセミコロンが含まれる設定は無効と見なされます。
func ParseQuery(query string) (Values, error)

// Encodeは値をキーでソートされた形式で「URLエンコード」します
// ("bar=baz&foo=quux")
func (v Values) Encode() string

<<<<<<< HEAD
// IsAbsはURLが絶対であるかどうかを報告します。
// 絶対とは、空ではないスキームを持っていることを意味します。
func (u *URL) IsAbs() bool

// ParseはレシーバのコンテキストでURLを解析します。提供されたURLは相対的または絶対的である可能性があります。Parseは解析の失敗時にはnil、errを返し、それ以外の場合はResolveReferenceと同じ値を返します。
func (u *URL) Parse(ref string) (*URL, error)

// ResolveReferenceは、RFC 3986 Section 5.2 に従って、絶対ベースURI uからURIリファレンスを絶対URIに解決します。URIリファレンスは相対または絶対のどちらでもかまいません。ResolveReferenceは常に新しいURLインスタンスを返しますが、返されたURLがベースまたはリファレンスと同じであってもです。refが絶対URLの場合、ResolveReferenceはbaseを無視してrefのコピーを返します。
func (u *URL) ResolveReference(ref *URL) *URL

// QueryはRawQueryを解析し、対応する値を返します。
// 不正な値の組み合わせは静かに破棄されます。
// エラーをチェックするにはParseQueryを使用してください。
=======
// IsAbs reports whether the [URL] is absolute.
// Absolute means that it has a non-empty scheme.
func (u *URL) IsAbs() bool

// Parse parses a [URL] in the context of the receiver. The provided URL
// may be relative or absolute. Parse returns nil, err on parse
// failure, otherwise its return value is the same as [URL.ResolveReference].
func (u *URL) Parse(ref string) (*URL, error)

// ResolveReference resolves a URI reference to an absolute URI from
// an absolute base URI u, per RFC 3986 Section 5.2. The URI reference
// may be relative or absolute. ResolveReference always returns a new
// [URL] instance, even if the returned URL is identical to either the
// base or reference. If ref is an absolute URL, then ResolveReference
// ignores base and returns a copy of ref.
func (u *URL) ResolveReference(ref *URL) *URL

// Query parses RawQuery and returns the corresponding values.
// It silently discards malformed value pairs.
// To check errors use [ParseQuery].
>>>>>>> upstream/release-branch.go1.22
func (u *URL) Query() Values

// RequestURIは、uのHTTPリクエストで使用される、エンコードされたpath?queryまたはopaque?queryの文字列を返します。
func (u *URL) RequestURI() string

// Hostnameは、存在する場合は有効なポート番号を削除してu.Hostを返します。
//
// 結果が角かっこで囲まれている場合、それはリテラルIPv6アドレスですので、
// 結果から角かっこは削除されます。
func (u *URL) Hostname() string

// Portはu.Hostのポート部分を返しますが、先頭のコロンは除かれます。
//
// もしu.Hostに有効な数値のポートが含まれていない場合、Portは空の文字列を返します。
func (u *URL) Port() string

func (u *URL) MarshalBinary() (text []byte, err error)

func (u *URL) UnmarshalBinary(text []byte) error

<<<<<<< HEAD
// JoinPath は、指定されたパス要素が既存のパスに結合され、
// 結果のパスが "./" や "../" の要素を除去された新しいURLを返します。
// 連続する複数の / 文字のシーケンスは、単一の / に縮小されます。
func (u *URL) JoinPath(elem ...string) *URL

// JoinPathは、指定されたパス要素が結合されたURL文字列を返します。
// ベースの既存パスと生成されたパスは、"./"や"../"要素が除去された状態でクリーンになります。
=======
// JoinPath returns a new [URL] with the provided path elements joined to
// any existing path and the resulting path cleaned of any ./ or ../ elements.
// Any sequences of multiple / characters will be reduced to a single /.
func (u *URL) JoinPath(elem ...string) *URL

// JoinPath returns a [URL] string with the provided path elements joined to
// the existing path of base and the resulting path cleaned of any ./ or ../ elements.
>>>>>>> upstream/release-branch.go1.22
func JoinPath(base string, elem ...string) (result string, err error)
