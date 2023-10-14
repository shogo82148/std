// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/cmd/internal/src"
)

// Object represents an ir.Node, but without needing to import cmd/compile/internal/ir,
// which would cause an import cycle. The uses in other packages must type assert
// values of type Object to ir.Node or a more specific type.
type Object interface {
	Pos() src.XPos
	Sym() *Sym
	Type() *Type
}

// Kind describes a kind of type.
type Kind uint8

const (
	Txxx Kind = iota

	TINT8
	TUINT8
	TINT16
	TUINT16
	TINT32
	TUINT32
	TINT64
	TUINT64
	TINT
	TUINT
	TUINTPTR

	TCOMPLEX64
	TCOMPLEX128

	TFLOAT32
	TFLOAT64

	TBOOL

	TPTR
	TFUNC
	TSLICE
	TARRAY
	TSTRUCT
	TCHAN
	TMAP
	TINTER
	TFORW
	TANY
	TSTRING
	TUNSAFEPTR

	// pseudo-types for literals
	TIDEAL
	TNIL
	TBLANK

	// pseudo-types used temporarily only during frame layout (CalcSize())
	TFUNCARGS
	TCHANARGS

	// SSA backend types
	TSSA
	TTUPLE
	TRESULTS

	NTYPE
)

// ChanDir is whether a channel can send, receive, or both.
type ChanDir uint8

func (c ChanDir) CanRecv() bool
func (c ChanDir) CanSend() bool

const (
	// types of channel
	// must match ../../../../reflect/type.go:/ChanDir
	Crecv ChanDir = 1 << 0
	Csend ChanDir = 1 << 1
	Cboth ChanDir = Crecv | Csend
)

// Types stores pointers to predeclared named types.
//
// It also stores pointers to several special types:
//   - Types[TANY] is the placeholder "any" type recognized by SubstArgTypes.
//   - Types[TBLANK] represents the blank variable's type.
//   - Types[TINTER] is the canonical "interface{}" type.
//   - Types[TNIL] represents the predeclared "nil" value's type.
//   - Types[TUNSAFEPTR] is package unsafe's Pointer type.
var Types [NTYPE]*Type

var (
	// Predeclared alias types. These are actually created as distinct
	// defined types for better error messages, but are then specially
	// treated as identical to their respective underlying types.
	AnyType  *Type
	ByteType *Type
	RuneType *Type

	// Predeclared error interface type.
	ErrorType *Type
	// Predeclared comparable interface type.
	ComparableType *Type

	// Types to represent untyped string and boolean constants.
	UntypedString = newType(TSTRING)
	UntypedBool   = newType(TBOOL)

	// Types to represent untyped numeric constants.
	UntypedInt     = newType(TIDEAL)
	UntypedRune    = newType(TIDEAL)
	UntypedFloat   = newType(TIDEAL)
	UntypedComplex = newType(TIDEAL)
)

// A Type represents a Go type.
//
// There may be multiple unnamed types with identical structure. However, there must
// be a unique Type object for each unique named (defined) type. After noding, a
// package-level type can be looked up by building its unique symbol sym (sym =
// package.Lookup(name)) and checking sym.Def. If sym.Def is non-nil, the type
// already exists at package scope and is available at sym.Def.(*ir.Name).Type().
// Local types (which may have the same name as a package-level type) are
// distinguished by the value of vargen.
type Type struct {
	// extra contains extra etype-specific fields.
	// As an optimization, those etype-specific structs which contain exactly
	// one pointer-shaped field are stored as values rather than pointers when possible.
	//
	// TMAP: *Map
	// TFORW: *Forward
	// TFUNC: *Func
	// TSTRUCT: *Struct
	// TINTER: *Interface
	// TFUNCARGS: FuncArgs
	// TCHANARGS: ChanArgs
	// TCHAN: *Chan
	// TPTR: Ptr
	// TARRAY: *Array
	// TSLICE: Slice
	// TSSA: string
	extra interface{}

	// width is the width of this Type in bytes.
	width int64

	// list of base methods (excluding embedding)
	methods Fields
	// list of all methods (including embedding)
	allMethods Fields

	// canonical OTYPE node for a named type (should be an ir.Name node with same sym)
	obj Object
	// the underlying type (type literal or predeclared type) for a defined type
	underlying *Type

	// Cache of composite types, with this type being the element type.
	cache struct {
		ptr   *Type
		slice *Type
	}

	vargen int32

	kind  Kind
	align uint8

	flags bitset8

	// For defined (named) generic types, a pointer to the list of type params
	// (in order) of this type that need to be instantiated. For instantiated
	// generic types, this is the targs used to instantiate them. These targs
	// may be typeparams (for re-instantiated types such as Value[T2]) or
	// concrete types (for fully instantiated types such as Value[int]).
	// rparams is only set for named types that are generic or are fully
	// instantiated from a generic type, and is otherwise set to nil.
	// TODO(danscales): choose a better name.
	rparams *[]*Type
}

