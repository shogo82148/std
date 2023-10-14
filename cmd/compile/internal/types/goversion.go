// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

// AllowsGoVersion reports whether local package is allowed
// to use Go version major.minor.
func AllowsGoVersion(major, minor int) bool

// ParseLangFlag verifies that the -lang flag holds a valid value, and
// exits if not. It initializes data used by langSupported.
func ParseLangFlag()
