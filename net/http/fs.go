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

// Dirは、特定のディレクトリツリーに制限されたネイティブファイルシステムを使用して [FileSystem] を実装します。
//
// [FileSystem.Open] メソッドは'/'で区切られたパスを取りますが、Dirの文字列値はURLではなくネイティブファイルシステム上のファイル名であるため、[filepath.Separator] で区切られます。これは必ずしも'/'ではありません。
//
// Dirは、機密ファイルやディレクトリを公開する可能性があります。Dirは、ディレクトリツリーから外部を指すシンボリックリンクを追跡します。これは、ユーザーが任意のシンボリックリンクを作成できるディレクトリからサービスを提供する場合に特に危険です。Dirは、ピリオドで始まるファイルやディレクトリにもアクセスを許可します。これには、.gitのような機密ディレクトリや.htpasswdのような機密ファイルが含まれます。ピリオドで始まるファイルを除外するには、ファイル/ディレクトリをサーバーから削除するか、カスタムFileSystem実装を作成してください。
//
// 空のDirは"."として扱われます。
type Dir string

// Openは、[os.Open] を使用して、ディレクトリdにルートされ、相対的なファイルを読み取るために [FileSystem] を実装します。
func (d Dir) Open(name string) (File, error)

// FileSystemは、名前付きファイルのコレクションへのアクセスを実装します。
// ファイルパスの要素は、ホストオペレーティングシステムの規約に関係なく、スラッシュ（'/'、U+002F）で区切られます。
// FileSystemを [Handler] に変換するには、[FileServer] 関数を参照してください。
//
// このインターフェースは、[fs.FS] インターフェースより前に存在しており、代わりに使用できます。
// [FS] アダプター関数は、fs.FSをFileSystemに変換します。
type FileSystem interface {
	Open(name string) (File, error)
}

// [FileSystem] のOpenメソッドによって返され、[FileServer] 実装によって提供されるファイルです。
//
// メソッドは、 [*os.File] と同じ動作をする必要があります。
type File interface {
	io.Closer
	io.Reader
	io.Seeker
	Readdir(count int) ([]fs.FileInfo, error)
	Stat() (fs.FileInfo, error)
}

// ServeContentは、提供されたReadSeekerの内容を使用してリクエストに応答します。
<<<<<<< HEAD
// [io.Copy] よりもServeContentの主な利点は、Rangeリクエストを適切に処理し、
=======
// ServeContentの [io.Copy] に対する主な利点は、Rangeリクエストを適切に処理し、
>>>>>>> release-branch.go1.22
// MIMEタイプを設定し、If-Match、If-Unmodified-Since、If-None-Match、
// If-Modified-Since、およびIf-Rangeリクエストを処理することです。
//
// レスポンスのContent-Typeヘッダーが設定されていない場合、ServeContentは
// 最初にnameのファイル拡張子からタイプを推測し、それが失敗した場合は、
<<<<<<< HEAD
// コンテンツの最初のブロックを読み取り、それを [DetectContentType] に渡すようにフォールバックします。
// それ以外の場合、nameは使用されません。特に、nameは空にでき、レスポンスで送信されることはありません。
=======
// コンテンツの最初のブロックを読み取り、それを [DetectContentType] に渡すことでフォールバックします。
// nameはそれ以外では使用されません。特に、空にすることができ、レスポンスで送信されることはありません。
>>>>>>> release-branch.go1.22
//
// modtimeがゼロ時またはUnixエポックでない場合、ServeContentは応答のLast-Modifiedヘッダーに含めます。
// リクエストにIf-Modified-Sinceヘッダーが含まれている場合、ServeContentはmodtimeを使用して、コンテンツを送信する必要があるかどうかを決定します。
//
// コンテンツのSeekメソッドは動作する必要があります。ServeContentは、コンテンツのサイズを決定するために、コンテンツの末尾にシークを使用します。
//
// 呼び出し元がRFC 7232、セクション2.3に従ってフォーマットされたwのETagヘッダーを設定している場合、ServeContentはそれを使用して、If-Match、If-None-Match、またはIf-Rangeを使用するリクエストを処理します。
//
// [*os.File] は [io.ReadSeeker] インターフェースを実装していることに注意してください。
func ServeContent(w ResponseWriter, req *Request, name string, modtime time.Time, content io.ReadSeeker)

