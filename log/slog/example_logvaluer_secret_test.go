// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"log/slog"
	"os"
)

<<<<<<< HEAD
// A token is a secret value that grants permissions.
type Token string

// この例では、自身を置き換えるValueを使用して、秘密を明らかにしないようにする方法を示します。
=======
// This example demonstrates a Value that replaces itself
// with an alternative representation to avoid revealing secrets.
>>>>>>> upstream/release-branch.go1.21
func ExampleLogValuer_secret() {
	t := Token("shhhh!")
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime}))
	logger.Info("permission granted", "user", "Perry", "token", t)

	// Output:
	// level=INFO msg="permission granted" user=Perry token=REDACTED_TOKEN
}
