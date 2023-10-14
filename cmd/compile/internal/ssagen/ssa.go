// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssagen

import (
	"github.com/shogo82148/std/cmd/compile/internal/abi"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/liveness"
	"github.com/shogo82148/std/cmd/compile/internal/objw"
	"github.com/shogo82148/std/cmd/compile/internal/ssa"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

func DumpInline(fn *ir.Func)

func InitEnv()

func InitConfig()

// AbiForBodylessFuncStackMap returns the ABI for a bodyless function's stack map.
// This is not necessarily the ABI used to call it.
// Currently (1.17 dev) such a stack map is always ABI0;
// any ABI wrapper that is present is nosplit, hence a precise
// stack map is not needed there (the parameters survive only long
// enough to call the wrapped assembly function).
// This always returns a freshly copied ABI.
func AbiForBodylessFuncStackMap(fn *ir.Func) *abi.ABIConfig

func InitTables()

func IsIntrinsicCall(n *ir.CallExpr) bool

// TypeOK reports whether variables of type t are SSA-able.
func TypeOK(t *types.Type) bool

// Branch is an unresolved branch.
type Branch struct {
	P *obj.Prog
	B *ssa.Block
}

// State contains state needed during Prog generation.
type State struct {
	ABI obj.ABI

	pp *objw.Progs

	// Branches remembers all the branch instructions we've seen
	// and where they would like to go.
	Branches []Branch

	// JumpTables remembers all the jump tables we've seen.
	JumpTables []*ssa.Block

	// bstart remembers where each block starts (indexed by block ID)
	bstart []*obj.Prog

	maxarg int64

	// Map from GC safe points to liveness index, generated by
	// liveness analysis.
	livenessMap liveness.Map

	// partLiveArgs includes arguments that may be partially live, for which we
	// need to generate instructions that spill the argument registers.
	partLiveArgs map[*ir.Name]bool

	// lineRunStart records the beginning of the current run of instructions
	// within a single block sharing the same line number
	// Used to move statement marks to the beginning of such runs.
	lineRunStart *obj.Prog

	// wasm: The number of values on the WebAssembly stack. This is only used as a safeguard.
	OnWasmStackSkipped int
}

func (s *State) FuncInfo() *obj.FuncInfo

// Prog appends a new Prog.
func (s *State) Prog(as obj.As) *obj.Prog

// Pc returns the current Prog.
func (s *State) Pc() *obj.Prog

// SetPos sets the current source position.
func (s *State) SetPos(pos src.XPos)

// Br emits a single branch instruction and returns the instruction.
// Not all architectures need the returned instruction, but otherwise
// the boilerplate is common to all.
func (s *State) Br(op obj.As, target *ssa.Block) *obj.Prog

// DebugFriendlySetPosFrom adjusts Pos.IsStmt subject to heuristics
// that reduce "jumpy" line number churn when debugging.
// Spill/fill/copy instructions from the register allocator,
// phi functions, and instructions with a no-pos position
// are examples of instructions that can cause churn.
func (s *State) DebugFriendlySetPosFrom(v *ssa.Value)

// emit argument info (locations on stack) of f for traceback.
func EmitArgInfo(f *ir.Func, abiInfo *abi.ABIParamResultInfo) *obj.LSym

// For generating consecutive jump instructions to model a specific branching
type IndexJump struct {
	Jump  obj.As
	Index int
}

// CombJump generates combinational instructions (2 at present) for a block jump,
// thereby the behaviour of non-standard condition codes could be simulated
func (s *State) CombJump(b, next *ssa.Block, jumps *[2][2]IndexJump)

// AddAux adds the offset in the aux fields (AuxInt and Aux) of v to a.
func AddAux(a *obj.Addr, v *ssa.Value)

func AddAux2(a *obj.Addr, v *ssa.Value, offset int64)

// CheckLoweredPhi checks that regalloc and stackalloc correctly handled phi values.
// Called during ssaGenValue.
func CheckLoweredPhi(v *ssa.Value)

// CheckLoweredGetClosurePtr checks that v is the first instruction in the function's entry block,
// except for incoming in-register arguments.
// The output of LoweredGetClosurePtr is generally hardwired to the correct register.
// That register contains the closure pointer on closure entry.
func CheckLoweredGetClosurePtr(v *ssa.Value)

// CheckArgReg ensures that v is in the function's entry block.
func CheckArgReg(v *ssa.Value)

func AddrAuto(a *obj.Addr, v *ssa.Value)

// Call returns a new CALL instruction for the SSA value v.
// It uses PrepareCall to prepare the call.
func (s *State) Call(v *ssa.Value) *obj.Prog

// TailCall returns a new tail call instruction for the SSA value v.
// It is like Call, but for a tail call.
func (s *State) TailCall(v *ssa.Value) *obj.Prog

// PrepareCall prepares to emit a CALL instruction for v and does call-related bookkeeping.
// It must be called immediately before emitting the actual CALL instruction,
// since it emits PCDATA for the stack map at the call (calls are safe points).
func (s *State) PrepareCall(v *ssa.Value)

// UseArgs records the fact that an instruction needs a certain amount of
// callee args space for its use.
func (s *State) UseArgs(n int64)

// SpillSlotAddr uses LocalSlot information to initialize an obj.Addr
// The resulting addr is used in a non-standard context -- in the prologue
// of a function, before the frame has been constructed, so the standard
// addressing for the parameters will be wrong.
func SpillSlotAddr(spill ssa.Spill, baseReg int16, extraOffset int64) obj.Addr

var (
	BoundsCheckFunc [ssa.BoundsKindCount]*obj.LSym
	ExtendCheckFunc [ssa.BoundsKindCount]*obj.LSym
)