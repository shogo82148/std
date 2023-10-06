// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"github.com/shogo82148/std/log/slog"
	"github.com/shogo82148/std/os"
)

// A token is a secret value that grants permissions.
type Token string

// この例では、自身を置き換えるValueを使用して、秘密を明らかにしないようにする方法を示します。
func ExampleLogValuer_secret() {
	t := Token("shhhh!")
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime}))
	logger.Info("permission granted", "user", "Perry", "token", t)

	// Output:
	// level=INFO msg="permission granted" user=Perry token=REDACTED_TOKEN
}
