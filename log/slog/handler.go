// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
)

// Handlerは、Loggerによって生成されたログレコードを処理します。
//
// 典型的なハンドラーは、ログレコードを標準エラーに出力したり、
// ファイルやデータベースに書き込んだり、
// または追加の属性を追加して別のハンドラーに渡すことができます。
//
// Handlerのメソッドのいずれかは、自身または他のメソッドと同時に呼び出される可能性があります。
// Handlerは、この並行性を管理する責任があります。
//
// slogパッケージのユーザーは、Handlerメソッドを直接呼び出すべきではありません。
// 代わりに [Logger] のメソッドを使用すべきです。
//
// 独自のハンドラーを実装する前に、https://go.dev/s/slog-handler-guide を参照してください。
type Handler interface {
	Enabled(context.Context, Level) bool

	Handle(context.Context, Record) error

	WithAttrs(attrs []Attr) Handler

	WithGroup(name string) Handler
}

// HandlerOptionsは、[TextHandler] または [JSONHandler] のオプションです。
// ゼロ値のHandlerOptionsは、完全にデフォルト値で構成されています。
type HandlerOptions struct {
	// AddSourceは、ハンドラーにログステートメントのソースコード位置を計算させ、
	// 出力にSourceKey属性を追加させます。
	AddSource bool

	// Levelは、ログに記録される最小レコードレベルを報告します。
	// ハンドラーは、より低いレベルのレコードを破棄します。
	// Levelがnilの場合、ハンドラーはLevelInfoを想定します。
	// ハンドラーは処理される各レコードに対してLevel.Levelを呼び出します。
	// 最小レベルを動的に調整するには、LevelVarを使用してください。
	Level Leveler

	// ReplaceAttrは、ログに記録される前に各非グループ属性を書き換えるために呼び出されます。
	// 属性の値は解決されています（[Value.Resolve]を参照）。
	// ReplaceAttrがゼロのAttrを返した場合、その属性は破棄されます。
	//
	// "time"、"level"、"source"、"msg"のキーを持つ組み込み属性は
	// この関数に渡されますが、timeがゼロの場合は省略され、
	// AddSourceがfalseの場合はsourceが省略されます。
	//
	// 最初の引数は、Attrを含む現在開いているグループのリストです。
	// これは保持または変更してはいけません。ReplaceAttrは
	// Group属性に対しては呼び出されず、その内容に対してのみ呼び出されます。例えば、属性
	// リスト
	//
	//     Int("a", 1), Group("g", Int("b", 2)), Int("c", 3)
	//
	// は、以下の引数でReplaceAttrを連続して呼び出します：
	//
	//     nil, Int("a", 1)
	//     []string{"g"}, Int("b", 2)
	//     nil, Int("c", 3)
	//
	// ReplaceAttrは、組み込み属性のデフォルトキーを変更したり、
	// 型を変換したり（例えば、`time.Time`をUnixエポックからの整数秒で置き換える）、
	// 個人情報をサニタイズしたり、出力から属性を削除するために使用できます。
	ReplaceAttr func(groups []string, a Attr) Attr
}

// "built-in"属性のキー。
const (
	// TimeKeyは、ログメソッドが呼び出されたときの時間を表すために、
	// 組み込みハンドラーによって使用されるキーです。
	// 関連する値は [time.Time] です。
	TimeKey = "time"
	// LevelKeyは、ログ呼び出しのレベルを表すために、
	// 組み込みハンドラーによって使用されるキーです。
	// 関連する値は [Level] です。
	LevelKey = "level"
	// MessageKeyは、ログ呼び出しのメッセージを表すために、
	// 組み込みハンドラーによって使用されるキーです。
	// 関連する値は文字列です。
	MessageKey = "msg"
	// SourceKey は、ログ呼び出しのソースファイルと行のためにビルトインハンドラーによって使用されるキーです。
	// 関連する値は *[Source] です。
	SourceKey = "source"
)

// DiscardHandlerはすべてのログ出力を破棄します。
// DiscardHandler.Enabledはすべてのレベルに対してfalseを返します。
var DiscardHandler Handler = discardHandler{}
