// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package filepath

import (
	"github.com/shogo82148/std/errors"
)

// ErrBadPattern はパターンが不正であることを示します。
var ErrBadPattern = errors.New("syntax error in pattern")

// Matchは、名前がシェルのファイル名パターンと一致するかどうかを報告します。
// パターンの構文は次の通りです:
//
//	pattern:
//		{ term }
//	term:
//		'*'         任意の非区切り文字のシーケンスに一致します
//		'?'         任意の単一の非区切り文字に一致します
//		'[' [ '^' ] { character-range } ']'
//		            文字クラス（空であってはなりません）
//		c           文字cに一致します（c != '*', '?', '\\', '['）
//		'\\' c      文字cに一致します（Windowsを除く）
//
//	character-range:
//		c           文字cに一致します（c != '\\', '-', ']'）
//		'\\' c      文字cに一致します（Windowsを除く）
//		lo '-' hi   lo <= c <= hi を満たす文字cに一致します
//
// パターン内のパスセグメントは [Separator] で区切られていなければなりません。
//
// Matchは、patternがname全体に一致することを要求し、部分文字列への一致ではありません。
// 返される可能性のある唯一のエラーは、patternが不正な場合の [ErrBadPattern] です。
//
// Matchはパターンがnameのすべてに一致することを要求し、部分文字列だけではありません。
// 返される可能性のある唯一のエラーは、パターンが不正な場合の [ErrBadPattern] です。
//
// Windowsでは、エスケープは無効になっています。代わりに'\\'はパスセパレータとして扱われます。
func Match(pattern, name string) (matched bool, err error)

// Globは、パターンに一致するすべてのファイルの名前を返します。一致するファイルがない場合はnilを返します。パターンの構文はMatchと同じです。パターンには、/usr/*/bin/ed（[Separator] が '/'と仮定）などの階層的な名前を記述することができます。
// Globは、ディレクトリを読み込む際のI/Oエラーなどのファイルシステムのエラーを無視します。返される唯一の可能性のあるエラーは、パターンが不正な場合の [ErrBadPattern] です。
func Glob(pattern string) (matches []string, err error)
