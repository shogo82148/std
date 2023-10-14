// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gccgoimporter

import (
	"github.com/shogo82148/std/go/types"
)

// Information about a specific installation of gccgo.
type GccgoInstallation struct {
	// Version of gcc (e.g. 4.8.0).
	GccVersion string

	// Target triple (e.g. x86_64-unknown-linux-gnu).
	TargetTriple string

	// Built-in library paths used by this installation.
	LibPaths []string
}

// Ask the driver at the given path for information for this GccgoInstallation.
// The given arguments are passed directly to the call of the driver.
func (inst *GccgoInstallation) InitFromDriver(gccgoPath string, args ...string) (err error)

// Return the list of export search paths for this GccgoInstallation.
func (inst *GccgoInstallation) SearchPaths() (paths []string)

// Return an importer that searches incpaths followed by the gcc installation's
// built-in search paths and the current directory.
func (inst *GccgoInstallation) GetImporter(incpaths []string, initmap map[*types.Package]InitData) Importer
