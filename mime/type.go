// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// mimeパッケージはMIME仕様の一部を実装します。
package mime

// TypeByExtensionは、ファイル拡張子extに関連付けられたMIMEタイプを返します。
// 拡張子extは、".html"のように先頭にドットを付けて開始する必要があります。
// extが関連付けられたタイプを持っていない場合、TypeByExtensionは""を返します。
//
// 拡張子は、最初に大文字と小文字を区別して検索され、次に大文字と小文字を区別せずに検索されます。
//
// 組み込みのテーブルは小さいですが、unixではローカルシステムのMIME-infoデータベースや
// mime.typesファイルによって補強されます（もし以下の名前の一つ以上で利用可能な場合）:
//
//	/usr/local/share/mime/globs2
//	/usr/share/mime/globs2
//	/etc/mime.types
//	/etc/apache2/mime.types
//	/etc/apache/mime.types
//
// Windowsでは、MIMEタイプはレジストリから抽出されます。
//
// テキストタイプは、デフォルトでcharsetパラメータが"utf-8"に設定されています。
func TypeByExtension(ext string) string

// ExtensionsByTypeは、MIMEタイプtypに関連付けられていると知られている拡張子を返します。
// 返される拡張子は、".html"のように先頭にドットが付いて始まります。
// typが関連付けられた拡張子を持っていない場合、ExtensionsByTypeはnilスライスを返します。
func ExtensionsByType(typ string) ([]string, error)

// AddExtensionTypeは、拡張子extに関連付けられたMIMEタイプをtypに設定します。
// 拡張子は、".html"のように先頭にドットを付けて開始する必要があります。
func AddExtensionType(ext, typ string) error
