// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort_test

import (
	"fmt"
	"sort"
)

func Example_sortWrapper() {
	s := []*Organ{
		{"brain", 1340},
		{"heart", 290},
		{"liver", 1494},
		{"pancreas", 131},
		{"prostate", 62},
		{"spleen", 162},
	}

	sort.Sort(ByWeight{s})
	fmt.Println("Organs by weight:")
	printOrgans(s)

	sort.Sort(ByName{s})
	fmt.Println("Organs by name:")
	printOrgans(s)

	// Output:
	// Organs by weight:
	// prostate (62g)
	// pancreas (131g)
	// spleen   (162g)
	// heart    (290g)
	// brain    (1340g)
	// liver    (1494g)
	// Organs by name:
	// brain    (1340g)
	// heart    (290g)
	// liver    (1494g)
	// pancreas (131g)
	// prostate (62g)
	// spleen   (162g)
}
