// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sgutil provides shared utilities for SIMD file
// generation across architectures.  This includes
//
// - "natural" comparison for better ordering
// - formatted-Go file saving
// - generic ops file generation
// - naming conventions and templates for the
//   bitwise vector reinterpretation no-op methods.

package sgutil

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/text/template"
)

// TemplateNamed returns a parsed template from temp, named name.
func TemplateNamed(name, temp string) *template.Template

// GenericOpsData holds one generic op entry for template rendering.
type GenericOpsData struct {
	OpName  string
	OpInLen int
	Comm    bool
	HasAux  bool
}

// WriteSIMDGenericOps generates the generic ops file content for archKey.
func WriteSIMDGenericOps(w io.Writer, ops []GenericOpsData, archKey string)
