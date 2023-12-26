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
// Known options:
//
// missingkey: Control the behavior during execution if a map is
// indexed with a key that is not present in the map.
//
//	"missingkey=default" or "missingkey=invalid"
//		The default behavior: Do nothing and continue execution.
//		If printed, the result of the index operation is the string
//		"<no value>".
//	"missingkey=zero"
//		The operation returns the zero value for the map type's element.
//	"missingkey=error"
//		Execution stops immediately with an error.
func (t *Template) Option(opt ...string) *Template
