// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asn1

// MarshalはvalのASN.1エンコーディングを返します。
//
// In addition to the struct tags recognized by Unmarshal, the following can be
// used:
//
//	ia5:         文字列をASN.1のIA5String値としてエンコードします。
//	omitempty:   空のスライスをスキップします。
//	printable:   文字列をASN.1のPrintableString値としてエンコードします。
//	utf8:        文字列をASN.1のUTF8String値としてエンコードします。
//	utc:         time.TimeをASN.1のUTCTime値としてエンコードします。
//	generalized: time.TimeをASN.1のGeneralizedTime値としてエンコードします。
func Marshal(val any) ([]byte, error)

// MarshalWithParamsは、トップレベルの要素にフィールドパラメータを指定することを可能にします。パラメータの形式は、フィールドタグと同じです。
func MarshalWithParams(val any, params string) ([]byte, error)
