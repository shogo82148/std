// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file tests types.Check by using it to
// typecheck the standard library and tests.

package types_test

import (
	. "go/types"
)

// The cmd/*/internal packages may have been deleted as part of a binary
// release. Import from source instead.
//
// (See https://golang.org/issue/43232 and
// https://github.com/golang/build/blob/df58bbac082bc87c4a3cdfe336d1ffe60bbaa916/cmd/release/release.go#L533-L545.)
//
// Use the same importer for all std lib tests to
// avoid repeated importing of the same packages.

// stdlibChecker implements concurrent type-checking of the packages defined by
// dirFiles, which must define a closed set of packages (such as GOROOT/src).

// A futurePackage is a future result of type-checking.

// Package paths of excluded packages.

// printPackageMu synchronizes the printing of type-checked package files in
// the typecheckFiles function.
//
// Without synchronization, package files may be interleaved during concurrent
// type-checking.
