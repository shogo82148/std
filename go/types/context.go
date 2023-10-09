// "go test -run=Generate -write=all" によって生成されたコードです。編集しないでください。

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package types

import (
	"github.com/shogo82148/std/sync"
)

// Contextは不透明な型チェックコンテキストです。これは、型チェックされたパッケージやInstantiateへの呼び出し間で同じ型のインスタンスを共有するために使用される場合があります。Contextは同時利用に関して安全です。
//
// 共有コンテキストの利用は、すべてのケースで同じインスタンスが重複削除されることを保証するものではありません。
type Context struct {
	mu        sync.Mutex
	typeMap   map[string][]ctxtEntry
	nextID    int
	originIDs map[Type]int
}

// NewContext は新しいコンテキストを作成します。
func NewContext() *Context
