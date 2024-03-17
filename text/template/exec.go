// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

import (
	"github.com/shogo82148/std/io"
)

// ExecErrorは、Executeがテンプレートを評価する際にエラーが発生したときに返される
// カスタムエラータイプです。（書き込みエラーが発生した場合、実際のエラーが返されます。
// それはExecErrorタイプではありません。）
type ExecError struct {
	Name string
	Err  error
}

func (e ExecError) Error() string

func (e ExecError) Unwrap() error

// ExecuteTemplateは、指定された名前を持つtに関連付けられたテンプレートを
// 指定されたデータオブジェクトに適用し、出力をwrに書き込みます。
// テンプレートの実行中またはその出力の書き込み中にエラーが発生した場合、
// 実行は停止しますが、部分的な結果がすでに出力ライターに書き込まれている可能性があります。
// テンプレートは並行して安全に実行することができますが、並行実行がライターを共有する場合、
// 出力が交互になる可能性があります。
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error

// Executeは、解析されたテンプレートを指定されたデータオブジェクトに適用し、
// 出力をwrに書き込みます。
// テンプレートの実行中またはその出力の書き込み中にエラーが発生した場合、
// 実行は停止しますが、部分的な結果がすでに出力ライターに書き込まれている可能性があります。
// テンプレートは並行して安全に実行することができますが、並行実行がライターを共有する場合、
// 出力が交互になる可能性があります。
//
<<<<<<< HEAD
// dataがreflect.Valueの場合、テンプレートはreflect.Valueが保持する具体的な
// 値に適用されます。これはfmt.Printと同様です。
func (t *Template) Execute(wr io.Writer, data any) error

// DefinedTemplatesは、定義されたテンプレートのリストを文字列として返します。
// これは文字列 "; defined templates are: " で始まります。もし定義されたテンプレートがなければ、
// 空の文字列を返します。こことhtml/templateでエラーメッセージを生成するために使用されます。
=======
// If data is a [reflect.Value], the template applies to the concrete
// value that the reflect.Value holds, as in [fmt.Print].
func (t *Template) Execute(wr io.Writer, data any) error

// DefinedTemplates returns a string listing the defined templates,
// prefixed by the string "; defined templates are: ". If there are none,
// it returns the empty string. For generating an error message here
// and in [html/template].
>>>>>>> upstream/master
func (t *Template) DefinedTemplates() string

// IsTrueは、値がその型のゼロでない「真」であるか、
// および値が意味のある真理値を持っているかどうかを報告します。
// これはifやその他のそのようなアクションで使用される真理の定義です。
func IsTrue(val any) (truth, ok bool)
