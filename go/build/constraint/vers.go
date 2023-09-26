// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package constraint

// GoVersion returns the minimum Go version implied by a given build expression.
// If the expression can be satisfied without any Go version tags, GoVersion returns an empty string.
//
// For example:
//
//	GoVersion(linux && go1.22) = "go1.22"
//	GoVersion((linux && go1.22) || (windows && go1.20)) = "go1.20" => go1.20
//	GoVersion(linux) = ""
//	GoVersion(linux || (windows && go1.22)) = ""
//	GoVersion(!go1.22) = ""
//
// GoVersion assumes that any tag or negated tag may independently be true,
// so that its analysis can be purely structural, without SAT solving.
// “Impossible” subexpressions may therefore affect the result.
//
// For example:
//
//	GoVersion((linux && !linux && go1.20) || go1.21) = "go1.20"
func GoVersion(x Expr) string
