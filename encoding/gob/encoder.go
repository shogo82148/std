// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/reflect"
	"github.com/shogo82148/std/sync"
)

// Encoderは、接続の他方に対する型とデータ情報の送信を管理します。
// 複数のgoroutineが同時に使用しても安全です。
type Encoder struct {
	mutex      sync.Mutex
	w          []io.Writer
	sent       map[reflect.Type]typeId
	countState *encoderState
	freeList   *encoderState
	byteBuf    encBuffer
	err        error
}

<<<<<<< HEAD
// NewEncoder returns a new encoder that will transmit on the [io.Writer].
=======
// NewEncoderは、io.Writer上で送信する新しいエンコーダを返します。
>>>>>>> release-branch.go1.21
func NewEncoder(w io.Writer) *Encoder

// Encodeは、空のインターフェース値で表されるデータ項目を送信します。
// 必要なすべての型情報が最初に送信されることを保証します。
// nilポインタをEncoderに渡すとパニックを引き起こします、なぜならそれらはgobによって送信できないからです。
func (enc *Encoder) Encode(e any) error

// EncodeValueは、リフレクション値によって表されるデータ項目を送信します。
// 必要なすべての型情報が最初に送信されることを保証します。
// nilポインタをEncodeValueに渡すとパニックを引き起こします、なぜならそれらはgobによって送信できないからです。
func (enc *Encoder) EncodeValue(value reflect.Value) error
