// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// downloadCache records the import paths we have already
// considered during the download, to avoid duplicate work when
// there is more than one dependency sequence leading to
// a particular package.

// downloadRootCache records the version control repository
// root directories we have already considered during the download.
// For example, all the packages in the github.com/google/codesearch repo
// share the same root (the directory for that path), and we only need
// to run the hg commands to consider each repository once.

// goTag matches go release tags such as go1 and go1.2.3.
// The numbers involved must be small (at most 4 digits),
// have no unnecessary leading zeros, and the version cannot
// end in .0 - it is go1, not go1.0 or go1.0.0.
