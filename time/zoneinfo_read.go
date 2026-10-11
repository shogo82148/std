// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Parse "zoneinfo" time zone file.
// This is a fairly standard file format used on OS X, Linux, BSD, Sun, and others.
// See tzfile(5), https://en.wikipedia.org/wiki/Zoneinfo,
// and ftp://munnari.oz.au/pub/oldtz/

package time

// LoadLocationFromTZDataは、IANAタイムゾーンデータベース形式のデータから初期化された、
// 指定された名前を持つ新しい [Location] を返します。
// データは標準のIANAタイムゾーンファイルの形式
// （例えば、Unixシステムの/etc/localtimeの内容）でなければなりません。
func LoadLocationFromTZData(name string, data []byte) (*Location, error)
