// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package html

// All entities that do not end with ';' are 6 or fewer bytes long.

// entity is a map from HTML entity names to their values. The semicolon matters:
// http://www.whatwg.org/specs/web-apps/current-work/multipage/named-character-references.html
// lists both "amp" and "amp;" as two separate entries.
//
// Note that the HTML5 list is larger than the HTML4 list at
// http://www.w3.org/TR/html4/sgml/entities.html

// HTML entities that are two unicode codepoints.
