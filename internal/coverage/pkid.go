// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package coverage

// NotHardCoded is a package pseudo-ID indicating that a given package
// is not part of the runtime and doesn't require a hard-coded ID.
const NotHardCoded = -1

// HardCodedPkgID returns the hard-coded ID for the specified package
// path, or -1 if we don't use a hard-coded ID. Hard-coded IDs start
// at -2 and decrease as we go down the list.
func HardCodedPkgID(pkgpath string) int
