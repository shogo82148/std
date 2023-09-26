// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package slog_test

import "log/slog"

type Name struct {
	First, Last string
}

func ExampleLogValuer_group() {
	n := Name{"Perry", "Platypus"}
	slog.Info("mission accomplished", "agent", n)

	// JSON Output would look in part like:
	// {
	//     ...
	//     "msg": "mission accomplished",
	//     "agent": {
	//         "first": "Perry",
	//         "last": "Platypus"
	//     }
	// }
}
