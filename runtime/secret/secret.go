// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.runtimesecret

package secret

// Do invokes f.
//
// Do ensures that any temporary storage used by f is erased in a
// timely manner. (In this context, "f" is shorthand for the
// entire call tree initiated by f.)
//   - Any registers used by f are erased before Do returns.
//   - Any stack used by f is erased before Do returns.
//   - Heap allocations done by f are erased as soon as the garbage
//     collector realizes that all allocated values are no longer reachable.
//   - Do works even if f panics or calls runtime.Goexit.  As part of
//     that, any panic raised by f will appear as if it originates from
//     Do itself.
//
// Users should be cautious of allocating inside Do.
// Erasing heap memory after Do returns may increase garbage collector sweep times and
// requires additional memory to keep track of allocations until they are to be erased.
// These costs can compound when an allocation is done in the service of growing a value,
// like appending to a slice or inserting into a map. In these cases, the entire new allocation is erased rather
// than just the secret parts of it.
//
// To reduce lifetimes of allocations and avoid unexpected performance issues,
// if a function invoked by Do needs to yield a result that shouldn't be erased,
// it should do so by copying the result into an allocation created by the caller.
//
// Limitations:
//   - Currently only supported on linux/amd64 and linux/arm64.  On unsupported
//     platforms, Do will invoke f directly.
//   - Protection does not extend to any global variables written by f.
//   - Protection does not extend to any new goroutines made by f.
//   - If f calls runtime.Goexit, erasure can be delayed by defers
//     higher up on the call stack.
//   - Heap allocations will only be erased if the program drops all
//     references to those allocations, and then the garbage collector
//     notices that those references are gone. The former is under
//     control of the program, but the latter is at the whim of the
//     runtime.
//   - Any value panicked by f may point to allocations from within
//     f. Those allocations will not be erased until (at least) the
//     panicked value is dead.
//   - Pointer addresses may leak into data buffers used by the runtime
//     to perform garbage collection. Users should not encode confidential
//     information into pointers. For example, if an offset into an array or
//     struct is confidential, then users should not create a pointer into
//     the object. Since this function is intended to be used with constant-time
//     cryptographic code, this requirement is usually fulfilled implicitly.
func Do(f func())

// Enabled reports whether [Do] appears anywhere on the call stack.
func Enabled() bool
