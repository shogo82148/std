// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
cgoパッケージは、cgoツールによって生成されたコードの実行時サポートを含んでいます。cgoの使用方法の詳細については、cgoコマンドのドキュメントを参照してください。
*/
package cgo

import "github.com/shogo82148/std/internal/runtime/sys"

// Incompleteは不完全なCの型のセマンティクスに特に使われます。
type Incomplete struct {
	_ sys.NotInHeap
}
