// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sgutil provides shared utilities for SIMD file
// generation across architectures.  This includes
//
// - "natural" comparison for better ordering
// - formatted-Go file saving
// - file merging for simdgenericOps.go
// - naming conventions and templates for the
//   bitwise vector reinterpretation no-op methods.

package sgutil

import (
	"github.com/shogo82148/std/bytes"
	"github.com/shogo82148/std/text/template"
)

// TemplateNamedreturns a parsed template from temp, named name.
func TemplateNamed(name, temp string) *template.Template

// GenericOpsData holds one generic op entry for template rendering.
type GenericOpsData struct {
	OpName  string
	OpInLen int
	Comm    bool
	HasAux  bool
	Archs   []string
}

// ArchTag returns the comma-separated arch list for the template comment.
func (d GenericOpsData) ArchTag() string

// MergeSIMDGenericOps merges a new set of generic ops with the existing ones read
// from goroot "/" file.
func MergeSIMDGenericOps(newOps []GenericOpsData, oldFile, currentArch string) *bytes.Buffer
