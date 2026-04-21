// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// simdgen is an experiment in generating Go <-> asm SIMD mappings.
//
// Usage: simdgen [-xedPath=path] [-q=query] input.yaml...
//
// If -xedPath is provided, one of the inputs is a sum of op-code definitions
// generated from the Intel XED data at path.
//
// If input YAML files are provided, each file is read as an input value. See
// [unify.Closure.UnmarshalYAML] or "go doc unify.Closure.UnmarshalYAML" for the
// format of these files.
//
// TODO: Example definitions and values.
//
// The command unifies across all of the inputs and prints all possible results
// of this unification.
//
// If the -q flag is provided, its string value is parsed as a value and treated
// as another input to unification. This is intended as a way to "query" the
// result, typically by narrowing it down to a small subset of results.
//
// Typical usage:
//
//	go run . -xedPath $XEDPATH *.yaml
//
// To see just the definitions generated from XED, run:
//
//	go run . -xedPath $XEDPATH
//
// (This works because if there's only one input, there's nothing to unify it
// with, so the result is simply itself.)
//
// To see just the definitions for VPADDQ:
//
//	go run . -xedPath $XEDPATH -q '{asm: VPADDQ}'
//
// simdgen can also generate Go definitions of SIMD mappings:
// To generate go files to the go root, run:
//
//	go run . -xedPath $XEDPATH -o godefs -goroot $PATH/TO/go go.yaml categories.yaml types.yaml
//
// types.yaml is already written, it specifies the shapes of vectors.
// categories.yaml and go.yaml contains definitions that unifies with types.yaml and XED
// data, you can find an example in ops/AddSub/.
//
// When generating Go definitions, simdgen do 3 "magic"s:
// - It splits masked operations(with op's [Masked] field set) to const and non const:
//   - One is a normal masked operation, the original
//   - The other has its mask operand's [Const] fields set to "K0".
//   - This way the user does not need to provide a separate "K0"-masked operation def.
//
// - It deduplicates intrinsic names that have duplicates:
//   - If there are two operations that shares the same signature, one is AVX512 the other
//     is before AVX512, the other will be selected.
//   - This happens often when some operations are defined both before AVX512 and after.
//     This way the user does not need to provide a separate "K0" operation for the
//     AVX512 counterpart.
//
// - It copies the op's [ConstImm] field to its immediate operand's [Const] field.
//   - This way the user does not need to provide verbose op definition while only
//     the const immediate field is different. This is useful to reduce verbosity of
//     compares with imm control predicates.
//
// These 3 magics could be disabled by enabling -nosplitmask, -nodedup or
// -noconstimmporting flags.
//
// simdgen right now only supports amd64, -arch=$OTHERARCH will trigger a fatal error.
package main

import (
	"github.com/shogo82148/std/flag"
)

var (
	FlagNoDedup           = flag.Bool("nodedup", false, "disable deduplicating godefs of 2 qualifying operations from different extensions")
	FlagNoConstImmPorting = flag.Bool("noconstimmporting", false, "disable const immediate porting from op to imm operand")
	FlagArch              = flag.String("arch", "amd64", "the target architecture")

	Verbose = flag.Bool("v", false, "verbose")

	FlagReportDup = flag.Bool("reportdup", false, "report the duplicate godefs")
)
