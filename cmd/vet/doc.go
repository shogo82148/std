// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
VetはGoのソースコードを検査し、Printfのような呼び出しがフォーマット文字列と一致しない場合に疑わしい構造を報告します。Vetは完全な報告を保証するわけではないヒューリスティックを使用しているため、すべての報告が本当の問題ではないかもしれませんが、コンパイラでは見つからないエラーを見つけることができます。

Vetは通常、goコマンドを通じて起動されます。
現在のディレクトリのパッケージを検査するためには、以下のコマンドを使用します:

	go vet

パスが指定されたパッケージを検査するためには、以下のコマンドを使用します:

	go vet my/project/...

パッケージを指定する他の方法については「go help packages」を参照してください。

Vetの終了コードは、ツールの誤った呼び出しまたは問題が報告された場合に非ゼロであり、それ以外の場合は0です。ツールはすべての可能な問題をチェックせず、信頼性の低いヒューリスティックに依存しているため、プログラムの正確性の厳密な指標ではなく、ガイダンスとして使用する必要があります。

利用可能なチェックをリストするには、「go tool vet help」と実行します:

	appends          appendの後の不足値をチェック
	asmdecl          アセンブリファイルとGoの宣言の不一致を報告
	assign           無駄な代入をチェック
	atomic           sync/atomicパッケージの一般的な間違いをチェック
	bools            真偽値演算子の一般的な間違いをチェック
	buildtag         //go:buildと// +buildディレクティブをチェック
	cgocall          cgoポインタ渡しルールの違反を検出
	composites       キーなしの複合リテラルをチェック
	copylocks        値で誤って渡されたロックをチェック
	defers           defer文の一般的な間違いを報告
	directive        //go:debugなどのGoツールチェーンディレクティブをチェック
	errorsas         errors.Asに非ポインタまたは非エラー値を渡すことを報告
	framepointer     フレームポインタを保存前に破壊するアセンブリを報告
	hostport         net.Dialに渡されるアドレスの形式をチェック
	httpresponse     HTTP応答使用の間違いをチェック
	ifaceassert      不可能なインターフェース間の型アサーションを検出
	loopclosure      ネストした関数内からのループ変数への参照をチェック
	lostcancel       context.WithCancelが返すcancel関数の呼び出しをチェック
	nilfunc          関数とnilの無駄な比較をチェック
	printf           Printfフォーマット文字列と引数の整合性をチェック
	shift            整数の幅以上のシフトをチェック
	sigchanyzer      os.Signalのバッファなしチャネルをチェック
	slog             無効な構造化ログ呼び出しをチェック
	stdmethods       よく知られたインターフェースのメソッドのシグネチャをチェック
	stdversion       新しすぎる標準ライブラリシンボルの使用を報告
	stringintconv    string(int)変換をチェック
	structtag        構造体フィールドタグがreflect.StructTag.Getに準拠するかチェック
	testinggoroutine テストが開始したgoroutineからの(*testing.T).Fatal呼び出しを報告
	tests            テストと例の一般的な誤用をチェック
	timeformat       2006-02-01を使った(time.Time).FormatやTime.Parse呼び出しをチェック
	unmarshal        unmarshalに非ポインタまたは非インターフェース値を渡すことを報告
	unreachable      到達不可能なコードをチェック
	unsafeptr        uintptrからunsafe.Pointerへの無効な変換をチェック
	unusedresult     一部の関数呼び出しの未使用結果をチェック
	waitgroup        sync.WaitGroupの誤用をチェック

printfなどの特定のチェックの詳細とフラグについての情報は、「go tool vet help printf」と実行してください。

デフォルトでは、すべてのチェックが実行されます。
フラグがtrueに明示的に設定されている場合、それらのテストのみが実行されます。
逆に、フラグが明示的にfalseに設定されている場合、それらのテストは無効になります。
したがって、-printf=trueはprintfチェックを実行し、
-printf=falseはprintfチェック以外のすべてのチェックを実行します。

新しいチェックの作成方法については、golang.org/x/tools/go/analysisを参照してください。

コアフラグ:

	-c=N
	  	エラーのある行とその周囲のN行を表示する
	-json
	  	分析診断（およびエラー）をJSON形式で出力する
*/
package main
