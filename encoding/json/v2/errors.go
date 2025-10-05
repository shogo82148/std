// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build goexperiment.jsonv2

package json

import (
	"github.com/shogo82148/std/errors"
	"github.com/shogo82148/std/reflect"

	"github.com/shogo82148/std/encoding/json/jsontext"
)

// ErrUnknownNameは、JSONオブジェクトのメンバー名が
// 対象のGo構造体で認識されずアンマーシャルできなかったことを示します。
// このエラーは生成時に [SemanticError] で直接ラップされます。
//
// 未知のJSONオブジェクトメンバー名は以下のように取得できます:
//
//	err := ...
//	var serr json.SemanticError
//	if errors.As(err, &serr) && serr.Err == json.ErrUnknownName {
//		ptr := serr.JSONPointer // 未知の名前へのJSONポインタ
//		name := ptr.LastToken() // 未知の名前そのもの
//		...
//	}
//
// このエラーは [RejectUnknownMembers] がtrueの場合のみ返されます。
var ErrUnknownName = errors.New("unknown object member name")

// SemanticErrorは、JSONデータをGoデータへ、またはその逆へ意味付けする際のエラーを表します。
//
// このパッケージによって生成されるこのエラーの内容は将来的に変更される可能性があります。
type SemanticError struct {
	requireKeyedLiterals
	nonComparable

	action string

	// ByteOffsetは、エラーがこのバイトオフセット以降で発生したことを示します。
	ByteOffset int64
	// JSONPointerは、RFC 6901で定義されたJSONポインタ表記を使って
	// このJSON値内でエラーが発生したことを示します。
	JSONPointer jsontext.Pointer

	// JSONKindは、処理できなかったJSONの種類を示します。
	JSONKind jsontext.Kind
	// JSONValueは、アンマーシャルできなかったJSONの数値または文字列です。
	// マーシャル時には設定されません。
	JSONValue jsontext.Value
	// GoTypeは、処理できなかったGoの型を示します。
	GoType reflect.Type

	// Errは、根本的なエラーです。
	Err error
}

func (e *SemanticError) Error() string

func (e *SemanticError) Unwrap() error
