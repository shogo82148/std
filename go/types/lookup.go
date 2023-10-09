// "go test -run=Generate -write=all" によって生成されたコードです。編集しないでください。

// Copyright 2013 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// このファイルは、さまざまなフィールドやメソッドの検索機能を実装しています。

package types

// LookupFieldOrMethodは、与えられたパッケージと名前でTというフィールドまたはメソッドを検索し、対応する*Varまたは*Func、インデックスのシーケンス、そしてパスにおいてポインタ間接参照があったかどうかを示すブール値を返します。addressableが設定されている場合、Tはアドレス可能な変数の型です（メソッドの検索にのみ関係します）。Tはnilであってはなりません。
// 最後のインデックスエントリは、エントリが見つかった（埋め込まれた可能性のある）型のフィールドまたはメソッドのインデックスであり、次のいずれかです：
//
//  1. 名前付き型の宣言されたメソッドのリスト
//  2. インターフェース型のすべてのメソッド（メソッドセット）のリスト
//  3. 構造体型のフィールドのリスト
//
// より早いインデックスエントリは、見つかったエントリに到達するためにトラバースされた埋め込まれた構造体フィールドのインデックスであり、depth 0から開始します。
// エントリが見つからない場合、nilオブジェクトが返されます。この場合、返されるインデックスとindirectの値の意味は次のとおりです：
//   - もしindex != nilなら、インデックスシーケンスは曖昧なエントリを指します（同じ名前が同じ埋め込みレベルで複数回現れました）。
//   - indirectが設定されている場合、ポインタレシーバータイプを持つメソッドが見つかりましたが、実際のレシーバータイプからメソッドの形式的なレシーバーベースタイプへのパスにポインタがなく、レシーバーがアドレス可能ではありませんでした。
func LookupFieldOrMethod(T Type, addressable bool, pkg *Package, name string) (obj Object, index []int, indirect bool)

// MissingMethodは、VがTを実装している場合、(nil, false)を返します。そうでない場合、Tに必要な欠落しているメソッドと、欠落しているか、または単に間違った型（ポインタレシーバーまたは間違ったシグネチャ）を返します。
// 非インターフェース型V、またはstaticが設定されている場合、VがTを実装するには、TのすべてのメソッドがVに存在する必要があります。それ以外の場合（Vがインターフェースであり、staticが設定されていない場合）、MissingMethodは、Vにも存在するTのメソッドの型が一致していることだけをチェックします（例：型アサーションx.(T)の場合、xがインターフェース型Vである場合）。
func MissingMethod(V Type, T *Interface, static bool) (method *Func, wrongType bool)
