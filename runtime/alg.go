// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// type algorithms - known to compiler

// runtime variable to check if the processor we're running on
// actually supports the instructions used by the AES-based
// hash implementation.

// used in asm_{386,amd64,arm64}.s to seed the hash function

// used in hash{32,64}.go to seed the hash function
