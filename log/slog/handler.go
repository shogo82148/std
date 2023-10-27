// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
)

type commonHandler struct{}

// A Handler handles log records produced by a Logger.
//
// 典型的なハンドラは、ログレコードを標準エラーに出力したり、
// ファイルやデータベースに書き込んだり、
// または追加の属性を追加して別のハンドラに渡すことができます。
//
// Handlerのメソッドのいずれかは、自身または他のメソッドと同時に呼び出される可能性があります。
// Handlerは、この並行性を管理する責任があります。
//
// slogパッケージのユーザーは、Handlerメソッドを直接呼び出すべきではありません。
// 代わりに、[Logger]のメソッドを使用する必要があります。
type Handler interface {
	Enabled(context.Context, Level) bool

	Handle(context.Context, Record) error

	WithAttrs(attrs []Attr) Handler

	WithGroup(name string) Handler
}

// HandlerOptionsは、[TextHandler] または [JSONHandler] のオプションです。
// ゼロ値のHandlerOptionsは、完全にデフォルト値で構成されています。
type HandlerOptions struct {
	// AddSource causes the handler to compute the source code position
	// of the log statement and add a SourceKey attribute to the output.
	AddSource bool

	// Level reports the minimum record level that will be logged.
	// The handler discards records with lower levels.
	// If Level is nil, the handler assumes LevelInfo.
	// The handler calls Level.Level for each record processed;
	// to adjust the minimum level dynamically, use a LevelVar.
	Level Leveler

	// ReplaceAttr is called to rewrite each non-group attribute before it is logged.
	// The attribute's value has been resolved (see [Value.Resolve]).
	// If ReplaceAttr returns a zero Attr, the attribute is discarded.
	//
	// The built-in attributes with keys "time", "level", "source", and "msg"
	// are passed to this function, except that time is omitted
	// if zero, and source is omitted if AddSource is false.
	//
	// The first argument is a list of currently open groups that contain the
	// Attr. It must not be retained or modified. ReplaceAttr is never called
	// for Group attributes, only their contents. For example, the attribute
	// list
	//
	//     Int("a", 1), Group("g", Int("b", 2)), Int("c", 3)
	//
	// results in consecutive calls to ReplaceAttr with the following arguments:
	//
	//     nil, Int("a", 1)
	//     []string{"g"}, Int("b", 2)
	//     nil, Int("c", 3)
	//
	// ReplaceAttr can be used to change the default keys of the built-in
	// attributes, convert types (for example, to replace a `time.Time` with the
	// integer seconds since the Unix epoch), sanitize personal information, or
	// remove attributes from the output.
	ReplaceAttr func(groups []string, a Attr) Attr
}

// "built-in"属性のキー。
const (
	// TimeKeyは、ログメソッドが呼び出されたときの時間を表すために、
	// 組み込みハンドラによって使用されるキーです。
	// 関連する値は[time.Time]です。
	TimeKey = "time"
	// LevelKeyは、ログ呼び出しのレベルを表すために、
	// 組み込みハンドラによって使用されるキーです。
	// 関連する値は[Level]です。
	LevelKey = "level"
	// MessageKeyは、ログ呼び出しのメッセージを表すために、
	// 組み込みハンドラによって使用されるキーです。
	// 関連する値は文字列です。
	MessageKey = "msg"
	// SourceKey は、ログ呼び出しのソースファイルと行のためにビルトインハンドラによって使用されるキーです。
	// 関連する値は *[Source] です。
	SourceKey = "source"
)
