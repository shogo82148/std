// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"log/slog"
	"os"
)

// A token is a secret value that grants permissions.
type Token string

// This example demonstrates a Value that replaces itself
// with an alternative representation to avoid revealing secrets.
func ExampleLogValuer_secret() {
	t := Token("shhhh!")
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime}))
	logger.Info("permission granted", "user", "Perry", "token", t)

	// Output:
	// level=INFO msg="permission granted" user=Perry token=REDACTED_TOKEN
}
