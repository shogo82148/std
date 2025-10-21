// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package modload

import (
	"github.com/shogo82148/std/context"
	"github.com/shogo82148/std/encoding/hex"

	"github.com/shogo82148/std/cmd/go/internal/modinfo"
)

var (
	_ = hex.DecodeString("3077af0c9274080241e1c107e6d618e6")
	_ = hex.DecodeString("f932433186182072008242104116d8f2")
)

// PackageModuleInfo returns information about the module that provides
// a given package. If modules are not enabled or if the package is in the
// standard library or if the package was not successfully loaded with
// LoadPackages or ImportFromFiles, nil is returned.
func PackageModuleInfo(loaderstate *State, ctx context.Context, pkgpath string) *modinfo.ModulePublic

// PackageModRoot returns the module root directory for the module that provides
// a given package. If modules are not enabled or if the package is in the
// standard library or if the package was not successfully loaded with
// LoadPackages or ImportFromFiles, the empty string is returned.
func PackageModRoot(loaderstate *State, ctx context.Context, pkgpath string) string

func ModuleInfo(loaderstate *State, ctx context.Context, path string) *modinfo.ModulePublic

func ModInfoProg(info string, isgccgo bool) []byte

func ModInfoData(info string) []byte
