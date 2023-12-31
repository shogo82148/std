// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package x509

import (
	"github.com/shogo82148/std/encoding/asn1"
)

// OIDはASN.1 OBJECT IDENTIFIERを表します。
type OID struct {
	der []byte
}

// OIDFromIntsは、整数を使用して新しいOIDを作成します。各整数は別々のコンポーネントです。
func OIDFromInts(oid []uint64) (OID, error)

// Equalは、oidとotherが同じオブジェクト識別子を表している場合にtrueを返します。
func (oid OID) Equal(other OID) bool

// EqualASN1OIDは、OIDがasn1.ObjectIdentifierと等しいかどうかを返します。もし
// asn1.ObjectIdentifierがoidによって指定されたOIDを表現できない場合、
// OIDのコンポーネントが31ビット以上必要とする場合、falseを返します。
func (oid OID) EqualASN1OID(other asn1.ObjectIdentifier) bool

// Stringsはオブジェクト識別子の文字列表現を返します。
func (oid OID) String() string
