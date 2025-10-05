// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package jsontext

// SyntacticErrorは、文法に従ってJSONをエンコードまたはデコードする際に発生した構文エラーの説明です。
//
// このパッケージによって生成されるこのエラーの内容は将来的に変更される可能性があります。
type SyntacticError struct {
	requireKeyedLiterals
	nonComparable

	// ByteOffsetは、このバイトオフセットの直後でエラーが発生したことを示します。
	ByteOffset int64
	// JSONPointerは、RFC 6901で定義されたJSONポインタ表記を用いて、
	// このJSON値内でエラーが発生したことを示します。
	JSONPointer Pointer

	// Errは基礎となるエラーです。
	Err error
}

func (e *SyntacticError) Error() string

func (e *SyntacticError) Unwrap() error
