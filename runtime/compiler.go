// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Compilerは、実行中のバイナリを構築したコンパイラツールチェーンの名前です。既知のツールチェーンは次のとおりです:
//
// gc       cmd/compileとしても知られています。
// gccgo    GCCコンパイラスイートの一部であるgccgoフロントエンドです。
const Compiler = "gc"
