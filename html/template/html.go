// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package template

// htmlReplacementTable contains the runes that need to be escaped
// inside a quoted attribute value or in a text node.

// htmlNormReplacementTable is like htmlReplacementTable but without '&' to
// avoid over-encoding existing entities.

// htmlNospaceReplacementTable contains the runes that need to be escaped
// inside an unquoted attribute value.
// The set of runes escaped is the union of the HTML specials and
// those determined by running the JS below in browsers:
// <div id=d></div>
// <script>(function () {
// var a = [], d = document.getElementById("d"), i, c, s;
// for (i = 0; i < 0x10000; ++i) {
//   c = String.fromCharCode(i);
//   d.innerHTML = "<span title=" + c + "lt" + c + "></span>"
//   s = d.getElementsByTagName("SPAN")[0];
//   if (!s || s.title !== c + "lt" + c) { a.push(i.toString(16)); }
// }
// document.write(a.join(", "));
// })()</script>

// htmlNospaceNormReplacementTable is like htmlNospaceReplacementTable but
// without '&' to avoid over-encoding existing entities.
