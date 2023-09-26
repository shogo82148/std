// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package reflect_test

import (
	. "reflect"
)

type SFG struct {
	F int
	G int
}

type SFG1 struct {
	SFG
}

type SFG2 struct {
	SFG1
}

type SFGH struct {
	F int
	G int
	H int
}

type SFGH1 struct {
	SFGH
}

type SFGH2 struct {
	SFGH1
}

type SFGH3 struct {
	SFGH2
}

type SF struct {
	F int
}

type SF1 struct {
	SF
}

type SF2 struct {
	SF1
}

type SG struct {
	G int
}

type SG1 struct {
	SG
}

type RS1 struct {
	i int
}

type RS2 struct {
	RS1
}

type RS3 struct {
	RS2
	RS1
}

type M map[string]interface{}

type Rec1 struct {
	*Rec2
}

type Rec2 struct {
	F string
	*Rec1
}