// ServeFileは、指定された名前の
// ファイルまたはディレクトリの内容でリクエストに応答します。
//
// 提供されたファイル名またはディレクトリ名が相対パスの場合、それは
// 現在のディレクトリに対して相対的に解釈され、親ディレクトリに昇格することができます。
// 提供された名前がユーザー入力から構築されている場合、ServeFileを呼び出す前に
// サニタイズする必要があります。
//
// 予防措置として、ServeFileはr.URL.Pathに".."パス要素が含まれているリクエストを拒否します。
// これは、r.URL.Pathをサニタイズせずに [filepath.Join] で安全でなく使用し、
// その結果をname引数として使用する可能性のある呼び出し元に対する保護です。
//
// 別の特殊なケースとして、ServeFileはr.URL.Pathが
// "/index.html"で終わる任意のリクエストを、最後の
// "index.html"なしで同じパスにリダイレクトします。そのようなリダイレクトを避けるためには、
// パスを変更するか、[ServeContent] を使用します。
//
// それらの2つの特殊なケースの外では、ServeFileは
// r.URL.Pathを使用して提供するファイルやディレクトリを選択しません。
// 名前引数で提供されたファイルやディレクトリのみが使用されます。
func ServeFile(w ResponseWriter, r *Request, name string)

// ServeFileFSは、ファイルシステムfsysから指定されたファイルまたはディレクトリの内容でリクエストに応答します。
//
// 提供されたファイルまたはディレクトリ名が相対パスの場合、現在のディレクトリを基準に解釈され、親ディレクトリに移動することができます。
// 提供された名前がユーザー入力から構築された場合、[ServeFile] を呼び出す前にサニタイズする必要があります。
//
// 予防措置として、ServeFileはr.URL.Pathに".."パス要素が含まれているリクエストを拒否します。
// これにより、r.URL.Pathに [filepath.Join] を安全に使用せずにサニタイズせずに使用し、そのfilepath.Joinの結果を名前引数として使用する可能性がある呼び出し元を保護します。
//
// もう1つの特別な場合として、ServeFileはr.URL.Pathが"/index.html"で終わるリクエストを、最後の"index.html"を除いた同じパスにリダイレクトします。
// そのようなリダイレクトを回避するには、パスを変更するか、ServeContentを使用してください。
//
// これら2つの特別な場合以外では、ServeFileはファイルまたはディレクトリを選択するためにr.URL.Pathを使用しません。
// 名前引数で提供されたファイルまたはディレクトリのみが使用されます。
func ServeFileFS(w ResponseWriter, r *Request, fsys fs.FS, name string)

<<<<<<< HEAD
// FSはfsysを [FileSystem] の実装に変換します。
// これは [FileServer] と [NewFileTransport] で使用するためのものです。
// fsysによって提供されるファイルは [io.Seeker] を実装しなければなりません。
=======
// FSは、io.Seekerを実装する必要があるfsysを [FileSystem] 実装に変換します。
// [FileServer] および [NewFileTransport] で使用するためです。
// fsysで提供されるファイルは、[io.Seeker] を実装する必要があります。
>>>>>>> release-branch.go1.22
func FS(fsys fs.FS) FileSystem

// FileServerは、ルートでルートされたファイルシステムの内容でHTTPリクエストを処理するハンドラーを返します。
//
// 特別な場合として、返されたファイルサーバーは、"/index.html"で終わるリクエストを、最後の"index.html"を除いた同じパスにリダイレクトします。
//
// オペレーティングシステムのファイルシステム実装を使用するには、[http.Dir] を使用してください。
//
//	http.Handle("/", http.FileServer(http.Dir("/tmp")))
//
// [fs.FS] の実装を使用するには、代わりに [http.FileServerFS] を使用します。
func FileServer(root FileSystem) Handler

// FileServerFSは、ファイルシステムfsysの内容でHTTPリクエストを処理するハンドラを返します。
//
// 特別なケースとして、返されたファイルサーバーは、"/index.html"で終わる任意のリクエストを、
// 最後の"index.html"なしの同じパスにリダイレクトします。
//
//	http.Handle("/", http.FileServerFS(fsys))
func FileServerFS(root fs.FS) Handler
