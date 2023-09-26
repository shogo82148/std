// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bzip2

// moveToFrontDecoder implements a move-to-front list. Such a list is an
// efficient way to transform a string with repeating elements into one with
// many small valued numbers, which is suitable for entropy encoding. It works
// by starting with an initial list of symbols and references symbols by their
// index into that list. When a symbol is referenced, it's moved to the front
// of the list. Thus, a repeated symbol ends up being encoded with many zeros,
// as the symbol will be at the front of the list after the first access.
