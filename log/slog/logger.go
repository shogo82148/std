// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/log"
)

<<<<<<< HEAD
// Defaultは、デフォルトのLoggerを返します。
func Default() *Logger

// SetDefaultは、lをデフォルトのLoggerに設定します。
// この呼び出しの後、logパッケージのデフォルトLoggerからの出力（[log.Print]など）は、
// lのHandlerを使用してLevelInfoでログに記録されます。
func SetDefault(l *Logger)

// Loggerは、各Log、Debug、Info、Warn、Errorメソッドの呼び出しについて、
// 構造化された情報を記録します。
// 各呼び出しに対して、Recordを作成し、Handlerに渡します。
=======
// Default returns the default [Logger].
func Default() *Logger

// SetDefault makes l the default [Logger].
// After this call, output from the log package's default Logger
// (as with [log.Print], etc.) will be logged at [LevelInfo] using l's Handler.
func SetDefault(l *Logger)

// A Logger records structured information about each call to its
// Log, Debug, Info, Warn, and Error methods.
// For each call, it creates a [Record] and passes it to a [Handler].
>>>>>>> upstream/master
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

<<<<<<< HEAD
// Withは、デフォルトのロガーでLogger.Withを呼び出します。
=======
// With calls [Logger.With] on the default logger.
>>>>>>> upstream/master
func With(args ...any) *Logger

// Enabledは、lが指定されたコンテキストとレベルでログレコードを生成するかどうかを報告します。
func (l *Logger) Enabled(ctx context.Context, level Level) bool

<<<<<<< HEAD
// NewLogLoggerは、指定されたハンドラにRecordをディスパッチするための新しいlog.Loggerを返します。
// ロガーは、古いログAPIから新しい構造化ログハンドラへのブリッジとして機能します。
=======
// NewLogLogger returns a new [log.Logger] such that each call to its Output method
// dispatches a Record to the specified handler. The logger acts as a bridge from
// the older log API to newer structured logging handlers.
>>>>>>> upstream/master
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

<<<<<<< HEAD
// Debugは、LevelDebugでログを記録します。
func (l *Logger) Debug(msg string, args ...any)

// DebugContextは、指定されたコンテキストでLevelDebugでログを記録します。
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any)

// Infoは、LevelInfoでログを記録します。
func (l *Logger) Info(msg string, args ...any)

// InfoContextは、指定されたコンテキストでLevelInfoでログを記録します。
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any)

// Warnは、LevelWarnでログを記録します。
func (l *Logger) Warn(msg string, args ...any)

// WarnContextは、指定されたコンテキストでLevelWarnでログを記録します。
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any)

// Errorは、LevelErrorでログを記録します。
func (l *Logger) Error(msg string, args ...any)

// ErrorContextは、指定されたコンテキストでLevelErrorでログを記録します。
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any)

// Debugは、デフォルトのロガーでLogger.Debugを呼び出します。
func Debug(msg string, args ...any)

// DebugContextは、デフォルトのロガーでLogger.DebugContextを呼び出します。
func DebugContext(ctx context.Context, msg string, args ...any)

// Infoは、デフォルトのロガーでLogger.Infoを呼び出します。
func Info(msg string, args ...any)

// InfoContextは、デフォルトのロガーでLogger.InfoContextを呼び出します。
func InfoContext(ctx context.Context, msg string, args ...any)

// Warnは、デフォルトのロガーでLogger.Warnを呼び出します。
func Warn(msg string, args ...any)

// WarnContextは、デフォルトのロガーでLogger.WarnContextを呼び出します。
func WarnContext(ctx context.Context, msg string, args ...any)

// Errorは、デフォルトのロガーでLogger.Errorを呼び出します。
func Error(msg string, args ...any)

// ErrorContextは、デフォルトのロガーでLogger.ErrorContextを呼び出します。
func ErrorContext(ctx context.Context, msg string, args ...any)

// Logは、デフォルトのロガーでLogger.Logを呼び出します。
func Log(ctx context.Context, level Level, msg string, args ...any)

// LogAttrsは、デフォルトのロガーでLogger.LogAttrsを呼び出します。
=======
// Debug logs at [LevelDebug].
func (l *Logger) Debug(msg string, args ...any)

// DebugContext logs at [LevelDebug] with the given context.
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any)

// Info logs at [LevelInfo].
func (l *Logger) Info(msg string, args ...any)

// InfoContext logs at [LevelInfo] with the given context.
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any)

// Warn logs at [LevelWarn].
func (l *Logger) Warn(msg string, args ...any)

// WarnContext logs at [LevelWarn] with the given context.
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any)

// Error logs at [LevelError].
func (l *Logger) Error(msg string, args ...any)

// ErrorContext logs at [LevelError] with the given context.
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any)

// Debug calls [Logger.Debug] on the default logger.
func Debug(msg string, args ...any)

// DebugContext calls [Logger.DebugContext] on the default logger.
func DebugContext(ctx context.Context, msg string, args ...any)

// Info calls [Logger.Info] on the default logger.
func Info(msg string, args ...any)

// InfoContext calls [Logger.InfoContext] on the default logger.
func InfoContext(ctx context.Context, msg string, args ...any)

// Warn calls [Logger.Warn] on the default logger.
func Warn(msg string, args ...any)

// WarnContext calls [Logger.WarnContext] on the default logger.
func WarnContext(ctx context.Context, msg string, args ...any)

// Error calls [Logger.Error] on the default logger.
func Error(msg string, args ...any)

// ErrorContext calls [Logger.ErrorContext] on the default logger.
func ErrorContext(ctx context.Context, msg string, args ...any)

// Log calls [Logger.Log] on the default logger.
func Log(ctx context.Context, level Level, msg string, args ...any)

// LogAttrs calls [Logger.LogAttrs] on the default logger.
>>>>>>> upstream/master
func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)
