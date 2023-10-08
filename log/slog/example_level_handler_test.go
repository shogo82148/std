// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"github.com/shogo82148/std/log/slog"
	"github.com/shogo82148/std/os"
)

// A LevelHandler wraps a Handler with an Enabled method
// that returns false for levels below a minimum.
type LevelHandler struct {
	level   slog.Leveler
	handler slog.Handler
}

// この例では、ログレベルを上げて、ロガーの出力を減らす方法を示しています。
//
// 別の一般的な使用方法は、（例えばLevelDebugに）ログレベルを下げて、
// バグが含まれていると疑われるプログラムの一部分でログを出力することです。
func ExampleHandler_levelHandler() {
	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime})
	logger := slog.New(NewLevelHandler(slog.LevelWarn, th))
	logger.Info("not printed")
	logger.Warn("printed")

	// Output:
	// level=WARN msg=printed
}
