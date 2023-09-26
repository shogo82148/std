// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/log"
	loginternal "log/internal"
)

// Default returns the default Logger.
func Default() *Logger

// SetDefault makes l the default Logger.
// After this call, output from the log package's default Logger
// (as with [log.Print], etc.) will be logged at LevelInfo using l's Handler.
func SetDefault(l *Logger)

// handlerWriter is an io.Writer that calls a Handler.
// It is used to link the default log.Logger to the default slog.Logger.

// A Logger records structured information about each call to its
// Log, Debug, Info, Warn, and Error methods.
// For each call, it creates a Record and passes it to a Handler.
//
// To create a new Logger, call [New] or a Logger method
// that begins "With".
type Logger struct {
	handler Handler
}

// Handler returns l's Handler.
func (l *Logger) Handler() Handler

// With returns a Logger that includes the given attributes
// in each output operation. Arguments are converted to
// attributes as if by [Logger.Log].
func (l *Logger) With(args ...any) *Logger

// WithGroup returns a Logger that starts a group, if name is non-empty.
// The keys of all attributes added to the Logger will be qualified by the given
// name. (How that qualification happens depends on the [Handler.WithGroup]
// method of the Logger's Handler.)
//
// If name is empty, WithGroup returns the receiver.
func (l *Logger) WithGroup(name string) *Logger

// New creates a new Logger with the given non-nil Handler.
func New(h Handler) *Logger

// With calls Logger.With on the default logger.
func With(args ...any) *Logger

// Enabled reports whether l emits log records at the given context and level.
func (l *Logger) Enabled(ctx context.Context, level Level) bool

// NewLogLogger returns a new log.Logger such that each call to its Output method
// dispatches a Record to the specified handler. The logger acts as a bridge from
// the older log API to newer structured logging handlers.
func NewLogLogger(h Handler, level Level) *log.Logger

// Log emits a log record with the current time and the given level and message.
// The Record's Attrs consist of the Logger's attributes followed by
// the Attrs specified by args.
//
// The attribute arguments are processed as follows:
//   - If an argument is an Attr, it is used as is.
//   - If an argument is a string and this is not the last argument,
//     the following argument is treated as the value and the two are combined
//     into an Attr.
//   - Otherwise, the argument is treated as a value with key "!BADKEY".
func (l *Logger) Log(ctx context.Context, level Level, msg string, args ...any)

// LogAttrs is a more efficient version of [Logger.Log] that accepts only Attrs.
func (l *Logger) LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)

// Debug logs at LevelDebug.
func (l *Logger) Debug(msg string, args ...any)

// DebugContext logs at LevelDebug with the given context.
func (l *Logger) DebugContext(ctx context.Context, msg string, args ...any)

// Info logs at LevelInfo.
func (l *Logger) Info(msg string, args ...any)

// InfoContext logs at LevelInfo with the given context.
func (l *Logger) InfoContext(ctx context.Context, msg string, args ...any)

// Warn logs at LevelWarn.
func (l *Logger) Warn(msg string, args ...any)

// WarnContext logs at LevelWarn with the given context.
func (l *Logger) WarnContext(ctx context.Context, msg string, args ...any)

// Error logs at LevelError.
func (l *Logger) Error(msg string, args ...any)

// ErrorContext logs at LevelError with the given context.
func (l *Logger) ErrorContext(ctx context.Context, msg string, args ...any)

// Debug calls Logger.Debug on the default logger.
func Debug(msg string, args ...any)

// DebugContext calls Logger.DebugContext on the default logger.
func DebugContext(ctx context.Context, msg string, args ...any)

// Info calls Logger.Info on the default logger.
func Info(msg string, args ...any)

// InfoContext calls Logger.InfoContext on the default logger.
func InfoContext(ctx context.Context, msg string, args ...any)

// Warn calls Logger.Warn on the default logger.
func Warn(msg string, args ...any)

// WarnContext calls Logger.WarnContext on the default logger.
func WarnContext(ctx context.Context, msg string, args ...any)

// Error calls Logger.Error on the default logger.
func Error(msg string, args ...any)

// ErrorContext calls Logger.ErrorContext on the default logger.
func ErrorContext(ctx context.Context, msg string, args ...any)

// Log calls Logger.Log on the default logger.
func Log(ctx context.Context, level Level, msg string, args ...any)

// LogAttrs calls Logger.LogAttrs on the default logger.
func LogAttrs(ctx context.Context, level Level, msg string, attrs ...Attr)
