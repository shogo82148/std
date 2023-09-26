// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// The file define a quoted-printable decoder, as specified in RFC 2045.
// Deviations:
// 1. in addition to "=\r\n", "=\n" is also treated as soft line break.
// 2. it will pass through a '\r' or '\n' not preceded by '=', consistent
//    with other broken QP encoders & decoders.

package multipart
