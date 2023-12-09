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
<<<<<<< HEAD
		importpathの名前がvalueになる文字列変数の値を設定します。
		これは、変数がソースコードで未初期化または定数の文字列式に初期化されている場合にのみ有効です。
		イニシャライザが関数呼び出しを行うか、他の変数に参照がある場合、-Xは機能しません。
		Go 1.5より前では、このオプションは2つの別々の引数をとりました。
	-a
		出力を逆アセンブルします。
	-asan
		C/C++のアドレスサニタイザーサポートをリンクします。
=======
		Set the value of the string variable in importpath named name to value.
		This is only effective if the variable is declared in the source code either uninitialized
		or initialized to a constant string expression. -X will not work if the initializer makes
		a function call or refers to other variables.
		Note that before Go 1.5 this option took two separate arguments.
	-asan
		Link with C/C++ address sanitizer support.
	-aslr
		Enable ASLR for buildmode=c-shared on windows (default true).
>>>>>>> upstream/master
	-buildid id
		GoツールチェインのビルドIDとしてidを記録します。
	-buildmode mode
		ビルドモードを設定します（デフォルトはexe）。
	-c
		コールグラフをダンプします。
	-compressdwarf
		可能な場合はDWARFを圧縮します（デフォルトはtrue）。
	-cpuprofile file
		CPUプロファイルをfileに書き込みます。
	-d
<<<<<<< HEAD
		動的実行可能ファイルの生成を無効にします。
		出力されるコードはどちらの場合も同じです。このオプションは、動的ヘッダが含まれるかどうかだけを制御します。
		動的ヘッダはデフォルトでオンになっており、動的ライブラリへの参照は必要ありませんが、
		多くの共通のシステムツールは、ヘッダが存在することを前提としています。
	-debugtramp int
		トランポリンのデバッグ。
=======
		Disable generation of dynamic executables.
		The emitted code is the same in either case; the option
		controls only whether a dynamic header is included.
		The dynamic header is on by default, even without any
		references to dynamic libraries, because many common
		system tools now assume the presence of the header.
>>>>>>> upstream/master
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
<<<<<<< HEAD
		C/C++のメモリサニタイザーサポートをリンクします。
	-n
		シンボルテーブルをダンプします。
=======
		Link with C/C++ memory sanitizer support.
>>>>>>> upstream/master
	-o file
		出力をfileに書き込みます（デフォルトはa.out、Windowsではa.out.exe）。
	-pluginpath path
		エクスポートされたプラグインシンボルの接頭辞として使用されるパス名です。
	-r dir1:dir2：...
		ELFダイナミックリンカの検索パスを設定します。
	-race
		レース検出ライブラリとリンクします。
	-s
<<<<<<< HEAD
		シンボルテーブルとデバッグ情報を省略します。
	-shared
		共有オブジェクトを生成します（-linkmode externalを含む; 実験的）。
	-tmpdir dir
		一時ファイルをdirに書き込みます。
		一時ファイルは外部リンキングモードでのみ使用されます。
	-u
		安全でないパッケージを拒否します。
=======
		Omit the symbol table and debug information.
	-tmpdir dir
		Write temporary files to dir.
		Temporary files are only used in external linking mode.
>>>>>>> upstream/master
	-v
		リンカ操作のトレースを表示します。
	-w
		DWARFシンボルテーブルを省略します。

*/package main
