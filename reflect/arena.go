// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.arenas

package reflect

import "github.com/shogo82148/std/arena"

// ArenaNewは、指定された型の新しいゼロ値へのポインタを表す [Value] を返し、
// 提供されたアリーナ内にそのためのストレージを割り当てます。つまり、
// 返されるValueのTypeは [PointerTo](typ)です。
func ArenaNew(a *arena.Arena, typ Type) Value
