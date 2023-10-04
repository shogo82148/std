// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"log/slog"
	"os"
)

<<<<<<< HEAD
// A LevelHandler wraps a Handler with an Enabled method
// that returns false for levels below a minimum.
type LevelHandler struct {
	level   slog.Leveler
	handler slog.Handler
}

// この例では、LevelHandlerを使用して、既存のHandlerのレベルを変更しながら、
// その他の動作を維持する方法を示します。
=======
// This example shows how to Use a LevelHandler to change the level of an
// existing Handler while preserving its other behavior.
>>>>>>> upstream/release-branch.go1.21
//
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
