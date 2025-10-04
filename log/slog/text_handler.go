// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/io"
)

// TextHandlerは、[io.Writer] にkey=valueペアのシーケンスと改行を続けて書き込むHandlerです。
type TextHandler struct {
	*commonHandler
}

// NewTextHandlerは、指定されたオプションを使用して、wに書き込む [TextHandler] を作成します。
// optsがnilの場合、デフォルトのオプションが使用されます。
func NewTextHandler(w io.Writer, opts *HandlerOptions) *TextHandler

// Enabledは、ハンドラが指定されたレベルのレコードを処理するかどうかを報告します。
// ハンドラは、レベルが低いレコードを無視します。
func (h *TextHandler) Enabled(_ context.Context, level Level) bool

// WithAttrsは、hの属性に続くattrsで構成される新しい [TextHandler] を返します。
func (h *TextHandler) WithAttrs(attrs []Attr) Handler

func (h *TextHandler) WithGroup(name string) Handler

// Handleは、引数 [Record] をスペースで区切られたkey=valueの1行としてフォーマットします。
//
// Recordのtimeがゼロの場合、timeは省略されます。
// そうでない場合、keyは"time"であり、RFC3339形式でミリ秒精度で出力されます。
//
<<<<<<< HEAD
// Recordのlevelがゼロの場合、levelは省略されます。
// そうでない場合、keyは"level"であり、 [Level.String] の値が出力されます。
=======
// The level's key is "level" and its value is the result of calling [Level.String].
>>>>>>> upstream/release-branch.go1.25
//
// AddSourceオプションが設定されており、ソース情報が利用可能な場合、
// keyは"source"であり、値はFILE:LINEとして出力されます。
//
// メッセージのkeyは"msg"です。
//
// これらまたは他の属性を変更したり、出力から削除するには、
// [HandlerOptions.ReplaceAttr] を使用します。
//
// 値が [encoding.TextMarshaler] を実装している場合、MarshalTextの結果が書き込まれます。
// そうでない場合、[fmt.Sprint] の結果が書き込まれます。
//
// キーと値は、Unicodeスペース文字、非表示文字、'"'、'='を含む場合、 [strconv.Quote] で引用符で囲まれます。
//
// グループ内のキーは、ドットで区切られたコンポーネント（キーまたはグループ名）で構成されます。
// さらにエスケープは行われません。
// したがって、キー"a.b.c"から、2つのグループ"a"と"b"とキー"c"があるか、
// 単一のグループ"a.b"とキー"c"があるか、単一のグループ"a"とキー"b.c"があるかを判断する方法はありません。
// コンポーネント内にドットがある場合でも、キーのグループ構造を再構築する必要がある場合は、
// [HandlerOptions.ReplaceAttr] を使用して、その情報をキーにエンコードします。
//
// Handleの各呼び出しは、io.Writer.Writeへの単一のシリアル化された呼び出しの結果を返します。
func (h *TextHandler) Handle(_ context.Context, r Record) error
