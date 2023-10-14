// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package loadpe implements a PE/COFF file reader.
package loadpe

import (
	"github.com/shogo82148/std/cmd/internal/bio"
	"github.com/shogo82148/std/cmd/internal/sys"
	"github.com/shogo82148/std/cmd/link/internal/loader"
)

const (
	IMAGE_SYM_UNDEFINED              = 0
	IMAGE_SYM_ABSOLUTE               = -1
	IMAGE_SYM_DEBUG                  = -2
	IMAGE_SYM_TYPE_NULL              = 0
	IMAGE_SYM_TYPE_VOID              = 1
	IMAGE_SYM_TYPE_CHAR              = 2
	IMAGE_SYM_TYPE_SHORT             = 3
	IMAGE_SYM_TYPE_INT               = 4
	IMAGE_SYM_TYPE_LONG              = 5
	IMAGE_SYM_TYPE_FLOAT             = 6
	IMAGE_SYM_TYPE_DOUBLE            = 7
	IMAGE_SYM_TYPE_STRUCT            = 8
	IMAGE_SYM_TYPE_UNION             = 9
	IMAGE_SYM_TYPE_ENUM              = 10
	IMAGE_SYM_TYPE_MOE               = 11
	IMAGE_SYM_TYPE_BYTE              = 12
	IMAGE_SYM_TYPE_WORD              = 13
	IMAGE_SYM_TYPE_UINT              = 14
	IMAGE_SYM_TYPE_DWORD             = 15
	IMAGE_SYM_TYPE_PCODE             = 32768
	IMAGE_SYM_DTYPE_NULL             = 0
	IMAGE_SYM_DTYPE_POINTER          = 1
	IMAGE_SYM_DTYPE_FUNCTION         = 2
	IMAGE_SYM_DTYPE_ARRAY            = 3
	IMAGE_SYM_CLASS_END_OF_FUNCTION  = -1
	IMAGE_SYM_CLASS_NULL             = 0
	IMAGE_SYM_CLASS_AUTOMATIC        = 1
	IMAGE_SYM_CLASS_EXTERNAL         = 2
	IMAGE_SYM_CLASS_STATIC           = 3
	IMAGE_SYM_CLASS_REGISTER         = 4
	IMAGE_SYM_CLASS_EXTERNAL_DEF     = 5
	IMAGE_SYM_CLASS_LABEL            = 6
	IMAGE_SYM_CLASS_UNDEFINED_LABEL  = 7
	IMAGE_SYM_CLASS_MEMBER_OF_STRUCT = 8
	IMAGE_SYM_CLASS_ARGUMENT         = 9
	IMAGE_SYM_CLASS_STRUCT_TAG       = 10
	IMAGE_SYM_CLASS_MEMBER_OF_UNION  = 11
	IMAGE_SYM_CLASS_UNION_TAG        = 12
	IMAGE_SYM_CLASS_TYPE_DEFINITION  = 13
	IMAGE_SYM_CLASS_UNDEFINED_STATIC = 14
	IMAGE_SYM_CLASS_ENUM_TAG         = 15
	IMAGE_SYM_CLASS_MEMBER_OF_ENUM   = 16
	IMAGE_SYM_CLASS_REGISTER_PARAM   = 17
	IMAGE_SYM_CLASS_BIT_FIELD        = 18
	IMAGE_SYM_CLASS_FAR_EXTERNAL     = 68
	IMAGE_SYM_CLASS_BLOCK            = 100
	IMAGE_SYM_CLASS_FUNCTION         = 101
	IMAGE_SYM_CLASS_END_OF_STRUCT    = 102
	IMAGE_SYM_CLASS_FILE             = 103
	IMAGE_SYM_CLASS_SECTION          = 104
	IMAGE_SYM_CLASS_WEAK_EXTERNAL    = 105
	IMAGE_SYM_CLASS_CLR_TOKEN        = 107
	IMAGE_REL_I386_ABSOLUTE          = 0x0000
	IMAGE_REL_I386_DIR16             = 0x0001
	IMAGE_REL_I386_REL16             = 0x0002
	IMAGE_REL_I386_DIR32             = 0x0006
	IMAGE_REL_I386_DIR32NB           = 0x0007
	IMAGE_REL_I386_SEG12             = 0x0009
	IMAGE_REL_I386_SECTION           = 0x000A
	IMAGE_REL_I386_SECREL            = 0x000B
	IMAGE_REL_I386_TOKEN             = 0x000C
	IMAGE_REL_I386_SECREL7           = 0x000D
	IMAGE_REL_I386_REL32             = 0x0014
	IMAGE_REL_AMD64_ABSOLUTE         = 0x0000
	IMAGE_REL_AMD64_ADDR64           = 0x0001
	IMAGE_REL_AMD64_ADDR32           = 0x0002
	IMAGE_REL_AMD64_ADDR32NB         = 0x0003
	IMAGE_REL_AMD64_REL32            = 0x0004
	IMAGE_REL_AMD64_REL32_1          = 0x0005
	IMAGE_REL_AMD64_REL32_2          = 0x0006
	IMAGE_REL_AMD64_REL32_3          = 0x0007
	IMAGE_REL_AMD64_REL32_4          = 0x0008
	IMAGE_REL_AMD64_REL32_5          = 0x0009
	IMAGE_REL_AMD64_SECTION          = 0x000A
	IMAGE_REL_AMD64_SECREL           = 0x000B
	IMAGE_REL_AMD64_SECREL7          = 0x000C
	IMAGE_REL_AMD64_TOKEN            = 0x000D
	IMAGE_REL_AMD64_SREL32           = 0x000E
	IMAGE_REL_AMD64_PAIR             = 0x000F
	IMAGE_REL_AMD64_SSPAN32          = 0x0010
	IMAGE_REL_ARM_ABSOLUTE           = 0x0000
	IMAGE_REL_ARM_ADDR32             = 0x0001
	IMAGE_REL_ARM_ADDR32NB           = 0x0002
	IMAGE_REL_ARM_BRANCH24           = 0x0003
	IMAGE_REL_ARM_BRANCH11           = 0x0004
	IMAGE_REL_ARM_SECTION            = 0x000E
	IMAGE_REL_ARM_SECREL             = 0x000F
	IMAGE_REL_ARM_MOV32              = 0x0010
	IMAGE_REL_THUMB_MOV32            = 0x0011
	IMAGE_REL_THUMB_BRANCH20         = 0x0012
	IMAGE_REL_THUMB_BRANCH24         = 0x0014
	IMAGE_REL_THUMB_BLX23            = 0x0015
	IMAGE_REL_ARM_PAIR               = 0x0016
	IMAGE_REL_ARM64_ABSOLUTE         = 0x0000
	IMAGE_REL_ARM64_ADDR32           = 0x0001
	IMAGE_REL_ARM64_ADDR32NB         = 0x0002
	IMAGE_REL_ARM64_BRANCH26         = 0x0003
	IMAGE_REL_ARM64_PAGEBASE_REL21   = 0x0004
	IMAGE_REL_ARM64_REL21            = 0x0005
	IMAGE_REL_ARM64_PAGEOFFSET_12A   = 0x0006
	IMAGE_REL_ARM64_PAGEOFFSET_12L   = 0x0007
	IMAGE_REL_ARM64_SECREL           = 0x0008
	IMAGE_REL_ARM64_SECREL_LOW12A    = 0x0009
	IMAGE_REL_ARM64_SECREL_HIGH12A   = 0x000A
	IMAGE_REL_ARM64_SECREL_LOW12L    = 0x000B
	IMAGE_REL_ARM64_TOKEN            = 0x000C
	IMAGE_REL_ARM64_SECTION          = 0x000D
	IMAGE_REL_ARM64_ADDR64           = 0x000E
	IMAGE_REL_ARM64_BRANCH19         = 0x000F
	IMAGE_REL_ARM64_BRANCH14         = 0x0010
	IMAGE_REL_ARM64_REL32            = 0x0011
)

