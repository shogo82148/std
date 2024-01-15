// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// HTTP file system request handler

package http

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
	"github.com/shogo82148/std/time"
)

<<<<<<< HEAD
// Dirは、特定のディレクトリツリーに制限されたネイティブファイルシステムを使用してFileSystemを実装します。
//
// FileSystem.Openメソッドは'/'で区切られたパスを取りますが、Dirの文字列値はURLではなくネイティブファイルシステム上のファイル名であるため、filepath.Separatorで区切られます。これは必ずしも'/'ではありません。
=======
// A Dir implements [FileSystem] using the native file system restricted to a
// specific directory tree.
//
// While the [FileSystem.Open] method takes '/'-separated paths, a Dir's string
// value is a filename on the native file system, not a URL, so it is separated
// by [filepath.Separator], which isn't necessarily '/'.
>>>>>>> upstream/master
//
// Dirは、機密ファイルやディレクトリを公開する可能性があります。Dirは、ディレクトリツリーから外部を指すシンボリックリンクを追跡します。これは、ユーザーが任意のシンボリックリンクを作成できるディレクトリからサービスを提供する場合に特に危険です。Dirは、ピリオドで始まるファイルやディレクトリにもアクセスを許可します。これには、.gitのような機密ディレクトリや.htpasswdのような機密ファイルが含まれます。ピリオドで始まるファイルを除外するには、ファイル/ディレクトリをサーバーから削除するか、カスタムFileSystem実装を作成してください。
//
// 空のDirは"."として扱われます。
type Dir string

<<<<<<< HEAD
// Openは、os.Openを使用して、ディレクトリdにルートされ、相対的なファイルを読み取るためにFileSystemを実装します。
func (d Dir) Open(name string) (File, error)

// FileSystemは、名前付きファイルのコレクションへのアクセスを実装します。
// ファイルパスの要素は、ホストオペレーティングシステムの規約に関係なく、スラッシュ（'/'、U+002F）で区切られます。
// FileSystemをHandlerに変換するには、FileServer関数を参照してください。
//
// このインターフェースは、fs.FSインターフェースより前に存在しており、代わりに使用できます。
// FSアダプター関数は、fs.FSをFileSystemに変換します。
=======
// Open implements [FileSystem] using [os.Open], opening files for reading rooted
// and relative to the directory d.
func (d Dir) Open(name string) (File, error)

// A FileSystem implements access to a collection of named files.
// The elements in a file path are separated by slash ('/', U+002F)
// characters, regardless of host operating system convention.
// See the [FileServer] function to convert a FileSystem to a [Handler].
//
// This interface predates the [fs.FS] interface, which can be used instead:
// the [FS] adapter function converts an fs.FS to a FileSystem.
>>>>>>> upstream/master
type FileSystem interface {
	Open(name string) (File, error)
}

<<<<<<< HEAD
// FileSystemのOpenメソッドによって返され、FileServer実装によって提供されるファイルです。
//
// メソッドは、 *os.File と同じ動作をする必要があります。
=======
// A File is returned by a [FileSystem]'s Open method and can be
// served by the [FileServer] implementation.
//
// The methods should behave the same as those on an [*os.File].
>>>>>>> upstream/master
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]fs.FileInfo, error)
	Stat() (fs.FileInfo, error)
}

<<<<<<< HEAD
// ServeContentは、提供されたReadSeeker内のコンテンツを使用してリクエストに応答します。
// ServeContentの主な利点は、Rangeリクエストを適切に処理し、MIMEタイプを設定し、If-Match、If-Unmodified-Since、If-None-Match、If-Modified-Since、およびIf-Rangeリクエストを処理することです。
//
// 応答のContent-Typeヘッダーが設定されていない場合、ServeContentはまず、名前のファイル拡張子からタイプを推測し、それでも失敗した場合は、コンテンツの最初のブロックを読み取ってDetectContentTypeに渡します。
// 名前はそれ以外では使用されず、特に空であっても、応答に送信されません。
=======
// ServeContent replies to the request using the content in the
// provided ReadSeeker. The main benefit of ServeContent over [io.Copy]
// is that it handles Range requests properly, sets the MIME type, and
// handles If-Match, If-Unmodified-Since, If-None-Match, If-Modified-Since,
// and If-Range requests.
//
// If the response's Content-Type header is not set, ServeContent
// first tries to deduce the type from name's file extension and,
// if that fails, falls back to reading the first block of the content
// and passing it to [DetectContentType].
// The name is otherwise unused; in particular it can be empty and is
// never sent in the response.
>>>>>>> upstream/master
//
// modtimeがゼロ時またはUnixエポックでない場合、ServeContentは応答のLast-Modifiedヘッダーに含めます。
// リクエストにIf-Modified-Sinceヘッダーが含まれている場合、ServeContentはmodtimeを使用して、コンテンツを送信する必要があるかどうかを決定します。
//
// コンテンツのSeekメソッドは動作する必要があります。ServeContentは、コンテンツのサイズを決定するために、コンテンツの末尾にシークを使用します。
//
// 呼び出し元がRFC 7232、セクション2.3に従ってフォーマットされたwのETagヘッダーを設定している場合、ServeContentはそれを使用して、If-Match、If-None-Match、またはIf-Rangeを使用するリクエストを処理します。
//
<<<<<<< HEAD
// *os.Fileはio.ReadSeekerインターフェースを実装していることに注意してください。
=======
// Note that [*os.File] implements the [io.ReadSeeker] interface.
>>>>>>> upstream/master
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)

