// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// “Abstract” syntax representation.

package ir

import (
	"github.com/shogo82148/std/fmt"
	"github.com/shogo82148/std/go/constant"

	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/src"
)

// A Node is the abstract interface to an IR node.
type Node interface {
	// Formatting
	Format(s fmt.State, verb rune)

	// Source position.
	Pos() src.XPos
	SetPos(x src.XPos)

	// For making copies. For Copy and SepCopy.
	copy() Node

	doChildren(func(Node) bool) bool
	editChildren(func(Node) Node)
	editChildrenWithHidden(func(Node) Node)

	// Abstract graph structure, for generic traversals.
	Op() Op
	Init() Nodes

	// Fields specific to certain Ops only.
	Type() *types.Type
	SetType(t *types.Type)
	Name() *Name
	Sym() *types.Sym
	Val() constant.Value
	SetVal(v constant.Value)

	// Storage for analysis passes.
	Esc() uint16
	SetEsc(x uint16)

	// Typecheck values:
	//  0 means the node is not typechecked
	//  1 means the node is completely typechecked
	//  2 means typechecking of the node is in progress
	Typecheck() uint8
	SetTypecheck(x uint8)
	NonNil() bool
	MarkNonNil()
}

// Line returns n's position as a string. If n has been inlined,
// it uses the outermost position where n has been inlined.
func Line(n Node) string

func IsSynthetic(n Node) bool

// IsAutoTmp indicates if n was created by the compiler as a temporary,
// based on the setting of the .AutoTemp flag in n's Name.
func IsAutoTmp(n Node) bool

// MayBeShared reports whether n may occur in multiple places in the AST.
// Extra care must be taken when mutating such a node.
func MayBeShared(n Node) bool

type InitNode interface {
	Node
	PtrInit() *Nodes
	SetInit(x Nodes)
}

func TakeInit(n Node) Nodes

type Op uint8

// Node ops.
const (
	OXXX Op = iota

	// names
	ONAME
	// Unnamed arg or return value: f(int, string) (int, error) { etc }
	// Also used for a qualified package identifier that hasn't been resolved yet.
	ONONAME
	OTYPE
	OLITERAL
	ONIL

	// expressions
	OADD
	OSUB
	OOR
	OXOR
	OADDSTR
	OADDR
	OANDAND
	OAPPEND
	OBYTES2STR
	OBYTES2STRTMP
	ORUNES2STR
	OSTR2BYTES
	OSTR2BYTESTMP
	OSTR2RUNES
	OSLICE2ARR
	OSLICE2ARRPTR
	// X = Y or (if Def=true) X := Y
	// If Def, then Init includes a DCL node for X.
	OAS
	// Lhs = Rhs (x, y, z = a, b, c) or (if Def=true) Lhs := Rhs
	// If Def, then Init includes DCL nodes for Lhs
	OAS2
	OAS2DOTTYPE
	OAS2FUNC
	OAS2MAPR
	OAS2RECV
	OASOP
	OCALL

	// OCALLFUNC, OCALLMETH, and OCALLINTER have the same structure.
	// Prior to walk, they are: X(Args), where Args is all regular arguments.
	// After walk, if any argument whose evaluation might requires temporary variable,
	// that temporary variable will be pushed to Init, Args will contains an updated
	// set of arguments.
	OCALLFUNC
	OCALLMETH
	OCALLINTER
	OCAP
	OCLEAR
	OCLOSE
	OCLOSURE
	OCOMPLIT
	OMAPLIT
	OSTRUCTLIT
	OARRAYLIT
	OSLICELIT
	OPTRLIT
	OCONV
	OCONVIFACE
	OCONVIDATA
	OCONVNOP
	OCOPY
	ODCL

	// Used during parsing but don't last.
	ODCLFUNC
	ODCLCONST
	ODCLTYPE

	ODELETE
	ODOT
	ODOTPTR
	ODOTMETH
	ODOTINTER
	OXDOT
	ODOTTYPE
	ODOTTYPE2
	OEQ
	ONE
	OLT
	OLE
	OGE
	OGT
	ODEREF
	OINDEX
	OINDEXMAP
	OKEY
	OSTRUCTKEY
	OLEN
	OMAKE
	OMAKECHAN
	OMAKEMAP
	OMAKESLICE
	OMAKESLICECOPY
	// OMAKESLICECOPY is created by the order pass and corresponds to:
	//  s = make(Type, Len); copy(s, Cap)
	//
	// Bounded can be set on the node when Len == len(Cap) is known at compile time.
	//
	// This node is created so the walk pass can optimize this pattern which would
	// otherwise be hard to detect after the order pass.
	OMUL
	ODIV
	OMOD
	OLSH
	ORSH
	OAND
	OANDNOT
	ONEW
	ONOT
	OBITNOT
	OPLUS
	ONEG
	OOROR
	OPANIC
	OPRINT
	OPRINTN
	OPAREN
	OSEND
	OSLICE
	OSLICEARR
	OSLICESTR
	OSLICE3
	OSLICE3ARR
	OSLICEHEADER
	OSTRINGHEADER
	ORECOVER
	ORECOVERFP
	ORECV
	ORUNESTR
	OSELRECV2
	OMIN
	OMAX
	OREAL
	OIMAG
	OCOMPLEX
	OALIGNOF
	OOFFSETOF
	OSIZEOF
	OUNSAFEADD
	OUNSAFESLICE
	OUNSAFESLICEDATA
	OUNSAFESTRING
	OUNSAFESTRINGDATA
	OMETHEXPR
	OMETHVALUE

	// statements
	OBLOCK
	OBREAK
	// OCASE:  case List: Body (List==nil means default)
	//   For OTYPESW, List is a OTYPE node for the specified type (or OLITERAL
	//   for nil) or an ODYNAMICTYPE indicating a runtime type for generics.
	//   If a type-switch variable is specified, Var is an
	//   ONAME for the version of the type-switch variable with the specified
	//   type.
	OCASE
	OCONTINUE
	ODEFER
	OFALL
	OFOR
	OGOTO
	OIF
	OLABEL
	OGO
	ORANGE
	ORETURN
	OSELECT
	OSWITCH
	// OTYPESW:  X := Y.(type) (appears as .Tag of OSWITCH)
	//   X is nil if there is no type-switch variable
	OTYPESW
	OFUNCINST

	// misc
	// intermediate representation of an inlined call.  Uses Init (assignments
	// for the captured variables, parameters, retvars, & INLMARK op),
	// Body (body of the inlined function), and ReturnVars (list of
	// return values)
	OINLCALL
	OEFACE
	OITAB
	OIDATA
	OSPTR
	OCFUNC
	OCHECKNIL
	ORESULT
	OINLMARK
	OLINKSYMOFFSET
	OJUMPTABLE

	// opcodes for generics
	ODYNAMICDOTTYPE
	ODYNAMICDOTTYPE2
	ODYNAMICTYPE

	// arch-specific opcodes
	OTAILCALL
	OGETG
	OGETCALLERPC
	OGETCALLERSP

	OEND
)

