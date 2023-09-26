// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

// transitionFunc is the array of context transition functions for text nodes.
// A transition function takes a context and template text input, and returns
// the updated context and the number of bytes consumed from the front of the
// input.

// specialTagEndMarkers maps element types to the character sequence that
// case-insensitively signals the end of the special tag body.
