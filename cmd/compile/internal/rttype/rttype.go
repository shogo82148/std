// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package rttype allows the compiler to share type information with
// the runtime. The shared type information is stored in
// internal/abi. This package translates those types from the host
// machine on which the compiler runs to the target machine on which
// the compiled program will run. In particular, this package handles
// layout differences between e.g. a 64 bit compiler and 32 bit
// target.
package rttype

import (
	"github.com/shogo82148/std/cmd/compile/internal/types"
	"github.com/shogo82148/std/cmd/internal/obj"
)

type RuntimeType struct {
	// A *types.Type representing a type used at runtime.
	t *types.Type
	// components maps from component names to their location in the type.
	components map[string]location
}

// Types shared with the runtime via internal/abi.
// TODO: add more
var Type *RuntimeType

func Init()

func (r *RuntimeType) Size() int64

func (r *RuntimeType) Alignment() int64

func (r *RuntimeType) Offset(name string) int64

// WritePtr writes a pointer "target" to the component named "name" in the
// static object "lsym".
func (r *RuntimeType) WritePtr(lsym *obj.LSym, name string, target *obj.LSym)

func (r *RuntimeType) WriteUintptr(lsym *obj.LSym, name string, val uint64)

func (r *RuntimeType) WriteUint32(lsym *obj.LSym, name string, val uint32)

func (r *RuntimeType) WriteUint8(lsym *obj.LSym, name string, val uint8)

func (r *RuntimeType) WriteSymPtrOff(lsym *obj.LSym, name string, target *obj.LSym, weak bool)
