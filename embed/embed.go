// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// パッケージembedは、実行中のGoプログラムに埋め込まれたファイルへのアクセスを提供します。
//
// "embed"をインポートするGoソースファイルは、//go:embedディレクティブを使用して、
// コンパイル時にパッケージディレクトリまたはサブディレクトリから読み取ったファイルの内容を
// string、[]byte、または [FS] 型の変数で初期化できます。
//
// たとえば、hello.txtという名前のファイルを埋め込んで、その内容を実行時に表示する3つの方法を示します。
//
// 文字列に1つのファイルを埋め込む場合：
//
//	import _ "embed"
//
//	//go:embed hello.txt
//	var s string
//	print(s)
//
// 1つのファイルをバイトのスライスに埋め込む場合：
//
//	import _ "embed"
//
//	//go:embed hello.txt
//	var b []byte
//	print(string(b))
//
// 1つ以上のファイルをファイルシステムに埋め込む：
//
//	import "embed"
//
//	//go:embed hello.txt
//	var f embed.FS
//	data, _ := f.ReadFile("hello.txt")
//	print(string(data))
//
// # ディレクティブ
//
// 変数宣言の上にある//go:embedディレクティブは、埋め込むファイルを1つ以上のpath.Matchパターンを使用して指定します。
//
// ディレクティブは、単一の変数の宣言を含む行の直前にある必要があります。
// ディレクティブと宣言の間には、空白行と'//'行コメントのみが許可されます。
//
// 変数の型は、文字列型、バイト型のスライス、または [FS] （または [FS] のエイリアス）である必要があります。
//
// 例：
//
//	package server
//
//	import "embed"
//
//	// content holds our static web server content.
//	//go:embed image/* template/*
//	//go:embed html/index.html
//	var content embed.FS
//
// Goビルドシステムは、ディレクティブを認識し、宣言された変数（上記の例ではcontent）を、
// ファイルシステムからマッチングファイルで埋め込むように配置します。
//
// //go:embedディレクティブは、簡潔さのために複数のスペースで区切られたパターンを受け入れますが、
// 多数のパターンがある場合に非常に長い行を避けるために、繰り返すこともできます。
// パターンは、ソースファイルを含むパッケージディレクトリに対して解釈されます。
// パスセパレータはスラッシュであり、Windowsシステムでも同様です。
// パターンには、'.'、'..'、または空のパス要素を含めることはできず、スラッシュで始まることも終わることもできません。
// 現在のディレクトリのすべてをマッチングするには、'.'の代わりに'*'を使用します。
// スペースを含むファイル名を許可するために、パターンはGoのダブルクォートまたはバッククォートの文字列リテラルとして書くことができます。
//
// パターンがディレクトリを指定している場合、そのディレクトリを根とするサブツリー内のすべてのファイルが再帰的に埋め込まれます。
// ただし、名前が「.」または「_」で始まるファイルは除外されます。
// したがって、上記の例の変数はほぼ次のようになります。
//
//	// content is our static web server content.
//	//go:embed image template html/index.html
//	var content embed.FS
//
// 違いは、 'image/*'が 'image/.tempfile'を埋め込むのに対して、 'image'はそうしないことです。
// どちらも 'image/dir/.tempfile'を埋め込みません。
//
// パターンがプレフィックス 'all:' で始まる場合、ディレクトリを走査するルールが変更され、
// '.'または'_'で始まるファイルも含まれるようになります。
// たとえば、 'all:image' は 'image/.tempfile' と 'image/dir/.tempfile' の両方を埋め込みます。
//
// //go:embedディレクティブは、パッケージがデータを他のパッケージで利用可能にするかどうかに応じて、
// エクスポートされた変数と非エクスポートされた変数の両方で使用できます。
// ローカル変数ではなく、パッケージスコープの変数でのみ使用できます。
//
// パターンは、'.git/*'やシンボリックリンクなど、パッケージのモジュール外のファイルに一致してはなりません。
// パターンは、特殊な句読点文字 " * < > ? ` ' | / \ および : を含むファイル名に一致してはなりません。
// 空のディレクトリの一致は無視されます。その後、//go:embed行の各パターンは、少なくとも1つのファイルまたは空でないディレクトリに一致する必要があります。
//
// パターンが無効であるか、無効な一致がある場合、ビルドは失敗します。
//
// # 文字列とバイト
//
// stringまたは[]byte型の変数の//go:embed行には、単一のパターンのみが含まれることができます。
// また、そのパターンは単一のファイルにのみ一致することができます。
// stringまたは[]byteは、そのファイルの内容で初期化されます。
//
// stringまたは[]byteを使用する場合でも、//go:embedディレクティブを使用するには、
// "embed"をインポートする必要があります。
// [embed.FS] を参照しないソースファイルでは、空のインポート（import _ "embed"）を使用してください。
//
// # ファイルシステム
//
// 単一のファイルを埋め込む場合、string型または[]byte型の変数が最適です。
// [FS] 型は、静的Webサーバーコンテンツのディレクトリなど、ファイルツリーを埋め込むことができます。
//
// FSは、 [io/fs] パッケージの [FS] インターフェースを実装しているため、
// [net/http]、[text/template]、[html/template] を含むファイルシステムを理解するパッケージで使用できます。
//
// たとえば、上記の例のcontent変数がある場合、次のように書くことができます。
//
//	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.FS(content))))
//
//	template.ParseFS(content, "*.tmpl")
//
// # ツール
//
// Goパッケージを分析するツールをサポートするために、//go:embed行で見つかったパターンは、
// "go list"出力で利用可能です。 "go help list"出力のEmbedPatterns、TestEmbedPatterns、
// XTestEmbedPatternsフィールドを参照してください。
package embed

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/io/fs"
)

