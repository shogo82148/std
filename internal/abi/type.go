// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package abi

import (
	"github.com/shogo82148/std/unsafe"
)

// Type is the runtime representation of a Go type.
//
// Be careful about accessing this type at build time, as the version
// of this type in the compiler/linker may not have the same layout
// as the version in the target binary, due to pointer width
// differences and any experiments. Use cmd/compile/internal/rttype
// or the functions in compiletype.go to access this type instead.
// (TODO: this admonition applies to every type in this package.
// Put it in some shared location?)
type Type struct {
	Size_       uintptr
	PtrBytes    uintptr
	Hash        uint32
	TFlag       TFlag
	Align_      uint8
	FieldAlign_ uint8
	Kind_       Kind
	// function for comparing objects of this type
	// (ptr to object A, ptr to object B) -> ==?
	Equal func(unsafe.Pointer, unsafe.Pointer) bool
	// GCData stores the GC type data for the garbage collector.
	// Normally, GCData points to a bitmask that describes the
	// ptr/nonptr fields of the type. The bitmask will have at
	// least PtrBytes/ptrSize bits.
	// If the TFlagGCMaskOnDemand bit is set, GCData is instead a
	// **byte and the pointer to the bitmask is one dereference away.
	// The runtime will build the bitmask if needed.
	// (See runtime/type.go:getGCMask.)
	// Note: multiple types may have the same value of GCData,
	// including when TFlagGCMaskOnDemand is set. The types will, of course,
	// have the same pointer layout (but not necessarily the same size).
	GCData    *byte
	Str       NameOff
	PtrToThis TypeOff
}

// A Kind represents the specific kind of type that a Type represents.
// The zero Kind is not a valid kind.
type Kind uint8

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Pointer
	Slice
	String
	Struct
	UnsafePointer
)

// TFlag is used by a Type to signal what extra type information is
// available in the memory directly following the Type value.
type TFlag uint8

const (
	// TFlagUncommon means that there is a data with a type, UncommonType,
	// just beyond the shared-per-type common data.  That is, the data
	// for struct types will store their UncommonType at one offset, the
	// data for interface types will store their UncommonType at a different
	// offset.  UncommonType is always accessed via a pointer that is computed
	// using trust-us-we-are-the-implementors pointer arithmetic.
	//
	// For example, if t.Kind() == Struct and t.tflag&TFlagUncommon != 0,
	// then t has UncommonType data and it can be accessed as:
	//
	//	type structTypeUncommon struct {
	//		structType
	//		u UncommonType
	//	}
	//	u := &(*structTypeUncommon)(unsafe.Pointer(t)).u
	TFlagUncommon TFlag = 1 << 0

	// TFlagExtraStar means the name in the str field has an
	// extraneous '*' prefix. This is because for most types T in
	// a program, the type *T also exists and reusing the str data
	// saves binary size.
	TFlagExtraStar TFlag = 1 << 1

	// TFlagNamed means the type has a name.
	TFlagNamed TFlag = 1 << 2

	// TFlagRegularMemory means that equal and hash functions can treat
	// this type as a single region of t.size bytes.
	TFlagRegularMemory TFlag = 1 << 3

	// TFlagGCMaskOnDemand means that the GC pointer bitmask will be
	// computed on demand at runtime instead of being precomputed at
	// compile time. If this flag is set, the GCData field effectively
	// has type **byte instead of *byte. The runtime will store a
	// pointer to the GC pointer bitmask in *GCData.
	TFlagGCMaskOnDemand TFlag = 1 << 4

	// TFlagDirectIface means that a value of this type is stored directly
	// in the data field of an interface, instead of indirectly. Normally
	// this means the type is pointer-ish.
	TFlagDirectIface TFlag = 1 << 5

	// Leaving this breadcrumb behind for dlv. It should not be used, and no
	// Kind should be big enough to set this bit.
	KindDirectIface Kind = 1 << 5
)

// NameOff is the offset to a name from moduledata.types.  See resolveNameOff in runtime.
type NameOff int32

// TypeOff is the offset to a type from moduledata.types.  See resolveTypeOff in runtime.
type TypeOff int32

// TextOff is an offset from the top of a text section.  See (rtype).textOff in runtime.
type TextOff int32

// String returns the name of k.
func (k Kind) String() string

