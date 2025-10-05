// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

// Validは、dataが有効なJSONエンコーディングかどうかを報告します。
func Valid(data []byte) bool

// SyntaxErrorは、JSON構文エラーの内容を表します。
// [Unmarshal]は、JSONをパースできない場合にSyntaxErrorを返します。
type SyntaxError struct {
	msg    string
	Offset int64
}

func (e *SyntaxError) Error() string