func (*Type) CanBeAnSSAAux()

func (t *Type) NotInHeap() bool
func (t *Type) Noalg() bool
func (t *Type) Deferwidth() bool
func (t *Type) Recur() bool
func (t *Type) IsShape() bool
func (t *Type) HasShape() bool

func (t *Type) SetNotInHeap(b bool)
func (t *Type) SetNoalg(b bool)
func (t *Type) SetDeferwidth(b bool)
func (t *Type) SetRecur(b bool)

// Should always do SetHasShape(true) when doing SetIsShape(true).
func (t *Type) SetIsShape(b bool)
func (t *Type) SetHasShape(b bool)

// Kind returns the kind of type t.
func (t *Type) Kind() Kind

// Sym returns the name of type t.
func (t *Type) Sym() *Sym

// Underlying returns the underlying type of type t.
func (t *Type) Underlying() *Type

// Pos returns a position associated with t, if any.
// This should only be used for diagnostics.
func (t *Type) Pos() src.XPos

func (t *Type) RParams() []*Type

func (t *Type) SetRParams(rparams []*Type)

// IsFullyInstantiated reports whether t is a fully instantiated generic type; i.e. an
// instantiated generic type where all type arguments are non-generic or fully
// instantiated generic types.
func (t *Type) IsFullyInstantiated() bool

// Map contains Type fields specific to maps.
type Map struct {
	Key  *Type
	Elem *Type

	Bucket *Type
	Hmap   *Type
	Hiter  *Type
}

// MapType returns t's extra map-specific fields.
func (t *Type) MapType() *Map

// Forward contains Type fields specific to forward types.
type Forward struct {
	Copyto      []*Type
	Embedlineno src.XPos
}

// ForwardType returns t's extra forward-type-specific fields.
func (t *Type) ForwardType() *Forward

// Func contains Type fields specific to func types.
type Func struct {
	Receiver *Type
	Results  *Type
	Params   *Type

	// Argwid is the total width of the function receiver, params, and results.
	// It gets calculated via a temporary TFUNCARGS type.
	// Note that TFUNC's Width is Widthptr.
	Argwid int64
}

// FuncType returns t's extra func-specific fields.
func (t *Type) FuncType() *Func

// StructType contains Type fields specific to struct types.
type Struct struct {
	fields Fields

	// Maps have three associated internal structs (see struct MapType).
	// Map links such structs back to their map type.
	Map *Type

	Funarg Funarg
}

// Funarg records the kind of function argument
type Funarg uint8

const (
	FunargNone Funarg = iota
	FunargRcvr
	FunargParams
	FunargResults
	FunargTparams
)

// StructType returns t's extra struct-specific fields.
func (t *Type) StructType() *Struct

// Interface contains Type fields specific to interface types.
type Interface struct {
}

// Ptr contains Type fields specific to pointer types.
type Ptr struct {
	Elem *Type
}

// ChanArgs contains Type fields specific to TCHANARGS types.
type ChanArgs struct {
	T *Type
}

// // FuncArgs contains Type fields specific to TFUNCARGS types.
type FuncArgs struct {
	T *Type
}

// Chan contains Type fields specific to channel types.
type Chan struct {
	Elem *Type
	Dir  ChanDir
}

// ChanType returns t's extra channel-specific fields.
func (t *Type) ChanType() *Chan

type Tuple struct {
	first  *Type
	second *Type
}

// Results are the output from calls that will be late-expanded.
type Results struct {
	Types []*Type
}

// Array contains Type fields specific to array types.
type Array struct {
	Elem  *Type
	Bound int64
}

// Slice contains Type fields specific to slice types.
type Slice struct {
	Elem *Type
}

// A Field is a (Sym, Type) pairing along with some other information, and,
// depending on the context, is used to represent:
//   - a field in a struct
//   - a method in an interface or associated with a named type
//   - a function parameter
type Field struct {
	flags bitset8

	Embedded uint8

	Pos src.XPos

	// Name of field/method/parameter. Can be nil for interface fields embedded
	// in interfaces and unnamed parameters.
	Sym  *Sym
	Type *Type
	Note string

	// For fields that represent function parameters, Nname points to the
	// associated ONAME Node. For fields that represent methods, Nname points to
	// the function name node.
	Nname Object

	// Offset in bytes of this field or method within its enclosing struct
	// or interface Type.  Exception: if field is function receiver, arg or
	// result, then this is BOGUS_FUNARG_OFFSET; types does not know the Abi.
	Offset int64
}

