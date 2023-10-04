// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Disabled for s390x because it uses assembly routines that are not
// accurate for huge arguments.

//go:build !s390x

package math_test

import (
	. "math"
)
