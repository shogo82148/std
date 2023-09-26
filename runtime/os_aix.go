// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build aix
// +build aix

package runtime

// funcDescriptor is a structure representing a function descriptor
// A variable with this type is always created in assembler

// tstart is a function descriptor to _tstart defined in assembly.

// sigtramp is a function descriptor to _sigtramp defined in assembly