func (f *Field) IsDDD() bool
func (f *Field) Nointerface() bool

func (f *Field) SetIsDDD(b bool)
func (f *Field) SetNointerface(b bool)

// End returns the offset of the first byte immediately after this field.
func (f *Field) End() int64

// IsMethod reports whether f represents a method rather than a struct field.
func (f *Field) IsMethod() bool

// Fields is a pointer to a slice of *Field.
// This saves space in Types that do not have fields or methods
// compared to a simple slice of *Field.
type Fields struct {
	s *[]*Field
}

// Len returns the number of entries in f.
func (f *Fields) Len() int

// Slice returns the entries in f as a slice.
// Changes to the slice entries will be reflected in f.
func (f *Fields) Slice() []*Field

// Index returns the i'th element of Fields.
// It panics if f does not have at least i+1 elements.
func (f *Fields) Index(i int) *Field

// Set sets f to a slice.
// This takes ownership of the slice.
func (f *Fields) Set(s []*Field)

// Append appends entries to f.
func (f *Fields) Append(s ...*Field)

// NewArray returns a new fixed-length array Type.
func NewArray(elem *Type, bound int64) *Type

// NewSlice returns the slice Type with element type elem.
func NewSlice(elem *Type) *Type

// NewChan returns a new chan Type with direction dir.
func NewChan(elem *Type, dir ChanDir) *Type

func NewTuple(t1, t2 *Type) *Type

func NewResults(types []*Type) *Type

// NewMap returns a new map Type with key type k and element (aka value) type v.
func NewMap(k, v *Type) *Type

// NewPtrCacheEnabled controls whether *T Types are cached in T.
// Caching is disabled just before starting the backend.
// This allows the backend to run concurrently.
var NewPtrCacheEnabled = true

// NewPtr returns the pointer type pointing to t.
func NewPtr(elem *Type) *Type

// NewChanArgs returns a new TCHANARGS type for channel type c.
func NewChanArgs(c *Type) *Type

// NewFuncArgs returns a new TFUNCARGS type for func type f.
func NewFuncArgs(f *Type) *Type

func NewField(pos src.XPos, sym *Sym, typ *Type) *Field

// SubstAny walks t, replacing instances of "any" with successive
// elements removed from types.  It returns the substituted type.
func SubstAny(t *Type, types *[]*Type) *Type

func (f *Field) Copy() *Field

func (t *Type) Recvs() *Type
func (t *Type) Params() *Type
func (t *Type) Results() *Type

func (t *Type) NumRecvs() int
func (t *Type) NumParams() int
func (t *Type) NumResults() int

// IsVariadic reports whether function type t is variadic.
func (t *Type) IsVariadic() bool

// Recv returns the receiver of function type t, if any.
func (t *Type) Recv() *Field

// RecvsParamsResults stores the accessor functions for a function Type's
// receiver, parameters, and result parameters, in that order.
// It can be used to iterate over all of a function's parameter lists.
var RecvsParamsResults = [3]func(*Type) *Type{
	(*Type).Recvs, (*Type).Params, (*Type).Results,
}

// RecvsParams is like RecvsParamsResults, but omits result parameters.
var RecvsParams = [2]func(*Type) *Type{
	(*Type).Recvs, (*Type).Params,
}

// ParamsResults is like RecvsParamsResults, but omits receiver parameters.
var ParamsResults = [2]func(*Type) *Type{
	(*Type).Params, (*Type).Results,
}

// Key returns the key type of map type t.
func (t *Type) Key() *Type

// Elem returns the type of elements of t.
// Usable with pointers, channels, arrays, slices, and maps.
func (t *Type) Elem() *Type

// ChanArgs returns the channel type for TCHANARGS type t.
func (t *Type) ChanArgs() *Type

// FuncArgs returns the func type for TFUNCARGS type t.
func (t *Type) FuncArgs() *Type

// IsFuncArgStruct reports whether t is a struct representing function parameters or results.
func (t *Type) IsFuncArgStruct() bool

// Methods returns a pointer to the base methods (excluding embedding) for type t.
// These can either be concrete methods (for non-interface types) or interface
// methods (for interface types).
func (t *Type) Methods() *Fields

// AllMethods returns a pointer to all the methods (including embedding) for type t.
// For an interface type, this is the set of methods that are typically iterated
// over. For non-interface types, AllMethods() only returns a valid result after
// CalcMethods() has been called at least once.
func (t *Type) AllMethods() *Fields

