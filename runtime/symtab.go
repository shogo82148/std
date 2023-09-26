// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Frames may be used to get function/file/line information for a
// slice of PC values returned by Callers.
type Frames struct {
	callers []uintptr

	stackExpander stackExpander

	elideWrapper bool
}

// Frame is the information returned by Frames for each call frame.
type Frame struct {
	PC uintptr

	Func *Func

	Function string

	File string
	Line int

	Entry uintptr
}

// stackExpander expands a call stack of PCs into a sequence of
// Frames. It tracks state across PCs necessary to perform this
// expansion.
//
// This is the core of the Frames implementation, but is a separate
// internal API to make it possible to use within the runtime without
// heap-allocating the PC slice. The only difference with the public
// Frames API is that the caller is responsible for threading the PC
// slice between expansion steps in this API. If escape analysis were
// smarter, we may not need this (though it may have to be a lot
// smarter).

// CallersFrames takes a slice of PC values returned by Callers and
// prepares to return function/file/line information.
// Do not change the slice until you are done with the Frames.
func CallersFrames(callers []uintptr) *Frames

// Next returns frame information for the next caller.
// If more is false, there are no more callers (the Frame value is valid).
func (ci *Frames) Next() (frame Frame, more bool)

// A pcExpander expands a single PC into a sequence of Frames.

// A Func represents a Go function in the running binary.
type Func struct {
	opaque struct{}
}

// PCDATA and FUNCDATA table indexes.
//
// See funcdata.h and ../cmd/internal/objabi/funcdata.go.

// A FuncID identifies particular functions that need to be treated
// specially by the runtime.
// Note that in some situations involving plugins, there may be multiple
// copies of a particular special runtime function.
// Note: this list must match the list in cmd/internal/objabi/funcid.go.

// moduledata records information about the layout of the executable
// image. It is written by the linker. Any changes here must be
// matched changes to the code in cmd/internal/ld/symtab.go:symtab.
// moduledata is stored in statically allocated non-pointer memory;
// none of the pointers here are visible to the garbage collector.

// A modulehash is used to compare the ABI of a new module or a
// package in a new module with the loaded program.
//
// For each shared library a module links against, the linker creates an entry in the
// moduledata.modulehashes slice containing the name of the module, the abi hash seen
// at link time and a pointer to the runtime abi hash. These are checked in
// moduledataverify1 below.
//
// For each loaded plugin, the pkghashes slice has a modulehash of the
// newly loaded package that can be used to check the plugin's version of
// a package against any previously loaded version of the package.
// This is done in plugin.lastmoduleinit.

// pinnedTypemaps are the map[typeOff]*_type from the moduledata objects.
//
// These typemap objects are allocated at run time on the heap, but the
// only direct reference to them is in the moduledata, created by the
// linker and marked SNOPTRDATA so it is ignored by the GC.
//
// To make sure the map isn't collected, we keep a second reference here.

// findfunctab is an array of these structures.
// Each bucket represents 4096 bytes of the text segment.
// Each subbucket represents 256 bytes of the text segment.
// To find a function given a pc, locate the bucket and subbucket for
// that pc. Add together the idx and subbucket value to obtain a
// function index. Then scan the functab array starting at that
// index to find the target function.
// This table uses 20 bytes for every 4096 bytes of code, or ~0.5% overhead.

// FuncForPC returns a *Func describing the function that contains the
// given program counter address, or else nil.
//
// If pc represents multiple functions because of inlining, it returns
// the *Func describing the outermost function.
func FuncForPC(pc uintptr) *Func

// Name returns the name of the function.
func (f *Func) Name() string

// Entry returns the entry address of the function.
func (f *Func) Entry() uintptr

// FileLine returns the file name and line number of the
// source code corresponding to the program counter pc.
// The result will not be accurate if pc is not a program
// counter within f.
func (f *Func) FileLine(pc uintptr) (file string, line int)

// inlinedCall is the encoding of entries in the FUNCDATA_InlTree table.
