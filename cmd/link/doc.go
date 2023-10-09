// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
「リンク（Link）」は、通常「go tool link」として呼び出され、パッケージmainのGoアーカイブまたはオブジェクトとその依存関係を読み込み、それらを実行可能バイナリに結合します。

# コマンドライン

使用法:

	go tool link [フラグ] main.a

フラグ:

<<<<<<< HEAD
	-B note
		Add an ELF_NT_GNU_BUILD_ID note when using ELF.
		The value should start with 0x and be an even number of hex digits.
		Alternatively, you can pass "gobuildid" in order to derive the
		GNU build ID from the Go build ID.
	-E entry
		Set entry symbol name.
	-H type
		Set executable format type.
		The default format is inferred from GOOS and GOARCH.
		On Windows, -H windowsgui writes a "GUI binary" instead of a "console binary."
	-I interpreter
		Set the ELF dynamic linker to use.
=======
	-B メモ
		ELFを使用している場合、ELF_NT_GNU_BUILD_IDノートを追加します。
		値は0xで始まり、16進数の偶数桁でなければなりません。
	-E エントリ
		エントリシンボル名を設定します。
	-H タイプ
		実行可能フォーマットタイプを設定します。
		デフォルトのフォーマットはGOOSおよびGOARCHから推測されます。
		Windowsでは、-H windowsguiは「GUIバイナリ」ではなく「コンソールバイナリ」を書き込みます。
	-I インタプリタ
		使用するELFダイナミックリンカを設定します。
>>>>>>> release-branch.go1.21
	-L dir1 -L dir2
		$GOROOT/pkg/$GOOS_$GOARCHを参照した後、dir1、dir2などでインポートされたパッケージを検索します。
	-R quantum
		アドレスの丸め量子を設定します。
	-T address
		テキストシンボルの開始アドレスを設定します。
	-V
		リンカのバージョンを表示して終了します。
	-X importpath.name=value
		importpathの名前がvalueになる文字列変数の値を設定します。
		これは、変数がソースコードで未初期化または定数の文字列式に初期化されている場合にのみ有効です。
		イニシャライザが関数呼び出しを行うか、他の変数に参照がある場合、-Xは機能しません。
		Go 1.5より前では、このオプションは2つの別々の引数をとりました。
	-a
		出力を逆アセンブルします。
	-asan
		C/C++のアドレスサニタイザーサポートをリンクします。
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
		動的実行可能ファイルの生成を無効にします。
		出力されるコードはどちらの場合も同じです。このオプションは、動的ヘッダが含まれるかどうかだけを制御します。
		動的ヘッダはデフォルトでオンになっており、動的ライブラリへの参照は必要ありませんが、
		多くの共通のシステムツールは、ヘッダが存在することを前提としています。
	-debugtramp int
		トランポリンのデバッグ。
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
	-n
		シンボルテーブルをダンプします。
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
	-shared
		共有オブジェクトを生成します（-linkmode externalを含む; 実験的）。
	-tmpdir dir
		一時ファイルをdirに書き込みます。
		一時ファイルは外部リンキングモードでのみ使用されます。
	-u
		安全でないパッケージを拒否します。
	-v
		リンカ操作のトレースを表示します。
	-w
		DWARFシンボルテーブルを省略します。

*/package main
