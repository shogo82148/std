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
		リンカー固有のオブジェクトをファイルに書き込み、コンパイラ固有の
		オブジェクトを通常の出力ファイルに書き込みます（-oで指定）。
		このフラグがない場合、-oの出力はリンカーとコンパイラの入力の両方の組み合わせです。
	-m
		最適化の決定を印刷します。より高い値または繰り返し
		詳細を生成します。
	-memprofile file
		コンパイルのメモリプロファイルをファイルに書き込みます。
	-memprofilerate rate
		コンパイルのruntime.MemProfileRateをrateに設定します。
	-msan
		C/C++メモリサニタイザーへの呼び出しを挿入します。
	-mutexprofile file
		コンパイルのミューテックスプロファイルをファイルに書き込みます。
	-nolocalimports
		ローカル（相対）インポートを禁止します。
	-o file
		オブジェクトをファイルに書き込みます（デフォルトはfile.oまたは、-packがある場合はfile.a）。
	-p path
		コンパイルされるコードの予想されるパッケージインポートパスを設定し、
		循環依存関係を引き起こすインポートを診断します。
	-pack
		オブジェクトファイルではなくパッケージ（アーカイブ）ファイルを書き込みます
	-race
		レースディテクターを有効にしてコンパイルします。
	-s
		簡略化できる複合リテラルについて警告します。
	-shared
		共有ライブラリにリンクできるコードを生成します。
	-spectre list
		リスト（all, index, ret）のスペクター軽減を有効にします。
	-traceprofile file
		実行トレースをファイルに書き込みます。
	-trimpath prefix
		記録されたソースファイルパスからプレフィックスを削除します。

デバッグ情報に関連するフラグ：

	-dwarf
		DWARFシンボルを生成します。
	-dwarflocationlists
		最適化モードでDWARFにロケーションリストを追加します。
	-gendwarfinl int
		DWARFインライン情報レコードを生成します（デフォルトは2）。

コンパイラ自体をデバッグするためのフラグ：

	-E
		シンボルエクスポートをデバッグします。
	-K
		行番号が欠落していることをデバッグします。
	-d list
		リスト内のアイテムについてのデバッグ情報を印刷します。詳細については、-d helpを試してみてください。
	-live
		ライブネス分析をデバッグします。
	-v
		デバッグの詳細度を増やします。
	-%
		静的初期化子ではないものをデバッグします。
	-W
		型チェック後のパースツリーをデバッグします。
	-f
		スタックフレームをデバッグします。
	-i
		行番号スタックをデバッグします。
	-j
		ランタイムで初期化された変数をデバッグします。
	-r
		生成されたラッパーをデバッグします。
	-w
		型チェックをデバッグします。

コンパイラディレクティブ

