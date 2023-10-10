// "go test -run=Generate -write=all" によって生成されたコードです。編集しないでください。

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルはジェネリック型の具現化を実装します
// 型パラメータを型引数による置換で行います。

package types

// Instantiateは、与えられた型引数targsで型origをインスタンス化します。
// origは*Named型または*Signature型でなければなりません。エラーがない場合、
// 結果の型は同じ種類（*Named型または*Signature型）のインスタンス化された型になります。
// *Named型に添付されたメソッドもインスタンス化され、元のメソッドと同じ位置を持つ
// nil関数スコープの新しい*Funcに関連付けられます。
// ctxtがnilでない場合、同じ識別子を持つ以前のインスタンスと重複排除するために使用できます。
// 特別な場合として、ジェネリックの*Signatureの元の型は、ポインタの等価性がある場合のみ同一視されるため、
// 異なる（しかし可能性としては同一の）シグネチャをインスタンス化すると、異なるインスタンスが生成されます。
// 共有されたコンテキストの使用は、同じインスタンスがすべての場合で重複排除されることを保証しません。
// validateが設定されている場合、Instantiateは型引数とパラメータの数が一致し、
// 型引数がそれに対応する型制約を満たしていることを検証します。
// 検証に失敗した場合、結果のエラーは、対応する型パラメータの制約を満たさなかった
// 型引数を示す*ArgumentErrorをラップし、その理由を示す可能性があります。
// validateが設定されていない場合、Instantiateは型引数の数や型引数が制約を満たしているかどうかを検証しません。
// Instantiateはエラーを返さないことが保証されていますが、パニックする可能性があります。
// 具体的には、*Signature型の場合、型引数の数が正しくない場合は即座にパニックします。
// *Named型の場合、パニックは後で*Named API内で発生する可能性があります。
func Instantiate(ctxt *Context, orig Type, targs []Type, validate bool) (Type, error)
