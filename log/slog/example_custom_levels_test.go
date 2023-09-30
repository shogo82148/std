// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"context"
	"log/slog"
	"os"
)

// This example demonstrates using custom log levels and custom log level names.
// In addition to the default log levels, it introduces Trace, Notice, and
// Emergency levels. The ReplaceAttr changes the way levels are printed for both
// the standard log levels and the custom log levels.
func ExampleHandlerOptions_customLevels() {
	// カスタムログパッケージからエクスポートされた定数。
	const (
		LevelTrace     = slog.Level(-8)
		LevelDebug     = slog.LevelDebug
		LevelInfo      = slog.LevelInfo
		LevelNotice    = slog.Level(2)
		LevelWarning   = slog.LevelWarn
		LevelError     = slog.LevelError
		LevelEmergency = slog.Level(12)
	)

	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		// すべてのログ出力を表示するために、カスタムレベルを設定します。
		// デフォルト値はLevelInfoであり、DebugとTraceログをドロップします。
		Level: LevelTrace,

		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// 予測可能なテスト出力のために、出力から時間を削除します。
			if a.Key == slog.TimeKey {
				return slog.Attr{}
			}

			// レベルキーの名前と出力文字列をカスタマイズします。
			// カスタムレベル値を含みます。
			if a.Key == slog.LevelKey {
				// レベルキーを "level" から "sev" に変更します。
				a.Key = "sev"

				// カスタムレベル値を処理します。
				level := a.Value.Any().(slog.Level)

				// これはマップや他の構造から名前を検索することもできますが、
				// この例ではswitch文を使用してレベルをリネームする方法を示しています。
				// 文字列値は定数であるべきですが、可読性のためにこの例では生の文字列を使用しています。
				switch {
				case level < LevelDebug:
					a.Value = slog.StringValue("TRACE")
				case level < LevelInfo:
					a.Value = slog.StringValue("DEBUG")
				case level < LevelNotice:
					a.Value = slog.StringValue("INFO")
				case level < LevelWarning:
					a.Value = slog.StringValue("NOTICE")
				case level < LevelError:
					a.Value = slog.StringValue("WARNING")
				case level < LevelEmergency:
					a.Value = slog.StringValue("ERROR")
				default:
					a.Value = slog.StringValue("EMERGENCY")
				}
			}

			return a
		},
	})

	logger := slog.New(th)
	ctx := context.Background()
	logger.Log(ctx, LevelEmergency, "missing pilots")
	logger.Error("failed to start engines", "err", "missing fuel")
	logger.Warn("falling back to default value")
	logger.Log(ctx, LevelNotice, "all systems are running")
	logger.Info("initiating launch")
	logger.Debug("starting background job")
	logger.Log(ctx, LevelTrace, "button clicked")

	// Output:
	// sev=EMERGENCY msg="missing pilots"
	// sev=ERROR msg="failed to start engines" err="missing fuel"
	// sev=WARNING msg="falling back to default value"
	// sev=NOTICE msg="all systems are running"
	// sev=INFO msg="initiating launch"
	// sev=DEBUG msg="starting background job"
	// sev=TRACE msg="button clicked"
}
