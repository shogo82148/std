// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.arenas

package reflect

import "github.com/shogo82148/std/arena"

<<<<<<< HEAD
// ArenaNewは指定された型の新しいゼロ値を表すポインタを返し、提供されたアリーナにそのためのストレージを割り当てます。つまり、返されるValueのTypeはtypのポインタです。
=======
// ArenaNew returns a [Value] representing a pointer to a new zero value for the
// specified type, allocating storage for it in the provided arena. That is,
// the returned Value's Type is [PointerTo](typ).
>>>>>>> upstream/master
func ArenaNew(a *arena.Arena, typ Type) Value