// IsCmp reports whether op is a comparison operation (==, !=, <, <=,
// >, or >=).
func (op Op) IsCmp() bool

// Nodes is a pointer to a slice of *Node.
// For fields that are not used in most nodes, this is used instead of
// a slice to save space.
type Nodes []Node

// Append appends entries to Nodes.
func (n *Nodes) Append(a ...Node)

// Prepend prepends entries to Nodes.
// If a slice is passed in, this will take ownership of it.
func (n *Nodes) Prepend(a ...Node)

// Take clears n, returning its former contents.
func (n *Nodes) Take() []Node

// Copy returns a copy of the content of the slice.
func (n Nodes) Copy() Nodes

// NameQueue is a FIFO queue of *Name. The zero value of NameQueue is
// a ready-to-use empty queue.
type NameQueue struct {
	ring       []*Name
	head, tail int
}

// Empty reports whether q contains no Names.
func (q *NameQueue) Empty() bool

// PushRight appends n to the right of the queue.
func (q *NameQueue) PushRight(n *Name)

// PopLeft pops a Name from the left of the queue. It panics if q is
// empty.
func (q *NameQueue) PopLeft() *Name

// NameSet is a set of Names.
type NameSet map[*Name]struct{}

// Has reports whether s contains n.
func (s NameSet) Has(n *Name) bool

// Add adds n to s.
func (s *NameSet) Add(n *Name)

// Sorted returns s sorted according to less.
func (s NameSet) Sorted(less func(*Name, *Name) bool) []*Name

type PragmaFlag uint16

const (
	// Func pragmas.
	Nointerface PragmaFlag = 1 << iota
	Noescape
	Norace
	Nosplit
	Noinline
	NoCheckPtr
	CgoUnsafeArgs
	UintptrKeepAlive
	UintptrEscapes

	// Runtime-only func pragmas.
	// See ../../../../runtime/HACKING.md for detailed descriptions.
	Systemstack
	Nowritebarrier
	Nowritebarrierrec
	Yeswritebarrierrec

	// Go command pragmas
	GoBuildPragma

	RegisterParams
)

func AsNode(n types.Object) Node

var BlankNode Node

func IsConst(n Node, ct constant.Kind) bool

// IsNil reports whether n represents the universal untyped zero value "nil".
func IsNil(n Node) bool

func IsBlank(n Node) bool

// IsMethod reports whether n is a method.
// n must be a function or a method.
func IsMethod(n Node) bool

func HasNamedResults(fn *Func) bool

// HasUniquePos reports whether n has a unique position that can be
// used for reporting error messages.
//
// It's primarily used to distinguish references to named objects,
// whose Pos will point back to their declaration position rather than
// their usage position.
func HasUniquePos(n Node) bool

func SetPos(n Node) src.XPos

// The result of InitExpr MUST be assigned back to n, e.g.
//
//	n.X = InitExpr(init, n.X)
func InitExpr(init []Node, expr Node) Node

// what's the outer value that a write to n affects?
// outer value means containing struct or array.
func OuterValue(n Node) Node

const (
	EscUnknown = iota
	EscNone
	EscHeap
	EscNever
)
