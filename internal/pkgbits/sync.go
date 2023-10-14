// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgbits

// SyncMarker is an enum type that represents markers that may be
// written to export data to ensure the reader and writer stay
// synchronized.
type SyncMarker int

const (
	_ SyncMarker = iota

	// Low-level coding markers.
	SyncEOF
	SyncBool
	SyncInt64
	SyncUint64
	SyncString
	SyncValue
	SyncVal
	SyncRelocs
	SyncReloc
	SyncUseReloc

	// Higher-level object and type markers.
	SyncPublic
	SyncPos
	SyncPosBase
	SyncObject
	SyncObject1
	SyncPkg
	SyncPkgDef
	SyncMethod
	SyncType
	SyncTypeIdx
	SyncTypeParamNames
	SyncSignature
	SyncParams
	SyncParam
	SyncCodeObj
	SyncSym
	SyncLocalIdent
	SyncSelector

	// Private markers (only known to cmd/compile).
	SyncPrivate

	SyncFuncExt
	SyncVarExt
	SyncTypeExt
	SyncPragma

	SyncExprList
	SyncExprs
	SyncExpr
	SyncExprType
	SyncAssign
	SyncOp
	SyncFuncLit
	SyncCompLit

	SyncDecl
	SyncFuncBody
	SyncOpenScope
	SyncCloseScope
	SyncCloseAnotherScope
	SyncDeclNames
	SyncDeclName

	SyncStmts
	SyncBlockStmt
	SyncIfStmt
	SyncForStmt
	SyncSwitchStmt
	SyncRangeStmt
	SyncCaseClause
	SyncCommClause
	SyncSelectStmt
	SyncDecls
	SyncLabeledStmt
	SyncUseObjLocal
	SyncAddLocal
	SyncLinkname
	SyncStmt1
	SyncStmtsEnd
	SyncLabel
	SyncOptLabel

	SyncMultiExpr
	SyncRType
	SyncConvRTTI
)
