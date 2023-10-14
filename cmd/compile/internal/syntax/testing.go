// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements testing support.

package syntax

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/regexp"
)

// CommentsDo parses the given source and calls the provided handler for each
// comment or error. If the text provided to handler starts with a '/' it is
// the comment text; otherwise it is the error message.
func CommentsDo(src io.Reader, handler func(line, col uint, text string))

// CommentMap collects all comments in the given src with comment text
// that matches the supplied regular expression rx and returns them as
// []Error lists in a map indexed by line number. The comment text is
// the comment with any comment markers ("//", "/*", or "*/") stripped.
// The position for each Error is the position of the token immediately
// preceding the comment and the Error message is the comment text,
// with all comments that are on the same line collected in a slice, in
// source order. If there is no preceding token (the matching comment
// appears at the beginning of the file), then the recorded position
// is unknown (line, col = 0, 0). If there are no matching comments,
// the result is nil.
func CommentMap(src io.Reader, rx *regexp.Regexp) (res map[uint][]Error)
