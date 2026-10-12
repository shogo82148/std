// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package version_test

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/go/version"
	"github.com/shogo82148/std/slices"
)

func ExampleCompare() {
	versions := []string{
		"go1.21.0",
		"go1.9",
		"go1.21",
		"go1.21rc1",
		"go1.10",
	}

	slices.SortFunc(versions, version.Compare)

	fmt.Println(versions)

	// Output:
	// [go1.9 go1.10 go1.21 go1.21rc1 go1.21.0]
}