// TypeOf returns the abi.Type of some value.
func TypeOf(a any) *Type

// TypeFor returns the abi.Type for a type parameter.
func TypeFor[T any]() *Type

func (t *Type) Kind() Kind

func (t *Type) HasName() bool

// Pointers reports whether t contains pointers.
func (t *Type) Pointers() bool

// IsDirectIface reports whether t is stored directly in an interface value.
func (t *Type) IsDirectIface() bool

func (t *Type) GcSlice(begin, end uintptr) []byte

// Method on non-interface type
type Method struct {
	Name NameOff
	Mtyp TypeOff
	Ifn  TextOff
	Tfn  TextOff
}

// UncommonType is present only for defined types or types with methods
// (if T is a defined type, the uncommonTypes for T and *T have methods).
// Using a pointer to this struct reduces the overall size required
// to describe a non-defined type with no methods.
type UncommonType struct {
	PkgPath NameOff
	Mcount  uint16
	Xcount  uint16
	Moff    uint32
	_       uint32
}

func (t *UncommonType) Methods() []Method

func (t *UncommonType) ExportedMethods() []Method

// Imethod represents a method on an interface type
type Imethod struct {
	Name NameOff
	Typ  TypeOff
}

// ArrayType represents a fixed array type.
type ArrayType struct {
	Type
	Elem  *Type
	Slice *Type
	Len   uintptr
}

// Len returns the length of t if t is an array type, otherwise 0
func (t *Type) Len() int

func (t *Type) Common() *Type

type ChanDir int

const (
	RecvDir ChanDir = 1 << iota
	SendDir
	BothDir            = RecvDir | SendDir
	InvalidDir ChanDir = 0
)

// ChanType represents a channel type
type ChanType struct {
	Type
	Elem *Type
	Dir  ChanDir
}

// ChanDir returns the direction of t if t is a channel type, otherwise InvalidDir (0).
func (t *Type) ChanDir() ChanDir

// Uncommon returns a pointer to T's "uncommon" data if there is any, otherwise nil
func (t *Type) Uncommon() *UncommonType

// Elem returns the element type for t if t is an array, channel, map, pointer, or slice, otherwise nil.
func (t *Type) Elem() *Type

// StructType returns t cast to a *StructType, or nil if its tag does not match.
func (t *Type) StructType() *StructType

// MapType returns t cast to a *OldMapType or *SwissMapType, or nil if its tag does not match.
func (t *Type) MapType() *mapType

// ArrayType returns t cast to a *ArrayType, or nil if its tag does not match.
func (t *Type) ArrayType() *ArrayType

// FuncType returns t cast to a *FuncType, or nil if its tag does not match.
func (t *Type) FuncType() *FuncType

// InterfaceType returns t cast to a *InterfaceType, or nil if its tag does not match.
func (t *Type) InterfaceType() *InterfaceType

// Size returns the size of data with type t.
func (t *Type) Size() uintptr

// Align returns the alignment of data with type t.
func (t *Type) Align() int

func (t *Type) FieldAlign() int

type InterfaceType struct {
	Type
	PkgPath Name
	Methods []Imethod
}

func (t *Type) ExportedMethods() []Method

func (t *Type) NumMethod() int

// NumMethod returns the number of interface methods in the type's method set.
func (t *InterfaceType) NumMethod() int

func (t *Type) Key() *Type

type SliceType struct {
	Type
	Elem *Type
}

// FuncType represents a function type.
//
// A *Type for each in and out parameter is stored in an array that
// directly follows the funcType (and possibly its uncommonType). So
// a function type with one method, one input, and one output is:
//
//	struct {
//		funcType
//		uncommonType
//		[2]*rtype    // [0] is in, [1] is out
//	}
type FuncType struct {
	Type
	InCount  uint16
	OutCount uint16
}

func (t *FuncType) In(i int) *Type

func (t *FuncType) NumIn() int

func (t *FuncType) NumOut() int

func (t *FuncType) Out(i int) *Type

func (t *FuncType) InSlice() []*Type

func (t *FuncType) OutSlice() []*Type

func (t *FuncType) IsVariadic() bool

type PtrType struct {
	Type
	Elem *Type
}

type StructField struct {
	Name   Name
	Typ    *Type
	Offset uintptr
}

func (f *StructField) Embedded() bool