コンパイラは、コメントの形式でディレクティブを受け入れます。
ディレクティブを非ディレクティブコメントと区別するために、ディレクティブの名前とコメントの開始の間にはスペースが必要ありません。しかし、
それらはコメントであるため、ディレクティブの規則や特定の
ディレクティブを知らないツールは、他のコメントと同様にディレクティブをスキップできます。
*/
// ラインディレクティブはいくつかの形式で存在します：
//
// 	//line :line
// 	//line :line:col
// 	//line filename:line
// 	//line filename:line:col
// 	/*line :line*/
// 	/*line :line:col*/
// 	/*line filename:line*/
// 	/*line filename:line:col*/
//
// ラインディレクティブとして認識されるためには、コメントは
// //lineまたは/*lineに続くスペースで始まり、少なくとも一つのコロンを含んでいなければなりません。
// //line形式は行の始まりでなければなりません。
// ラインディレクティブは、コメントの直後の文字のソース位置を指定したファイル、行、列から来たものとして指定します：
// //lineコメントの場合、これは次の行の最初の文字であり、
// /*lineコメントの場合、これは閉じる*/の直後の文字位置です。
// ファイル名が指定されていない場合、記録されたファイル名は列番号もない場合は空であり、
// それ以外の場合は最近記録されたファイル名（実際のファイル名または前のラインディレクティブで指定されたファイル名）です。
// ラインディレクティブが列番号を指定していない場合、列は"未知"であり、
// 次のディレクティブまでコンパイラはその範囲の列番号を報告しません。
// ラインディレクティブのテキストは後ろから解釈されます：まず、dddが有効な数値> 0である場合、
// ディレクティブテキストから末尾の:dddが剥がされます。次に、それが有効である場合、
// 同じ方法で2番目の:dddが剥がされます。それ以前のものはすべてファイル名と見なされます
// （空白とコロンを含む可能性があります）。無効な行または列の値はエラーとして報告されます。
//
// 例：
//
//	//line foo.go:10      ファイル名はfoo.goで、次の行の行番号は10です
//	//line C:foo.go:10    ファイル名にはコロンが許可されています。ここではファイル名はC:foo.goで、行は10です
//	//line  a:100 :10     ファイル名には空白が許可されています。ここではファイル名は " a:100 "（引用符を除く）
//	/*line :10:20*/x      xの位置は現在のファイル内で、行番号は10、列番号は20です
//	/*line foo: 10 */     このコメントは無効な行ディレクティブとして認識されます（行番号の周囲に余分な空白があります）
//
// ラインディレクティブは通常、機械生成されたコードに現れます。これにより、コンパイラとデバッガは
// ジェネレータへの元の入力の位置を報告します。
/*
ラインディレクティブは歴史的な特例であり、他のすべてのディレクティブは
//go:nameの形式で、それらがGoツールチェーンによって定義されていることを示しています。
各ディレクティブは自身の行に配置する必要があり、コメントの前には先頭のスペースとタブのみが許可されます。
各ディレクティブは、それに直後に続くGoコードに適用され、
通常は宣言でなければなりません。

	//go:noescape

//go:noescape ディレクティブは、本体を持たない関数宣言（つまり、Goで書かれていない実装を持つ関数）に続く必要があります。
これは、関数が引数として渡されたポインタをヒープに逃がしたり、関数から返される値に逃がしたりしないことを指定します。
この情報は、関数を呼び出すGoコードのコンパイラのエスケープ解析中に使用できます。

	//go:uintptrescapes

//go:uintptrescapes ディレクティブは、関数宣言に続く必要があります。
これは、関数のuintptr引数がuintptrに変換されたポインタ値であり、
呼び出しの期間中、型だけから見てオブジェクトが呼び出し中には必要ないように見える場合でも、
ヒープ上に保持され、生きていなければならないことを指定します。
ポインタからuintptrへの変換は、この関数への任意の呼び出しの引数リストに現れなければなりません。
このディレクティブは、一部の低レベルシステムコールの実装に必要であり、それ以外の場合は避けるべきです。

	//go:noinline

//go:noinline ディレクティブは、関数宣言に続く必要があります。
これは、関数への呼び出しがインライン化されないように指定し、
コンパイラの通常の最適化ルールを上書きします。これは通常、
特別なランタイム関数やコンパイラのデバッグ時にのみ必要です。

	//go:norace

//go:norace ディレクティブは、関数宣言に続く必要があります。
これは、関数のメモリアクセスがレース検出器によって無視されるべきであることを指定します。
これは最も一般的に、レース検出器ランタイムに呼び出すことが安全でない時期に呼び出される低レベルのコードで使用されます。

	//go:nosplit

//go:nosplitディレクティブは、関数宣言に続く必要があります。
これは、関数が通常のスタックオーバーフローチェックを省略する必要があることを指定します。
これは最も一般的に、呼び出し元のゴルーチンがプリエンプトされるのが安全でない時期に呼び出される低レベルのランタイムコードで使用されます。

	//go:linkname localname [importpath.name]

//go:linknameディレクティブは、通常、「localname」で指定されたvarまたはfunc宣言の前に配置されますが、
その位置はその効果を変えません。
このディレクティブは、Goのvarまたはfunc宣言に使用されるオブジェクトファイルシンボルを決定し、
2つのGoシンボルが同じオブジェクトファイルシンボルをエイリアスとして使用できるようにします。
これにより、一つのパッケージが、通常は未エクスポート宣言のカプセル化を侵害する、
または型安全性を侵害する場合でも、別のパッケージのシンボルにアクセスできます。
そのため、"unsafe"をインポートしたファイルでのみ有効になります。

二つのシナリオで使用することができます。パッケージupperがパッケージlowerを
インポートしていると仮定しましょう、おそらく間接的に。最初のシナリオでは、
パッケージlowerは、そのオブジェクトファイル名がパッケージupperに属するシンボルを定義します。
両方のパッケージにはlinknameディレクティブが含まれています：パッケージlowerは
二つの引数形式を使用し、パッケージupperは一つの引数形式を使用します。
以下の例では、lower.fは関数upper.gのエイリアスです：

    package upper
    import _ "unsafe"
    //go:linkname g
    func g()

    package lower
    import _ "unsafe"
    //go:linkname f upper.g
    func f() { ... }

パッケージupperのlinknameディレクティブは、本体を持たない関数に対する通常のエラーを抑制します。
（そのチェックは、パッケージに.sファイル（空でも可）を含めることで、代わりに抑制することもできます。）

二つ目のシナリオでは、パッケージupperが一方的にパッケージlowerのシンボルのエイリアスを作成します。
以下の例では、upper.gは関数lower.fのエイリアスです。

    package upper
    import _ "unsafe"
    //go:linkname g lower.f
    func g()

    package lower
    func f() { ... }

lower.fの宣言には、単一の引数fを持つlinknameディレクティブも含まれているかもしれません。
これはオプションですが、関数がパッケージ外からアクセスされることを読者に警告するのに役立ちます。

	//go:wasmimport importmodule importname

//go:wasmimportディレクティブはwasm専用で、関数宣言に続く必要があります。
これは、関数が``importmodule``と``importname``で識別されるwasmモジュールによって提供されることを指定します。

	//go:wasmimport a_module f
	func g()

Go関数のパラメータと戻り値の型は、以下の表に従ってWasmに変換されます：

    Go types        Wasm types
    int32, uint32   i32
    int64, uint64   i64
    float32         f32
    float64         f64
    unsafe.Pointer  i32

コンパイラは他のすべてのパラメータ型を許可しません。

*/
package main
