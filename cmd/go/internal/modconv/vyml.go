// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modconv

import (
	"golang.org/x/mod/modfile"
)

func ParseVendorYML(file string, data []byte) (*modfile.File, error)
