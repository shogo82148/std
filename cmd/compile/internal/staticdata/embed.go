// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package staticdata

import (
	"github.com/shogo82148/std/cmd/compile/internal/ir"
)

// WriteEmbed emits the init data for a //go:embed variable,
// which is either a string, a []byte, or an embed.FS.
func WriteEmbed(v *ir.Name)
