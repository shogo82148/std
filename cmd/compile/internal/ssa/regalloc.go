// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Register allocation.
//
// We use a version of a linear scan register allocator. We treat the
// whole function as a single long basic block and run through
// it using a greedy register allocator. Then all merge edges
// (those targeting a block with len(Preds)>1) are processed to
// shuffle data into the place that the target of the edge expects.
//
// The greedy allocator moves values into registers just before they
// are used, spills registers only when necessary, and spills the
// value whose next use is farthest in the future.
//
// The register allocator requires that a block is not scheduled until
// at least one of its predecessors have been scheduled. The most recent
// such predecessor provides the starting register state for a block.
//
// It also requires that there are no critical edges (critical =
// comes from a block with >1 successor and goes to a block with >1
// predecessor).  This makes it easy to add fixup code on merge edges -
// the source of a merge edge has only one successor, so we can add
// fixup code to the end of that block.

// Spilling
//
// During the normal course of the allocator, we might throw a still-live
// value out of all registers. When that value is subsequently used, we must
// load it from a slot on the stack. We must also issue an instruction to
// initialize that stack location with a copy of v.
//
// pre-regalloc:
//   (1) v = Op ...
//   (2) x = Op ...
//   (3) ... = Op v ...
//
// post-regalloc:
//   (1) v = Op ...    : AX // computes v, store result in AX
//       s = StoreReg v     // spill v to a stack slot
//   (2) x = Op ...    : AX // some other op uses AX
//       c = LoadReg s : CX // restore v from stack slot
//   (3) ... = Op c ...     // use the restored value
//
// Allocation occurs normally until we reach (3) and we realize we have
// a use of v and it isn't in any register. At that point, we allocate
// a spill (a StoreReg) for v. We can't determine the correct place for
// the spill at this point, so we allocate the spill as blockless initially.
// The restore is then generated to load v back into a register so it can
// be used. Subsequent uses of v will use the restored value c instead.
//
// What remains is the question of where to schedule the spill.
// During allocation, we keep track of the dominator of all restores of v.
// The spill of v must dominate that block. The spill must also be issued at
// a point where v is still in a register.
//
// To find the right place, start at b, the block which dominates all restores.
//  - If b is v.Block, then issue the spill right after v.
//    It is known to be in a register at that point, and dominates any restores.
//  - Otherwise, if v is in a register at the start of b,
//    put the spill of v at the start of b.
//  - Otherwise, set b = immediate dominator of b, and repeat.
//
// Phi values are special, as always. We define two kinds of phis, those
// where the merge happens in a register (a "register" phi) and those where
// the merge happens in a stack location (a "stack" phi).
//
// A register phi must have the phi and all of its inputs allocated to the
// same register. Register phis are spilled similarly to regular ops.
//
// A stack phi must have the phi and all of its inputs allocated to the same
// stack location. Stack phis start out life already spilled - each phi
// input must be a store (using StoreReg) at the end of the corresponding
// predecessor block.
//     b1: y = ... : AX        b2: z = ... : BX
//         y2 = StoreReg y         z2 = StoreReg z
//         goto b3                 goto b3
//     b3: x = phi(y2, z2)
// The stack allocator knows that StoreReg args of stack-allocated phis
// must be allocated to the same stack slot as the phi that uses them.
// x is now a spilled value and a restore must appear before its first use.

// TODO

// Use an affinity graph to mark two values which should use the
// same register. This affinity graph will be used to prefer certain
// registers for allocation. This affinity helps eliminate moves that
// are required for phi implementations and helps generate allocations
// for 2-register architectures.

// Note: regalloc generates a not-quite-SSA output. If we have:
//
//             b1: x = ... : AX
//                 x2 = StoreReg x
//                 ... AX gets reused for something else ...
//                 if ... goto b3 else b4
//
//   b3: x3 = LoadReg x2 : BX       b4: x4 = LoadReg x2 : CX
//       ... use x3 ...                 ... use x4 ...
//
//             b2: ... use x3 ...
//
// If b3 is the primary predecessor of b2, then we use x3 in b2 and
// add a x4:CX->BX copy at the end of b4.
// But the definition of x3 doesn't dominate b2.  We should really
// insert an extra phi at the start of b2 (x5=phi(x3,x4):BX) to keep
// SSA form. For now, we ignore this problem as remaining in strict
// SSA form isn't needed after regalloc. We'll just leave the use
// of x3 not dominated by the definition of x3, and the CX->BX copy
// will have no use (so don't run deadcode after regalloc!).
// TODO: maybe we should introduce these extra phis?

package ssa
