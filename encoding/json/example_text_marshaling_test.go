// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json_test

import (
	"github.com/shogo82148/std/encoding/json"
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/log"
)

const (
	Unrecognized Size = iota
	Small
	Large
)

func Example_textMarshalJSON() {
	blob := `["small","regular","large","unrecognized","small","normal","small","large"]`
	var inventory []Size
	if err := json.Unmarshal([]byte(blob), &inventory); err != nil {
		log.Fatal(err)
	}

	counts := make(map[Size]int)
	for _, size := range inventory {
		counts[size] += 1
	}

	fmt.Printf("Inventory Counts:\n* Small:        %d\n* Large:        %d\n* Unrecognized: %d\n",
		counts[Small], counts[Large], counts[Unrecognized])

	// Output:
	// Inventory Counts:
	// * Small:        3
	// * Large:        2
	// * Unrecognized: 3
}
