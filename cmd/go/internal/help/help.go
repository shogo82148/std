// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package help implements the “go help” command.
package help

import (
	"github.com/shogo82148/std/io"

	"github.com/shogo82148/std/cmd/go/internal/base"
)

// Help implements the 'help' command.
func Help(w io.Writer, args []string)

func PrintUsage(w io.Writer, cmd *base.Command)
