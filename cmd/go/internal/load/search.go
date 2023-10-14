// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package load

// MatchPackage(pattern, cwd)(p) reports whether package p matches pattern in the working directory cwd.
func MatchPackage(pattern, cwd string) func(*Package) bool
