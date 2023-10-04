// Copyright 2020 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// tzdata パッケージは、タイムゾーンデータベースの埋め込みコピーを提供します。
// このパッケージがプログラムのどこかでインポートされている場合、
// タイムパッケージがシステム上の tzdata ファイルを見つけられない場合、
// この埋め込み情報を使用します。
//
// このパッケージをインポートすると、プログラムのサイズが約 450 KB 増加します。
//
// このパッケージは、通常、プログラムのメインパッケージによってインポートされるべきです。
// ライブラリは通常、プログラムにタイムゾーンデータベースを含めるかどうかを決定すべきではありません。
//
// このパッケージは、-tags timetzdata でビルドすると自動的にインポートされます。
package tzdata
