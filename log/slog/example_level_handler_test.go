// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"github.com/shogo82148/std/log/slog"
	"github.com/shogo82148/std/os"
)

// This example shows how to Use a LevelHandler to change the level of an
// existing Handler while preserving its other behavior.
//
// This example demonstrates increasing the log level to reduce a logger's
// output.
//
// Another typical use would be to decrease the log level (to LevelDebug, say)
// during a part of the program that was suspected of containing a bug.
func ExampleHandler_levelHandler() {
	th := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime})
	logger := slog.New(NewLevelHandler(slog.LevelWarn, th))
	logger.Info("not printed")
	logger.Warn("printed")

	// Output:
	// level=WARN msg=printed
}
