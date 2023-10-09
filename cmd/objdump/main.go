// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Objdumpは実行可能ファイルを逆アセンブルします。
//
// 使用法：
//
//	go tool objdump [-s symregexp] binary
//
// Objdumpはバイナリのすべてのテキストシンボル（コード）の逆アセンブルを表示します。
// -sオプションが指定されている場合、objdumpは正規表現に一致する名前のシンボルのみを逆アセンブルします。
//
// 代替の使用法：
//
//	go tool objdump binary start end
//
// このモードでは、objdumpは開始アドレスから終了アドレスまでのバイナリを逆アセンブルします。
// 開始アドレスと終了アドレスは16進数形式で、オプションの0xプレフィックスを付けて書かれたプログラムカウンタです。
// このモードでは、objdumpは次の形式の連続したアドレス範囲の逆アセンブルを出力します：
//
//	file:line
//	 address: assembly
//	 address: assembly
//	 ...
//
// 各節は、元のソースファイルと行番号にマップされた連続したアドレス範囲の逆アセンブルを示します。
// このモードはpprofによる使用を想定しています。
package main
