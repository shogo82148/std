// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package asn1

// MarshalはvalのASN.1エンコーディングを返します。
//
// In addition to the struct tags recognized by Unmarshal, the following can be
// used:
//
//	ia5:         文字列をASN.1 IA5String値としてマーシャルします
//	omitempty:   空のスライスをスキップします
//	printable:   文字列をASN.1 PrintableString値としてマーシャルします
//	utf8:        文字列をASN.1 UTF8String値としてマーシャルします
//	numeric:     文字列をASN.1 NumericString値としてマーシャルします
//	utc:         time.TimeをASN.1 UTCTime値としてマーシャルします
//	generalized: time.TimeをASN.1 GeneralizedTime値としてマーシャルします
func Marshal(val any) ([]byte, error)

// MarshalWithParamsは、トップレベルの要素にフィールドパラメータを指定することを可能にします。パラメータの形式は、フィールドタグと同じです。
func MarshalWithParams(val any, params string) ([]byte, error)
