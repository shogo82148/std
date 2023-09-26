// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package coverage

// emitState holds useful state information during the emit process.
//
// When an instrumented program finishes execution and starts the
// process of writing out coverage data, it's possible that an
// existing meta-data file already exists in the output directory. In
// this case openOutputFiles() below will leave the 'mf' field below
// as nil. If a new meta-data file is needed, field 'mfname' will be
// the final desired path of the meta file, 'mftmp' will be a
// temporary file, and 'mf' will be an open os.File pointer for
// 'mftmp'. The meta-data file payload will be written to 'mf', the
// temp file will be then closed and renamed (from 'mftmp' to
// 'mfname'), so as to insure that the meta-data file is created
// atomically; we want this so that things work smoothly in cases
// where there are several instances of a given instrumented program
// all terminating at the same time and trying to create meta-data
// files simultaneously.
//
// For counter data files there is less chance of a collision, hence
// the openOutputFiles() stores the counter data file in 'cfname' and
// then places the *io.File into 'cf'.

// fileType is used to select between counter-data files and
// meta-data files.
