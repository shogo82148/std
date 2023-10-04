// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// These examples demonstrate more intricate uses of the flag package.
package flag_test

func Example() {
	// All the interesting pieces are with the variables declared above, but
	// to enable the flag package to see the flags defined there, one must
	// execute, typically at the start of main (not init!):
	//	flag.Parse()
	// We don't call it here because this code is a function called "Example"
	// that is part of the testing suite for the package, which has already
	// parsed the flags. When viewed at pkg.go.dev, however, the function is
	// renamed to "main" and it could be run as a standalone example.
}
