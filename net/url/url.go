// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// urlパッケージはURLを解析し、クエリのエスケープを実装します。
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

// QueryUnescapeは [QueryEscape] の逆変換を行います
// "%AB"のような形式の3バイトエンコードされた部分文字列を
// 16進数でデコードされたバイト0xABに変換します
// もし%の後に2桁の16進数が続かない場合、エラーが返されます。
func QueryUnescape(s string) (string, error)

// PathUnescapeは [PathEscape] の逆の変換を行います。形式が"%AB"の各3バイトエンコードされた部分文字列をhexデコードされたバイト0xABに変換します。もし%の後に2桁の16進数が続かない場合、エラーが返されます。
//
// PathUnescapeは [QueryUnescape] と同じですが、'+'を' '（スペース）に変換しない点が異なります。
func PathUnescape(s string) (string, error)

// QueryEscapeは、文字列をエスケープして、安全に [URL] クエリ内に配置できるようにします。
func QueryEscape(s string) string

// PathEscapeは、文字列を安全に [URL] パスセグメント内に配置できるようにエスケープします。
// 必要に応じて特殊文字（/を含む）を%XXシーケンスで置き換えます。
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
// Pathフィールドは、デコードされた形式で保存されます：/%47%6f%2fは/Go/になります。
// 結果として、Path内のどのスラッシュが生のURL内のスラッシュであり、どのスラッシュが%2fであるかを区別することはできません。
// この区別はほとんど重要ではありませんが、重要な場合は、コードは [URL.EscapedPath] メソッドを使用する必要があります。
// このメソッドは、Pathの元のエンコーディングを保持します。
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

// Userは、提供されたユーザー名を含む [Userinfo] を返します
// パスワードは設定されていません。
func User(username string) *Userinfo

// UserPasswordは提供されたユーザー名とパスワードを含む [Userinfo] を返します。
// この機能は、レガシーウェブサイトでのみ使用するべきです。
// RFC 2396は、この方法でUserinfoを解釈することを「推奨されない」と警告しています。
// 「URIなどで平文で認証情報を渡すことは、ほとんどの場合セキュリティリスクとなっている」と述べています。
func UserPassword(username, password string) *Userinfo

// Userinfo型は、[URL] のユーザー名とパスワードの詳細を不変なカプセル化します。既存のUserinfo値には、ユーザー名が設定されていることが保証されています（RFC 2396で許可されているように、空にすることも可能です）、また、オプションでパスワードも持つことができます。
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

// Parseは生のURLを [URL] 構造に解析します。
//
// URLは相対的なもの（ホストなしのパス）または絶対的なもの（スキームで始まる）である可能性があります。
// スキームなしでホスト名とパスを解析しようとすることは無効ですが、解析の曖昧さにより、
// エラーを返さない場合があります。
func Parse(rawURL string) (*URL, error)

// ParseRequestURIは生のURLを [URL] 構造体に解析します。これは、URLがHTTPリクエストで受け取られたものであることを前提としており、urlは絶対URIまたは絶対パスとしてのみ解釈されます。
// 文字列urlには#fragmentの接尾辞がないことが前提とされています。
// （ウェブブラウザはURLをウェブサーバーに送信する前に#fragmentを取り除きます。）
func ParseRequestURI(rawURL string) (*URL, error)

// EscapedPathはu.Pathのエスケープされた形式を返します。
// 一般的には、任意のパスには複数のエスケープされた形式が存在します。
// EscapedPathはu.RawPathがu.Pathの有効なエスケープである場合にはu.RawPathを返します。
// そうでない場合、EscapedPathはu.RawPathを無視し、独自のエスケープ形式を計算します。
// [URL.String] メソッドと [URL.RequestURI] メソッドは、それぞれの結果を構築するためにEscapedPathを使用します。
// 一般的に、コードはu.RawPathを直接読むのではなく、EscapedPathを呼び出すべきです。
func (u *URL) EscapedPath() string

// EscapedFragmentはu.Fragmentのエスケープ形式を返します。
// 一般的には、任意のフラグメントには複数のエスケープ形式が存在します。
// u.Fragmentが有効なエスケープである場合、EscapedFragmentはu.RawFragmentを返します。
// そうでない場合、EscapedFragmentはu.RawFragmentを無視し、独自のエスケープ形式を計算します。
// [URL.String] メソッドは、その結果を構築するためにEscapedFragmentを使用します。
// 一般的には、コードはu.RawFragmentを直接読む代わりにEscapedFragmentを呼び出すべきです。
func (u *URL) EscapedFragment() string

// Stringは [URL] を有効なURL文字列に再構築します。
// 結果の一般的な形式は次のいずれかです:
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

// Redactedは [URL.String] と似ていますが、パスワードを「xxxxx」で置き換えます。
// u.User内のパスワードのみが伏せられます。
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

// IsAbsは [URL] が絶対であるかどうかを報告します。
// 絶対とは、空ではないスキームを持っていることを意味します。
func (u *URL) IsAbs() bool

// Parseはレシーバのコンテキストで [URL] を解析します。提供されたURLは相対的または絶対的である可能性があります。Parseは解析の失敗時にはnil、errを返し、それ以外の場合は [URL.ResolveReference] と同じ値を返します。
func (u *URL) Parse(ref string) (*URL, error)

// ResolveReferenceは、RFC 3986 Section 5.2 に従って、絶対ベースURI uからURIリファレンスを絶対URIに解決します。URIリファレンスは相対または絶対のどちらでもかまいません。ResolveReferenceは常に新しい [URL] インスタンスを返しますが、返されたURLがベースまたはリファレンスと同じであってもです。refが絶対URLの場合、ResolveReferenceはbaseを無視してrefのコピーを返します。
func (u *URL) ResolveReference(ref *URL) *URL

// QueryはRawQueryを解析し、対応する値を返します。
// 不正な値の組み合わせは静かに破棄されます。
// エラーをチェックするには [ParseQuery] を使用してください。
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

// JoinPath は、指定されたパス要素が既存のパスに結合され、
// 結果のパスが "./" や "../" の要素を除去された新しい [URL] を返します。
// 連続する複数の / 文字のシーケンスは、単一の / に縮小されます。
func (u *URL) JoinPath(elem ...string) *URL

// JoinPathは、指定されたパス要素が結合された [URL] 文字列を返します。
// ベースの既存パスと生成されたパスは、"./"や"../"要素が除去された状態でクリーンになります。
func JoinPath(base string, elem ...string) (result string, err error)
