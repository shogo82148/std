// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import (
	"github.com/shogo82148/std/log/slog"
	"github.com/shogo82148/std/net/http"
	"github.com/shogo82148/std/os"
	"github.com/shogo82148/std/time"
)

func ExampleGroup() {
	r, _ := http.NewRequest("GET", "localhost", nil)
	// ...

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{ReplaceAttr: slogtest.RemoveTime}))
	logger.Info("finished",
		slog.Group("req",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String())),
		slog.Int("status", http.StatusOK),
		slog.Duration("duration", time.Second))

	// Output:
	// level=INFO msg=finished req.method=GET req.url=localhost status=200 duration=1s
}
