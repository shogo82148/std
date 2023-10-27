// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
)

<<<<<<< HEAD
// JSONHandlerは、レコードを行区切りのJSONオブジェクトとしてio.Writerに書き込むHandlerです。
=======
// JSONHandler is a [Handler] that writes Records to an [io.Writer] as
// line-delimited JSON objects.
>>>>>>> upstream/master
type JSONHandler struct {
	*commonHandler
}

<<<<<<< HEAD
// NewJSONHandlerは、指定されたオプションを使用して、
// wに書き込むJSONHandlerを作成します。
// optsがnilの場合、デフォルトのオプションが使用されます。
=======
// NewJSONHandler creates a [JSONHandler] that writes to w,
// using the given options.
// If opts is nil, the default options are used.
>>>>>>> upstream/master
func NewJSONHandler(w io.Writer, opts *HandlerOptions) *JSONHandler

// Enabledは、ハンドラが指定されたレベルのレコードを処理するかどうかを報告します。
// ハンドラは、レベルが低いレコードを無視します。
func (h *JSONHandler) Enabled(_ context.Context, level Level) bool

<<<<<<< HEAD
// WithAttrsは、hの属性に続く属性で構成される新しいJSONHandlerを返します。
=======
// WithAttrs returns a new [JSONHandler] whose attributes consists
// of h's attributes followed by attrs.
>>>>>>> upstream/master
func (h *JSONHandler) WithAttrs(attrs []Attr) Handler

func (h *JSONHandler) WithGroup(name string) Handler

<<<<<<< HEAD
// Handleは、引数のRecordをJSONオブジェクトとして1行にフォーマットします。
=======
// Handle formats its argument [Record] as a JSON object on a single line.
>>>>>>> upstream/master
//
// Recordの時間がゼロの場合、時間は省略されます。
// そうでない場合、キーは "time" であり、値はjson.Marshalと同様に出力されます。
//
// Recordのレベルがゼロの場合、レベルは省略されます。
// そうでない場合、キーは "level" であり、 [Level.String] の値が出力されます。
//
// AddSourceオプションが設定されており、ソース情報が利用可能な場合、
// キーは "source" であり、値は [Source] 型のレコードです。
//
// メッセージのキーは "msg" です。
//
// これらまたは他の属性を変更したり、出力から削除したりするには、
// [HandlerOptions.ReplaceAttr] を使用します。
//
// 値は、SetEscapeHTML(false)を使用して [encoding/json.Encoder] と同様にフォーマットされます。
// ただし、2つの例外があります。
//
<<<<<<< HEAD
// 1つ目は、Valueがerror型のAttrは、そのErrorメソッドを呼び出すことで文字列としてフォーマットされます。
// エラーは、構造体、スライス、マップなどの他のデータ構造に埋め込まれたエラーではなく、Attrにのみこの特別な処理が適用されます。
=======
// First, an Attr whose Value is of type error is formatted as a string, by
// calling its Error method. Only errors in Attrs receive this special treatment,
// not errors embedded in structs, slices, maps or other data structures that
// are processed by the [encoding/json] package.
>>>>>>> upstream/master
//
// 2つ目は、エンコードの失敗がHandleからエラーを返すことはありません。
// 代わりに、エラーメッセージが文字列としてフォーマットされます。
//
// Handleの各呼び出しは、io.Writer.Writeに対して1回のシリアル化された呼び出しを生成します。
func (h *JSONHandler) Handle(_ context.Context, r Record) error
