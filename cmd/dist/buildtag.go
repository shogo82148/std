// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

// exprParser is a //go:build expression parser and evaluator.
// The parser is a trivial precedence-based parser which is still
// almost overkill for these very simple expressions.

// val is the value type result of parsing.
// We don't keep a parse tree, just the value of the expression.

// exprToken describes a single token in the input.
// Prefix operators define a prefix func that parses the
// upcoming value. Binary operators define an infix func
// that combines two values according to the operator.
// In that case, the parsing loop parses the two values.
