// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gob

import (
	"github.com/shogo82148/std/io"
	"github.com/shogo82148/std/reflect"
	"github.com/shogo82148/std/sync"
)

// Decoderは、接続のリモート側から読み取られた型とデータ情報の受信を管理します。
// 複数のゴルーチンによる並行使用が安全です。
//
// Decoderは、デコードされた入力サイズに対して基本的な健全性チェックのみを行い、
// その制限は設定可能ではありません。信頼できないソースからのgobデータをデコードする際は注意が必要です。
type Decoder struct {
	mutex        sync.Mutex
	r            io.Reader
	buf          decBuffer
	wireType     map[typeId]*wireType
	decoderCache map[reflect.Type]map[typeId]**decEngine
	ignorerCache map[typeId]**decEngine
	freeList     *decoderState
	countBuf     []byte
	err          error
}

// NewDecoderは、io.Readerから読み取る新しいデコーダを返します。
// もしrがio.ByteReaderも実装していない場合、それはbufio.Readerでラップされます。
func NewDecoder(r io.Reader) *Decoder

// Decodeは、入力ストリームから次の値を読み取り、
// 空のインターフェース値で表されるデータに格納します。
// もしeがnilの場合、値は破棄されます。それ以外の場合、
// eの下にある値は、受け取った次のデータ項目の
// 正しい型へのポインタでなければなりません。
// 入力がEOFにある場合、Decodeはio.EOFを返し、
// eを変更しません。
func (dec *Decoder) Decode(e any) error

// DecodeValueは、入力ストリームから次の値を読み取ります。
// もしvがゼロのreflect.Value（v.Kind() == Invalid）の場合、DecodeValueは値を破棄します。
// それ以外の場合、値はvに格納されます。その場合、vは
// 非nilのデータへのポインタを表すか、または代入可能なreflect.Value（v.CanSet()）でなければなりません。
// 入力がEOFにある場合、DecodeValueはio.EOFを返し、
// vを変更しません。
func (dec *Decoder) DecodeValue(v reflect.Value) error
