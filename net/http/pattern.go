// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Patterns for ServeMux routing.

package http

// A pattern is something that can be matched against an HTTP request.
// It has an optional method, an optional host, and a path.

// A segment is a pattern piece that matches one or more path segments, or
// a trailing slash.
//
// If wild is false, it matches a literal segment, or, if s == "/", a trailing slash.
// Examples:
//
//	"a" => segment{s: "a"}
//	"/{$}" => segment{s: "/"}
//
// If wild is true and multi is false, it matches a single path segment.
// Example:
//
//	"{x}" => segment{s: "x", wild: true}
//
// If both wild and multi are true, it matches all remaining path segments.
// Example:
//
//	"{rest...}" => segment{s: "rest", wild: true, multi: true}

// relationship is a relationship between two patterns, p1 and p2.
