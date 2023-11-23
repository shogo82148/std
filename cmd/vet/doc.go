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

	appends          append後の欠損値を確認する
	asmdecl          アセンブリファイルとGo宣言の不一致を報告する
	assign           不要な代入をチェックする
	atomic           sync/atomicパッケージを使用した一般的なミスをチェックする
	bools            ブール演算子に関連する一般的なミスをチェックする
	buildtag         正常に配置された+buildタグを確認する
	cgocall          cgoのポインタ渡しルールの違反を検出する
	composites       キーの指定されていない組み合わせリテラルをチェックする
	copylocks        値によって誤って渡されたロックをチェックする
	defers           defer文での一般的な間違いを報告します
	directive        //go:debugなどのGoツールチェーンディレクティブをチェックします
	errorsas         errors.Asに非ポインタまたは非エラー値を渡すことを報告します
	framepointer     それを保存する前にフレームポインタを破壊するアセンブリを報告します
	httpresponse     HTTPレスポンスのミスをチェックする
	ifaceassert      実行不可能なインターフェース間の型アサーションを検出します
	loopclosure      ネストされた関数内からのループ変数への参照をチェックする
	lostcancel       context.WithCancelが返すキャンセル関数が呼び出されたかをチェックする
	nilfunc          関数とnilの無駄な比較をチェックする
	printf           Printfのフォーマット文字列と引数の整合性をチェックする
	shift            整数の幅と等しいかそれを上回るシフトをチェックする
	sigchanyzer      os.Signalの非バッファードチャネルをチェックします
	slog             log/slog関数への不正な引数をチェックする
	stdmethods       cよく知られたインターフェースのメソッドのシグネチャをチェックする
	stringintconv    string(int)変換をチェックします
	structtag        構造体のフィールドタグがreflect.StructTag.Getに合致しているかをチェックする
	testinggoroutine テストによって開始されたゴルーチンからの(*testing.T).Fatalへの呼び出しを報告します
	tests            テストと例の一般的な誤った使用法をチェックする
	timeformat       (time.Time).Formatまたはtime.Parseの呼び出しで2006-02-01をチェックします
	unmarshal        ポインタでない値やインターフェースでない値がunmarshalに渡されていることを報告する
	unreachable      到達不可能なコードをチェックする
	unsafeptr        不正なuintptrからunsafe.Pointerへの変換をチェックする
	unusedresult     いくつかの関数呼び出しの未使用結果をチェックする

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
