// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Addr2lineはGNU addr2lineツールの最低限のシミュレーションであり、pprofをサポートするために必要なだけの機能を有しています。
// 使用方法：
// 	go tool addr2line バイナリ
// Addr2lineは標準入力から16進数のアドレスを読み取ります。各入力アドレスに対して、addr2lineは2つの出力行を表示します。まず、アドレスを含む関数の名前が表示され、次にそのアドレスに対応するソースコードのファイルと行が表示されます。
// このツールはpprofの内部でのみ使用することを目的としており、将来のリリースでインターフェースが変更される可能性があるか、または完全に削除される可能性があります。
package main
