// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// Map predeclared Go types to Type.

// Prologue defining TSAN functions in C.

// This must match the TSAN code in runtime/cgo/libcgo.h.
// This is used when the code is built with the C/C++ Thread SANitizer,
// which is not the same as the Go race detector.
// __tsan_acquire tells TSAN that we are acquiring a lock on a variable,
// in this case _cgo_sync. __tsan_release releases the lock.
// (There is no actual lock, we are just telling TSAN that there is.)
//
// When we call from Go to C we call _cgo_tsan_acquire.
// When the C function returns we call _cgo_tsan_release.
// Similarly, when C calls back into Go we call _cgo_tsan_release
// and then call _cgo_tsan_acquire when we return to C.
// These calls tell TSAN that there is a serialization point at the C call.
//
// This is necessary because TSAN, which is a C/C++ tool, can not see
// the synchronization in the Go code. Without these calls, when
// multiple goroutines call into C code, TSAN does not understand
// that the calls are properly synchronized on the Go side.
//
// To be clear, if the calls are not properly synchronized on the Go side,
// we will be hiding races. But when using TSAN on mixed Go C/C++ code
// it is more important to avoid false positives, which reduce confidence
// in the tool, than to avoid false negatives.

// Set to yesTsanProlog if we see -fsanitize=thread in the flags for gcc.

// cMallocDefC defines the C version of C.malloc for the gc compiler.
// It is defined here because C.CString and friends need a definition.
// We define it by hand, rather than simply inventing a reference to
// C.malloc, because <stdlib.h> may not have been included.
// This is approximately what writeOutputFunc would generate, but
// skips the cgo_topofstack code (which is only needed if the C code
// calls back into Go). This also avoids returning nil for an
// allocation of 0 bytes.

// builtinExportProlog is a shorter version of builtinProlog,
// to be put into the _cgo_export.h file.
// For historical reasons we can't use builtinProlog in _cgo_export.h,
// because _cgo_export.h defines GoString as a struct while builtinProlog
// defines it as a function. We don't change this to avoid unnecessarily
// breaking existing code.

// gccExportHeaderEpilog goes at the end of the generated header file.

// gccgoExportFileProlog is written to the _cgo_export.c file when
// using gccgo.
// We use weak declarations, and test the addresses, so that this code
// works with older versions of gccgo.
