// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssa

import (
	"github.com/shogo82148/std/cmd/compile/internal/abi"
	"github.com/shogo82148/std/cmd/compile/internal/ir"
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
	"github.com/shogo82148/std/cmd/internal/src"
)

// A Func represents a Go func declaration (or function literal) and its body.
// This package compiles each Func independently.
// Funcs are single-use; a new Func must be created for every compiled function.
type Func struct {
	Config *Config
	Cache  *Cache
	fe     Frontend
	pass   *pass
	Name   string
	Type   *types.Type
	Blocks []*Block
	Entry  *Block

	bid idAlloc
	vid idAlloc

	HTMLWriter     *HTMLWriter
	PrintOrHtmlSSA bool
	ruleMatches    map[string]int
	ABI0           *abi.ABIConfig
	ABI1           *abi.ABIConfig
	ABISelf        *abi.ABIConfig
	ABIDefault     *abi.ABIConfig

	scheduled   bool
	laidout     bool
	NoSplit     bool
	dumpFileSeq uint8
	IsPgoHot    bool
	DeferReturn *Block

	// when register allocation is done, maps value ids to locations
	RegAlloc []Location

	// temporary registers allocated to rare instructions
	tempRegs map[ID]*Register

	// map from LocalSlot to set of Values that we want to store in that slot.
	NamedValues map[LocalSlot][]*Value
	// Names is a copy of NamedValues.Keys. We keep a separate list
	// of keys to make iteration order deterministic.
	Names []*LocalSlot
	// Canonicalize root/top-level local slots, and canonicalize their pieces.
	// Because LocalSlot pieces refer to their parents with a pointer, this ensures that equivalent slots really are equal.
	CanonicalLocalSlots  map[LocalSlot]*LocalSlot
	CanonicalLocalSplits map[LocalSlotSplitKey]*LocalSlot

	// RegArgs is a slice of register-memory pairs that must be spilled and unspilled in the uncommon path of function entry.
	RegArgs []Spill
	// OwnAux describes parameters and results for this function.
	OwnAux *AuxCall
	// CloSlot holds the compiler-synthesized name (".closureptr")
	// where we spill the closure pointer for range func bodies.
	CloSlot *ir.Name

	freeValues *Value
	freeBlocks *Block

	cachedPostorder  []*Block
	cachedIdom       []*Block
	cachedSdom       SparseTree
	cachedLoopnest   *loopnest
	cachedLineStarts *xposmap

	auxmap    auxmap
	constants map[int64][]*Value
}

type LocalSlotSplitKey struct {
	parent *LocalSlot
	Off    int64
	Type   *types.Type
}

// NewFunc returns a new, empty function object.
// Caller must reset cache before calling NewFunc.
func (c *Config) NewFunc(fe Frontend, cache *Cache) *Func

// NumBlocks returns an integer larger than the id of any Block in the Func.
func (f *Func) NumBlocks() int

// NumValues returns an integer larger than the id of any Value in the Func.
func (f *Func) NumValues() int

// NameABI returns the function name followed by comma and the ABI number.
// This is intended for use with GOSSAFUNC and HTML dumps, and differs from
// the linker's "<1>" convention because "<" and ">" require shell quoting
// and are not legal file names (for use with GOSSADIR) on Windows.
func (f *Func) NameABI() string

// FuncNameABI returns n followed by a comma and the value of a.
// This is a separate function to allow a single point encoding
// of the format, which is used in places where there's not a Func yet.
func FuncNameABI(n string, a obj.ABI) string

func (f *Func) SplitString(name *LocalSlot) (*LocalSlot, *LocalSlot)

func (f *Func) SplitInterface(name *LocalSlot) (*LocalSlot, *LocalSlot)

func (f *Func) SplitSlice(name *LocalSlot) (*LocalSlot, *LocalSlot, *LocalSlot)

func (f *Func) SplitComplex(name *LocalSlot) (*LocalSlot, *LocalSlot)

func (f *Func) SplitInt64(name *LocalSlot) (*LocalSlot, *LocalSlot)

func (f *Func) SplitStruct(name *LocalSlot, i int) *LocalSlot

func (f *Func) SplitArray(name *LocalSlot) *LocalSlot

func (f *Func) SplitSlot(name *LocalSlot, sfx string, offset int64, t *types.Type) *LocalSlot

// LogStat writes a string key and int value as a warning in a
// tab-separated format easily handled by spreadsheets or awk.
// file names, lines, and function names are included to provide enough (?)
// context to allow item-by-item comparisons across runs.
// For example:
// awk 'BEGIN {FS="\t"} $3~/TIME/{sum+=$4} END{print "t(ns)=",sum}' t.log
func (f *Func) LogStat(key string, args ...any)

// NewBlock allocates a new Block of the given kind and places it at the end of f.Blocks.
func (f *Func) NewBlock(kind BlockKind) *Block

// NewValue0 returns a new value in the block with no arguments and zero aux values.
func (b *Block) NewValue0(pos src.XPos, op Op, t *types.Type) *Value

// NewValue0I returns a new value in the block with no arguments and an auxint value.
func (b *Block) NewValue0I(pos src.XPos, op Op, t *types.Type, auxint int64) *Value

