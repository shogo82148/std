// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains some utility functions to help define Funcs for testing.
// As an example, the following func
//
//   b1:
//     v1 = InitMem <mem>
//     Plain -> b2
//   b2:
//     Exit v1
//   b3:
//     v2 = Const <bool> [true]
//     If v2 -> b3 b2
//
// can be defined as
//
//   fun := Fun("entry",
//       Bloc("entry",
//           Valu("mem", OpInitMem, types.TypeMem, 0, nil),
//           Goto("exit")),
//       Bloc("exit",
//           Exit("mem")),
//       Bloc("deadblock",
//          Valu("deadval", OpConstBool, c.config.Types.Bool, 0, true),
//          If("deadval", "deadblock", "exit")))
//
// and the Blocks or Values used in the Func can be accessed
// like this:
//   fun.blocks["entry"] or fun.values["deadval"]

package ssa
