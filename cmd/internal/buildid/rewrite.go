// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package buildid

import (
	"github.com/shogo82148/std/io"
)

// FindAndHash reads all of r and returns the offsets of occurrences of id.
// While reading, findAndHash also computes and returns
// a hash of the content of r, but with occurrences of id replaced by zeros.
// FindAndHash reads bufSize bytes from r at a time.
// If bufSize == 0, FindAndHash uses a reasonable default.
func FindAndHash(r io.Reader, id string, bufSize int) (matches []int64, hash [32]byte, err error)

func Rewrite(w io.WriterAt, pos []int64, id string) error
