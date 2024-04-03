// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Helper functions to make constructing templates easier.

package template

import (
	"github.com/shogo82148/std/io/fs"
)

<<<<<<< HEAD
// Mustは、(*Template, error)を返す関数への呼び出しをラップし、
// エラーが非nilの場合にパニックを起こすヘルパーです。これは変数の
// 初期化などで使用することを意図しています。
=======
// Must is a helper that wraps a call to a function returning ([*Template], error)
// and panics if the error is non-nil. It is intended for use in variable
// initializations such as
>>>>>>> upstream/master
//
//	var t = template.Must(template.New("name").Parse("text"))
func Must(t *Template, err error) *Template

// ParseFilesは新しい [Template] を作成し、指定されたファイルからテンプレート定義を解析します。
// 返されるテンプレートの名前は、最初のファイルのベース名と解析された内容を持つことになります。
// ファイルは少なくとも1つ必要です。
// エラーが発生した場合、解析は停止し、返される*Templateはnilになります。
//
// 異なるディレクトリにある同名の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
// 例えば、ParseFiles("a/foo", "b/foo")は "b/foo" を "foo" という名前のテンプレートとして保存し、
// "a/foo" は利用できません。
func ParseFiles(filenames ...string) (*Template, error)

<<<<<<< HEAD
// ParseFilesは指定されたファイルを解析し、結果となるテンプレートを
// tと関連付けます。エラーが発生した場合、解析は停止し、返されるテンプレートはnilになります。
// それ以外の場合はtです。ファイルは少なくとも1つ必要です。
// ParseFilesによって作成されたテンプレートは、引数のファイルのベース名によって名付けられるため、
// tは通常、ファイルの（ベース）名の1つを持つべきです。そうでない場合、tの
// 内容によっては、t.Executeが失敗する可能性があります。その場合は、
// t.ExecuteTemplateを使用して有効なテンプレートを実行します。
=======
// ParseFiles parses the named files and associates the resulting templates with
// t. If an error occurs, parsing stops and the returned template is nil;
// otherwise it is t. There must be at least one file.
// Since the templates created by ParseFiles are named by the base
// (see [filepath.Base]) names of the argument files, t should usually have the
// name of one of the (base) names of the files. If it does not, depending on
// t's contents before calling ParseFiles, t.Execute may fail. In that
// case use t.ExecuteTemplate to execute a valid template.
>>>>>>> upstream/master
//
// 異なるディレクトリにある同名の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
func (t *Template) ParseFiles(filenames ...string) (*Template, error)

<<<<<<< HEAD
// ParseGlobは新しい [Template] を作成し、パターンによって識別された
// ファイルからテンプレート定義を解析します。ファイルはfilepath.Matchの
// セマンティクスに従ってマッチし、パターンは少なくとも1つのファイルと
// マッチしなければなりません。返されるテンプレートは、パターンによって
// マッチした最初のファイルの（ベース）名と（解析された）内容を持つことになります。
// ParseGlobは、パターンによってマッチしたファイルのリストで [ParseFiles] を
// 呼び出すのと同等です。
=======
// ParseGlob creates a new [Template] and parses the template definitions from
// the files identified by the pattern. The files are matched according to the
// semantics of [filepath.Match], and the pattern must match at least one file.
// The returned template will have the [filepath.Base] name and (parsed)
// contents of the first file matched by the pattern. ParseGlob is equivalent to
// calling [ParseFiles] with the list of files matched by the pattern.
>>>>>>> upstream/master
//
// 異なるディレクトリにある同名の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
func ParseGlob(pattern string) (*Template, error)

<<<<<<< HEAD
// ParseGlobは、パターンによって識別されたファイルのテンプレート定義を解析し、
// 結果となるテンプレートをtと関連付けます。ファイルはfilepath.Matchの
// セマンティクスに従ってマッチし、パターンは少なくとも1つのファイルと
// マッチしなければなりません。ParseGlobは、パターンによってマッチした
// ファイルのリストでt.ParseFilesを呼び出すのと同等です。
=======
// ParseGlob parses the template definitions in the files identified by the
// pattern and associates the resulting templates with t. The files are matched
// according to the semantics of [filepath.Match], and the pattern must match at
// least one file. ParseGlob is equivalent to calling [Template.ParseFiles] with
// the list of files matched by the pattern.
>>>>>>> upstream/master
//
// 異なるディレクトリにある同名の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
func (t *Template) ParseGlob(pattern string) (*Template, error)

<<<<<<< HEAD
// ParseFSは、[Template.ParseFiles] や [Template.ParseGlob] と似ていますが、ホストオペレーティングシステムの
// ファイルシステムの代わりにファイルシステムfsysから読み取ります。
// それはグロブパターンのリストを受け入れます。
// （ほとんどのファイル名は、自分自身のみにマッチするグロブパターンとして機能します。）
func ParseFS(fsys fs.FS, patterns ...string) (*Template, error)

// ParseFSは、[Template.ParseFiles] や [Template.ParseGlob] と似ていますが、ホストオペレーティングシステムの
// ファイルシステムの代わりにファイルシステムfsysから読み取ります。
// それはグロブパターンのリストを受け入れます。
// （ほとんどのファイル名は、自分自身のみにマッチするグロブパターンとして機能します。）
=======
// ParseFS is like [Template.ParseFiles] or [Template.ParseGlob] but reads from the file system fsys
// instead of the host operating system's file system.
// It accepts a list of glob patterns (see [path.Match]).
// (Note that most file names serve as glob patterns matching only themselves.)
func ParseFS(fsys fs.FS, patterns ...string) (*Template, error)

// ParseFS is like [Template.ParseFiles] or [Template.ParseGlob] but reads from the file system fsys
// instead of the host operating system's file system.
// It accepts a list of glob patterns (see [path.Match]).
// (Note that most file names serve as glob patterns matching only themselves.)
>>>>>>> upstream/master
func (t *Template) ParseFS(fsys fs.FS, patterns ...string) (*Template, error)
