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

// Dirは、特定のディレクトリツリーに制限されたネイティブファイルシステムを使用してFileSystemを実装します。

// FileSystem.Openメソッドは'/'で区切られたパスを取りますが、Dirの文字列値はURLではなくネイティブファイルシステム上のファイル名であるため、filepath.Separatorで区切られます。これは必ずしも'/'ではありません。

// Dirは、機密ファイルやディレクトリを公開する可能性があります。Dirは、ディレクトリツリーから外部を指すシンボリックリンクを追跡します。これは、ユーザーが任意のシンボリックリンクを作成できるディレクトリからサービスを提供する場合に特に危険です。Dirは、ピリオドで始まるファイルやディレクトリにもアクセスを許可します。これには、.gitのような機密ディレクトリや.htpasswdのような機密ファイルが含まれます。ピリオドで始まるファイルを除外するには、ファイル/ディレクトリをサーバーから削除するか、カスタムFileSystem実装を作成してください。

// 空のDirは"."として扱われます。
type Dir string

// Openは、os.Openを使用して、ディレクトリdにルートされ、相対的なファイルを読み取るためにFileSystemを実装します。
func (d Dir) Open(name string) (File, error)

// FileSystemは、名前付きファイルのコレクションへのアクセスを実装します。
// ファイルパスの要素は、ホストオペレーティングシステムの規約に関係なく、スラッシュ（'/'、U+002F）で区切られます。
// FileSystemをHandlerに変換するには、FileServer関数を参照してください。
//
// このインターフェースは、fs.FSインターフェースより前に存在しており、代わりに使用できます。
// FSアダプター関数は、fs.FSをFileSystemに変換します。
type FileSystem interface {
	Open(name string) (File, error)
}

// FileSystemのOpenメソッドによって返され、FileServer実装によって提供されるファイルです。

// メソッドは、 *os.File と同じ動作をする必要があります。
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]fs.FileInfo, error)
	Stat() (fs.FileInfo, error)
}

// ServeContentは、提供されたReadSeeker内のコンテンツを使用してリクエストに応答します。
// ServeContentの主な利点は、Rangeリクエストを適切に処理し、MIMEタイプを設定し、If-Match、If-Unmodified-Since、If-None-Match、If-Modified-Since、およびIf-Rangeリクエストを処理することです。
//
// 応答のContent-Typeヘッダーが設定されていない場合、ServeContentはまず、名前のファイル拡張子からタイプを推測し、それでも失敗した場合は、コンテンツの最初のブロックを読み取ってDetectContentTypeに渡します。
// 名前はそれ以外では使用されず、特に空であっても、応答に送信されません。
//
// modtimeがゼロ時またはUnixエポックでない場合、ServeContentは応答のLast-Modifiedヘッダーに含めます。
// リクエストにIf-Modified-Sinceヘッダーが含まれている場合、ServeContentはmodtimeを使用して、コンテンツを送信する必要があるかどうかを決定します。
//
// コンテンツのSeekメソッドは動作する必要があります。ServeContentは、コンテンツのサイズを決定するために、コンテンツの末尾にシークを使用します。
//
// 呼び出し元がRFC 7232、セクション2.3に従ってフォーマットされたwのETagヘッダーを設定している場合、ServeContentはそれを使用して、If-Match、If-None-Match、またはIf-Rangeを使用するリクエストを処理します。
//
// *os.Fileはio.ReadSeekerインターフェースを実装していることに注意してください。
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)

// ServeFile replies to the request with the contents of the named
// file or directory.
//
// If the provided file or directory name is a relative path, it is
// interpreted relative to the current directory and may ascend to
// parent directories. If the provided name is constructed from user
// input, it should be sanitized before calling ServeFile.
//
// As a precaution, ServeFile will reject requests where r.URL.Path
// contains a ".." path element; this protects against callers who
// might unsafely use filepath.Join on r.URL.Path without sanitizing
// it and then use that filepath.Join result as the name argument.
//
// As another special case, ServeFile redirects any request where r.URL.Path
// ends in "/index.html" to the same path, without the final
// "index.html". To avoid such redirects either modify the path or
// use ServeContent.
//
// Outside of those two special cases, ServeFile does not use
// r.URL.Path for selecting the file or directory to serve; only the
// file or directory provided in the name argument is used.
func ServeFile(w ResponseWriter, r *Request, name string)

// FSは、fsysをFileSystem実装に変換し、FileServerおよびNewFileTransportで使用するために使用されます。
// fsysによって提供されるファイルは、io.Seekerを実装する必要があります。
func FS(fsys fs.FS) FileSystem

// FileServerは、ルートでルートされたファイルシステムの内容でHTTPリクエストを処理するハンドラーを返します。
//
// 特別な場合として、返されたファイルサーバーは、"/index.html"で終わるリクエストを、最後の"index.html"を除いた同じパスにリダイレクトします。
//
// オペレーティングシステムのファイルシステム実装を使用するには、http.Dirを使用してください。
//
//	http.Handle("/", http.FileServer(http.Dir("/tmp")))
//
// fs.FS実装を使用するには、http.FSを使用して変換してください。
//
//	http.Handle("/", http.FileServer(http.FS(fsys)))
func FileServer(root FileSystem) Handler
