// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package strings

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/sync"
)

// Replacerは、置換物のリストで文字列を置換します。
// 複数のゴルーチンによる同時使用に対して安全です。
type Replacer struct {
	once   sync.Once
	r      replacer
	oldnew []string
}

<<<<<<< HEAD
// NewReplacerは、古い文字列と新しい文字列のペアのリストから新しいReplacerを返します。
// 置換は、対象文字列に現れる順序で実行され、重複するマッチングは行われません。
// 古い文字列の比較は引数の順序で行われます。
=======
// NewReplacer returns a new [Replacer] from a list of old, new string
// pairs. Replacements are performed in the order they appear in the
// target string, without overlapping matches. The old string
// comparisons are done in argument order.
>>>>>>> upstream/master
//
// NewReplacerは、奇数の引数が与えられた場合にパニックを引き起こします。
func NewReplacer(oldnew ...string) *Replacer

// Replaceは、すべての置換を実行したsのコピーを返します。
func (r *Replacer) Replace(s string) string

// WriteStringは、すべての置換を実行したsをwに書き込みます。
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
