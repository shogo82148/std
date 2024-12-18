// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin

package issue24161e2

import "github.com/shogo82148/std/C"
import (
	"github.com/shogo82148/std/testing"
)

var _ C.CFStringRef

func Test(t *testing.T)
