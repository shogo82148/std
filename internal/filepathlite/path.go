// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package filepathlite implements a subset of path/filepath,
// only using packages which may be imported by "os".
//
// Tests for these functions are in path/filepath.
package filepathlite

// Clean is filepath.Clean.
func Clean(path string) string

// IsLocal is filepath.IsLocal.
func IsLocal(path string) bool

// Localize is filepath.Localize.
func Localize(path string) (string, error)

// ToSlash is filepath.ToSlash.
func ToSlash(path string) string

// FromSlash is filepath.ToSlash.
func FromSlash(path string) string

// Split is filepath.Split.
func Split(path string) (dir, file string)

// Ext is filepath.Ext.
func Ext(path string) string

// Base is filepath.Base.
func Base(path string) string

// Dir is filepath.Dir.
func Dir(path string) string

// VolumeName is filepath.VolumeName.
func VolumeName(path string) string

// VolumeNameLen returns the length of the leading volume name on Windows.
// It returns 0 elsewhere.
func VolumeNameLen(path string) int
