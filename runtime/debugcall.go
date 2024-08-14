// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Though the debug call function feature is not enabled on
// ppc64, inserted ppc64 to avoid missing Go declaration error
// for debugCallPanicked while building runtime.test
//go:build amd64 || arm64 || loong64 || ppc64le || ppc64

package runtime
