// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sort_test

import (
	"github.com/shogo82148/std/fmt"
)

// Example_sortMultiKeysは、構造体型を複数のフィールドセットで比較してソートする手法を示します。
// "Less"関数を連鎖させ、それぞれが1つのフィールドを比較します。
func Example_sortMultiKeys() {
	// Change構造体を順序付けるクロージャー。
	user := func(c1, c2 *Change) bool {
		return c1.user < c2.user
	}
	language := func(c1, c2 *Change) bool {
		return c1.language < c2.language
	}
	increasingLines := func(c1, c2 *Change) bool {
		return c1.lines < c2.lines
	}
	decreasingLines := func(c1, c2 *Change) bool {
		return c1.lines > c2.lines // 注意：> は下方向に並べ替えます。
	}

	// 簡単な使い方: ユーザーでソートする。
	OrderedBy(user).Sort(changes)
	fmt.Println("By user:", changes)

	// もっと多くの例。
	OrderedBy(user, increasingLines).Sort(changes)
	fmt.Println("By user,<lines:", changes)

	OrderedBy(user, decreasingLines).Sort(changes)
	fmt.Println("By user,>lines:", changes)

	OrderedBy(language, increasingLines).Sort(changes)
	fmt.Println("By language,<lines:", changes)

	OrderedBy(language, increasingLines, user).Sort(changes)
	fmt.Println("By language,<lines,user:", changes)

	// Output:
	// By user: [{dmr C 100} {glenda Go 200} {gri Go 100} {gri Smalltalk 80} {ken C 150} {ken Go 200} {r Go 100} {r C 150} {rsc Go 200}]
	// By user,<lines: [{dmr C 100} {glenda Go 200} {gri Smalltalk 80} {gri Go 100} {ken C 150} {ken Go 200} {r Go 100} {r C 150} {rsc Go 200}]
	// By user,>lines: [{dmr C 100} {glenda Go 200} {gri Go 100} {gri Smalltalk 80} {ken Go 200} {ken C 150} {r C 150} {r Go 100} {rsc Go 200}]
	// By language,<lines: [{dmr C 100} {ken C 150} {r C 150} {gri Go 100} {r Go 100} {glenda Go 200} {ken Go 200} {rsc Go 200} {gri Smalltalk 80}]
	// By language,<lines,user: [{dmr C 100} {ken C 150} {r C 150} {gri Go 100} {r Go 100} {glenda Go 200} {ken Go 200} {rsc Go 200} {gri Smalltalk 80}]

}
