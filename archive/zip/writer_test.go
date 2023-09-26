// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package zip

import (
	"io/fs"
)

type WriteTest struct {
	Name   string
	Data   []byte
	Method uint16
	Mode   fs.FileMode
}
