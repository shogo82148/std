// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xml

import (
	"github.com/shogo82148/std/reflect"
)

// TagPathErrorは、競合するパスを持つフィールドタグの使用によって
// アンマーシャル処理中に発生したエラーを表します。
type TagPathError struct {
	Struct       reflect.Type
	Field1, Tag1 string
	Field2, Tag2 string
}

func (e *TagPathError) Error() string
