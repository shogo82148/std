// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/log"
)

// SetLogLoggerLevelは、[log] パッケージへのブリッジのレベルを制御します。
//
// [SetDefault] が呼び出される前は、slogのトップレベルのログ関数はデフォルトの [log.Logger] を呼び出します。
// そのモードでは、SetLogLoggerLevelはそれらの呼び出しの最小レベルを設定します。
// デフォルトでは、最小レベルはInfoなので、[Debug] への呼び出し
// （およびより低いレベルでのトップレベルのログ呼び出し）
// はlog.Loggerに渡されません。次の呼び出し後
//
//	slog.SetLogLoggerLevel(slog.LevelDebug)
//
// [Debug] への呼び出しはlog.Loggerに渡されます。
//
// [SetDefault] が呼び出された後、デフォルトの [log.Logger] への呼び出しは
// slogのデフォルトハンドラーに渡されます。そのモードでは、
// SetLogLoggerLevelはそれらの呼び出しがログに記録されるレベルを設定します。
// つまり、次の呼び出し後
//
//	slog.SetLogLoggerLevel(slog.LevelDebug)
//
// [log.Printf]への呼び出しは、[LevelDebug] レベルでの出力を結果とします。
//
// SetLogLoggerLevelは前の値を返します。
func SetLogLoggerLevel(level Level) (oldLevel Level)

// Defaultは、デフォルトの [Logger] を返します。
func Default() *Logger

// SetDefault makes l the default [Logger], which is used by
// the top-level functions [Info], [Debug] and so on.
// After this call, output from the log package's default Logger
// (as with [log.Print], etc.) will be logged using l's Handler,
// at a level controlled by [SetLogLoggerLevel].
func SetDefault(l *Logger)

// Loggerは、各Log、Debug、Info、Warn、Errorメソッドの呼び出しについて、
// 構造化された情報を記録します。
// 各呼び出しに対して、[Record] を作成し、[Handler] に渡します。
//
// 新しいLoggerを作成するには、[New]または"With"で始まるLoggerメソッドを呼び出します。
type Logger struct {
	handler Handler
}

// Handlerは、lのHandlerを返します。
func (l *Logger) Handler() Handler

// Withは、各出力操作で指定された属性を含むLoggerを返します。
// 引数は、[Logger.Log]によって属性に変換されます。
func (l *Logger) With(args ...any) *Logger

// WithGroupは、nameが空でない場合、グループを開始するLoggerを返します。
// 追加されたすべての属性のキーは、指定された名前で修飾されます。
// (修飾方法は、LoggerのHandlerの[Handler.WithGroup]メソッドに依存します。)
//
// nameが空の場合、WithGroupはレシーバーを返します。
func (l *Logger) WithGroup(name string) *Logger

// Newは、指定された非nil Handlerを持つ新しいLoggerを作成します。
func New(h Handler) *Logger

// Withは、デフォルトのロガーで [Logger.With] を呼び出します。
func With(args ...any) *Logger

// Enabledは、lが指定されたコンテキストとレベルでログレコードを生成するかどうかを報告します。
func (l *Logger) Enabled(ctx context.Context, level Level) bool

// NewLogLoggerは、指定されたハンドラにRecordをディスパッチするための新しい [log.Logger] を返します。
// ロガーは、古いログAPIから新しい構造化ログハンドラへのブリッジとして機能します。
func NewLogLogger(h Handler, level Level) *log.Logger

// Logは、現在の時刻と指定されたレベルとメッセージでログレコードを生成します。
// RecordのAttrsは、Loggerの属性に続くAttrsで構成されます。
//
// 属性引数は、次のように処理されます。
//   - 引数がAttrの場合、そのまま使用されます。
//   - 引数が文字列で、これが最後の引数でない場合、
//     次の引数が値として扱われ、2つがAttrに結合されます。
//   - それ以外の場合、引数はキー "!BADKEY" を持つ値として扱われます。
func (l *Logger) Log(ctx context.Context, level Level, msg string, args ...any)

// LogAttrsは、Attrのみを受け入れるより効率的な[Logger.Log]のバージョンです。
func (l *Logger) LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)

// Debugは、[LevelDebug] でログを記録します。
func (l *Logger) Debug(msg string, args ...any)

// DebugContextは、指定されたコンテキストで [LevelDebug] でログを記録します。
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any)

// Infoは、[LevelInfo] でログを記録します。
func (l *Logger) Info(msg string, args ...any)

// InfoContextは、指定されたコンテキストで [LevelInfo] でログを記録します。
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any)

// Warnは、[LevelWarn] でログを記録します。
func (l *Logger) Warn(msg string, args ...any)

// WarnContextは、指定されたコンテキストで [LevelWarn] でログを記録します。
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any)

// Errorは、[LevelError] でログを記録します。
func (l *Logger) Error(msg string, args ...any)

// ErrorContextは、指定されたコンテキストで [LevelError] でログを記録します。
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any)

// Debugは、デフォルトのロガーで [Logger.Debug] を呼び出します。
func Debug(msg string, args ...any)

// DebugContextは、デフォルトのロガーで [Logger.DebugContext] を呼び出します。
func DebugContext(ctx context.Context, msg string, args ...any)

// Infoは、デフォルトのロガーで [Logger.Info] を呼び出します。
func Info(msg string, args ...any)

// InfoContextは、デフォルトのロガーで [Logger.InfoContext] を呼び出します。
func InfoContext(ctx context.Context, msg string, args ...any)

// Warnは、デフォルトのロガーで [Logger.Warn] を呼び出します。
func Warn(msg string, args ...any)

// WarnContextは、デフォルトのロガーで [Logger.WarnContext] を呼び出します。
func WarnContext(ctx context.Context, msg string, args ...any)

// Errorは、デフォルトのロガーで [Logger.Error] を呼び出します。
func Error(msg string, args ...any)

// ErrorContextは、デフォルトのロガーで [Logger.ErrorContext] を呼び出します。
func ErrorContext(ctx context.Context, msg string, args ...any)

// Logは、デフォルトのロガーで [Logger.Log] を呼び出します。
func Log(ctx context.Context, level Level, msg string, args ...any)

// LogAttrsは、デフォルトのロガーで [Logger.LogAttrs] を呼び出します。
func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)
