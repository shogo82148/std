// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Type conversions for Scan.

package sql

import (
	"github.com/shogo82148/std/database/sql/driver"
)

// ConvertAssignは、srcの値をdestが指す値にコピーします。
// 変換の詳細については、[Rows.Scan] のドキュメントを参照してください。
// destはポインタであるか、[Scanner] を実装していなければなりません。
//
// [driver.RowsColumnScanner] の実装では、
// [driver.ScanContext] パラメータをそのまま渡す必要があります。
// それ以外の場合は、コンテキストとして driver.ScanContext{} を渡します。
//
// ConvertAssignは、ドライバ実装での使用を意図しています。
// ほとんどのユーザーは、これを直接使う必要はありません。
func ConvertAssign(scanCtx driver.ScanContext, dest any, src driver.Value) error
