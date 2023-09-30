// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
)

type commonHandler struct{}

// Handlerは、Loggerによって生成されたログレコードを処理します。
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

// HandlerOptionsは、TextHandlerまたはJSONHandlerのオプションです。
// ゼロ値のHandlerOptionsは、完全にデフォルト値で構成されています。
type HandlerOptions struct {
	AddSource bool

	Level Leveler

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
	// SourceKeyは、ログ呼び出しのソースファイルと行を表すために、
	// 組み込みハンドラによって使用されるキーです。
	// 関連する値は文字列です。
	SourceKey = "source"
)

// handleState holds state for a single call to commonHandler.handle.
// The initial value of sep determines whether to emit a separator
// before the next key, after which it stays true.

// Separator for group names and keys.
