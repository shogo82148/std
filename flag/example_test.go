// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// These examples demonstrate more intricate uses of the flag package.
package flag_test

// Example 1: A single string flag called "species" with default value "gopher".

// Example 2: Two flags sharing a variable, so we can have a shorthand.
// The order of initialization is undefined, so make sure both use the
// same default value. They must be set up with an init function.

// Example 3: A user-defined flag type, a slice of durations.

func Example() {
	// All the interesting pieces are with the variables declared above, but
	// to enable the flag package to see the flags defined there, one must
	// execute, typically at the start of main (not init!):
	//	flag.Parse()
	// We don't run it here because this is not a main function and
	// the testing suite has already parsed the flags.
}