// SetAllMethods sets the set of all methods (including embedding) for type t.
// Use this method instead of t.AllMethods().Set(), which might call CalcSize() on
// an uninitialized interface type.
func (t *Type) SetAllMethods(fs []*Field)

// Fields returns the fields of struct type t.
func (t *Type) Fields() *Fields

// Field returns the i'th field of struct type t.
func (t *Type) Field(i int) *Field

// FieldSlice returns a slice of containing all fields of
// a struct type t.
func (t *Type) FieldSlice() []*Field

// SetFields sets struct type t's fields to fields.
func (t *Type) SetFields(fields []*Field)

// SetInterface sets the base methods of an interface type t.
func (t *Type) SetInterface(methods []*Field)

// ArgWidth returns the total aligned argument size for a function.
// It includes the receiver, parameters, and results.
func (t *Type) ArgWidth() int64

func (t *Type) Size() int64

func (t *Type) Alignment() int64

func (t *Type) SimpleString() string

// Cmp is a comparison between values a and b.
//
//	-1 if a < b
//	 0 if a == b
//	 1 if a > b
type Cmp int8

const (
	CMPlt = Cmp(-1)
	CMPeq = Cmp(0)
	CMPgt = Cmp(1)
)

// Compare compares types for purposes of the SSA back
// end, returning a Cmp (one of CMPlt, CMPeq, CMPgt).
// The answers are correct for an optimizer
// or code generator, but not necessarily typechecking.
// The order chosen is arbitrary, only consistency and division
// into equivalence classes (Types that compare CMPeq) matters.
func (t *Type) Compare(x *Type) Cmp

// IsKind reports whether t is a Type of the specified kind.
func (t *Type) IsKind(et Kind) bool

func (t *Type) IsBoolean() bool

// ToUnsigned returns the unsigned equivalent of integer type t.
func (t *Type) ToUnsigned() *Type

func (t *Type) IsInteger() bool

func (t *Type) IsSigned() bool

func (t *Type) IsUnsigned() bool

func (t *Type) IsFloat() bool

func (t *Type) IsComplex() bool

// IsPtr reports whether t is a regular Go pointer type.
// This does not include unsafe.Pointer.
func (t *Type) IsPtr() bool

// IsPtrElem reports whether t is the element of a pointer (to t).
func (t *Type) IsPtrElem() bool

// IsUnsafePtr reports whether t is an unsafe pointer.
func (t *Type) IsUnsafePtr() bool

// IsUintptr reports whether t is a uintptr.
func (t *Type) IsUintptr() bool

// IsPtrShaped reports whether t is represented by a single machine pointer.
// In addition to regular Go pointer types, this includes map, channel, and
// function types and unsafe.Pointer. It does not include array or struct types
// that consist of a single pointer shaped type.
// TODO(mdempsky): Should it? See golang.org/issue/15028.
func (t *Type) IsPtrShaped() bool

// HasNil reports whether the set of values determined by t includes nil.
func (t *Type) HasNil() bool

func (t *Type) IsString() bool

func (t *Type) IsMap() bool

func (t *Type) IsChan() bool

func (t *Type) IsSlice() bool

func (t *Type) IsArray() bool

func (t *Type) IsStruct() bool

func (t *Type) IsInterface() bool

// IsEmptyInterface reports whether t is an empty interface type.
func (t *Type) IsEmptyInterface() bool

// IsScalar reports whether 't' is a scalar Go type, e.g.
// bool/int/float/complex. Note that struct and array types consisting
// of a single scalar element are not considered scalar, likewise
// pointer types are also not considered scalar.
func (t *Type) IsScalar() bool

func (t *Type) PtrTo() *Type

func (t *Type) NumFields() int

func (t *Type) FieldType(i int) *Type

func (t *Type) FieldOff(i int) int64

func (t *Type) FieldName(i int) string

func (t *Type) NumElem() int64

const (
	IgnoreBlankFields componentsIncludeBlankFields = false
	CountBlankFields  componentsIncludeBlankFields = true
)

// NumComponents returns the number of primitive elements that compose t.
// Struct and array types are flattened for the purpose of counting.
// All other types (including string, slice, and interface types) count as one element.
// If countBlank is IgnoreBlankFields, then blank struct fields
// (and their comprised elements) are excluded from the count.
// struct { x, y [3]int } has six components; [10]struct{ x, y string } has twenty.
func (t *Type) NumComponents(countBlank componentsIncludeBlankFields) int64

// SoleComponent returns the only primitive component in t,
// if there is exactly one. Otherwise, it returns nil.
// Components are counted as in NumComponents, including blank fields.
// Keep in sync with cmd/compile/internal/walk/convert.go:soleComponent.
func (t *Type) SoleComponent() *Type

