// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pe

// StringTableは、COFFの文字列テーブルです。
type StringTable []byte

// Stringは、オフセットstartでCOFF文字列テーブルstから文字列を抽出します。
func (st StringTable) String(start uint32) (string, error)
