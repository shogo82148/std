// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains test cases for short valid and invalid programs.

package parser

// validWithTParamsOnly holds source code examples that are valid if
// parseTypeParams is set, but invalid if not. When checking with the
// parseTypeParams set, errors are ignored.

// invalidNoTParamErrs holds invalid source code examples annotated with the
// error messages produced when ParseTypeParams is not set.

// invalidTParamErrs holds invalid source code examples annotated with the
// error messages produced when ParseTypeParams is set.
