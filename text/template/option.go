// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains the code to handle template options.

package template

// Optionはテンプレートのオプションを設定します。オプションは
// 文字列によって記述され、単純な文字列または "key=value" の形式を取ります。
// オプション文字列には最大で一つの等号が含まれていることができます。
// オプション文字列が認識できないものや無効なものである場合、Optionはパニックを起こします。
//
// 既知のオプション:
//
// missingkey: マップに存在しないキーでマップがインデックスされた場合の、
// 実行時の動作を制御します。
//
//	"missingkey=default" または "missingkey=invalid"
//		デフォルトの動作: 何もせずに実行を続けます。
//		出力された場合、インデックス操作の結果は文字列
//		"<no value>" になります。
//	"missingkey=zero"
//		操作はマップ型の要素のゼロ値を返します。
//	"missingkey=error"
//		エラーで即座に実行を停止します。
func (t *Template) Option(opt ...string) *Template
