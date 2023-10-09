// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Asmは通常「go tool asm」として呼び出され、ソースファイルをオブジェクトファイルにアセンブルします。
オブジェクトファイルは他のオブジェクトと結合してパッケージアーカイブにできます。

# コマンドライン

使用法:

	go tool asm [フラグ] ファイル

指定されたファイルはGoアセンブリファイルでなければなりません。
同じアセンブラがすべてのターゲットオペレーティングシステムとアーキテクチャで使用されます。
GOOSとGOARCHの環境変数によって目的のターゲットが設定されます。

フラグ:

	-D name[=value]
		省略可能な単純な値を持つ記号名を予め定義します。
		複数の記号を定義するために繰り返すことができます。
	-I dir1 -I dir2
		dir1、dir2などのディレクトリで#includeファイルを検索します。
		検索前に$GOROOT/pkg/$GOOS_$GOARCHを参照します。
	-S
		アセンブリとマシンコードを表示します。
	-V
		アセンブラのバージョンを表示して終了します。
	-debug
		パースされる命令をダンプします。
	-dynlink
		他の共有ライブラリで定義されたGoシンボルへの参照をサポートします。
	-e
		エラーの報告数に制限はありません。
	-gensymabis
		シンボルABI情報を出力ファイルに書き込みます。アセンブルを行いません。
	-o ファイル
		出力をファイルに書き込みます。/a/b/c/foo.sの場合、デフォルトはfoo.oです。
	-p pkgpath
		予想されるパッケージのインポートをpkgpathに設定します。
	-shared
		共有ライブラリにリンクできるコードを生成します。
	-spectre list
		リスト（all、ret）のスペクトル緩和を有効にします。
	-trimpath prefix
		記録されたソースファイルパスからプレフィックスを削除します。
	-v
		デバッグ出力を表示します。

入力言語:

アセンブラは基本的にすべてのアーキテクチャに対してほぼ同じ構文を使用しますが、
アドレッシングモードに関しては主に変化があります。入力は
#include、#define、#ifdef/endifを実装した簡略化されたCプリプロセッサを経由しますが、
#ifや##は使用されません。

詳細については、https://golang.org/doc/asmを参照してください。
*/package main
