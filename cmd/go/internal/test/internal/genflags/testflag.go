// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package genflags

// ShortTestFlags returns the set of "-test." flag shorthand names that end
// users may pass to 'go test'.
func ShortTestFlags() []string
