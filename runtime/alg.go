// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// type algorithms - known to compiler

// typeAlg is also copied/used in reflect/type.go.
// keep them in sync.

// used in asm_{386,amd64,arm64}.s to seed the hash function

// used in hash{32,64}.go to seed the hash function
