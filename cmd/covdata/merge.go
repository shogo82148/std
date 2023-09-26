// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// mstate encapsulates state and provides methods for implementing the
// merge operation. This type implements the CovDataVisitor interface,
// and is designed to be used in concert with the CovDataReader
// utility, which abstracts away most of the grubby details of reading
// coverage data files. Most of the heavy lifting for merging is done
// using apis from 'metaMerge' (this is mainly a wrapper around that
// functionality).
