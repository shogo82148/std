// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !js && !wasip1

package base

// Exec replaces the current process with an invocation of exe with the current
// command-line arguments and environment, or runs it as a subprocess and exits
// on platforms without syscall.Exec. name is used in error messages.
func Exec(name, exe string)
