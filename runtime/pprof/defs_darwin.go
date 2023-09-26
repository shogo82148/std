// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file is used as input to cgo --godefs (GOOS=arm64 or amd64) to
// generate the types used in viminfo_darwin_{arm64,amd64}.go which are
// hand edited as appropriate, primarily to avoid exporting the types.

//go:build ignore

package pprof

/*
#include <sys/param.h>
#include <mach/vm_prot.h>
#include <mach/vm_region.h>
*/
