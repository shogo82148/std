// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build race
// +build race

// This program is used to verify the race detector
// by running the tests and parsing their output.
// It does not check stack correctness, completeness or anything else:
// it merely verifies that if a test is expected to be racy
// then the race is detected.
package race_test
