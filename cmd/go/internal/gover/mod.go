// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gover

import (
	"golang.org/x/mod/module"
)

// IsToolchain reports whether the module path corresponds to the
// virtual, non-downloadable module tracking go or toolchain directives in the go.mod file.
//
// Note that IsToolchain only matches "go" and "toolchain", not the
// real, downloadable module "golang.org/toolchain" containing toolchain files.
//
//	IsToolchain("go") = true
//	IsToolchain("toolchain") = true
//	IsToolchain("golang.org/x/tools") = false
//	IsToolchain("golang.org/toolchain") = false
func IsToolchain(path string) bool

// ModCompare returns the result of comparing the versions x and y
// for the module with the given path.
// The path is necessary because the "go" and "toolchain" modules
// use a different version syntax and semantics (gover, this package)
// than most modules (semver).
func ModCompare(path string, x, y string) int

// ModSort is like module.Sort but understands the "go" and "toolchain"
// modules and their version ordering.
func ModSort(list []module.Version)

// ModIsValid reports whether vers is a valid version syntax for the module with the given path.
func ModIsValid(path, vers string) bool

// ModIsPrefix reports whether v is a valid version syntax prefix for the module with the given path.
// The caller is assumed to have checked that ModIsValid(path, vers) is true.
func ModIsPrefix(path, vers string) bool

// ModIsPrerelease reports whether v is a prerelease version for the module with the given path.
// The caller is assumed to have checked that ModIsValid(path, vers) is true.
func ModIsPrerelease(path, vers string) bool

// ModMajorMinor returns the "major.minor" truncation of the version v,
// for use as a prefix in "@patch" queries.
func ModMajorMinor(path, vers string) string