type StructType struct {
	Type
	PkgPath Name
	Fields  []StructField
}

type Name struct {
	Bytes *byte
}

// DataChecked does pointer arithmetic on n's Bytes, and that arithmetic is asserted to
// be safe for the reason in whySafe (which can appear in a backtrace, etc.)
func (n Name) DataChecked(off int, whySafe string) *byte

// Data does pointer arithmetic on n's Bytes, and that arithmetic is asserted to
// be safe because the runtime made the call (other packages use DataChecked)
func (n Name) Data(off int) *byte

// IsExported returns "is n exported?"
func (n Name) IsExported() bool

// HasTag returns true iff there is tag data following this name
func (n Name) HasTag() bool

// IsEmbedded returns true iff n is embedded (an anonymous field).
func (n Name) IsEmbedded() bool

// ReadVarint parses a varint as encoded by encoding/binary.
// It returns the number of encoded bytes and the encoded value.
func (n Name) ReadVarint(off int) (int, int)

// IsBlank indicates whether n is "_".
func (n Name) IsBlank() bool

// Name returns the tag string for n, or empty if there is none.
func (n Name) Name() string

// Tag returns the tag string for n, or empty if there is none.
func (n Name) Tag() string

func NewName(n, tag string, exported, embedded bool) Name

const (
	TraceArgsLimit    = 10
	TraceArgsMaxDepth = 5

	// maxLen is a (conservative) upper bound of the byte stream length. For
	// each arg/component, it has no more than 2 bytes of data (size, offset),
	// and no more than one {, }, ... at each level (it cannot have both the
	// data and ... unless it is the last one, just be conservative). Plus 1
	// for _endSeq.
	TraceArgsMaxLen = (TraceArgsMaxDepth*3+2)*TraceArgsLimit + 1
)

// Populate the data.
// The data is a stream of bytes, which contains the offsets and sizes of the
// non-aggregate arguments or non-aggregate fields/elements of aggregate-typed
// arguments, along with special "operators". Specifically,
//   - for each non-aggregate arg/field/element, its offset from FP (1 byte) and
//     size (1 byte)
//   - special operators:
//   - 0xff - end of sequence
//   - 0xfe - print { (at the start of an aggregate-typed argument)
//   - 0xfd - print } (at the end of an aggregate-typed argument)
//   - 0xfc - print ... (more args/fields/elements)
//   - 0xfb - print _ (offset too large)
const (
	TraceArgsEndSeq         = 0xff
	TraceArgsStartAgg       = 0xfe
	TraceArgsEndAgg         = 0xfd
	TraceArgsDotdotdot      = 0xfc
	TraceArgsOffsetTooLarge = 0xfb
	TraceArgsSpecial        = 0xf0
)

// MaxPtrmaskBytes is the maximum length of a GC ptrmask bitmap,
// which holds 1-bit entries describing where pointers are in a given type.
// Above this length, the GC information is recorded as a GC program,
// which can express repetition compactly. In either form, the
// information is used by the runtime to initialize the heap bitmap,
// and for large types (like 128 or more words), they are roughly the
// same speed. GC programs are never much larger and often more
// compact. (If large arrays are involved, they can be arbitrarily
// more compact.)
//
// The cutoff must be large enough that any allocation large enough to
// use a GC program is large enough that it does not share heap bitmap
// bytes with any other objects, allowing the GC program execution to
// assume an aligned start and not use atomic operations. In the current
// runtime, this means all malloc size classes larger than the cutoff must
// be multiples of four words. On 32-bit systems that's 16 bytes, and
// all size classes >= 16 bytes are 16-byte aligned, so no real constraint.
// On 64-bit systems, that's 32 bytes, and 32-byte alignment is guaranteed
// for size classes >= 256 bytes. On a 64-bit system, 256 bytes allocated
// is 32 pointers, the bits for which fit in 4 bytes. So MaxPtrmaskBytes
// must be >= 4.
//
// We used to use 16 because the GC programs do have some constant overhead
// to get started, and processing 128 pointers seems to be enough to
// amortize that overhead well.
//
// To make sure that the runtime's chansend can call typeBitsBulkBarrier,
// we raised the limit to 2048, so that even 32-bit systems are guaranteed to
// use bitmaps for objects up to 64 kB in size.
const MaxPtrmaskBytes = 2048
