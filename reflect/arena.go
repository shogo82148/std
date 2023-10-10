// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.arenas

package reflect

import "github.com/shogo82148/std/arena"

// ArenaNewは指定された型の新しいゼロ値を表すポインタを返し、提供されたアリーナにそのためのストレージを割り当てます。つまり、返されるValueのTypeはtypのポインタです。
func ArenaNew(a *arena.Arena, typ Type) Value
