// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package http

// A mapping is a collection of key-value pairs where the keys are unique.
// A zero mapping is empty and ready to use.
// A mapping tries to pick a representation that makes [mapping.find] most efficient.

// maxSlice is the maximum number of pairs for which a slice is used.
// It is a variable for benchmarking.
