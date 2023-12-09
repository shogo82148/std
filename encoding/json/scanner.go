// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

// Validは、dataが有効なJSONエンコーディングであるかどうかを報告します。
func Valid(data []byte) bool

// SyntaxErrorは、JSON構文エラーの説明です。
// JSONを解析できない場合、[Unmarshal] はSyntaxErrorを返します。
type SyntaxError struct {
	msg    string
	Offset int64
}

func (e *SyntaxError) Error() string