const (
	// When stored into the PLT value for a symbol, this token tells
	// windynrelocsym to redirect direct references to this symbol to a stub
	// that loads from the corresponding import symbol and then does
	// a jump to the loaded value.
	CreateImportStubPltToken = -2

	// When stored into the GOT value for an import symbol __imp_X this
	// token tells windynrelocsym to redirect references to the
	// underlying DYNIMPORT symbol X.
	RedirectToDynImportGotToken = -2
)

// Load loads the PE file pn from input.
// Symbols from the object file are created via the loader 'l',
// and a slice of the text symbols is returned.
// If an .rsrc section or set of .rsrc$xx sections is found, its symbols are
// returned as rsrc.
func Load(l *loader.Loader, arch *sys.Arch, localSymVersion int, input *bio.Reader, pkg string, length int64, pn string) (textp []loader.Sym, rsrc []loader.Sym, err error)

// PostProcessImports works to resolve inconsistencies with DLL import
// symbols; it is needed when building with more "modern" C compilers
// with internal linkage.
//
// Background: DLL import symbols are data (SNOPTRDATA) symbols whose
// name is of the form "__imp_XXX", which contain a pointer/reference
// to symbol XXX. It's possible to have import symbols for both data
// symbols ("__imp__fmode") and text symbols ("__imp_CreateEventA").
// In some case import symbols are just references to some external
// thing, and in other cases we see actual definitions of import
// symbols when reading host objects.
//
// Previous versions of the linker would in most cases immediately
// "forward" import symbol references, e.g. treat a references to
// "__imp_XXX" a references to "XXX", however this doesn't work well
// with more modern compilers, where you can sometimes see import
// symbols that are defs (as opposed to external refs).
//
// The main actions taken below are to search for references to
// SDYNIMPORT symbols in host object text/data sections and flag the
// symbols for later fixup. When we see a reference to an import
// symbol __imp_XYZ where XYZ corresponds to some SDYNIMPORT symbol,
// we flag the symbol (via GOT setting) so that it can be redirected
// to XYZ later in windynrelocsym. When we see a direct reference to
// an SDYNIMPORT symbol XYZ, we also flag the symbol (via PLT setting)
// to indicated that the reference will need to be redirected to a
// stub.
func PostProcessImports() error

// LookupBaseFromImport examines the symbol "s" to see if it
// corresponds to an import symbol (name of the form "__imp_XYZ") and
// if so, it looks up the underlying target of the import symbol and
// returns it. An error is returned if the symbol is of the form
// "__imp_XYZ" but no XYZ can be found.
func LookupBaseFromImport(s loader.Sym, ldr *loader.Loader, arch *sys.Arch) (loader.Sym, error)
