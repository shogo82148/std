// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package load

import (
	"github.com/shogo82148/std/errors"
)

var ErrNotGoDebug = errors.New("not //go:debug line")

func ParseGoDebug(text string) (key, value string, err error)
