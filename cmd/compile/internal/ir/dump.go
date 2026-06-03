// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file implements textual dumping of arbitrary data structures
// for debugging purposes. The code is customized for Node graphs
// and may be used for an alternative view of the node structure.

package ir

import (
	"github.com/shogo82148/std/io"
)

// DumpAny is like FDumpAny but prints to stderr.
func DumpAny(root any, filter string, depth int)

// FDumpAny prints the structure of a rooted data structure
// to w by depth-first traversal of the data structure.
//
// The filter parameter is a regular expression. If it is
// non-empty, only struct fields whose names match filter
// are printed.
//
// The depth parameter controls how deep traversal recurses
// before it returns (higher value means greater depth).
// If an empty field filter is given, a good depth default value
// is 4. A negative depth means no depth limit, which may be fine
// for small data structures or if there is a non-empty filter.
//
// In the output, Node structs are identified by their Op name
// rather than their type; struct fields with zero values or
// non-matching field names are omitted, and "…" means recursion
// depth has been reached or struct fields have been omitted.
func FDumpAny(w io.Writer, root any, filter string, depth int)

// MatchAstDump returns true if the fn matches the value
// of the astdump debug flag.  Fn matches in the following
// cases:
//
//   - astdump == name(fn)
//   - astdump == pkgname(fn).name(fn)
//   - astdump == afterslash(pkgname(fn)).name(fn)
//   - astdump begins with a "~" and what follows "~" is a
//     regular expression matching pkgname(fn).name(fn)
//
// If MatchAstDump returns true, it also prints to os.Stderr
//
//	\nir.Match(<fn>, <astdump>) for <where>\n
func MatchAstDump(fn *Func, where string) bool

// MatchPkgFn returns true if pkg and fnName "match" toMatch.
// "~REGEXP" matches REGEXP against pkgName + "." + fnName
// "aFunc" matches "aFunc" (in any package)
// "aPkg.aFunc" matches "aPkg.aFunc"
// "aPkg/subPkg.aFunc" matches "subPkg.aFunc"
func MatchPkgFn(pkgName, fnName, toMatch string) bool

// AstDump appends the ast dump for fn to the ast dump file for fn.
// The generated file name is
//
//	url.PathEscape(PkgFuncName(fn)) + ".ast"
//
// It also prints
//
//	Writing ast output to <astfilename>\n
//
// to os.Stderr.
func AstDump(fn *Func, why string)

// EscapedFileName constructs a file name from fn and suffix,
// url-path-escaping the function part of the name and replacing it
// with a hash if it is too long.  The suffix is neither escaped
// nor including in the length calculation, so an excessively
// creative suffix will result in problems.
func EscapedFileName(fn, suffix string) string

// DumpNodeHTML dumps the node n to the HTML writer for fn.
// It uses the same phase name as the text dump.
func DumpNodeHTML(fn *Func, why string, n Node)

// CloseHTMLWriters closes the HTML writer for fn, if one exists.
func CloseHTMLWriters()
