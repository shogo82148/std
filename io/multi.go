// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package io

// MultiReader returns a Reader that's the logical concatenation of
// the provided input readers.  They're read sequentially.  Once all
// inputs have returned EOF, Read will return EOF.  If any of the readers
// return a non-nil, non-EOF error, Read will return that error.
func MultiReader(readers ...Reader) Reader

var _ stringWriter = (*multiWriter)(nil)

// MultiWriter creates a writer that duplicates its writes to all the
// provided writers, similar to the Unix tee(1) command.
func MultiWriter(writers ...Writer) Writer
