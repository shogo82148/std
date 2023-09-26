// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testing

// matcher sanitizes, uniques, and filters names of subtests and subbenchmarks.

// TODO: fix test_main to avoid race and improve caching, also allowing to
// eliminate this Mutex.
