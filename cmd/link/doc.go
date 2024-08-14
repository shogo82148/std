// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
「リンク（Link）」は、通常「go tool link」として呼び出され、パッケージmainのGoアーカイブまたはオブジェクトとその依存関係を読み込み、それらを実行可能バイナリに結合します。

# コマンドライン

使用法:

	go tool link [フラグ] main.a

フラグ:

	- Bノート
		ELFを使用する場合、ELF_NT_GNU_BUILD_IDノートを追加します。
		値は0xで始まり、偶数桁の16進数である必要があります。
		代わりに、GoビルドIDからGNUビルドIDを派生させるために「gobuildid」を渡すこともできます。
	-E エントリ
		エントリシンボル名を設定します。
	-H タイプ
		実行可能フォーマットタイプを設定します。
		デフォルトのフォーマットはGOOSおよびGOARCHから推測されます。
		Windowsでは、-H windowsguiは「GUIバイナリ」ではなく「コンソールバイナリ」を書き込みます。
	-I インタプリタ
		使用するELFダイナミックリンカを設定します。
	-L dir1 -L dir2
		$GOROOT/pkg/$GOOS_$GOARCHを参照した後、dir1、dir2などでインポートされたパッケージを検索します。
	-R quantum
		アドレスの丸め量子を設定します。
	-T address
		テキストシンボルの開始アドレスを設定します。
	-V
		リンカのバージョンを表示して終了します。
	-X importpath.name=value
		importpath内のnameという名前の文字列変数の値をvalueに設定します。
		これは、変数がソースコード内で未初期化または定数文字列式に初期化されて宣言されている場合にのみ有効です。
		-Xは、初期化子が関数呼び出しを行うか、他の変数を参照する場合には機能しません。
		Go 1.5より前では、このオプションは2つの別々の引数を取りました。
	-asan
		C/C++アドレスサニタイザーサポートとリンクします。
	-aslr
		Enable ASLR for buildmode=c-shared on windows (default true).
	-bindnow
		Mark a dynamically linked ELF object for immediate function binding (default false).
	-buildid id
		GoツールチェインのビルドIDとしてidを記録します。
	-buildmode mode
		ビルドモードを設定します（デフォルトはexe）。
	-c
		コールグラフをダンプします。
	-checklinkname=value
		If value is 0, all go:linkname directives are permitted.
		If value is 1 (the default), only a known set of widely-used
		linknames are permitted.
	-compressdwarf
		可能な場合はDWARFを圧縮します（デフォルトはtrue）。
	-cpuprofile file
		CPUプロファイルをfileに書き込みます。
	-d
		動的実行可能ファイルの生成を無効にします。
		出力されるコードはどちらの場合も同じです。このオプションは
		動的ヘッダーが含まれるかどうかだけを制御します。
		動的ヘッダーは、多くの一般的な
		システムツールがヘッダーの存在を前提としているため、
		動的ライブラリへの参照がなくてもデフォルトでオンになっています。
	-dumpdep
		シンボル依存関係グラフをダンプします。
	-extar ar
		外部アーカイブプログラムを設定します（デフォルトは「ar」）。
		-buildmode=c-archiveにのみ使用されます。
	-extld linker
		外部リンカを設定します（デフォルトは「clang」または「gcc」）。
	-extldflags flags
		外部リンカに渡すスペース区切りのフラグを設定します。
	-f
		リンクされたアーカイブのバージョンの不一致を無視します。
	-g
		Goパッケージデータのチェックを無効にします。
	-importcfg file
		ファイルからインポート構成を読み込みます。
		ファイルで、packagefile、packageshlibを設定してインポート解決を指定します。
	-installsuffix suffix
		$GOROOT/pkg/$GOOS_$GOARCH_suffixでパッケージを検索します。
		$GOROOT/pkg/$GOOS_$GOARCHではなく。
	-k symbol
		フィールドトラッキングシンボルを設定します。GOEXPERIMENT=fieldtrackが設定されている場合にこのフラグを使用します。
	-libgcc file
		コンパイラサポートライブラリの名前を設定します。
		これは内部リンクモードでのみ使用されます。
		設定されていない場合、デフォルトの値はコンパイラの実行結果から取得されます。
		-extldオプションで設定できます。
		サポートライブラリを使用しない場合は「none」を設定します。
	-linkmode mode
		リンクモードを設定します（internal、external、auto）。
		これはcmd/cgo/doc.goで説明されているリンクモードを設定します。
	-linkshared
		インストールされたGo共有ライブラリとリンクします（実験的）。
	-memprofile file
		メモリプロファイルをfileに書き込みます。
	-memprofilerate rate
		runtime.MemProfileRateをrateに設定します。
	-msan
		C/C++のメモリサニタイザーサポートをリンクします。
	-o file
		出力をfileに書き込みます（デフォルトはa.out、Windowsではa.out.exe）。
	-pluginpath path
		エクスポートされたプラグインシンボルの接頭辞として使用されるパス名です。
	-r dir1:dir2：...
		ELFダイナミックリンカの検索パスを設定します。
	-race
		レース検出ライブラリとリンクします。
	-s
		シンボルテーブルとデバッグ情報を省略します。
	-tmpdir dir
		一時ファイルをdirに書き込みます。
		一時ファイルは外部リンキングモードでのみ使用されます。
	-v
		リンカ操作のトレースを表示します。
	-w
		DWARFシンボルテーブルを省略します。

*/package main