// NewValue0A returns a new value in the block with no arguments and an aux value.
func (b *Block) NewValue0A(pos src.XPos, op Op, t *types.Type, aux Aux) *Value

// NewValue0IA returns a new value in the block with no arguments and both an auxint and aux values.
func (b *Block) NewValue0IA(pos src.XPos, op Op, t *types.Type, auxint int64, aux Aux) *Value

// NewValue1 returns a new value in the block with one argument and zero aux values.
func (b *Block) NewValue1(pos src.XPos, op Op, t *types.Type, arg *Value) *Value

// NewValue1I returns a new value in the block with one argument and an auxint value.
func (b *Block) NewValue1I(pos src.XPos, op Op, t *types.Type, auxint int64, arg *Value) *Value

// NewValue1A returns a new value in the block with one argument and an aux value.
func (b *Block) NewValue1A(pos src.XPos, op Op, t *types.Type, aux Aux, arg *Value) *Value

// NewValue1IA returns a new value in the block with one argument and both an auxint and aux values.
func (b *Block) NewValue1IA(pos src.XPos, op Op, t *types.Type, auxint int64, aux Aux, arg *Value) *Value

// NewValue2 returns a new value in the block with two arguments and zero aux values.
func (b *Block) NewValue2(pos src.XPos, op Op, t *types.Type, arg0, arg1 *Value) *Value

// NewValue2A returns a new value in the block with two arguments and one aux values.
func (b *Block) NewValue2A(pos src.XPos, op Op, t *types.Type, aux Aux, arg0, arg1 *Value) *Value

// NewValue2I returns a new value in the block with two arguments and an auxint value.
func (b *Block) NewValue2I(pos src.XPos, op Op, t *types.Type, auxint int64, arg0, arg1 *Value) *Value

// NewValue2IA returns a new value in the block with two arguments and both an auxint and aux values.
func (b *Block) NewValue2IA(pos src.XPos, op Op, t *types.Type, auxint int64, aux Aux, arg0, arg1 *Value) *Value

// NewValue3 returns a new value in the block with three arguments and zero aux values.
func (b *Block) NewValue3(pos src.XPos, op Op, t *types.Type, arg0, arg1, arg2 *Value) *Value

// NewValue3I returns a new value in the block with three arguments and an auxint value.
func (b *Block) NewValue3I(pos src.XPos, op Op, t *types.Type, auxint int64, arg0, arg1, arg2 *Value) *Value

// NewValue3A returns a new value in the block with three argument and an aux value.
func (b *Block) NewValue3A(pos src.XPos, op Op, t *types.Type, aux Aux, arg0, arg1, arg2 *Value) *Value

// NewValue4 returns a new value in the block with four arguments and zero aux values.
func (b *Block) NewValue4(pos src.XPos, op Op, t *types.Type, arg0, arg1, arg2, arg3 *Value) *Value

// NewValue4I returns a new value in the block with four arguments and auxint value.
func (b *Block) NewValue4I(pos src.XPos, op Op, t *types.Type, auxint int64, arg0, arg1, arg2, arg3 *Value) *Value

// ConstBool returns an int constant representing its argument.
func (f *Func) ConstBool(t *types.Type, c bool) *Value

func (f *Func) ConstInt8(t *types.Type, c int8) *Value

func (f *Func) ConstInt16(t *types.Type, c int16) *Value

func (f *Func) ConstInt32(t *types.Type, c int32) *Value

func (f *Func) ConstInt64(t *types.Type, c int64) *Value

func (f *Func) ConstFloat32(t *types.Type, c float64) *Value

func (f *Func) ConstFloat64(t *types.Type, c float64) *Value

func (f *Func) ConstSlice(t *types.Type) *Value

func (f *Func) ConstInterface(t *types.Type) *Value

func (f *Func) ConstNil(t *types.Type) *Value

func (f *Func) ConstEmptyString(t *types.Type) *Value

func (f *Func) ConstOffPtrSP(t *types.Type, c int64, sp *Value) *Value

func (f *Func) Frontend() Frontend
func (f *Func) Warnl(pos src.XPos, msg string, args ...any)
func (f *Func) Logf(msg string, args ...any)
func (f *Func) Log() bool

func (f *Func) Fatalf(msg string, args ...any)

func (f *Func) Postorder() []*Block

// Idom returns a map from block ID to the immediate dominator of that block.
// f.Entry.ID maps to nil. Unreachable blocks map to nil as well.
func (f *Func) Idom() []*Block

// Sdom returns a sparse tree representing the dominator relationships
// among the blocks of f.
func (f *Func) Sdom() SparseTree

// DebugHashMatch returns
//
//	base.DebugHashMatch(this function's package.name)
//
// for use in bug isolation.  The return value is true unless
// environment variable GOCOMPILEDEBUG=gossahash=X is set, in which case "it depends on X".
// See [base.DebugHashMatch] for more information.
func (f *Func) DebugHashMatch() bool

// NewLocal returns a new anonymous local variable of the given type.
func (f *Func) NewLocal(pos src.XPos, typ *types.Type) *ir.Name

// IsMergeCandidate returns true if variable n could participate in
// stack slot merging. For now we're restricting the set to things to
// items larger than what CanSSA would allow (approximateky, we disallow things
// marked as open defer slots so as to avoid complicating liveness
// analysis.
func IsMergeCandidate(n *ir.Name) bool
