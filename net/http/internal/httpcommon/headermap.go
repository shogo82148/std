// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package httpcommon

// LowerHeader returns the lowercase form of a header name,
// used on the wire for HTTP/2 and HTTP/3 requests.
func LowerHeader(v string) (lower string, ascii bool)

// CanonicalHeader canonicalizes a header name. (For example, "host" becomes "Host".)
func CanonicalHeader(v string) string

// CachedCanonicalHeader returns the canonical form of a well-known header name.
func CachedCanonicalHeader(v string) (string, bool)
