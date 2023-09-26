// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// sstate holds state needed to implement subtraction and intersection
// operations on code coverage data files. This type provides methods
// to implement the CovDataVisitor interface, and is designed to be
// used in concert with the CovDataReader utility, which abstracts
// away most of the grubby details of reading coverage data files.
