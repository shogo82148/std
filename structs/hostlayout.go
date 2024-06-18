// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package structs

// HostLayout marks a struct as using host memory layout. A struct with a
// field of type HostLayout will be laid out in memory according to host
// expectations, generally following the host's C ABI.
//
// HostLayout does not affect layout within any other struct-typed fields
// of the containing struct, nor does it affect layout of structs
// containing the struct marked as host layout.
//
// By convention, HostLayout should be used as the type of a field
// named "_", placed at the beginning of the struct type definition.
type HostLayout struct {
	_ hostLayout
}