// ServeFileは、指定された名前の
// ファイルまたはディレクトリの内容でリクエストに応答します。
//
// 提供されたファイル名またはディレクトリ名が相対パスの場合、それは
// 現在のディレクトリに対して相対的に解釈され、親ディレクトリに昇格することができます。
// 提供された名前がユーザー入力から構築されている場合、ServeFileを呼び出す前に
// サニタイズする必要があります。
//
<<<<<<< HEAD
// 予防措置として、ServeFileはr.URL.Pathに".."パス要素が含まれているリクエストを拒否します。
// これは、r.URL.Pathをサニタイズせずにfilepath.Joinで安全でなく使用し、
// その結果をname引数として使用する可能性のある呼び出し元に対する保護です。
//
// 別の特殊なケースとして、ServeFileはr.URL.Pathが
// "/index.html"で終わる任意のリクエストを、最後の
// "index.html"なしで同じパスにリダイレクトします。そのようなリダイレクトを避けるためには、
// パスを変更するか、ServeContentを使用します。
=======
// As a precaution, ServeFile will reject requests where r.URL.Path
// contains a ".." path element; this protects against callers who
// might unsafely use [filepath.Join] on r.URL.Path without sanitizing
// it and then use that filepath.Join result as the name argument.
//
// As another special case, ServeFile redirects any request where r.URL.Path
// ends in "/index.html" to the same path, without the final
// "index.html". To avoid such redirects either modify the path or
// use [ServeContent].
>>>>>>> upstream/master
//
// それらの2つの特殊なケースの外では、ServeFileは
// r.URL.Pathを使用して提供するファイルやディレクトリを選択しません。
// 名前引数で提供されたファイルやディレクトリのみが使用されます。
func ServeFile(w ResponseWriter, r *Request, name string)

// ServeFileFSは、ファイルシステムfsysから指定されたファイルまたはディレクトリの内容でリクエストに応答します。
//
<<<<<<< HEAD
// 提供されたファイルまたはディレクトリ名が相対パスの場合、現在のディレクトリを基準に解釈され、親ディレクトリに移動することができます。
// 提供された名前がユーザー入力から構築された場合、ServeFileを呼び出す前にサニタイズする必要があります。
//
// 予防措置として、ServeFileはr.URL.Pathに".."パス要素が含まれているリクエストを拒否します。
// これにより、r.URL.Pathにfilepath.Joinを安全に使用せずにサニタイズせずに使用し、そのfilepath.Joinの結果を名前引数として使用する可能性がある呼び出し元を保護します。
=======
// If the provided file or directory name is a relative path, it is
// interpreted relative to the current directory and may ascend to
// parent directories. If the provided name is constructed from user
// input, it should be sanitized before calling [ServeFile].
//
// As a precaution, ServeFile will reject requests where r.URL.Path
// contains a ".." path element; this protects against callers who
// might unsafely use [filepath.Join] on r.URL.Path without sanitizing
// it and then use that filepath.Join result as the name argument.
>>>>>>> upstream/master
//
// もう1つの特別な場合として、ServeFileはr.URL.Pathが"/index.html"で終わるリクエストを、最後の"index.html"を除いた同じパスにリダイレクトします。
// そのようなリダイレクトを回避するには、パスを変更するか、ServeContentを使用してください。
//
// これら2つの特別な場合以外では、ServeFileはファイルまたはディレクトリを選択するためにr.URL.Pathを使用しません。
// 名前引数で提供されたファイルまたはディレクトリのみが使用されます。
func ServeFileFS(w ResponseWriter, r *Request, fsys fs.FS, name string)

<<<<<<< HEAD
// FSは、io.Seekerを実装する必要があるfsysをFileSystem実装に変換します。
// FileServerおよびNewFileTransportで使用するためです。
=======
// FS converts fsys to a [FileSystem] implementation,
// for use with [FileServer] and [NewFileTransport].
// The files provided by fsys must implement [io.Seeker].
>>>>>>> upstream/master
func FS(fsys fs.FS) FileSystem

// FileServerは、ルートでルートされたファイルシステムの内容でHTTPリクエストを処理するハンドラーを返します。
//
// 特別な場合として、返されたファイルサーバーは、"/index.html"で終わるリクエストを、最後の"index.html"を除いた同じパスにリダイレクトします。
//
<<<<<<< HEAD
// オペレーティングシステムのファイルシステム実装を使用するには、http.Dirを使用してください。
//
//	http.Handle("/", http.FileServer(http.Dir("/tmp")))
//
// fs.FSの実装を使用するには、代わりにhttp.FileServerFSを使用します。
=======
// To use the operating system's file system implementation,
// use [http.Dir]:
//
//	http.Handle("/", http.FileServer(http.Dir("/tmp")))
//
// To use an [fs.FS] implementation, use [http.FileServerFS] instead.
>>>>>>> upstream/master
func FileServer(root FileSystem) Handler

// FileServerFSは、ファイルシステムfsysの内容でHTTPリクエストを処理するハンドラを返します。
//
// 特別なケースとして、返されたファイルサーバーは、"/index.html"で終わる任意のリクエストを、
// 最後の"index.html"なしの同じパスにリダイレクトします。
//
//	http.Handle("/", http.FileServerFS(fsys))
func FileServerFS(root fs.FS) Handler