// ChanDir returns the direction of a channel type t.
// The direction will be one of Crecv, Csend, or Cboth.
func (t *Type) ChanDir() ChanDir

func (t *Type) IsMemory() bool

func (t *Type) IsFlags() bool
func (t *Type) IsVoid() bool
func (t *Type) IsTuple() bool
func (t *Type) IsResults() bool

// IsUntyped reports whether t is an untyped type.
func (t *Type) IsUntyped() bool

// HasPointers reports whether t contains a heap pointer.
// Note that this function ignores pointers to not-in-heap types.
func (t *Type) HasPointers() bool

// FakeRecvType returns the singleton type used for interface method receivers.
func FakeRecvType() *Type

func FakeRecv() *Field

var (
	// TSSA types. HasPointers assumes these are pointer-free.
	TypeInvalid   = newSSA("invalid")
	TypeMem       = newSSA("mem")
	TypeFlags     = newSSA("flags")
	TypeVoid      = newSSA("void")
	TypeInt128    = newSSA("int128")
	TypeResultMem = newResults([]*Type{TypeMem})
)

// NewNamed returns a new named type for the given type name. obj should be an
// ir.Name. The new type is incomplete (marked as TFORW kind), and the underlying
// type should be set later via SetUnderlying(). References to the type are
// maintained until the type is filled in, so those references can be updated when
// the type is complete.
func NewNamed(obj Object) *Type

// Obj returns the canonical type name node for a named type t, nil for an unnamed type.
func (t *Type) Obj() Object

// SetVargen assigns a unique generation number to type t, which must
// be a defined type declared within function scope. The generation
// number is used to distinguish it from other similarly spelled
// defined types from the same package.
//
// TODO(mdempsky): Come up with a better solution.
func (t *Type) SetVargen()

// SetUnderlying sets the underlying type of an incomplete type (i.e. type whose kind
// is currently TFORW). SetUnderlying automatically updates any types that were waiting
// for this type to be completed.
func (t *Type) SetUnderlying(underlying *Type)

// NewInterface returns a new interface for the given methods and
// embedded types. Embedded types are specified as fields with no Sym.
func NewInterface(methods []*Field) *Type

const BOGUS_FUNARG_OFFSET = -1000000000

// NewSignature returns a new function type for the given receiver,
// parameters, and results, any of which may be nil.
func NewSignature(recv *Field, params, results []*Field) *Type

// NewStruct returns a new struct with the given fields.
func NewStruct(fields []*Field) *Type

var (
	IsInt     [NTYPE]bool
	IsFloat   [NTYPE]bool
	IsComplex [NTYPE]bool
	IsSimple  [NTYPE]bool
)

var IsOrdered [NTYPE]bool

// IsReflexive reports whether t has a reflexive equality operator.
// That is, if x==x for all x of type t.
func IsReflexive(t *Type) bool

// Can this type be stored directly in an interface word?
// Yes, if the representation is a single pointer.
func IsDirectIface(t *Type) bool

// IsInterfaceMethod reports whether (field) m is
// an interface method. Such methods have the
// special receiver type types.FakeRecvType().
func IsInterfaceMethod(f *Type) bool

// IsMethodApplicable reports whether method m can be called on a
// value of type t. This is necessary because we compute a single
// method set for both T and *T, but some *T methods are not
// applicable to T receivers.
func IsMethodApplicable(t *Type, m *Field) bool

// IsRuntimePkg reports whether p is package runtime.
func IsRuntimePkg(p *Pkg) bool

// IsReflectPkg reports whether p is package reflect.
func IsReflectPkg(p *Pkg) bool

// IsTypePkg reports whether p is pesudo package type.
func IsTypePkg(p *Pkg) bool

// IsNoInstrumentPkg reports whether p is a package that
// should not be instrumented.
func IsNoInstrumentPkg(p *Pkg) bool

// IsNoRacePkg reports whether p is a package that
// should not be race instrumented.
func IsNoRacePkg(p *Pkg) bool

// ReceiverBaseType returns the underlying type, if any,
// that owns methods with receiver parameter t.
// The result is either a named type or an anonymous struct.
func ReceiverBaseType(t *Type) *Type

func FloatForComplex(t *Type) *Type

func ComplexForFloat(t *Type) *Type

func TypeSym(t *Type) *Sym

func TypeSymLookup(name string) *Sym

func TypeSymName(t *Type) string

var SimType [NTYPE]Kind

// Fake package for shape types (see typecheck.Shapify()).
var ShapePkg = NewPkg("go.shape", "go.shape")
