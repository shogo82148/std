// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
)

// JSONHandlerは、レコードを行区切りのJSONオブジェクトとして [io.Writer] に書き込む [Handler] です。
type JSONHandler struct {
	*commonHandler
}

// NewJSONHandlerは、指定されたオプションを使用して、
// wに書き込む [JSONHandler] を作成します。
// optsがnilの場合、デフォルトのオプションが使用されます。
func NewJSONHandler(w io.Writer, opts *HandlerOptions) *JSONHandler

// Enabledは、ハンドラが指定されたレベルのレコードを処理するかどうかを報告します。
// ハンドラは、レベルが低いレコードを無視します。
func (h *JSONHandler) Enabled(_ context.Context, level Level) bool

// WithAttrsは、hの属性に続く属性で構成される新しい [JSONHandler] を返します。
func (h *JSONHandler) WithAttrs(attrs []Attr) Handler

func (h *JSONHandler) WithGroup(name string) Handler

// Handleは、引数の [Record] をJSONオブジェクトとして1行にフォーマットします。
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
// 最初に、値がエラー型であるAttrは、そのErrorメソッドを呼び出すことによって文字列としてフォーマットされます。
// この特別な処理は、[encoding/json] パッケージによって処理される構造体、スライス、マップ、その他のデータ構造に埋め込まれたエラーではなく、
// Attrs内のエラーのみが受け取ります。
//
// 2つ目は、エンコードの失敗がHandleからエラーを返すことはありません。
// 代わりに、エラーメッセージが文字列としてフォーマットされます。
//
// Handleの各呼び出しは、io.Writer.Writeに対して1回のシリアル化された呼び出しを生成します。
func (h *JSONHandler) Handle(_ context.Context, r Record) error
