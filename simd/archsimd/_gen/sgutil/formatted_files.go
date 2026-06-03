// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sgutil

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/os"
)

func CreatePath(newFile string) (*os.File, error)

// FormatWriteAndClose formats the Go source code in source and writes it
// to newFile.  If there is a problem with the formatting, the
// entire source is numbered and emitted along with the error message,
// to help figure out where the source code generation went wrong.
func FormatWriteAndClose(source *bytes.Buffer, newFile string)

// WriteAndClose creates newFile, writes g to it,
// and closes the file.
func WriteAndClose(b []byte, newFile string)

// NumberLines takes a slice of bytes, and returns a string where each line
// is numbered, starting from 1.
func NumberLines(data []byte) string
