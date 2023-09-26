// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !netgo && cgo && darwin

package net

/*
#include <resolv.h>
*/

// This will cause a compile error when the size of
// unix.ResState is too small.
