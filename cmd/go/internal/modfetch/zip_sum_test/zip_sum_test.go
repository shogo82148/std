// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package zip_sum_test tests that the module zip files produced by modfetch
// have consistent content sums. Ideally the zip files themselves are also
// stable over time, though this is not strictly necessary.
//
// This test loads a table from testdata/zip_sums.csv. The table has columns
// for module path, version, content sum, and zip file hash. The table
// includes a large number of real modules. The test downloads these modules
// in direct mode and verifies the zip files.
//
// This test is very slow, and it depends on outside modules that change
// frequently, so this is a manual test. To enable it, pass the -zipsum flag.
package zip_sum_test
