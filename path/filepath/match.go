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
//		            キャラクタークラス（非空である必要があります）
//		c           文字cに一致します（c != '*', '?', '\\', '['）
//		'\\' c      文字cに一致します
//
//	character-range:
//		c           文字cに一致します（c != '\\', '-', ']'）
//		'\\' c      文字cに一致します
//		lo '-' hi   lo <= c <= hi の条件で文字cに一致します
//
// Matchは、パターンが名前全体ではなく、部分文字列ではないことを要求します。
// 返される唯一の可能なエラーは、パターンが異常である場合の [ErrBadPattern] です。
//
// Windowsでは、エスケープは無効になっています。代わりに'\\'はパスセパレータとして扱われます。
func Match(pattern, name string) (matched bool, err error)

// Globは、パターンに一致するすべてのファイルの名前を返します。一致するファイルがない場合はnilを返します。パターンの構文はMatchと同じです。パターンには、/usr/*/bin/ed（[Separator] が '/'と仮定）などの階層的な名前を記述することができます。
// Globは、ディレクトリを読み込む際のI/Oエラーなどのファイルシステムのエラーを無視します。返される唯一の可能性のあるエラーは、パターンが不正な場合の [ErrBadPattern] です。
func Glob(pattern string) (matches []string, err error)
