// Copyright 2025 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// ArchInfo contains all architecture-specific naming conventions.
type ArchInfo struct {
	Arch            string
	ArchUpper       string
	ObjArch         string
	RegInfoKeys     []string
	RegInfoSet      map[string]bool
	RegInfoParams   string
	GeneratedHeader string
	Arrangements    []string
}

// GetArchInfo returns architecture-specific information based on the target architecture.
func GetArchInfo(arch string) (ArchInfo, error)

// CurrentArch returns the ArchInfo for the current FlagArch setting.
func CurrentArch() ArchInfo
