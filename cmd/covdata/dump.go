// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// dstate encapsulates state and provides methods for implementing
// various dump operations. Specifically, dstate implements the
// CovDataVisitor interface, and is designed to be used in
// concert with the CovDataReader utility, which abstracts away most
// of the grubby details of reading coverage data files.
