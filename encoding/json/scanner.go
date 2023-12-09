// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package json

// Validは、dataが有効なJSONエンコーディングであるかどうかを報告します。
func Valid(data []byte) bool

<<<<<<< HEAD
// A SyntaxError is a description of a JSON syntax error.
// [Unmarshal] will return a SyntaxError if the JSON can't be parsed.
=======
// SyntaxErrorは、JSON構文エラーの説明です。
// JSONを解析できない場合、UnmarshalはSyntaxErrorを返します。
>>>>>>> release-branch.go1.21
type SyntaxError struct {
	msg    string
	Offset int64
}

func (e *SyntaxError) Error() string
