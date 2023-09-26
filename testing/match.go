// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

// matcher sanitizes, uniques, and filters names of subtests and subbenchmarks.

// simpleMatch matches a test name if all of the pattern strings match in
// sequence.

// alternationMatch matches a test name if one of the alternations match.

// TODO: fix test_main to avoid race and improve caching, also allowing to
// eliminate this Mutex.
