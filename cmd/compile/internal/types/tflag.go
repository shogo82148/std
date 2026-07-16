// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/internal/abi"
)

// TFlag returns the abi.TFlag value for t's runtime type. Callers
// must have run typecheck.CalcMethods on ReceiverBaseType(t).
func (t *Type) TFlag() abi.TFlag
