// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build unix || (js && wasm) || wasip1

package filepath

// IsAbsはパスが絶対パスかどうかを報告します。
func IsAbs(path string) bool

// HasPrefixは歴史的な互換性のために存在しており、使用するべきではありません。
//
// Deprecated: HasPrefixはパスの境界を尊重せず、
// 必要な場合に大文字と小文字を無視しません。
func HasPrefix(p, prefix string) bool
