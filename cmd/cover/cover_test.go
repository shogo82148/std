// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main_test

import (
	cmdcover "cmd/cover"
)

// testTempDir is a temporary directory created in TestMain.

// If set, this will preserve all the tmpdir files from the test run.

// lineDupContents becomes linedup.go in testFuncWithDuplicateLines.

// lineDupTestContents becomes linedup_test.go in testFuncWithDuplicateLines.
