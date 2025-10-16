// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Build initialization (after flag parsing).

package work

import (
	"github.com/shogo82148/std/cmd/go/internal/modload"
)

func BuildInit(loaderstate *modload.State)
