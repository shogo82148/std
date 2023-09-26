// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package mime

// FormatMediaType serializes mediatype t and the parameters
// param as a media type conforming to RFC 2045 and RFC 2616.
// The type and parameter names are written in lower-case.
// When any of the arguments result in a standard violation then
// FormatMediaType returns the empty string.
func FormatMediaType(t string, param map[string]string) string

// ParseMediaType parses a media type value and any optional
// parameters, per RFC 1521.  Media types are the values in
// Content-Type and Content-Disposition headers (RFC 2183).
// On success, ParseMediaType returns the media type converted
// to lowercase and trimmed of white space and a non-nil map.
// The returned map, params, maps from the lowercase
// attribute to the attribute value with its case preserved.
func ParseMediaType(v string) (mediatype string, params map[string]string, err error)
