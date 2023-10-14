// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package bug implements the “go bug” command.
package bug

import (
	"github.com/shogo82148/std/cmd/go/internal/base"
)

var CmdBug = &base.Command{
	Run:       runBug,
	UsageLine: "go bug",
	Short:     "start a bug report",
	Long: `
Bug opens the default browser and starts a new bug report.
The report includes useful system information.
	`,
}
