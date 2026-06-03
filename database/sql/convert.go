// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Type conversions for Scan.

package sql

import (
	"github.com/shogo82148/std/database/sql/driver"
)

// ConvertAssign copies the value in src to the value pointed at by dest.
// See the documentation on [Rows.Scan] for details on conversions.
// dest must be a pointer or must implement [Scanner].
//
// Implementations of [driver.RowsColumnScanner] should pass through
// their [driver.ScanContext] parameter.
// In other cases, pass driver.ScanContext{} as the context.
//
// ConvertAssign is intended for use by driver implementations.
// Most users should not need to use it directly.
func ConvertAssign(scanCtx driver.ScanContext, dest any, src driver.Value) error