// FSは、通常//go:embedディレクティブで初期化された読み取り専用のファイルコレクションです。
// //go:embedディレクティブがない場合、FSは空のファイルシステムです。
//
// FSは読み取り専用の値であるため、複数のゴルーチンから同時に使用することが安全であり、
// FS型の値を互いに割り当てることも安全です。
//
// FSはfs.FSを実装しているため、ファイルシステムインターフェースを理解する
// net/http、text/template、html/templateを含むパッケージで使用できます。
//
// FSの初期化についての詳細については、パッケージのドキュメントを参照してください。
type FS struct {
	// The compiler knows the layout of this struct.
	// See cmd/compile/internal/staticdata's WriteEmbed.
	//
	// The files list is sorted by name but not by simple string comparison.
	// Instead, each file's name takes the form "dir/elem" or "dir/elem/".
	// The optional trailing slash indicates that the file is itself a directory.
	// The files list is sorted first by dir (if dir is missing, it is taken to be ".")
	// and then by base, so this list of files:
	//
	//	p
	//	q/
	//	q/r
	//	q/s/
	//	q/s/t
	//	q/s/u
	//	q/v
	//	w
	//
	// is actually sorted as:
	//
	//	p       # dir=.    elem=p
	//	q/      # dir=.    elem=q
	//	w/      # dir=.    elem=w
	//	q/r     # dir=q    elem=r
	//	q/s/    # dir=q    elem=s
	//	q/v     # dir=q    elem=v
	//	q/s/t   # dir=q/s  elem=t
	//	q/s/u   # dir=q/s  elem=u
	//
	// This order brings directory contents together in contiguous sections
	// of the list, allowing a directory read to use binary search to find
	// the relevant sequence of entries.
	files *[]file
}

var (
	_ fs.ReadDirFS  = FS{}
	_ fs.ReadFileFS = FS{}
)

var (
	_ fs.FileInfo = (*file)(nil)
	_ fs.DirEntry = (*file)(nil)
)

// Openは、指定されたファイルを読み取り用に開き、 [fs.File] として返します。
//
// ファイルがディレクトリでない場合、返されたファイルは [io.Seeker] と [io.ReaderAt] を実装します。
func (f FS) Open(name string) (fs.File, error)

// ReadDirは、指定されたディレクトリ全体を読み取り、返します。
func (f FS) ReadDir(name string) ([]fs.DirEntry, error)

// ReadFileは、指定されたファイルの内容を読み取り、返します。
func (f FS) ReadFile(name string) ([]byte, error)

var (
	_ io.Seeker   = (*openFile)(nil)
	_ io.ReaderAt = (*openFile)(nil)
)
