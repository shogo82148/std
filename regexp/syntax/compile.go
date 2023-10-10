// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

// Compileは正規表現を実行するためのプログラムにコンパイルします。
// 正規表現はすでに簡素化されている必要があります（re.Simplifyから戻されたもの）。
func Compile(re *Regexp) (*Prog, error)
