// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package runtime

// Error インターフェースはランタイムエラーを識別します。
type Error interface {
	error

	RuntimeError()
}

// TypeAssertionErrorは、型アサーションの失敗を説明します。
type TypeAssertionError struct {
	_interface    *_type
	concrete      *_type
	asserted      *_type
	missingMethod string
}

func (*TypeAssertionError) RuntimeError()

func (e *TypeAssertionError) Error() string
