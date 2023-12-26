// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Helper functions to make constructing templates easier.

package template

import (
	"github.com/shogo82148/std/io/fs"
)

// Mustは、(*Template, error)を返す関数への呼び出しをラップし、
// エラーが非nilの場合にパニックを起こすヘルパーです。これは変数の
// 初期化などで使用することを意図しています。
//
//	var t = template.Must(template.New("name").Parse("text"))
func Must(t *Template, err error) *Template

// ParseFilesは新しいテンプレートを作成し、指定されたファイルからテンプレート定義を解析します。
// 返されるテンプレートの名前は、最初のファイルのベース名と解析された内容を持つことになります。
// ファイルは少なくとも1つ必要です。
// エラーが発生した場合、解析は停止し、返される*Templateはnilになります。
//
// 異なるディレクトリにある同名の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
// 例えば、ParseFiles("a/foo", "b/foo")は "b/foo" を "foo" という名前のテンプレートとして保存し、
// "a/foo" は利用できません。
func ParseFiles(filenames ...string) (*Template, error)

// ParseFilesは指定されたファイルを解析し、結果となるテンプレートを
// tと関連付けます。エラーが発生した場合、解析は停止し、返されるテンプレートはnilになります。
// それ以外の場合はtです。ファイルは少なくとも1つ必要です。
// ParseFilesによって作成されたテンプレートは、引数のファイルのベース名によって名付けられるため、
// tは通常、ファイルの（ベース）名の1つを持つべきです。そうでない場合、tの
// 内容によっては、t.Executeが失敗する可能性があります。その場合は、
// t.ExecuteTemplateを使用して有効なテンプレートを実行します。
//
// 異なるディレクトリにある同名の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
func (t *Template) ParseFiles(filenames ...string) (*Template, error)

// ParseGlobは新しいテンプレートを作成し、パターンによって識別された
// ファイルからテンプレート定義を解析します。ファイルはfilepath.Matchの
// セマンティクスに従ってマッチし、パターンは少なくとも1つのファイルと
// マッチしなければなりません。返されるテンプレートは、パターンによって
// マッチした最初のファイルの（ベース）名と（解析された）内容を持つことになります。
// ParseGlobは、パターンによってマッチしたファイルのリストでParseFilesを
// 呼び出すのと同等です。
//
// 異なるディレクトリにある同名の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
func ParseGlob(pattern string) (*Template, error)

// ParseGlobは、パターンによって識別されたファイルのテンプレート定義を解析し、
// 結果となるテンプレートをtと関連付けます。ファイルはfilepath.Matchの
// セマンティクスに従ってマッチし、パターンは少なくとも1つのファイルと
// マッチしなければなりません。ParseGlobは、パターンによってマッチした
// ファイルのリストでt.ParseFilesを呼び出すのと同等です。
//
// 異なるディレクトリにある同名の複数のファイルを解析するとき、
// 最後に指定されたものが結果となります。
func (t *Template) ParseGlob(pattern string) (*Template, error)

// ParseFSは、ParseFilesやParseGlobと似ていますが、ホストオペレーティングシステムの
// ファイルシステムの代わりにファイルシステムfsysから読み取ります。
// それはグロブパターンのリストを受け入れます。
// （ほとんどのファイル名は、自分自身のみにマッチするグロブパターンとして機能します。）
func ParseFS(fsys fs.FS, patterns ...string) (*Template, error)

// ParseFSは、ParseFilesやParseGlobと似ていますが、ホストオペレーティングシステムの
// ファイルシステムの代わりにファイルシステムfsysから読み取ります。
// それはグロブパターンのリストを受け入れます。
// （ほとんどのファイル名は、自分自身のみにマッチするグロブパターンとして機能します。）
func (t *Template) ParseFS(fsys fs.FS, patterns ...string) (*Template, error)
