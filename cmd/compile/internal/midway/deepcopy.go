// Copyright 2026 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package midway

import (
	"github.com/shogo82148/std/cmd/compile/internal/syntax"
	"github.com/shogo82148/std/cmd/compile/internal/types2"
)

// DeepCopier clones syntax nodes and maintains types2.Info mappings.
type DeepCopier struct {
	VecLen   int
	info     *types2.Info
	pkg      *types2.Package
	analyzer *Analyzer
	suffix   string

	vars map[*types2.Var]*types2.Var
}

func NewDeepCopier(pkg *types2.Package, info *types2.Info, vecLen int, analyzer *Analyzer, suffix string) *DeepCopier

// OnName rewrites "dependent" and SIMD names to their architecture-specific version.
func (c *DeepCopier) OnName(id *syntax.Name) *syntax.Name

// OnNameExpr rewrites references to simd.<simd type> into
// <bridge package>.<size-dependent-type>.
func (c *DeepCopier) OnNameExpr(id *syntax.Name) syntax.Expr

// OnSelector is looking for simd.Something, to be rewritten into
// appropriately.  Note that this will not work properly within the simd
// package because there is no "simd." selection there.
func (c *DeepCopier) OnSelector(se *syntax.SelectorExpr) syntax.Expr

func (c *DeepCopier) CopyDecl(d syntax.Decl) syntax.Decl

func (c *DeepCopier) CopyVarDecl(d *syntax.VarDecl) *syntax.VarDecl

func (c *DeepCopier) CopyTypeDecl(d *syntax.TypeDecl) *syntax.TypeDecl

func (c *DeepCopier) CopyConstDecl(d *syntax.ConstDecl) *syntax.ConstDecl

func (c *DeepCopier) CopyFuncDecl(d *syntax.FuncDecl) *syntax.FuncDecl

func (c *DeepCopier) CopyName(id *syntax.Name, isDef bool) *syntax.Name

func (c *DeepCopier) CopyNameExpr(id *syntax.Name) syntax.Expr

func (c *DeepCopier) CopyExpr(e syntax.Expr) syntax.Expr

func (c *DeepCopier) CopyStmt(s syntax.Stmt) syntax.Stmt

func (c *DeepCopier) CopySimpleStmt(s syntax.SimpleStmt) syntax.SimpleStmt

func (c *DeepCopier) CopyCaseClauses(list []*syntax.CaseClause) []*syntax.CaseClause

func (c *DeepCopier) CopyCommClauses(list []*syntax.CommClause) []*syntax.CommClause

func (c *DeepCopier) CopyBlockStmt(b *syntax.BlockStmt) *syntax.BlockStmt

func (c *DeepCopier) CopyFieldList(f []*syntax.Field) []*syntax.Field

func (c *DeepCopier) CopyField(f *syntax.Field) *syntax.Field
