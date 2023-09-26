// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file tests types.Check by using it to
// typecheck the standard library and tests.

package types_test

import (
	. "go/types"
)

// Use the same importer for all std lib tests to
// avoid repeated importing of the same packages.

// Package paths of excluded packages.
