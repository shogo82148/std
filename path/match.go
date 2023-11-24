// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package path

import (
	"github.com/shogo82148/std/errors"
)

// ErrBadPatternは、パターンが不正であることを示します。
var ErrBadPattern = errors.New("syntax error in pattern")

// Matchはnameがシェルパターンと一致するかどうかを判定します。
// パターンの構文は以下のようになります：
//
// パターン:
//
//	{ 要素 }
//
// 要素:
//
//	'*'         非/の任意のシーケンスに一致します
//	'?'         非/の任意の1文字に一致します
//	'[' [ '^' ] { 文字範囲 } ']'
//	           文字のクラス（空でなければならない）
//	c           文字cに一致します（c != '*', '?', '\\', '['）
//	'\\' c      文字cに一致します
//
// 文字範囲:
//
//	c           文字cに一致します（c != '\\', '-', ']')
//	'\\' c      文字cに一致します
//	lo '-' hi   lo <= c <= hiの範囲の文字cに一致します
//
// Matchはパターンがname全体と一致する必要があります。部分一致ではありません。
// 返される可能性のある唯一のエラーは、patternが正しくない場合の [ErrBadPattern] です。
func Match(pattern, name string) (matched bool, err error)
