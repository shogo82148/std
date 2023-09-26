// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

// regexPrecederKeywords is a set of reserved JS keywords that can precede a
// regular expression in JS source.

// jsStrNormReplacementTable is like jsStrReplacementTable but does not
// overencode existing escapes since this table has no entry for `\`.
