// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package load

import (
	"github.com/shogo82148/std/cmd/go/internal/modload"
)

// MatchPackage(pattern, cwd)(p) reports whether package p matches pattern in the working directory cwd.
func MatchPackage(pattern, cwd string) func(*modload.State, *Package) bool
