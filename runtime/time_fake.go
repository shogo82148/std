// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build faketime && !windows

// Faketime isn't currently supported on Windows. This would require
// modifying syscall.Write to call syscall.faketimeWrite,
// translating the Stdout and Stderr handles into FDs 1 and 2.
// (See CL 192739 PS 3.)

package runtime
