// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
コンパイルは通常、 ``go tool compile'' として呼び出され、コマンドラインで指定されたファイルの名前を持つ単一のGoパッケージをコンパイルします。それはその後、最初のソースファイルのベース名に.oの接尾辞を付けた単一のオブジェクトファイルを書き込みます。オブジェクトファイルは、他のオブジェクトと組み合わせてパッケージアーカイブに結合するか、直接リンカ（ ``go tool link''）に渡すことができます。-packを使用して呼び出された場合、コンパイラは中間オブジェクトファイルを経由せずに直接アーカイブを書き込みます。

生成されたファイルには、パッケージがエクスポートするシンボルに関する型情報と、パッケージが他のパッケージからインポートされたシンボルで使用される型に関する情報が含まれています。したがって、パッケージPのクライアントCをコンパイルする場合、Pの依存関係のファイルを読み込む必要はありません。コンパイルされたPの出力のみが必要です。

コマンドライン

使用法：

	go tool compile [フラグ] ファイル...

指定されたファイルはGoのソースファイルでなければなりません。すべて同じパッケージの一部です。
すべてのターゲットオペレーティングシステムとアーキテクチャには同じコンパイラが使用されます。
GOOSとGOARCHの環境変数が目標となるものを設定します。

フラグ：

	-D パス
		ローカルインポートの相対パスを設定します。
	-I dir1 -I dir2
		dir1、dir2などのインポートされたパッケージを検索します。
		この後、$GOROOT/pkg/$GOOS_$GOARCHを参照します。
	-L
		エラーメッセージに完全なファイルパスを表示します。
	-N
		最適化を無効にします。
	-S
		アセンブリリストを標準出力に表示します（コードのみ）。
	-S -S
		アセンブリリスト（コードとデータ）を標準出力に表示します。
	-V
		コンパイラのバージョンを表示して終了します。
	-asmhdr ファイル
		アセンブリヘッダをファイルに書き込みます。
	-asan
		C/C++アドレスサニタイザへの呼び出しを挿入します。
	-buildid ID
		エクスポートメタデータのビルドIDとしてIDを記録します。
	-blockprofile ファイル
		コンパイルのためのブロックプロファイルをファイルに書き込みます。
	-c int
		コンパイル中の並行性を設定します。並行性を行わない場合は1を設定します（デフォルトは1）。
	-complete
		パッケージに非Goコンポーネントがないと想定します。
	-cpuprofile ファイル
		コンパイルのためのCPUプロファイルをファイルに書き込みます。
	-dynlink
		共有ライブラリ内のGoシンボルへの参照を許可します（実験的）。
	-e
		エラーの数に制限を解除します（デフォルトの制限は10です）。
	-goversion string
		ランタイムの必要なgoツールバージョンを指定します。
		ランタイムのgoバージョンがgoversionと一致しない場合は終了します。
	-h
		最初のエラーが検出されたときにスタックトレースで停止します。
	-importcfg ファイル
		インポート構成をファイルから読み取ります。
		ファイルでは、importmap、packagefileを設定してインポートの解決を指定します。
	-installsuffix suffix
		$GOROOT/pkg/$GOOS_$GOARCH_suffixのかわりに$GOROOT/pkg/$GOOS_$GOARCHでパッケージを検索します。
	-l
		インライニングを無効にします。
	-lang version
		コンパイルするための言語バージョンを設定します（-lang=go1.12など）。
		デフォルトは現在のバージョンです。
	-linkobj file
		リンカ固有のオブジェクトをファイルに書き込み、コンパイラ固有のオブジェクト
		オブジェクトを通常の出力ファイル（-oで指定されたファ*/package main
