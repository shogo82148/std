// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objabi

import (
	"github.com/shogo82148/std/internal/abi"
)

// Get the function ID for the named function in the named file.
// The function should be package-qualified.
func GetFuncID(name string, isWrapper bool) abi.FuncID
