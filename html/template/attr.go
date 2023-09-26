// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

// attrTypeMap[n] describes the value of the given attribute.
// If an attribute affects (or can mask) the encoding or interpretation of
// other content, or affects the contents, idempotency, or credentials of a
// network message, then the value in this map is contentTypeUnsafe.
// This map is derived from HTML5, specifically
// http://www.w3.org/TR/html5/Overview.html#attributes-1
// as well as "%URI"-typed attributes from
// http://www.w3.org/TR/html4/index/attributes.html
