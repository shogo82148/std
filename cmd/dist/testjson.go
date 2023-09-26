// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// lockedWriter serializes Write calls to an underlying Writer.

// testJSONFilter is an io.Writer filter that replaces the Package field in
// test2json output.
