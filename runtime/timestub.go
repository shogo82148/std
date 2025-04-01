// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Declarations for operating systems implementing time.now
// indirectly, in terms of walltime and nanotime assembly.

//go:build !faketime && !windows && !(linux && amd64) && !plan9

package runtime
