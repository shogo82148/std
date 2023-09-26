// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// metaMerge provides state and methods to help manage the process
// of selecting or merging meta data files. There are three cases
// of interest here: the "-pcombine" flag provided by merge, the
// "-pkg" option provided by all merge/subtract/intersect, and
// a regular vanilla merge with no package selection
//
// In the -pcombine case, we're essentially glomming together all the
// meta-data for all packages and all functions, meaning that
// everything we see in a given package needs to be added into the
// meta-data file builder; we emit a single meta-data file at the end
// of the run.
//
// In the -pkg case, we will typically emit a single meta-data file
// per input pod, where that new meta-data file contains entries for
// just the selected packages.
//
// In the third case (vanilla merge with no combining or package
// selection) we can carry over meta-data files without touching them
// at all (only counter data files will be merged).

// pkstate

// pcombinestate
